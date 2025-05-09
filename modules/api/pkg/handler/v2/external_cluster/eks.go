/*
Copyright 2021 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package externalcluster

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	ekstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/go-kit/kit/endpoint"

	apiv1 "k8c.io/dashboard/v2/pkg/api/v1"
	apiv2 "k8c.io/dashboard/v2/pkg/api/v2"
	handlercommon "k8c.io/dashboard/v2/pkg/handler/common"
	providercommon "k8c.io/dashboard/v2/pkg/handler/common/provider"
	"k8c.io/dashboard/v2/pkg/handler/v1/common"
	"k8c.io/dashboard/v2/pkg/provider"
	awsprovider "k8c.io/dashboard/v2/pkg/provider/cloud/aws"
	eksprovider "k8c.io/dashboard/v2/pkg/provider/cloud/eks"
	kubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/kubermatic/v1"
	kuberneteshelper "k8c.io/kubermatic/v2/pkg/kubernetes"
	"k8c.io/kubermatic/v2/pkg/resources"
	utilerrors "k8c.io/kubermatic/v2/pkg/util/errors"
	clusterv1alpha1 "k8c.io/machine-controller/sdk/apis/cluster/v1alpha1"
	"k8c.io/machine-controller/sdk/providerconfig"

	"k8s.io/utils/ptr"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	RegionEndpoint   = "eu-central-1"
	EKSAMITypes      = "Amazon Linux 2"
	EKSCustomAMIType = "CUSTOM"
	EKSCapacityTypes = "SPOT"
)

// EKSProjectReq represents a request with a project context.
// swagger:parameters listProjectEKSVersions listProjectEKSAMITypes listProjectEKSCapacityTypes
type EKSProjectReq struct {
	common.ProjectReq
}

func DecodeEKSProjectReq(c context.Context, r *http.Request) (interface{}, error) {
	req, err := common.DecodeProjectRequest(c, r)
	if err != nil {
		return nil, err
	}

	return EKSProjectReq{
		ProjectReq: req.(common.ProjectReq),
	}, nil
}

// EKSCommonReq represent a request with common parameters for EKS.
// swagger:parameters listEKSRegions listEKSClusterRoles
type EKSCommonReq struct {
	// in: header
	// name: AccessKeyID
	AccessKeyID string
	// in: header
	// name: SecretAccessKey
	SecretAccessKey string
	// in: header
	// name: Credential
	Credential string
	// in: header
	// name: AssumeRoleARN
	AssumeRoleARN string
	// in: header
	// name: AssumeRoleExternalID
	AssumeRoleExternalID string
}

// EKSProjectCommonReq represents a request with common parameters and project context for EKS.
// swagger:parameters listProjectEKSRegions listProjectEKSClusterRoles
type EKSProjectCommonReq struct {
	common.ProjectReq
	EKSCommonReq
}

func DecodeEKSCommonReq(c context.Context, r *http.Request) (interface{}, error) {
	var req EKSCommonReq

	req.AccessKeyID = r.Header.Get("AccessKeyID")
	req.SecretAccessKey = r.Header.Get("SecretAccessKey")
	req.Credential = r.Header.Get("Credential")
	req.AssumeRoleARN = r.Header.Get("AssumeRoleARN")
	req.AssumeRoleExternalID = r.Header.Get("AssumeRoleExternalID")

	return req, nil
}

func DecodeEKSProjectCommonReq(c context.Context, r *http.Request) (interface{}, error) {
	commonReq, err := DecodeEKSCommonReq(c, r)
	if err != nil {
		return nil, err
	}

	projectReq, err := common.DecodeProjectRequest(c, r)
	if err != nil {
		return nil, err
	}

	return EKSProjectCommonReq{
		ProjectReq:   projectReq.(common.ProjectReq),
		EKSCommonReq: commonReq.(EKSCommonReq),
	}, nil
}

// EKSRegionReq represent a request that includes region information.
// swagger:parameters validateEKSCredentials listEKSVPCS
type EKSRegionReq struct {
	EKSCommonReq
	// in: header
	// name: Region
	Region string
}

// EKSProjectRegionReq represents a request that includes region information with project context.
// swagger:parameters validateProjectEKSCredentials listProjectEKSVPCs listEKSClusters listProjectEKSClusterRoles
type EKSProjectRegionReq struct {
	common.ProjectReq
	EKSRegionReq
}

func DecodeEKSRegionReq(c context.Context, r *http.Request) (interface{}, error) {
	var req EKSRegionReq

	commonReq, err := DecodeEKSCommonReq(c, r)
	if err != nil {
		return nil, err
	}
	req.Region = r.Header.Get("Region")
	req.EKSCommonReq = commonReq.(EKSCommonReq)

	return req, nil
}

func DecodeEKSProjectRegionReq(c context.Context, r *http.Request) (interface{}, error) {
	regionReq, err := DecodeEKSRegionReq(c, r)
	if err != nil {
		return nil, err
	}

	projectReq, err := common.DecodeProjectRequest(c, r)
	if err != nil {
		return nil, err
	}

	return EKSProjectRegionReq{
		ProjectReq:   projectReq.(common.ProjectReq),
		EKSRegionReq: regionReq.(EKSRegionReq),
	}, nil
}

func (req EKSCommonReq) Validate() error {
	if len(req.Credential) != 0 {
		return nil
	}
	if len(req.AccessKeyID) == 0 || len(req.SecretAccessKey) == 0 {
		return fmt.Errorf("EKS Credentials cannot be empty")
	}
	return nil
}

func (req EKSRegionReq) Validate() error {
	if len(req.Region) == 0 {
		return fmt.Errorf("region cannot be empty")
	}
	if err := req.EKSCommonReq.Validate(); err != nil {
		return err
	}
	return nil
}

func (req EKSVPCReq) Validate() error {
	if len(req.VpcId) == 0 {
		return fmt.Errorf("EKS VPC ID cannot be empty")
	}
	if err := req.EKSRegionReq.Validate(); err != nil {
		return err
	}
	return nil
}

// eksNoCredentialReq represent a request for EKS resources
// swagger:parameters listEKSVPCsNoCredentials listEKSInstanceTypesNoCredentials listEKSNodeRolesNoCredentials
type eksNoCredentialReq struct {
	GetClusterReq
}

func DecodeEKSNoCredentialReq(c context.Context, r *http.Request) (interface{}, error) {
	var req eksNoCredentialReq
	re, err := DecodeGetReq(c, r)
	if err != nil {
		return nil, err
	}
	req.GetClusterReq = re.(GetClusterReq)

	return req, nil
}

// eksNoCredentialSizeReq represent a request for EKS VM sizes.
// swagger:parameters listEKSInstanceTypesNoCredentials
type eksNoCredentialSizeReq struct {
	GetClusterReq

	// architecture query parameter. Supports: arm64 and x86_64 types.
	// in: query
	Architecture string `json:"architecture,omitempty"`
}

func DecodeEKSNoCredentialSizeReq(c context.Context, r *http.Request) (interface{}, error) {
	var req eksNoCredentialSizeReq
	re, err := DecodeGetReq(c, r)
	if err != nil {
		return nil, err
	}
	req.GetClusterReq = re.(GetClusterReq)

	req.Architecture = r.URL.Query().Get("architecture")
	if len(req.Architecture) > 0 {
		if req.Architecture == handlercommon.EKSARM64Architecture || req.Architecture == handlercommon.EKSX86_64Architecture {
			return req, nil
		}
		return nil, fmt.Errorf("wrong query parameter, unsupported architecture: %s", req.Architecture)
	}

	return req, nil
}

func (req eksNoCredentialSizeReq) Validate() error {
	if err := req.GetClusterReq.Validate(); err != nil {
		return err
	}

	return nil
}

// eksSubnetsNoCredentialReq represent a request for EKS resources
// swagger:parameters listEKSSubnetsNoCredentials
type eksSubnetsNoCredentialReq struct {
	eksNoCredentialReq
	// in: header
	// name: VpcId
	VpcId string
}

func DecodeEKSSubnetsNoCredentialReq(c context.Context, r *http.Request) (interface{}, error) {
	var req eksSubnetsNoCredentialReq
	re, err := DecodeEKSNoCredentialReq(c, r)
	if err != nil {
		return nil, err
	}
	req.eksNoCredentialReq = re.(eksNoCredentialReq)
	req.VpcId = r.Header.Get("VpcId")

	return req, nil
}

// Validate validates eksSubnetsNoCredentialReq request.
func (req eksSubnetsNoCredentialReq) Validate() error {
	if err := req.GetClusterReq.Validate(); err != nil {
		return err
	}
	if len(req.VpcId) == 0 {
		return fmt.Errorf("EKS VPC ID cannot be empty")
	}
	return nil
}

func ListEKSClustersEndpoint(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider, privilegedProjectProvider provider.PrivilegedProjectProvider, clusterProvider provider.ExternalClusterProvider, presetProvider provider.PresetProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(EKSProjectRegionReq)
		if !ok {
			return nil, utilerrors.NewBadRequest("invalid request")
		}
		if err := req.Validate(); err != nil {
			return nil, utilerrors.NewBadRequest("%v", err)
		}

		credential, err := getEKSCredentialsFromReq(ctx, req.EKSRegionReq, userInfoGetter, presetProvider, req.ProjectID)
		if err != nil {
			return nil, err
		}

		return providercommon.ListEKSClusters(ctx, projectProvider, privilegedProjectProvider, userInfoGetter, clusterProvider, *credential, req.ProjectID)
	}
}

func ListEKSVPCEndpoint(userInfoGetter provider.UserInfoGetter, presetProvider provider.PresetProvider, withProject bool) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req       EKSRegionReq
			projectID string
		)

		if !withProject {
			regionReq, ok := request.(EKSRegionReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}
			req = regionReq
		} else {
			projectReq, ok := request.(EKSProjectRegionReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}

			projectID = projectReq.GetProjectID()
			req = projectReq.EKSRegionReq
		}

		if err := req.Validate(); err != nil {
			return nil, utilerrors.NewBadRequest("%v", err)
		}

		credential, err := getEKSCredentialsFromReq(ctx, req, userInfoGetter, presetProvider, projectID)
		if err != nil {
			return nil, err
		}

		return providercommon.ListEKSVPC(ctx, *credential)
	}
}

func ListEKSSubnetsEndpoint(userInfoGetter provider.UserInfoGetter, presetProvider provider.PresetProvider, withProject bool) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req       EKSVPCReq
			projectID string
		)

		if !withProject {
			vpcReq, ok := request.(EKSVPCReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}
			req = vpcReq
		} else {
			projectReq, ok := request.(EKSProjectVPCReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}

			projectID = projectReq.GetProjectID()
			req = projectReq.EKSVPCReq
		}

		if err := req.Validate(); err != nil {
			return nil, utilerrors.NewBadRequest("%v", err)
		}

		credential, err := getEKSCredentialsFromReq(ctx, req.EKSRegionReq, userInfoGetter, presetProvider, projectID)
		if err != nil {
			return nil, err
		}

		return providercommon.ListEKSSubnetIDs(ctx, *credential, req.VpcId)
	}
}

func ListEKSSecurityGroupsEndpoint(userInfoGetter provider.UserInfoGetter, presetProvider provider.PresetProvider, withProject bool) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req       EKSVPCReq
			projectID string
		)

		if !withProject {
			vpcReq, ok := request.(EKSVPCReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}
			req = vpcReq
		} else {
			projectReq, ok := request.(EKSProjectVPCReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}

			projectID = projectReq.GetProjectID()
			req = projectReq.EKSVPCReq
		}

		if err := req.Validate(); err != nil {
			return nil, utilerrors.NewBadRequest("%v", err)
		}

		credential, err := getEKSCredentialsFromReq(ctx, req.EKSRegionReq, userInfoGetter, presetProvider, projectID)
		if err != nil {
			return nil, err
		}

		return providercommon.ListEKSSecurityGroup(ctx, *credential, req.VpcId)
	}
}

func ListEKSRegionsEndpoint(userInfoGetter provider.UserInfoGetter, presetProvider provider.PresetProvider, withProject bool) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req       EKSCommonReq
			projectID string
		)

		if !withProject {
			commonReq, ok := request.(EKSCommonReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}
			req = commonReq
		} else {
			projectReq, ok := request.(EKSProjectCommonReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}

			projectID = projectReq.GetProjectID()
			req = projectReq.EKSCommonReq
		}

		if err := req.Validate(); err != nil {
			return nil, utilerrors.NewBadRequest("%v", err)
		}

		eksTypesReq := EKSRegionReq{
			EKSCommonReq: req,
			Region:       RegionEndpoint,
		}
		credential, err := getEKSCredentialsFromReq(ctx, eksTypesReq, userInfoGetter, presetProvider, projectID)
		if err != nil {
			return nil, err
		}

		return providercommon.ListEKSRegions(ctx, *credential)
	}
}

func ListEKSClusterRolesEndpoint(userInfoGetter provider.UserInfoGetter, presetProvider provider.PresetProvider, withProject bool) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req       EKSCommonReq
			projectID string
		)

		if !withProject {
			commonReq, ok := request.(EKSCommonReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}
			req = commonReq
		} else {
			projectReq, ok := request.(EKSProjectCommonReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}

			projectID = projectReq.GetProjectID()
			req = projectReq.EKSCommonReq
		}

		if err := req.Validate(); err != nil {
			return nil, utilerrors.NewBadRequest("%v", err)
		}

		eksTypesReq := EKSRegionReq{
			EKSCommonReq: req,
			Region:       RegionEndpoint,
		}
		credential, err := getEKSCredentialsFromReq(ctx, eksTypesReq, userInfoGetter, presetProvider, projectID)
		if err != nil {
			return nil, err
		}

		return providercommon.ListEKSClusterRoles(ctx, *credential)
	}
}

func ListEKSNodeRolesEndpoint(userInfoGetter provider.UserInfoGetter, presetProvider provider.PresetProvider, withProject bool) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req       EKSRegionReq
			projectID string
		)

		if !withProject {
			regionReq, ok := request.(EKSRegionReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}
			req = regionReq
		} else {
			projectReq, ok := request.(EKSProjectRegionReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}

			projectID = projectReq.GetProjectID()
			req = projectReq.EKSRegionReq
		}

		if err := req.Validate(); err != nil {
			return nil, utilerrors.NewBadRequest("%v", err)
		}

		credential, err := getEKSCredentialsFromReq(ctx, req, userInfoGetter, presetProvider, projectID)
		if err != nil {
			return nil, err
		}

		return providercommon.ListEKSNodeRoles(ctx, *credential)
	}
}

func EKSValidateCredentialsEndpoint(presetProvider provider.PresetProvider, userInfoGetter provider.UserInfoGetter, withProject bool) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req       EKSRegionReq
			projectID string
		)

		if !withProject {
			regionReq, ok := request.(EKSRegionReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}
			req = regionReq
		} else {
			projectReq, ok := request.(EKSProjectRegionReq)
			if !ok {
				return nil, utilerrors.NewBadRequest("invalid request")
			}

			projectID = projectReq.GetProjectID()
			req = projectReq.EKSRegionReq
		}

		if err := req.Validate(); err != nil {
			return nil, utilerrors.NewBadRequest("%v", err)
		}

		credential, err := getEKSCredentialsFromReq(ctx, req, userInfoGetter, presetProvider, projectID)
		if err != nil {
			return nil, err
		}

		err = eksprovider.ValidateCredentials(ctx, *credential)
		if err != nil {
			err = fmt.Errorf("invalid credentials!: %w", err)
		}
		return nil, err
	}
}

func getEKSCredentialsFromReq(ctx context.Context, req EKSRegionReq, userInfoGetter provider.UserInfoGetter, presetProvider provider.PresetProvider, projectID string) (*resources.EKSCredential, error) {
	accessKeyID := req.AccessKeyID
	secretAccessKey := req.SecretAccessKey
	region := req.Region
	assumeRoleARN := req.AssumeRoleARN
	assumeRoleExternalID := req.AssumeRoleExternalID

	userInfo, err := userInfoGetter(ctx, projectID)
	if err != nil {
		return nil, common.KubernetesErrorToHTTPError(err)
	}

	presetName := req.Credential
	if len(presetName) > 0 {
		preset, err := presetProvider.GetPreset(ctx, userInfo, ptr.To(projectID), presetName)
		if err != nil {
			return nil, utilerrors.New(http.StatusInternalServerError, fmt.Sprintf("can not get preset %s for user %s", presetName, userInfo.Email))
		}
		if credentials := preset.Spec.EKS; credentials != nil {
			accessKeyID = credentials.AccessKeyID
			secretAccessKey = credentials.SecretAccessKey
		}
	}

	return &resources.EKSCredential{
		AccessKeyID:          accessKeyID,
		SecretAccessKey:      secretAccessKey,
		Region:               region,
		AssumeRoleARN:        assumeRoleARN,
		AssumeRoleExternalID: assumeRoleExternalID,
	}, nil
}

// EKSVPCReq represents a request for resources within a VPC.
// swagger:parameters listEKSSubnets listEKSSecurityGroups
type EKSVPCReq struct {
	EKSRegionReq
	// in: header
	// name: VpcId
	VpcId string
}

// EKSProjectVPCReq represents a request for resources within a VPC with project context.
// swagger:parameters listProjectEKSSubnets listProjectEKSSecurityGroups
type EKSProjectVPCReq struct {
	common.ProjectReq
	EKSVPCReq
}

func DecodeEKSVPCReq(c context.Context, r *http.Request) (interface{}, error) {
	var req EKSVPCReq

	typesReq, err := DecodeEKSRegionReq(c, r)
	if err != nil {
		return nil, err
	}
	req.EKSRegionReq = typesReq.(EKSRegionReq)
	req.VpcId = r.Header.Get("VpcId")

	return req, nil
}
func DecodeEKSProjectVPCReq(c context.Context, r *http.Request) (interface{}, error) {
	vpcReq, err := DecodeEKSVPCReq(c, r)
	if err != nil {
		return nil, err
	}

	projectReq, err := common.DecodeProjectRequest(c, r)
	if err != nil {
		return nil, err
	}

	return EKSProjectVPCReq{
		ProjectReq: projectReq.(common.ProjectReq),
		EKSVPCReq:  vpcReq.(EKSVPCReq),
	}, nil
}

func createNewEKSCluster(ctx context.Context, eksClusterSpec *apiv2.EKSClusterSpec, eksCloudSpec *apiv2.EKSCloudSpec) error {
	client, err := awsprovider.GetClientSet(ctx, eksCloudSpec.AccessKeyID, eksCloudSpec.SecretAccessKey, eksCloudSpec.AssumeRoleARN, eksCloudSpec.AssumeRoleExternalID, eksCloudSpec.Region)
	if err != nil {
		return err
	}

	clusterSpec := eksClusterSpec

	fields := reflect.ValueOf(clusterSpec).Elem()
	for i := 0; i < fields.NumField(); i++ {
		yourjsonTags := fields.Type().Field(i).Tag.Get("required")
		if strings.Contains(yourjsonTags, "true") && fields.Field(i).IsZero() {
			return fmt.Errorf("required field is missing %v", fields.Type().Field(i).Tag)
		}
	}

	return eksprovider.CreateCluster(ctx, client, clusterSpec, eksCloudSpec.Name)
}

func createOrImportEKSCluster(ctx context.Context, name string, userInfoGetter provider.UserInfoGetter, project *kubermaticv1.Project, spec *apiv2.ExternalClusterSpec, cloud *apiv2.ExternalClusterCloudSpec, clusterProvider provider.ExternalClusterProvider, privilegedClusterProvider provider.PrivilegedExternalClusterProvider) (*kubermaticv1.ExternalCluster, error) {
	isImported := resources.ExternalClusterIsImportedTrue

	fields := reflect.ValueOf(cloud.EKS).Elem()
	for i := 0; i < fields.NumField(); i++ {
		yourjsonTags := fields.Type().Field(i).Tag.Get("required")
		if strings.Contains(yourjsonTags, "true") && fields.Field(i).IsZero() {
			return nil, utilerrors.NewBadRequest("required field is missing: %v", fields.Type().Field(i).Name)
		}
	}

	if spec != nil && spec.EKSClusterSpec != nil {
		if err := createNewEKSCluster(ctx, spec.EKSClusterSpec, cloud.EKS); err != nil {
			return nil, err
		}
		isImported = resources.ExternalClusterIsImportedFalse
	}

	newCluster := genExternalCluster(name, project.Name, isImported)
	newCluster.Spec.CloudSpec = kubermaticv1.ExternalClusterCloudSpec{
		EKS: &kubermaticv1.ExternalClusterEKSCloudSpec{
			Name:   cloud.EKS.Name,
			Region: cloud.EKS.Region,
		},
	}

	keyRef, err := clusterProvider.CreateOrUpdateCredentialSecretForCluster(ctx, cloud, project.Name, newCluster.Name)
	if err != nil {
		return nil, common.KubernetesErrorToHTTPError(err)
	}
	newCluster.Spec.CloudSpec.EKS.CredentialsReference = keyRef

	return createNewCluster(ctx, userInfoGetter, clusterProvider, privilegedClusterProvider, newCluster, project)
}

func patchEKSCluster(ctx context.Context, oldCluster, newCluster *apiv2.ExternalCluster, secretKeySelector provider.SecretKeySelectorValueFunc, cloudSpec *kubermaticv1.ExternalClusterEKSCloudSpec) (*apiv2.ExternalCluster, error) {
	creds, err := eksprovider.GetCredentialsForCluster(cloudSpec, secretKeySelector)
	if err != nil {
		return nil, err
	}

	client, err := getClientSet(ctx, creds, cloudSpec.Region)
	if err != nil {
		return nil, err
	}

	newVersion := newCluster.Spec.Version.Semver()
	err = eksprovider.UpgradeClusterVersion(ctx, client, newVersion, cloudSpec.Name)
	if err != nil {
		return nil, err
	}

	return newCluster, nil
}

func getEKSNodeGroups(ctx context.Context, masterClient ctrlruntimeclient.Client, cluster *kubermaticv1.ExternalCluster, secretKeySelector provider.SecretKeySelectorValueFunc, clusterProvider provider.ExternalClusterProvider) ([]apiv2.ExternalClusterMachineDeployment, error) {
	cloudSpec := cluster.Spec.CloudSpec.EKS
	creds, err := eksprovider.GetCredentialsForCluster(cloudSpec, secretKeySelector)
	if err != nil {
		return nil, err
	}

	client, err := getClientSet(ctx, creds, cloudSpec.Region)
	if err != nil {
		return nil, err
	}

	clusterName := cloudSpec.Name

	nodeGroups, err := eksprovider.ListNodegroups(ctx, client, clusterName)
	if err != nil {
		return nil, err
	}

	machineDeployments := make([]apiv2.ExternalClusterMachineDeployment, 0, len(nodeGroups))

	nodes, err := clusterProvider.ListNodes(ctx, masterClient, cluster)
	if err != nil {
		return nil, common.KubernetesErrorToHTTPError(err)
	}

	for _, nodeGroupName := range nodeGroups {
		readyReplicasCount := kuberneteshelper.GetNodeGroupReadyCount(nodes, resources.EKSNodeGroupNameLabel, nodeGroupName)

		nodeGroup, err := eksprovider.DescribeNodeGroup(ctx, client, clusterName, nodeGroupName)
		if err != nil {
			return nil, err
		}
		machineDeployments = append(machineDeployments, createMachineDeploymentFromEKSNodePool(nodeGroup, readyReplicasCount))
	}

	return machineDeployments, err
}

func getEKSNodeGroup(ctx context.Context, masterClient ctrlruntimeclient.Client, cluster *kubermaticv1.ExternalCluster, nodeGroupName string, secretKeySelector provider.SecretKeySelectorValueFunc, clusterProvider provider.ExternalClusterProvider) (*apiv2.ExternalClusterMachineDeployment, error) {
	cloudSpec := cluster.Spec.CloudSpec.EKS

	creds, err := eksprovider.GetCredentialsForCluster(cloudSpec, secretKeySelector)
	if err != nil {
		return nil, err
	}

	client, err := getClientSet(ctx, creds, cloudSpec.Region)
	if err != nil {
		return nil, err
	}

	return getEKSMachineDeployment(ctx, client, masterClient, cluster, nodeGroupName, clusterProvider)
}

func getEKSMachineDeployment(ctx context.Context,
	awsClient *awsprovider.ClientSet,
	masterClient ctrlruntimeclient.Client,
	cluster *kubermaticv1.ExternalCluster,
	nodeGroupName string,
	clusterProvider provider.ExternalClusterProvider) (*apiv2.ExternalClusterMachineDeployment, error) {
	clusterName := cluster.Spec.CloudSpec.EKS.Name

	nodeGroup, err := eksprovider.DescribeNodeGroup(ctx, awsClient, clusterName, nodeGroupName)
	if err != nil {
		return nil, err
	}

	nodes, err := clusterProvider.ListNodes(ctx, masterClient, cluster)
	if err != nil {
		return nil, common.KubernetesErrorToHTTPError(err)
	}

	readyReplicasCount := kuberneteshelper.GetNodeGroupReadyCount(nodes, resources.EKSNodeGroupNameLabel, nodeGroupName)
	machineDeployment := createMachineDeploymentFromEKSNodePool(nodeGroup, readyReplicasCount)

	return &machineDeployment, err
}

func createMachineDeploymentFromEKSNodePool(nodeGroup *ekstypes.Nodegroup, readyReplicas int32) apiv2.ExternalClusterMachineDeployment {
	md := apiv2.ExternalClusterMachineDeployment{
		NodeDeployment: apiv1.NodeDeployment{
			ObjectMeta: apiv1.ObjectMeta{
				ID:   ptr.Deref(nodeGroup.NodegroupName, ""),
				Name: ptr.Deref(nodeGroup.NodegroupName, ""),
			},
			Spec: apiv1.NodeDeploymentSpec{
				Template: apiv1.NodeSpec{
					Versions: apiv1.NodeVersionInfo{
						Kubelet: ptr.Deref(nodeGroup.Version, ""),
					},
				},
			},
			Status: clusterv1alpha1.MachineDeploymentStatus{
				ReadyReplicas: readyReplicas,
			},
		},
		Cloud: &apiv2.ExternalClusterMachineDeploymentCloudSpec{},
	}

	md.Cloud.EKS = &apiv2.EKSMachineDeploymentCloudSpec{
		Subnets:       nodeGroup.Subnets,
		NodeRole:      ptr.Deref(nodeGroup.NodeRole, ""),
		AmiType:       string(nodeGroup.AmiType),
		CapacityType:  string(nodeGroup.CapacityType),
		DiskSize:      ptr.Deref(nodeGroup.DiskSize, 0),
		InstanceTypes: nodeGroup.InstanceTypes,
		Labels:        nodeGroup.Labels,
		Tags:          nodeGroup.Tags,
		Version:       ptr.Deref(nodeGroup.Version, ""),
	}

	if nodeGroup.CreatedAt != nil {
		md.Cloud.EKS.CreatedAt = *nodeGroup.CreatedAt
	}

	scalingConfig := nodeGroup.ScalingConfig
	if scalingConfig != nil {
		md.Status.Replicas = ptr.Deref(scalingConfig.DesiredSize, 0)
		md.Spec.Replicas = ptr.Deref(scalingConfig.DesiredSize, 0)
		md.Cloud.EKS.ScalingConfig = apiv2.EKSNodegroupScalingConfig{
			DesiredSize: ptr.Deref(scalingConfig.DesiredSize, 0),
			MaxSize:     ptr.Deref(scalingConfig.MaxSize, 0),
			MinSize:     ptr.Deref(scalingConfig.MinSize, 0),
		}
	}

	md.Phase = apiv2.ExternalClusterMDPhase{
		State: eksprovider.ConvertMDStatus(nodeGroup.Status),
	}

	return md
}

func patchEKSMachineDeployment(ctx context.Context, oldMD, newMD *apiv2.ExternalClusterMachineDeployment, secretKeySelector provider.SecretKeySelectorValueFunc, cluster *kubermaticv1.ExternalCluster) (*apiv2.ExternalClusterMachineDeployment, error) {
	cloudSpec := cluster.Spec.CloudSpec.EKS

	creds, err := eksprovider.GetCredentialsForCluster(cloudSpec, secretKeySelector)
	if err != nil {
		return nil, err
	}

	client, err := getClientSet(ctx, creds, cloudSpec.Region)
	if err != nil {
		return nil, err
	}

	// The EKS can update Node Group size or NodeGroup object. Can't change both till one of the update is in progress.
	// It's required to update NodeGroup size separately.

	clusterName := cloudSpec.Name
	nodeGroupName := newMD.Name

	currentReplicas := oldMD.Spec.Replicas
	desiredReplicas := newMD.Spec.Replicas
	currentVersion := oldMD.Spec.Template.Versions.Kubelet
	desiredVersion := newMD.Spec.Template.Versions.Kubelet
	if desiredReplicas != currentReplicas {
		_, err = eksprovider.ResizeNodeGroup(ctx, client, clusterName, nodeGroupName, currentReplicas, desiredReplicas)
		if err != nil {
			return nil, err
		}
		newMD.Status.Replicas = desiredReplicas
		newMD.Spec.Template.Versions.Kubelet = currentVersion
		return newMD, nil
	}

	if desiredVersion != currentVersion {
		_, err = eksprovider.UpgradeNodeGroup(ctx, client, &clusterName, &nodeGroupName, &currentVersion, &desiredVersion)
		if err != nil {
			return nil, err
		}
		newMD.Spec.Replicas = currentReplicas
		return newMD, nil
	}

	return newMD, nil
}

func deleteEKSNodeGroup(ctx context.Context, cluster *kubermaticv1.ExternalCluster, nodeGroupName string, secretKeySelector provider.SecretKeySelectorValueFunc, credentialsReference *providerconfig.GlobalSecretKeySelector, clusterProvider provider.ExternalClusterProvider) error {
	creds, err := eksprovider.GetCredentialsForCluster(cluster.Spec.CloudSpec.EKS, secretKeySelector)
	if err != nil {
		return err
	}

	clusterCloudSpec := cluster.Spec.CloudSpec
	eksClusterCloudSpec := clusterCloudSpec.EKS
	client, err := getClientSet(ctx, creds, eksClusterCloudSpec.Region)
	if err != nil {
		return err
	}

	return eksprovider.DeleteNodegroup(ctx, client, eksClusterCloudSpec.Name, nodeGroupName)
}

func EKSInstanceTypesWithClusterCredentialsEndpoint(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider, privilegedProjectProvider provider.PrivilegedProjectProvider, clusterProvider provider.ExternalClusterProvider, privilegedClusterProvider provider.PrivilegedExternalClusterProvider, settingsProvider provider.SettingsProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(eksNoCredentialSizeReq)
		if !ok {
			return nil, utilerrors.NewBadRequest("invalid request")
		}
		project, err := common.GetProject(ctx, userInfoGetter, projectProvider, privilegedProjectProvider, req.ProjectID, nil)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		cluster, err := getCluster(ctx, userInfoGetter, clusterProvider, privilegedClusterProvider, project.Name, req.ClusterID)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		secretKeySelector := provider.SecretKeySelectorValueFuncFactory(ctx, privilegedClusterProvider.GetMasterClient())

		cloudSpec := cluster.Spec.CloudSpec.EKS
		if cloudSpec == nil {
			return nil, utilerrors.NewNotFound("cloud spec for %s", cluster.Name)
		}

		creds, err := eksprovider.GetCredentialsForCluster(cloudSpec, secretKeySelector)
		if err != nil {
			return nil, err
		}

		if cloudSpec.Region == "" {
			return nil, errors.New("no region provided in externalcluster spec")
		}
		credential := resources.EKSCredential{
			AccessKeyID:          creds.AccessKeyID,
			SecretAccessKey:      creds.SecretAccessKey,
			AssumeRoleARN:        creds.AssumeRoleARN,
			AssumeRoleExternalID: creds.AssumeRoleExternalID,
			Region:               cloudSpec.Region,
		}

		return providercommon.ListInstanceTypes(ctx, credential, req.Architecture)
	}
}

func EKSVPCsWithClusterCredentialsEndpoint(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider, privilegedProjectProvider provider.PrivilegedProjectProvider, clusterProvider provider.ExternalClusterProvider, privilegedClusterProvider provider.PrivilegedExternalClusterProvider, settingsProvider provider.SettingsProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(eksNoCredentialReq)
		if !ok {
			return nil, utilerrors.NewBadRequest("invalid request")
		}
		project, err := common.GetProject(ctx, userInfoGetter, projectProvider, privilegedProjectProvider, req.ProjectID, nil)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		cluster, err := getCluster(ctx, userInfoGetter, clusterProvider, privilegedClusterProvider, project.Name, req.ClusterID)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		secretKeySelector := provider.SecretKeySelectorValueFuncFactory(ctx, privilegedClusterProvider.GetMasterClient())

		cloudSpec := cluster.Spec.CloudSpec.EKS
		if cloudSpec == nil {
			return nil, utilerrors.NewNotFound("cloud spec for %s", cluster.Name)
		}

		creds, err := eksprovider.GetCredentialsForCluster(cloudSpec, secretKeySelector)
		if err != nil {
			return nil, err
		}

		if cloudSpec.Region == "" {
			return nil, errors.New("no region provided in externalcluster spec")
		}
		credential := resources.EKSCredential{
			AccessKeyID:          creds.AccessKeyID,
			SecretAccessKey:      creds.SecretAccessKey,
			AssumeRoleARN:        creds.AssumeRoleARN,
			AssumeRoleExternalID: creds.AssumeRoleExternalID,
			Region:               cloudSpec.Region,
		}
		return providercommon.ListEKSVPC(ctx, credential)
	}
}

func EKSSubnetsWithClusterCredentialsEndpoint(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider, privilegedProjectProvider provider.PrivilegedProjectProvider, clusterProvider provider.ExternalClusterProvider, privilegedClusterProvider provider.PrivilegedExternalClusterProvider, settingsProvider provider.SettingsProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(eksSubnetsNoCredentialReq)
		if !ok {
			return nil, utilerrors.NewBadRequest("invalid request")
		}
		if err := req.Validate(); err != nil {
			return nil, utilerrors.NewBadRequest("%v", err)
		}

		project, err := common.GetProject(ctx, userInfoGetter, projectProvider, privilegedProjectProvider, req.ProjectID, nil)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		cluster, err := getCluster(ctx, userInfoGetter, clusterProvider, privilegedClusterProvider, project.Name, req.ClusterID)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		secretKeySelector := provider.SecretKeySelectorValueFuncFactory(ctx, privilegedClusterProvider.GetMasterClient())

		cloudSpec := cluster.Spec.CloudSpec.EKS
		if cloudSpec == nil {
			return nil, utilerrors.NewNotFound("cloud spec for %s", cluster.Name)
		}

		creds, err := eksprovider.GetCredentialsForCluster(cloudSpec, secretKeySelector)
		if err != nil {
			return nil, err
		}

		if cloudSpec.Region == "" {
			return nil, errors.New("no region provided in externalcluster spec")
		}
		credential := resources.EKSCredential{
			AccessKeyID:          creds.AccessKeyID,
			SecretAccessKey:      creds.SecretAccessKey,
			AssumeRoleARN:        creds.AssumeRoleARN,
			AssumeRoleExternalID: creds.AssumeRoleExternalID,
			Region:               cloudSpec.Region,
		}
		return providercommon.ListEKSSubnetIDs(ctx, credential, req.VpcId)
	}
}

func EKSNodeRolesWithClusterCredentialsEndpoint(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider, privilegedProjectProvider provider.PrivilegedProjectProvider, clusterProvider provider.ExternalClusterProvider, privilegedClusterProvider provider.PrivilegedExternalClusterProvider, settingsProvider provider.SettingsProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(eksNoCredentialReq)
		if !ok {
			return nil, utilerrors.NewBadRequest("invalid request")
		}
		project, err := common.GetProject(ctx, userInfoGetter, projectProvider, privilegedProjectProvider, req.ProjectID, nil)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		cluster, err := getCluster(ctx, userInfoGetter, clusterProvider, privilegedClusterProvider, project.Name, req.ClusterID)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		secretKeySelector := provider.SecretKeySelectorValueFuncFactory(ctx, privilegedClusterProvider.GetMasterClient())

		cloudSpec := cluster.Spec.CloudSpec.EKS
		if cloudSpec == nil {
			return nil, utilerrors.NewNotFound("cloud spec for %s", cluster.Name)
		}

		creds, err := eksprovider.GetCredentialsForCluster(cloudSpec, secretKeySelector)
		if err != nil {
			return nil, err
		}

		credential := resources.EKSCredential{
			AccessKeyID:          creds.AccessKeyID,
			SecretAccessKey:      creds.SecretAccessKey,
			AssumeRoleARN:        creds.AssumeRoleARN,
			AssumeRoleExternalID: creds.AssumeRoleExternalID,
			Region:               cloudSpec.Region,
		}
		return providercommon.ListEKSNodeRoles(ctx, credential)
	}
}

func getEKSClusterDetails(ctx context.Context, apiCluster *apiv2.ExternalCluster, secretKeySelector provider.SecretKeySelectorValueFunc, cloudSpec *kubermaticv1.ExternalClusterEKSCloudSpec) (*apiv2.ExternalCluster, error) {
	creds, err := eksprovider.GetCredentialsForCluster(cloudSpec, secretKeySelector)
	if err != nil {
		return nil, err
	}

	client, err := getClientSet(ctx, creds, cloudSpec.Region)
	if err != nil {
		return nil, err
	}
	cluster, err := eksprovider.GetCluster(ctx, client, cloudSpec.Name)
	if err != nil {
		return nil, err
	}
	if cluster == nil {
		return apiCluster, nil
	}

	clusterSpec := &apiv2.EKSClusterSpec{
		Version:   ptr.Deref(cluster.Version, ""),
		CreatedAt: cluster.CreatedAt,
	}

	if cluster.KubernetesNetworkConfig != nil {
		clusterSpec.KubernetesNetworkConfig = &apiv2.EKSKubernetesNetworkConfigResponse{
			IpFamily:        string(cluster.KubernetesNetworkConfig.IpFamily),
			ServiceIpv4Cidr: cluster.KubernetesNetworkConfig.ServiceIpv4Cidr,
			ServiceIpv6Cidr: cluster.KubernetesNetworkConfig.ServiceIpv6Cidr,
		}
	}

	if cluster.ResourcesVpcConfig != nil {
		clusterSpec.ResourcesVpcConfig = apiv2.VpcConfigRequest{
			SecurityGroupIds: cluster.ResourcesVpcConfig.SecurityGroupIds,
			SubnetIds:        cluster.ResourcesVpcConfig.SubnetIds,
			VpcId:            cluster.ResourcesVpcConfig.VpcId,
		}
	}

	apiCluster.Spec.EKSClusterSpec = clusterSpec
	return apiCluster, nil
}

func checkCreatePoolReqValid(machineDeployment apiv2.ExternalClusterMachineDeployment) error {
	eksMD := machineDeployment.Cloud.EKS
	if eksMD == nil {
		return fmt.Errorf("EKS MachineDeployment Spec cannot be nil")
	}

	if len(eksMD.Subnets) == 0 || len(eksMD.NodeRole) == 0 {
		return utilerrors.NewBadRequest("required field is missing: Subnets or NodeRole cannot be empty")
	}
	return nil
}

func createEKSNodePool(ctx context.Context, cloudSpec *kubermaticv1.ExternalClusterEKSCloudSpec, machineDeployment apiv2.ExternalClusterMachineDeployment, secretKeySelector provider.SecretKeySelectorValueFunc, credentialsReference *providerconfig.GlobalSecretKeySelector) (*apiv2.ExternalClusterMachineDeployment, error) {
	if err := checkCreatePoolReqValid(machineDeployment); err != nil {
		return nil, err
	}
	eksMDCloudSpec := machineDeployment.Cloud.EKS

	creds, err := eksprovider.GetCredentialsForCluster(cloudSpec, secretKeySelector)
	if err != nil {
		return nil, err
	}
	client, err := getClientSet(ctx, creds, cloudSpec.Region)
	if err != nil {
		return nil, err
	}

	err = eksprovider.CreateNodeGroup(ctx, client, cloudSpec.Name, machineDeployment.Name, eksMDCloudSpec)
	if err != nil {
		return nil, err
	}
	machineDeployment.Phase = apiv2.ExternalClusterMDPhase{
		State: apiv2.ProvisioningExternalClusterMDState,
	}

	return &machineDeployment, nil
}

func deleteEKSCluster(ctx context.Context, secretKeySelector provider.SecretKeySelectorValueFunc, cloudSpec *kubermaticv1.ExternalClusterEKSCloudSpec) error {
	creds, err := eksprovider.GetCredentialsForCluster(cloudSpec, secretKeySelector)
	if err != nil {
		return err
	}

	client, err := getClientSet(ctx, creds, cloudSpec.Region)
	if err != nil {
		return err
	}

	return eksprovider.DeleteCluster(ctx, client, cloudSpec.Name)
}

func EKSAMITypesEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var ami ekstypes.AMITypes = EKSAMITypes
		var amiTypes apiv2.EKSAMITypeList

		for _, amiType := range ami.Values() {
			// AMI type Custom is not valid
			if amiType == EKSCustomAMIType {
				continue
			}
			amiTypes = append(amiTypes, string(amiType))
		}
		return amiTypes, nil
	}
}

func EKSCapacityTypesEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var capacityType ekstypes.CapacityTypes = EKSCapacityTypes
		var capacityTypes apiv2.EKSCapacityTypeList

		for _, c := range capacityType.Values() {
			capacityTypes = append(capacityTypes, string(c))
		}
		return capacityTypes, nil
	}
}

func EKSVersionsEndpoint(configGetter provider.KubermaticConfigurationGetter,
	clusterProvider provider.ExternalClusterProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return clusterProvider.VersionsEndpoint(ctx, configGetter, kubermaticv1.EKSProviderType)
	}
}

func getClientSet(ctx context.Context, creds eksprovider.EKSCredentials, region string) (*awsprovider.ClientSet, error) {
	return awsprovider.GetClientSet(ctx, creds.AccessKeyID, creds.SecretAccessKey, creds.AssumeRoleARN, creds.AssumeRoleExternalID, region)
}
