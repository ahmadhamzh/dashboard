/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package cluster_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	apiv1 "k8c.io/dashboard/v2/pkg/api/v1"
	apiv2 "k8c.io/dashboard/v2/pkg/api/v2"
	"k8c.io/dashboard/v2/pkg/handler/test"
	"k8c.io/dashboard/v2/pkg/handler/test/hack"
	kubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/kubermatic/v1"
	"k8c.io/kubermatic/sdk/v2/semver"
	"k8c.io/kubermatic/v2/pkg/cni"
	"k8c.io/kubermatic/v2/pkg/defaulting"
	kuberneteshelper "k8c.io/kubermatic/v2/pkg/kubernetes"
	"k8c.io/kubermatic/v2/pkg/resources"
	clusterv1alpha1 "k8c.io/machine-controller/sdk/apis/cluster/v1alpha1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/metrics/pkg/apis/metrics/v1beta1"
	"k8s.io/utils/ptr"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	usCentral1 = "us-central1"
)

var (
	healthUp = kubermaticv1.HealthStatusUp
)

func TestCreateClusterEndpoint(t *testing.T) {
	version := defaulting.DefaultKubernetesVersioning.Default.String()

	t.Parallel()
	testcases := []struct {
		Name                   string
		Body                   string
		ExpectedResponse       string
		HTTPStatus             int
		ProjectToSync          string
		ExistingProject        *kubermaticv1.Project
		ExistingAPIUser        *apiv1.User
		ExistingKubermaticObjs []ctrlruntimeclient.Object
		RewriteClusterID       bool
	}{
		// scenario 1
		{
			Name:             "scenario 1: a cluster with invalid spec is rejected",
			Body:             `{"cluster":{"name":"keen-snyder","spec":{"cloud":{"fake":{"token":"dummy_token"},"dc":"fake-dc"}, "version":""}}}`,
			ExpectedResponse: `{"error":{"code":400,"message":"invalid cluster: invalid cloud spec \"Version\" is required but was not specified"}}`,
			HTTPStatus:       http.StatusBadRequest,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
			),
			ProjectToSync:   test.GenDefaultProject().Name,
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 2
		{
			Name:             "scenario 2: cluster is created when valid spec and ssh key are passed",
			Body:             fmt.Sprintf(`{"cluster":{"name":"keen-snyder","spec":{"version":"%s","cloud":{"fake":{"token":"dummy_token"},"dc":"fake-dc"}}}}`, version),
			ExpectedResponse: fmt.Sprintf(`{"id":"%%s","name":"keen-snyder","annotations":{"kubermatic.io/initial-application-installations-request":"[]"},"creationTimestamp":"0001-01-01T00:00:00Z","type":"kubernetes","spec":{"cloud":{"dc":"fake-dc","fake":{}},"version":"%s","oidc":{},"enableUserSSHKeyAgent":true,"kubernetesDashboard":{"enabled":true},"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["10.240.16.0/20"]},"pods":{"cidrBlocks":["172.25.0.0/16"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true},"nodeLocalDNSCacheEnabled":true,"konnectivityEnabled":true},"cniPlugin":{"type":"cilium","version":"1.16.9"},"exposeStrategy":"NodePort"},"status":{"version":"","url":"","externalCCMMigration":"Unsupported"}}`, version),
			RewriteClusterID: true,
			HTTPStatus:       http.StatusCreated,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add an ssh key
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-c08aa5c7abf34504f18552846485267d-yafn",
						OwnerReferences: []metav1.OwnerReference{
							{
								APIVersion: "kubermatic.k8c.io/v1",
								Kind:       "Project",
								UID:        "",
								Name:       test.GenDefaultProject().Name,
							},
						},
					},
				},
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 3
		{
			Name:             "scenario 3: unable to create a cluster when the user doesn't belong to the project",
			Body:             fmt.Sprintf(`{"cluster":{"name":"keen-snyder","pause":false,"spec":{"version":"%s","cloud":{"version":"%s","fake":{"token":"dummy_token"},"dc":"fake-dc","enableUserSSHKeyAgent":true,"containerRuntime":"containerd"}}},"sshKeys":["key-c08aa5c7abf34504f18552846485267d-yafn"]}`, version, version),
			ExpectedResponse: `{"error":{"code":403,"message":"forbidden: \"john@acme.com\" doesn't belong to project my-first-project-ID"}}`,
			HTTPStatus:       http.StatusForbidden,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenUser("", "John", "john@acme.com"),
			),
			ExistingAPIUser: func() *apiv1.User {
				defaultUser := test.GenDefaultAPIUser()
				defaultUser.Email = "john@acme.com"
				return defaultUser
			}(),
		},
		// scenario 4
		{
			Name:             "scenario 4: unable to create a cluster when project is not ready",
			Body:             fmt.Sprintf(`{"cluster":{"name":"keen-snyder","pause":false,"spec":{"version":"%s","cloud":{"fake":{"token":"dummy_token"},"dc":"fake-dc"}}},"sshKeys":["key-c08aa5c7abf34504f18552846485267d-yafn"]}`, version),
			ExpectedResponse: `{"error":{"code":503,"message":"Project is not initialized yet"}}`,
			HTTPStatus:       http.StatusServiceUnavailable,
			ExistingProject: func() *kubermaticv1.Project {
				project := test.GenDefaultProject()
				project.Status.Phase = kubermaticv1.ProjectInactive
				return project
			}(),
			ProjectToSync: test.GenDefaultProject().Name,
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				test.GenTestSeed(),
				test.GenDefaultUser(),
				test.GenDefaultOwnerBinding(),
			},
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		{
			Name:             "scenario 9a: rejected an attempt to create a cluster in email-restricted datacenter - legacy single domain restriction with requiredEmailDomains",
			Body:             fmt.Sprintf(`{"cluster":{"name":"keen-snyder","spec":{"version":"%s","cloud":{"fake":{"token":"dummy_token"},"dc":"restricted-fake-dc"}}}}`, version),
			ExpectedResponse: `{"error":{"code":403,"message":"cannot access restricted-fake-dc datacenter due to email requirements"}}`,
			RewriteClusterID: false,
			HTTPStatus:       http.StatusForbidden,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		{
			Name:             "scenario 9b: rejected an attempt to create a cluster in email-restricted datacenter - domain array restriction with `requiredEmailDomains`",
			Body:             fmt.Sprintf(`{"cluster":{"name":"keen-snyder","spec":{"version":"%s","cloud":{"fake":{"token":"dummy_token"},"dc":"restricted-fake-dc2"}}}}`, version),
			ExpectedResponse: `{"error":{"code":403,"message":"cannot access restricted-fake-dc2 datacenter due to email requirements"}}`,
			RewriteClusterID: false,
			HTTPStatus:       http.StatusForbidden,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		{
			Name:             "scenario 10a: create a cluster in email-restricted datacenter, to which the user does have access - legacy single domain restriction with requiredEmailDomains",
			Body:             fmt.Sprintf(`{"cluster":{"name":"keen-snyder","spec":{"version":"%s","cloud":{"fake":{"token":"dummy_token"},"dc":"restricted-fake-dc"}}}}`, version),
			ExpectedResponse: fmt.Sprintf(`{"id":"%%s","name":"keen-snyder","annotations":{"kubermatic.io/initial-application-installations-request":"[]"},"creationTimestamp":"0001-01-01T00:00:00Z","type":"kubernetes","spec":{"cloud":{"dc":"restricted-fake-dc","fake":{}},"version":"%s","oidc":{},"enableUserSSHKeyAgent":true,"kubernetesDashboard":{"enabled":true},"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["10.240.16.0/20"]},"pods":{"cidrBlocks":["172.25.0.0/16"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true},"nodeLocalDNSCacheEnabled":true,"konnectivityEnabled":true},"cniPlugin":{"type":"cilium","version":"1.16.9"},"exposeStrategy":"NodePort"},"status":{"version":"","url":"","externalCCMMigration":"Unsupported"}}`, version),
			RewriteClusterID: true,
			HTTPStatus:       http.StatusCreated,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenUser(test.UserID2, test.UserName2, test.UserEmail2),
				test.GenBinding(test.GenDefaultProject().Name, test.UserEmail2, "editors"),
			),
			ExistingAPIUser: test.GenAPIUser(test.UserName2, test.UserEmail2),
		},
		{
			Name:             "scenario 10b: create a cluster in email-restricted datacenter, to which the user does have access - domain array restriction with `requiredEmailDomains`",
			Body:             fmt.Sprintf(`{"cluster":{"name":"keen-snyder","spec":{"version":"%s","cloud":{"fake":{"token":"dummy_token"},"dc":"restricted-fake-dc2"}}}}`, version),
			ExpectedResponse: fmt.Sprintf(`{"id":"%%s","name":"keen-snyder","annotations":{"kubermatic.io/initial-application-installations-request":"[]"},"creationTimestamp":"0001-01-01T00:00:00Z","type":"kubernetes","spec":{"cloud":{"dc":"restricted-fake-dc2","fake":{}},"version":"%s","oidc":{},"enableUserSSHKeyAgent":true,"kubernetesDashboard":{"enabled":true},"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["10.240.16.0/20"]},"pods":{"cidrBlocks":["172.25.0.0/16"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true},"nodeLocalDNSCacheEnabled":true,"konnectivityEnabled":true},"cniPlugin":{"type":"cilium","version":"1.16.9"},"exposeStrategy":"NodePort"},"status":{"version":"","url":"","externalCCMMigration":"Unsupported"}}`, version),
			RewriteClusterID: true,
			HTTPStatus:       http.StatusCreated,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenUser(test.UserID2, test.UserName2, test.UserEmail2),
				test.GenBinding(test.GenDefaultProject().Name, test.UserEmail2, "editors"),
			),
			ExistingAPIUser: test.GenAPIUser(test.UserName2, test.UserEmail2),
		},
		{
			Name:             "scenario 11: create a cluster in audit-logging-enforced datacenter, without explicitly enabling audit logging",
			Body:             fmt.Sprintf(`{"cluster":{"name":"keen-snyder","spec":{"version":"%s","cloud":{"fake":{"token":"dummy_token"},"dc":"audited-dc"}}}}`, version),
			ExpectedResponse: fmt.Sprintf(`{"id":"%%s","name":"keen-snyder","annotations":{"kubermatic.io/initial-application-installations-request":"[]"},"creationTimestamp":"0001-01-01T00:00:00Z","type":"kubernetes","spec":{"cloud":{"dc":"audited-dc","fake":{}},"version":"%s","oidc":{},"enableUserSSHKeyAgent":true,"kubernetesDashboard":{"enabled":true},"auditLogging":{"enabled":true},"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["10.240.16.0/20"]},"pods":{"cidrBlocks":["172.25.0.0/16"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true},"nodeLocalDNSCacheEnabled":true,"konnectivityEnabled":true},"cniPlugin":{"type":"cilium","version":"1.16.9"},"exposeStrategy":"NodePort"},"status":{"version":"","url":"","externalCCMMigration":"Unsupported"}}`, version),
			RewriteClusterID: true,
			HTTPStatus:       http.StatusCreated,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenUser(test.UserID2, test.UserName2, test.UserEmail2),
				test.GenBinding(test.GenDefaultProject().Name, test.UserEmail2, "editors"),
			),
			ExistingAPIUser: test.GenAPIUser(test.UserName2, test.UserEmail2),
		},
		{
			Name:             "scenario 12: the admin user can create cluster for any project",
			Body:             fmt.Sprintf(`{"cluster":{"name":"keen-snyder","spec":{"version":"%s","cloud":{"fake":{"token":"dummy_token"},"dc":"fake-dc"}}}}`, version),
			ExpectedResponse: fmt.Sprintf(`{"id":"%%s","name":"keen-snyder","annotations":{"kubermatic.io/initial-application-installations-request":"[]"},"creationTimestamp":"0001-01-01T00:00:00Z","type":"kubernetes","spec":{"cloud":{"dc":"fake-dc","fake":{}},"version":"%s","oidc":{},"enableUserSSHKeyAgent":true,"kubernetesDashboard":{"enabled":true},"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["10.240.16.0/20"]},"pods":{"cidrBlocks":["172.25.0.0/16"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true},"nodeLocalDNSCacheEnabled":true,"konnectivityEnabled":true},"cniPlugin":{"type":"cilium","version":"1.16.9"},"exposeStrategy":"NodePort"},"status":{"version":"","url":"","externalCCMMigration":"Unsupported"}}`, version),
			RewriteClusterID: true,
			HTTPStatus:       http.StatusCreated,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add admin user
				genUser("John", "john@acme.com", true),
				// add an ssh key
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-c08aa5c7abf34504f18552846485267d-yafn",
						OwnerReferences: []metav1.OwnerReference{
							{
								APIVersion: "kubermatic.k8c.io/v1",
								Kind:       "Project",
								UID:        "",
								Name:       test.GenDefaultProject().Name,
							},
						},
					},
				},
			),
			ExistingAPIUser: test.GenAPIUser("John", "john@acme.com"),
		},
		// scenario 13
		{
			Name:             "scenario 13: a cluster with invalid version",
			Body:             `{"cluster":{"name":"keen-snyder","spec":{"cloud":{"fake":{"token":"dummy_token"},"dc":"fake-dc"}, "version":"1.2.3"}}}`,
			ExpectedResponse: `{"error":{"code":400,"message":"invalid cluster: invalid cloud spec: unsupported version 1.2.3"}}`,
			HTTPStatus:       http.StatusBadRequest,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
			),
			ProjectToSync:   test.GenDefaultProject().Name,
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 14
		{
			Name:             "scenario 14: a cluster without version",
			Body:             `{"cluster":{"name":"keen-snyder","spec":{"cloud":{"fake":{"token":"dummy_token"},"dc":"fake-dc"}}}}`,
			ExpectedResponse: `{"error":{"code":400,"message":"invalid cluster: invalid cloud spec \"Version\" is required but was not specified"}}`,
			HTTPStatus:       http.StatusBadRequest,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
			),
			ProjectToSync:   test.GenDefaultProject().Name,
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 15
		{
			Name:             "scenario 15: cluster is created with preset annotation",
			Body:             fmt.Sprintf(`{"cluster":{"name":"keen-snyder","credential":"fake","spec":{"version":"%s","cloud":{"fake":{},"dc":"fake-dc"}}}}`, version),
			ExpectedResponse: fmt.Sprintf(`{"id":"%%s","name":"keen-snyder","annotations":{"kubermatic.io/initial-application-installations-request":"[]","presetName":"fake"},"creationTimestamp":"0001-01-01T00:00:00Z","type":"kubernetes","spec":{"cloud":{"dc":"fake-dc","fake":{}},"version":"%s","oidc":{},"enableUserSSHKeyAgent":true,"kubernetesDashboard":{"enabled":true},"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["10.240.16.0/20"]},"pods":{"cidrBlocks":["172.25.0.0/16"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true},"nodeLocalDNSCacheEnabled":true,"konnectivityEnabled":true},"cniPlugin":{"type":"cilium","version":"1.16.9"},"exposeStrategy":"NodePort"},"status":{"version":"","url":"","externalCCMMigration":"Unsupported"}}`, version),
			RewriteClusterID: true,
			HTTPStatus:       http.StatusCreated,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add an ssh key
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-c08aa5c7abf34504f18552846485267d-yafn",
						OwnerReferences: []metav1.OwnerReference{
							{
								APIVersion: "kubermatic.k8c.io/v1",
								Kind:       "Project",
								UID:        "",
								Name:       test.GenDefaultProject().Name,
							},
						},
					},
				},
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
	}

	dummyKubermaticConfiguration := &kubermaticv1.KubermaticConfiguration{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "kubermatic",
			Namespace: resources.KubermaticNamespace,
		},
		Spec: kubermaticv1.KubermaticConfigurationSpec{
			Versions: kubermaticv1.KubermaticVersioningConfiguration{
				Versions: defaulting.DefaultKubernetesVersioning.Versions,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v2/projects/%s/clusters", tc.ProjectToSync), strings.NewReader(tc.Body))
			res := httptest.NewRecorder()
			var kubermaticObj []ctrlruntimeclient.Object
			if tc.ExistingProject != nil {
				kubermaticObj = append(kubermaticObj, tc.ExistingProject)
			}
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)

			ep, err := test.CreateTestEndpoint(*tc.ExistingAPIUser, []ctrlruntimeclient.Object{}, kubermaticObj, dummyKubermaticConfiguration, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			ep.ServeHTTP(res, req)

			if res.Code != tc.HTTPStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.HTTPStatus, res.Code, res.Body.String())
			}

			expectedResponse := tc.ExpectedResponse
			// since Cluster.Name is automatically generated by the system just rewrite it.
			if tc.RewriteClusterID {
				actualCluster := &apiv1.Cluster{}
				err = json.Unmarshal(res.Body.Bytes(), actualCluster)
				if err != nil {
					t.Fatal(err)
				}
				expectedResponse = fmt.Sprintf(tc.ExpectedResponse, actualCluster.ID)
			}

			test.CompareWithResult(t, res, expectedResponse)
		})
	}
}

func TestListClusters(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		Name                   string
		ExpectedClusters       apiv2.ProjectClusterList
		HTTPStatus             int
		ExistingAPIUser        *apiv1.User
		ExistingKubermaticObjs []ctrlruntimeclient.Object
	}{
		// scenario 1
		{
			Name: "scenario 1: list clusters that belong to the given project",
			ExpectedClusters: apiv2.ProjectClusterList{
				Clusters: apiv1.ClusterList{
					{
						ObjectMeta: apiv1.ObjectMeta{
							ID:                "clusterAbcID",
							Name:              "clusterAbc",
							CreationTimestamp: apiv1.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC),
						},
						Spec: apiv1.ClusterSpec{
							Cloud: kubermaticv1.CloudSpec{
								DatacenterName: "private-do1",
								Fake:           &kubermaticv1.FakeCloudSpec{},
							},
							ContainerRuntime: "containerd",
							ClusterNetwork: &kubermaticv1.ClusterNetworkingConfig{
								IPFamily:             kubermaticv1.IPFamilyIPv4,
								Pods:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"1.2.3.4/8"}},
								Services:             kubermaticv1.NetworkRanges{CIDRBlocks: []string{"5.6.7.8/8"}},
								NodeCIDRMaskSizeIPv4: ptr.To[int32](24),
								ProxyMode:            resources.IPVSProxyMode,
								IPVS: &kubermaticv1.IPVSConfiguration{
									StrictArp: ptr.To(true),
								},
								DNSDomain: "cluster.local",
							},
							CNIPlugin: &kubermaticv1.CNIPluginSettings{
								Type:    kubermaticv1.CNIPluginTypeCanal,
								Version: cni.GetDefaultCNIPluginVersion(kubermaticv1.CNIPluginTypeCanal),
							},
							ExposeStrategy:        kubermaticv1.ExposeStrategyNodePort,
							Version:               *semver.NewSemverOrDie("9.9.9"),
							EnableUserSSHKeyAgent: ptr.To(false),
						},
						Status: apiv1.ClusterStatus{
							Version:              *semver.NewSemverOrDie("9.9.9"),
							URL:                  "https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885",
							ExternalCCMMigration: apiv1.ExternalCCMMigrationUnsupported,
						},
						Type: "kubernetes",
					},
					{
						ObjectMeta: apiv1.ObjectMeta{
							ID:                "clusterDefID",
							Name:              "clusterDef",
							CreationTimestamp: apiv1.Date(2013, 02, 04, 01, 54, 0, 0, time.UTC),
						},
						Spec: apiv1.ClusterSpec{
							Cloud: kubermaticv1.CloudSpec{
								DatacenterName: "private-do1",
								Fake:           &kubermaticv1.FakeCloudSpec{},
							},
							ContainerRuntime: "containerd",
							ClusterNetwork: &kubermaticv1.ClusterNetworkingConfig{
								IPFamily:             kubermaticv1.IPFamilyIPv4,
								Pods:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"1.2.3.4/8"}},
								Services:             kubermaticv1.NetworkRanges{CIDRBlocks: []string{"5.6.7.8/8"}},
								NodeCIDRMaskSizeIPv4: ptr.To[int32](24),
								ProxyMode:            resources.IPVSProxyMode,
								IPVS: &kubermaticv1.IPVSConfiguration{
									StrictArp: ptr.To(true),
								},
								DNSDomain: "cluster.local",
							},
							CNIPlugin: &kubermaticv1.CNIPluginSettings{
								Type:    kubermaticv1.CNIPluginTypeCanal,
								Version: cni.GetDefaultCNIPluginVersion(kubermaticv1.CNIPluginTypeCanal),
							},
							ExposeStrategy:        kubermaticv1.ExposeStrategyNodePort,
							Version:               *semver.NewSemverOrDie("9.9.9"),
							EnableUserSSHKeyAgent: ptr.To(false),
						},
						Status: apiv1.ClusterStatus{
							Version:              *semver.NewSemverOrDie("9.9.9"),
							URL:                  "https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885",
							ExternalCCMMigration: apiv1.ExternalCCMMigrationUnsupported,
						},
						Type: "kubernetes",
					},
					{
						ObjectMeta: apiv1.ObjectMeta{
							ID:                "clusterOpenstackID",
							Name:              "clusterOpenstack",
							CreationTimestamp: apiv1.Date(2013, 02, 04, 03, 54, 0, 0, time.UTC),
						},
						Spec: apiv1.ClusterSpec{
							Cloud: kubermaticv1.CloudSpec{
								DatacenterName: "OpenstackDatacenter",
								Openstack: &kubermaticv1.OpenstackCloudSpec{
									FloatingIPPool: "floatingIPPool",
									SubnetID:       "subnetID",
									Domain:         "domain",
									Network:        "network",
									RouterID:       "routerID",
									SecurityGroups: "securityGroups",
									Project:        "project",
								},
							},
							ContainerRuntime: "containerd",
							ClusterNetwork: &kubermaticv1.ClusterNetworkingConfig{
								IPFamily:             kubermaticv1.IPFamilyIPv4,
								Pods:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"1.2.3.4/8"}},
								Services:             kubermaticv1.NetworkRanges{CIDRBlocks: []string{"5.6.7.8/8"}},
								NodeCIDRMaskSizeIPv4: ptr.To[int32](24),
								ProxyMode:            resources.IPVSProxyMode,
								IPVS: &kubermaticv1.IPVSConfiguration{
									StrictArp: ptr.To(true),
								},
								DNSDomain: "cluster.local",
							},
							CNIPlugin: &kubermaticv1.CNIPluginSettings{
								Type:    kubermaticv1.CNIPluginTypeCanal,
								Version: cni.GetDefaultCNIPluginVersion(kubermaticv1.CNIPluginTypeCanal),
							},
							ExposeStrategy:        kubermaticv1.ExposeStrategyNodePort,
							Version:               *semver.NewSemverOrDie("9.9.9"),
							EnableUserSSHKeyAgent: ptr.To(false),
						},
						Status: apiv1.ClusterStatus{
							Version:              *semver.NewSemverOrDie("9.9.9"),
							URL:                  "https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885",
							ExternalCCMMigration: apiv1.ExternalCCMMigrationSupported,
						},
						Type: "kubernetes",
					},
				},
			},
			HTTPStatus: http.StatusOK,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(func(seed *kubermaticv1.Seed) {
					seed.Spec.Datacenters["OpenstackDatacenter"] = kubermaticv1.Datacenter{
						Spec: kubermaticv1.DatacenterSpec{
							Openstack: &kubermaticv1.DatacenterSpecOpenstack{},
						},
					}
				}),
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
				test.GenCluster("clusterDefID", "clusterDef", test.GenDefaultProject().Name, time.Date(2013, 02, 04, 01, 54, 0, 0, time.UTC)),
				test.GenClusterWithOpenstack(test.GenCluster("clusterOpenstackID", "clusterOpenstack", test.GenDefaultProject().Name, time.Date(2013, 02, 04, 03, 54, 0, 0, time.UTC))),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 2
		{
			Name: "scenario 2: the admin John can list Bob's clusters",
			ExpectedClusters: apiv2.ProjectClusterList{
				Clusters: apiv1.ClusterList{
					{
						ObjectMeta: apiv1.ObjectMeta{
							ID:                "clusterAbcID",
							Name:              "clusterAbc",
							CreationTimestamp: apiv1.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC),
						},
						Spec: apiv1.ClusterSpec{
							Cloud: kubermaticv1.CloudSpec{
								DatacenterName: "private-do1",
								Fake:           &kubermaticv1.FakeCloudSpec{},
							},
							ContainerRuntime: "containerd",
							ClusterNetwork: &kubermaticv1.ClusterNetworkingConfig{
								IPFamily:             kubermaticv1.IPFamilyIPv4,
								Pods:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"1.2.3.4/8"}},
								Services:             kubermaticv1.NetworkRanges{CIDRBlocks: []string{"5.6.7.8/8"}},
								NodeCIDRMaskSizeIPv4: ptr.To[int32](24),
								ProxyMode:            resources.IPVSProxyMode,
								IPVS: &kubermaticv1.IPVSConfiguration{
									StrictArp: ptr.To(true),
								},
								DNSDomain: "cluster.local",
							},
							CNIPlugin: &kubermaticv1.CNIPluginSettings{
								Type:    kubermaticv1.CNIPluginTypeCanal,
								Version: cni.GetDefaultCNIPluginVersion(kubermaticv1.CNIPluginTypeCanal),
							},
							ExposeStrategy:        kubermaticv1.ExposeStrategyNodePort,
							Version:               *semver.NewSemverOrDie("9.9.9"),
							EnableUserSSHKeyAgent: ptr.To(false),
						},
						Status: apiv1.ClusterStatus{
							Version:              *semver.NewSemverOrDie("9.9.9"),
							URL:                  "https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885",
							ExternalCCMMigration: apiv1.ExternalCCMMigrationUnsupported,
						},
						Type: "kubernetes",
					},
					{
						ObjectMeta: apiv1.ObjectMeta{
							ID:                "clusterDefID",
							Name:              "clusterDef",
							CreationTimestamp: apiv1.Date(2013, 02, 04, 01, 54, 0, 0, time.UTC),
						},
						Spec: apiv1.ClusterSpec{
							Cloud: kubermaticv1.CloudSpec{
								DatacenterName: "private-do1",
								Fake:           &kubermaticv1.FakeCloudSpec{},
							},
							ContainerRuntime: "containerd",
							ClusterNetwork: &kubermaticv1.ClusterNetworkingConfig{
								IPFamily:             kubermaticv1.IPFamilyIPv4,
								Pods:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"1.2.3.4/8"}},
								Services:             kubermaticv1.NetworkRanges{CIDRBlocks: []string{"5.6.7.8/8"}},
								NodeCIDRMaskSizeIPv4: ptr.To[int32](24),
								ProxyMode:            resources.IPVSProxyMode,
								IPVS: &kubermaticv1.IPVSConfiguration{
									StrictArp: ptr.To(true),
								},
								DNSDomain: "cluster.local",
							},
							CNIPlugin: &kubermaticv1.CNIPluginSettings{
								Type:    kubermaticv1.CNIPluginTypeCanal,
								Version: cni.GetDefaultCNIPluginVersion(kubermaticv1.CNIPluginTypeCanal),
							},
							ExposeStrategy:        kubermaticv1.ExposeStrategyNodePort,
							Version:               *semver.NewSemverOrDie("9.9.9"),
							EnableUserSSHKeyAgent: ptr.To(false),
						},
						Status: apiv1.ClusterStatus{
							Version:              *semver.NewSemverOrDie("9.9.9"),
							URL:                  "https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885",
							ExternalCCMMigration: apiv1.ExternalCCMMigrationUnsupported,
						},
						Type: "kubernetes",
					},
					{
						ObjectMeta: apiv1.ObjectMeta{
							ID:                "clusterOpenstackID",
							Name:              "clusterOpenstack",
							CreationTimestamp: apiv1.Date(2013, 02, 04, 03, 54, 0, 0, time.UTC),
						},
						Spec: apiv1.ClusterSpec{
							Cloud: kubermaticv1.CloudSpec{
								DatacenterName: "OpenstackDatacenter",
								Openstack: &kubermaticv1.OpenstackCloudSpec{
									FloatingIPPool: "floatingIPPool",
									SubnetID:       "subnetID",
									Domain:         "domain",
									Network:        "network",
									RouterID:       "routerID",
									SecurityGroups: "securityGroups",
									Project:        "project",
								},
							},
							ContainerRuntime: "containerd",
							ClusterNetwork: &kubermaticv1.ClusterNetworkingConfig{
								IPFamily:             kubermaticv1.IPFamilyIPv4,
								Pods:                 kubermaticv1.NetworkRanges{CIDRBlocks: []string{"1.2.3.4/8"}},
								Services:             kubermaticv1.NetworkRanges{CIDRBlocks: []string{"5.6.7.8/8"}},
								NodeCIDRMaskSizeIPv4: ptr.To[int32](24),
								ProxyMode:            resources.IPVSProxyMode,
								IPVS: &kubermaticv1.IPVSConfiguration{
									StrictArp: ptr.To(true),
								},
								DNSDomain: "cluster.local",
							},
							CNIPlugin: &kubermaticv1.CNIPluginSettings{
								Type:    kubermaticv1.CNIPluginTypeCanal,
								Version: cni.GetDefaultCNIPluginVersion(kubermaticv1.CNIPluginTypeCanal),
							},
							ExposeStrategy:        kubermaticv1.ExposeStrategyNodePort,
							Version:               *semver.NewSemverOrDie("9.9.9"),
							EnableUserSSHKeyAgent: ptr.To(false),
						},
						Status: apiv1.ClusterStatus{
							Version:              *semver.NewSemverOrDie("9.9.9"),
							URL:                  "https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885",
							ExternalCCMMigration: apiv1.ExternalCCMMigrationSupported,
						},
						Type: "kubernetes",
					},
				},
			},
			HTTPStatus: http.StatusOK,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(func(seed *kubermaticv1.Seed) {
					seed.Spec.Datacenters["OpenstackDatacenter"] = kubermaticv1.Datacenter{
						Spec: kubermaticv1.DatacenterSpec{
							Openstack: &kubermaticv1.DatacenterSpecOpenstack{},
						},
					}
				}),
				genUser("John", "john@acme.com", true),
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
				test.GenCluster("clusterDefID", "clusterDef", test.GenDefaultProject().Name, time.Date(2013, 02, 04, 01, 54, 0, 0, time.UTC)),
				test.GenClusterWithOpenstack(test.GenCluster("clusterOpenstackID", "clusterOpenstack", test.GenDefaultProject().Name, time.Date(2013, 02, 04, 03, 54, 0, 0, time.UTC))),
			),
			ExistingAPIUser: test.GenAPIUser("John", "john@acme.com"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v2/projects/%s/clusters", test.ProjectName), strings.NewReader(""))
			res := httptest.NewRecorder()
			var kubermaticObj []ctrlruntimeclient.Object
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
			ep, err := test.CreateTestEndpoint(*tc.ExistingAPIUser, []ctrlruntimeclient.Object{}, kubermaticObj, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			ep.ServeHTTP(res, req)

			if res.Code != tc.HTTPStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.HTTPStatus, res.Code, res.Body.String())
			}

			actualClusters := decodeProjectClusterListOrDie(res.Body, t)

			wrappedExpectedClusters := test.NewClusterV1SliceWrapper(tc.ExpectedClusters.Clusters)
			wrappedExpectedClusters.Sort()

			actualClusters.EqualOrDie(wrappedExpectedClusters, t)
		})
	}
}

func TestGetCluster(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		Name                   string
		Body                   string
		ExpectedResponse       string
		HTTPStatus             int
		ClusterToGet           string
		ExistingAPIUser        *apiv1.User
		ExistingKubermaticObjs []ctrlruntimeclient.Object
	}{
		// scenario 1
		{
			Name:             "scenario 1: gets cluster with the given name that belongs to the given project",
			Body:             ``,
			ExpectedResponse: `{"id":"defClusterID","name":"defClusterName","creationTimestamp":"2013-02-03T19:54:00Z","type":"kubernetes","spec":{"cloud":{"dc":"private-do1","fake":{}},"version":"9.9.9","oidc":{},"enableUserSSHKeyAgent":false,"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["5.6.7.8/8"]},"pods":{"cidrBlocks":["1.2.3.4/8"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true}},"cniPlugin":{"type":"canal","version":"v3.29"},"exposeStrategy":"NodePort"},"status":{"version":"9.9.9","url":"https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885","externalCCMMigration":"Unsupported"}}`,
			ClusterToGet:     test.GenDefaultCluster().Name,
			HTTPStatus:       http.StatusOK,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 2
		{
			Name:             "scenario 2: gets cluster for Openstack auth with project and no sensitive data (credentials) are returned",
			Body:             ``,
			ExpectedResponse: `{"id":"defClusterID","name":"defClusterName","creationTimestamp":"2013-02-03T19:54:00Z","type":"kubernetes","spec":{"cloud":{"dc":"OpenstackDatacenter","openstack":{"floatingIPPool":"floatingIPPool","project":"project","domain":"domain","network":"network","securityGroups":"securityGroups","routerID":"routerID","subnetID":"subnetID"}},"version":"9.9.9","oidc":{},"enableUserSSHKeyAgent":false,"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["5.6.7.8/8"]},"pods":{"cidrBlocks":["1.2.3.4/8"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true}},"cniPlugin":{"type":"canal","version":"v3.29"},"exposeStrategy":"NodePort"},"status":{"version":"9.9.9","url":"https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885","externalCCMMigration":"Supported"}}`,
			ClusterToGet:     test.GenDefaultCluster().Name,
			HTTPStatus:       http.StatusOK,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(func(seed *kubermaticv1.Seed) {
					seed.Spec.Datacenters["OpenstackDatacenter"] = kubermaticv1.Datacenter{
						Spec: kubermaticv1.DatacenterSpec{
							Openstack: &kubermaticv1.DatacenterSpecOpenstack{},
						},
					}
				}),
				test.GenClusterWithOpenstack(test.GenDefaultCluster()),
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 3
		{
			Name:             "scenario 3: gets cluster for Openstack auth with project and no sensitive data (credentials) are returned",
			Body:             ``,
			ExpectedResponse: `{"id":"defClusterID","name":"defClusterName","creationTimestamp":"2013-02-03T19:54:00Z","type":"kubernetes","spec":{"cloud":{"dc":"OpenstackDatacenter","openstack":{"floatingIPPool":"floatingIPPool","project":"project","projectID":"projectID","domain":"domain","network":"network","securityGroups":"securityGroups","routerID":"routerID","subnetID":"subnetID"}},"version":"9.9.9","oidc":{},"enableUserSSHKeyAgent":false,"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["5.6.7.8/8"]},"pods":{"cidrBlocks":["1.2.3.4/8"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true}},"cniPlugin":{"type":"canal","version":"v3.29"},"exposeStrategy":"NodePort"},"status":{"version":"9.9.9","url":"https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885","externalCCMMigration":"Supported"}}`,
			ClusterToGet:     test.GenDefaultCluster().Name,
			HTTPStatus:       http.StatusOK,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(func(seed *kubermaticv1.Seed) {
					seed.Spec.Datacenters["OpenstackDatacenter"] = kubermaticv1.Datacenter{
						Spec: kubermaticv1.DatacenterSpec{
							Openstack: &kubermaticv1.DatacenterSpecOpenstack{},
						},
					}
				}),
				test.GenClusterWithOpenstackProjectAuth(test.GenDefaultCluster()),
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 4
		{
			Name:             "scenario 4: the admin John can get Bob's cluster",
			Body:             ``,
			ExpectedResponse: `{"id":"defClusterID","name":"defClusterName","creationTimestamp":"2013-02-03T19:54:00Z","type":"kubernetes","spec":{"cloud":{"dc":"OpenstackDatacenter","openstack":{"floatingIPPool":"floatingIPPool","project":"project","domain":"domain","network":"network","securityGroups":"securityGroups","routerID":"routerID","subnetID":"subnetID"}},"version":"9.9.9","oidc":{},"enableUserSSHKeyAgent":false,"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["5.6.7.8/8"]},"pods":{"cidrBlocks":["1.2.3.4/8"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true}},"cniPlugin":{"type":"canal","version":"v3.29"},"exposeStrategy":"NodePort"},"status":{"version":"9.9.9","url":"https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885","externalCCMMigration":"Supported"}}`,
			ClusterToGet:     test.GenDefaultCluster().Name,
			HTTPStatus:       http.StatusOK,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(func(seed *kubermaticv1.Seed) {
					seed.Spec.Datacenters["OpenstackDatacenter"] = kubermaticv1.Datacenter{
						Spec: kubermaticv1.DatacenterSpec{
							Openstack: &kubermaticv1.DatacenterSpecOpenstack{},
						},
					}
				}),
				genUser("John", "john@acme.com", true),
				test.GenClusterWithOpenstack(test.GenDefaultCluster()),
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
			),
			ExistingAPIUser: test.GenAPIUser("John", "john@acme.com"),
		},
		// scenario 5
		{
			Name:             "scenario 5: the regular user John can not get Bob's cluster",
			Body:             ``,
			ExpectedResponse: `{"error":{"code":403,"message":"forbidden: \"john@acme.com\" doesn't belong to project my-first-project-ID"}}`,
			ClusterToGet:     test.GenDefaultCluster().Name,
			HTTPStatus:       http.StatusForbidden,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				genUser("John", "john@acme.com", false),
				test.GenClusterWithOpenstack(test.GenDefaultCluster()),
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
			),
			ExistingAPIUser: test.GenAPIUser("John", "john@acme.com"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v2/projects/%s/clusters/%s", test.ProjectName, tc.ClusterToGet), strings.NewReader(tc.Body))
			res := httptest.NewRecorder()
			var kubermaticObj []ctrlruntimeclient.Object
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
			ep, err := test.CreateTestEndpoint(*tc.ExistingAPIUser, []ctrlruntimeclient.Object{}, kubermaticObj, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			ep.ServeHTTP(res, req)

			if res.Code != tc.HTTPStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.HTTPStatus, res.Code, res.Body.String())
			}

			test.CompareWithResult(t, res, tc.ExpectedResponse)
		})
	}
}

func TestDeleteClusterEndpoint(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		Name                          string
		Body                          string
		ExpectedResponse              string
		HTTPStatus                    int
		ProjectToSync                 string
		ClusterToSync                 string
		ExistingKubermaticObjs        []ctrlruntimeclient.Object
		ExistingAPIUser               *apiv1.User
		ExpectedListClusterKeysStatus int
	}{
		{
			Name:             "scenario 1: tests deletion of a cluster and its dependent resources",
			Body:             ``,
			ExpectedResponse: `{}`,
			HTTPStatus:       http.StatusOK,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add a cluster
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
				// add ssh keys
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-c08aa5c7abf34504f18552846485267d-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{"clusterAbcID"},
					},
				},
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-abc-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{"clusterAbcID"},
					},
				},
			),
			ClusterToSync:                 "clusterAbcID",
			ExistingAPIUser:               test.GenDefaultAPIUser(),
			ExpectedListClusterKeysStatus: http.StatusNotFound,
		},
		{
			Name:             "scenario 2: the admin John can delete Bob's cluster",
			Body:             ``,
			ExpectedResponse: `{}`,
			HTTPStatus:       http.StatusOK,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				genUser("John", "john@acme.com", true),
				// add a cluster
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
				// add ssh keys
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-c08aa5c7abf34504f18552846485267d-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{"clusterAbcID"},
					},
				},
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-abc-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{"clusterAbcID"},
					},
				},
			),
			ClusterToSync:                 "clusterAbcID",
			ExistingAPIUser:               test.GenAPIUser("John", "john@acme.com"),
			ExpectedListClusterKeysStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			// validate if deletion was successful
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v2/projects/%s/clusters/%s", tc.ProjectToSync, tc.ClusterToSync), strings.NewReader(tc.Body))
			res := httptest.NewRecorder()
			var kubermaticObj []ctrlruntimeclient.Object
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
			ep, err := test.CreateTestEndpoint(*tc.ExistingAPIUser, []ctrlruntimeclient.Object{}, kubermaticObj, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			ep.ServeHTTP(res, req)

			if res.Code != tc.HTTPStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.HTTPStatus, res.Code, res.Body.String())
			}
			test.CompareWithResult(t, res, tc.ExpectedResponse)

			// validate if the cluster was deleted
			req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/projects/%s/dc/us-central1/clusters/abcd/sshkeys", tc.ProjectToSync), strings.NewReader(tc.Body))
			res = httptest.NewRecorder()
			ep.ServeHTTP(res, req)
			if res.Code != tc.ExpectedListClusterKeysStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.ExpectedListClusterKeysStatus, res.Code, res.Body.String())
			}
		})
	}
}

func TestPatchCluster(t *testing.T) {
	t.Parallel()

	const fakeDC = "fake-dc"
	cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
	cluster.Spec.Cloud.DatacenterName = usCentral1

	testcases := []struct {
		Name                      string
		Body                      string
		ExpectedResponse          string
		HTTPStatus                int
		cluster                   string
		project                   string
		ExistingAPIUser           *apiv1.User
		ExistingMachines          []*clusterv1alpha1.Machine
		ExistingKubermaticObjects []ctrlruntimeclient.Object
	}{
		// scenario 1
		{
			Name:             "scenario 1: update the cluster version",
			Body:             `{"spec":{"version":"9.9.10"}}`,
			ExpectedResponse: `{"id":"keen-snyder","name":"clusterAbc","creationTimestamp":"2013-02-03T19:54:00Z","type":"kubernetes","spec":{"cloud":{"dc":"fake-dc","fake":{}},"version":"9.9.10","oidc":{},"enableUserSSHKeyAgent":false,"kubernetesDashboard":{"enabled":true},"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["5.6.7.8/8"]},"pods":{"cidrBlocks":["1.2.3.4/8"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true},"nodeLocalDNSCacheEnabled":true},"cniPlugin":{"type":"canal","version":"v3.29"},"exposeStrategy":"NodePort"},"status":{"version":"9.9.9","url":"https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885","externalCCMMigration":"Unsupported"}}`,
			cluster:          "keen-snyder",
			HTTPStatus:       http.StatusOK,
			project:          test.GenDefaultProject().Name,
			ExistingAPIUser:  test.GenDefaultAPIUser(),
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Spec.Cloud.DatacenterName = fakeDC
					return cluster
				}(),
			),
		},
		// scenario 2
		{
			Name:             "scenario 2: fail on invalid patch json",
			Body:             `{"spec":{"cloud":{"dc":"dc1"`,
			ExpectedResponse: `{"error":{"code":400,"message":"cannot patch cluster: Invalid JSON Patch"}}`,
			cluster:          "keen-snyder",
			HTTPStatus:       http.StatusBadRequest,
			project:          test.GenDefaultProject().Name,
			ExistingAPIUser:  test.GenDefaultAPIUser(),
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
			),
		},
		// scenario 3
		{
			Name:             "scenario 3: tried to update cluster with older but compatible nodes",
			Body:             `{"spec":{"version":"9.11.3"}}`, // kubelet is 9.9.9, maximum compatible master is 9.11.x
			ExpectedResponse: `{"id":"keen-snyder","name":"clusterAbc","creationTimestamp":"2013-02-03T19:54:00Z","type":"kubernetes","spec":{"cloud":{"dc":"fake-dc","fake":{}},"version":"9.11.3","oidc":{},"enableUserSSHKeyAgent":false,"kubernetesDashboard":{"enabled":true},"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["5.6.7.8/8"]},"pods":{"cidrBlocks":["1.2.3.4/8"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true},"nodeLocalDNSCacheEnabled":true},"cniPlugin":{"type":"canal","version":"v3.29"},"exposeStrategy":"NodePort"},"status":{"version":"9.9.9","url":"https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885","externalCCMMigration":"Unsupported"}}`,
			cluster:          "keen-snyder",
			HTTPStatus:       http.StatusOK,
			project:          test.GenDefaultProject().Name,
			ExistingAPIUser:  test.GenDefaultAPIUser(),
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Spec.Cloud.DatacenterName = fakeDC
					return cluster
				}(),
			),
			ExistingMachines: []*clusterv1alpha1.Machine{
				test.GenTestMachine("venus", `{"cloudProvider":"digitalocean","cloudProviderSpec":{"token":"dummy-token","region":"fra1","size":"2GB"},"operatingSystem":"ubuntu","containerRuntimeInfo":{"name":"docker","version":"1.13"},"operatingSystemSpec":{"distUpgradeOnBoot":true}}`, map[string]string{"md-id": "123", "some-other": "xyz"}, nil),
				test.GenTestMachine("mars", `{"cloudProvider":"aws","cloudProviderSpec":{"token":"dummy-token","region":"eu-central-1","availabilityZone":"eu-central-1a","vpcId":"vpc-819f62e9","subnetId":"subnet-2bff4f43","instanceType":"t2.micro","diskSize":50}, "containerRuntimeInfo":{"name":"docker","version":"1.12"},"operatingSystem":"ubuntu", "operatingSystemSpec":{"distUpgradeOnBoot":false}}`, map[string]string{"md-id": "123", "some-other": "xyz"}, nil),
			},
		},
		// scenario 4
		{
			Name:             "scenario 4: tried to update cluster with old nodes",
			Body:             `{"spec":{"version":"9.12.3"}}`, // kubelet is 9.9.9, maximum compatible master is 9.11.x
			ExpectedResponse: `{"error":{"code":400,"message":"Cluster contains nodes running the following incompatible kubelet versions: [9.9.9]. Upgrade your nodes before you upgrade the cluster."}}`,
			cluster:          "keen-snyder",
			HTTPStatus:       http.StatusBadRequest,
			project:          test.GenDefaultProject().Name,
			ExistingAPIUser:  test.GenDefaultAPIUser(),
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(func(seed *kubermaticv1.Seed) {
					seed.Spec.Datacenters[usCentral1] = kubermaticv1.Datacenter{}
				}),
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Spec.Cloud.DatacenterName = usCentral1
					return cluster
				}(),
			),
			ExistingMachines: []*clusterv1alpha1.Machine{
				test.GenTestMachine("venus", `{"cloudProvider":"digitalocean","cloudProviderSpec":{"token":"dummy-token","region":"fra1","size":"2GB"},"operatingSystem":"ubuntu","containerRuntimeInfo":{"name":"docker","version":"1.13"},"operatingSystemSpec":{"distUpgradeOnBoot":true}}`, map[string]string{"md-id": "123", "some-other": "xyz"}, nil),
				test.GenTestMachine("mars", `{"cloudProvider":"aws","cloudProviderSpec":{"token":"dummy-token","region":"eu-central-1","availabilityZone":"eu-central-1a","vpcId":"vpc-819f62e9","subnetId":"subnet-2bff4f43","instanceType":"t2.micro","diskSize":50}, "containerRuntimeInfo":{"name":"docker","version":"1.12"},"operatingSystem":"ubuntu", "operatingSystemSpec":{"distUpgradeOnBoot":false}}`, map[string]string{"md-id": "123", "some-other": "xyz"}, nil),
			},
		},
		// scenario 5
		{
			Name:             "scenario 5: tried to downgrade cluster to version older than its nodes",
			Body:             `{"spec":{"version":"9.8.12"}}`, // kubelet is 9.9.9, cluster cannot be older
			ExpectedResponse: `{"error":{"code":400,"message":"Cluster contains nodes running the following incompatible kubelet versions: [9.9.9]. Upgrade your nodes before you upgrade the cluster."}}`,
			cluster:          "keen-snyder",
			HTTPStatus:       http.StatusBadRequest,
			project:          test.GenDefaultProject().Name,
			ExistingAPIUser:  test.GenDefaultAPIUser(),
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(func(seed *kubermaticv1.Seed) {
					seed.Spec.Datacenters[usCentral1] = kubermaticv1.Datacenter{}
				}),
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Spec.Cloud.DatacenterName = usCentral1
					return cluster
				}(),
			),
			ExistingMachines: []*clusterv1alpha1.Machine{
				test.GenTestMachine("venus", `{"cloudProvider":"digitalocean","cloudProviderSpec":{"token":"dummy-token","region":"fra1","size":"2GB"},"operatingSystem":"ubuntu","containerRuntimeInfo":{"name":"docker","version":"1.13"},"operatingSystemSpec":{"distUpgradeOnBoot":true}}`, map[string]string{"md-id": "123", "some-other": "xyz"}, nil),
				test.GenTestMachine("mars", `{"cloudProvider":"aws","cloudProviderSpec":{"token":"dummy-token","region":"eu-central-1","availabilityZone":"eu-central-1a","vpcId":"vpc-819f62e9","subnetId":"subnet-2bff4f43","instanceType":"t2.micro","diskSize":50}, "containerRuntimeInfo":{"name":"docker","version":"1.12"},"operatingSystem":"ubuntu", "operatingSystemSpec":{"distUpgradeOnBoot":false}}`, map[string]string{"md-id": "123", "some-other": "xyz"}, nil),
			},
		},
		// scenario 6
		{
			Name:             "scenario 6: the admin John can update Bob's cluster version",
			Body:             `{"spec":{"version":"9.9.10"}}`,
			ExpectedResponse: `{"id":"keen-snyder","name":"clusterAbc","creationTimestamp":"2013-02-03T19:54:00Z","type":"kubernetes","spec":{"cloud":{"dc":"fake-dc","fake":{}},"version":"9.9.10","oidc":{},"enableUserSSHKeyAgent":false,"kubernetesDashboard":{"enabled":true},"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["5.6.7.8/8"]},"pods":{"cidrBlocks":["1.2.3.4/8"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true},"nodeLocalDNSCacheEnabled":true},"cniPlugin":{"type":"canal","version":"v3.29"},"exposeStrategy":"NodePort"},"status":{"version":"9.9.9","url":"https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885","externalCCMMigration":"Unsupported"}}`,
			cluster:          "keen-snyder",
			HTTPStatus:       http.StatusOK,
			project:          test.GenDefaultProject().Name,
			ExistingAPIUser:  test.GenAPIUser("John", "john@acme.com"),
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Spec.Cloud.DatacenterName = fakeDC
					return cluster
				}(),
				genUser("John", "john@acme.com", true),
			),
		},
		// scenario 7
		{
			Name:             "scenario 7: the regular user John can not update Bob's cluster version",
			Body:             `{"spec":{"version":"9.9.9"}}`,
			ExpectedResponse: `{"error":{"code":403,"message":"forbidden: \"john@acme.com\" doesn't belong to project my-first-project-ID"}}`,
			cluster:          "keen-snyder",
			HTTPStatus:       http.StatusForbidden,
			project:          test.GenDefaultProject().Name,
			ExistingAPIUser:  test.GenAPIUser("John", "john@acme.com"),
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Spec.Cloud.DatacenterName = fakeDC
					return cluster
				}(),
				genUser("John", "john@acme.com", false),
			),
		},
		// scenario 8
		{
			Name:             "scenario 8: update the pod node selector admission plugin config",
			Body:             `{"spec":{"version":"9.9.9","podNodeSelectorAdmissionPluginConfig":{"clusterDefaultNodeSelector":"env=development"}}}`,
			ExpectedResponse: `{"id":"keen-snyder","name":"clusterAbc","creationTimestamp":"2013-02-03T19:54:00Z","type":"kubernetes","spec":{"cloud":{"dc":"fake-dc","fake":{}},"version":"9.9.9","oidc":{},"enableUserSSHKeyAgent":false,"kubernetesDashboard":{"enabled":true},"podNodeSelectorAdmissionPluginConfig":{"clusterDefaultNodeSelector":"env=development"},"containerRuntime":"containerd","clusterNetwork":{"ipFamily":"IPv4","services":{"cidrBlocks":["5.6.7.8/8"]},"pods":{"cidrBlocks":["1.2.3.4/8"]},"nodeCidrMaskSizeIPv4":24,"dnsDomain":"cluster.local","proxyMode":"ipvs","ipvs":{"strictArp":true},"nodeLocalDNSCacheEnabled":true},"cniPlugin":{"type":"canal","version":"v3.29"},"exposeStrategy":"NodePort"},"status":{"version":"9.9.9","url":"https://w225mx4z66.asia-east1-a-1.cloud.kubermatic.io:31885","externalCCMMigration":"Unsupported"}}`,
			cluster:          "keen-snyder",
			HTTPStatus:       http.StatusOK,
			project:          test.GenDefaultProject().Name,
			ExistingAPIUser:  test.GenDefaultAPIUser(),
			ExistingKubermaticObjects: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Spec.Cloud.DatacenterName = fakeDC
					return cluster
				}(),
			),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			var machineObj []ctrlruntimeclient.Object
			for _, existingMachine := range tc.ExistingMachines {
				machineObj = append(machineObj, existingMachine)
			}
			// test data
			req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/api/v2/projects/%s/clusters/%s", tc.project, tc.cluster), strings.NewReader(tc.Body))
			res := httptest.NewRecorder()
			ep, _, err := test.CreateTestEndpointAndGetClients(*tc.ExistingAPIUser, nil, []ctrlruntimeclient.Object{}, machineObj, tc.ExistingKubermaticObjects, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			// act
			ep.ServeHTTP(res, req)

			// validate
			if res.Code != tc.HTTPStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.HTTPStatus, res.Code, res.Body.String())
			}

			test.CompareWithResult(t, res, tc.ExpectedResponse)
		})
	}
}

func TestGetClusterEventsEndpoint(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		Name                   string
		HTTPStatus             int
		ExpectedResult         string
		ProjectIDToSync        string
		ClusterIDToSync        string
		ExistingProject        *kubermaticv1.Project
		ExistingKubermaticUser *kubermaticv1.User
		ExistingAPIUser        *apiv1.User
		ExistingKubermaticObjs []ctrlruntimeclient.Object
		ExistingEvents         []*corev1.Event
		NodeDeploymentID       string
		QueryParams            string
	}{
		// scenario 1
		{
			Name:            "scenario 1: list all events",
			HTTPStatus:      http.StatusOK,
			ClusterIDToSync: test.GenDefaultCluster().Name,
			ProjectIDToSync: test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
			ExistingEvents: []*corev1.Event{
				test.GenTestEvent("event-1", corev1.EventTypeNormal, "Started", "message started", "Cluster", "venus-1-machine", test.GenDefaultCluster().Name),
				test.GenTestEvent("event-2", corev1.EventTypeWarning, "Killed", "message killed", "Cluster", "venus-1-machine", test.GenDefaultCluster().Name),
			},
			ExpectedResult: `[{"name":"event-1","creationTimestamp":"0001-01-01T00:00:00Z","message":"message started","type":"Normal","involvedObject":{"type":"Cluster","namespace":"kube-system","name":"defClusterID"},"lastTimestamp":"0001-01-01T00:00:00Z","count":1},{"name":"event-2","creationTimestamp":"0001-01-01T00:00:00Z","message":"message killed","type":"Warning","involvedObject":{"type":"Cluster","namespace":"kube-system","name":"defClusterID"},"lastTimestamp":"0001-01-01T00:00:00Z","count":1}]`,
		},
		// scenario 2
		{
			Name:            "scenario 2: list all warning events",
			QueryParams:     "?type=warning",
			HTTPStatus:      http.StatusOK,
			ClusterIDToSync: test.GenDefaultCluster().Name,
			ProjectIDToSync: test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
			ExistingEvents: []*corev1.Event{
				test.GenTestEvent("event-1", corev1.EventTypeNormal, "Started", "message started", "Cluster", "venus-1-machine", test.GenDefaultCluster().Name),
				test.GenTestEvent("event-2", corev1.EventTypeWarning, "Killed", "message killed", "Cluster", "venus-1-machine", test.GenDefaultCluster().Name),
			},
			ExpectedResult: `[{"name":"event-2","creationTimestamp":"0001-01-01T00:00:00Z","message":"message killed","type":"Warning","involvedObject":{"type":"Cluster","namespace":"kube-system","name":"defClusterID"},"lastTimestamp":"0001-01-01T00:00:00Z","count":1}]`,
		},
		// scenario 3
		{
			Name:            "scenario 3: list all normal events",
			QueryParams:     "?type=normal",
			HTTPStatus:      http.StatusOK,
			ClusterIDToSync: test.GenDefaultCluster().Name,
			ProjectIDToSync: test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
			ExistingEvents: []*corev1.Event{
				test.GenTestEvent("event-1", corev1.EventTypeNormal, "Started", "message started", "Cluster", "venus-1-machine", test.GenDefaultCluster().Name),
				test.GenTestEvent("event-2", corev1.EventTypeWarning, "Killed", "message killed", "Cluster", "venus-1-machine", test.GenDefaultCluster().Name),
			},
			ExpectedResult: `[{"name":"event-1","creationTimestamp":"0001-01-01T00:00:00Z","message":"message started","type":"Normal","involvedObject":{"type":"Cluster","namespace":"kube-system","name":"defClusterID"},"lastTimestamp":"0001-01-01T00:00:00Z","count":1}]`,
		},
		// scenario 4
		{
			Name:            "scenario 4: the admin John can list Bob's cluster events",
			HTTPStatus:      http.StatusOK,
			ClusterIDToSync: test.GenDefaultCluster().Name,
			ProjectIDToSync: test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				genUser("John", "john@acme.com", true),
			),
			ExistingAPIUser: test.GenAPIUser("John", "john@acme.com"),
			ExistingEvents: []*corev1.Event{
				test.GenTestEvent("event-1", corev1.EventTypeNormal, "Started", "message started", "Cluster", "venus-1-machine", test.GenDefaultCluster().Name),
				test.GenTestEvent("event-2", corev1.EventTypeWarning, "Killed", "message killed", "Cluster", "venus-1-machine", test.GenDefaultCluster().Name),
			},
			ExpectedResult: `[{"name":"event-1","creationTimestamp":"0001-01-01T00:00:00Z","message":"message started","type":"Normal","involvedObject":{"type":"Cluster","namespace":"kube-system","name":"defClusterID"},"lastTimestamp":"0001-01-01T00:00:00Z","count":1},{"name":"event-2","creationTimestamp":"0001-01-01T00:00:00Z","message":"message killed","type":"Warning","involvedObject":{"type":"Cluster","namespace":"kube-system","name":"defClusterID"},"lastTimestamp":"0001-01-01T00:00:00Z","count":1}]`,
		},
		// scenario 5
		{
			Name:            "scenario 5: the user John can not list Bob's cluster events",
			HTTPStatus:      http.StatusForbidden,
			ClusterIDToSync: test.GenDefaultCluster().Name,
			ProjectIDToSync: test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				genUser("John", "john@acme.com", false),
			),
			ExistingAPIUser: test.GenAPIUser("John", "john@acme.com"),
			ExistingEvents: []*corev1.Event{
				test.GenTestEvent("event-1", corev1.EventTypeNormal, "Started", "message started", "Cluster", "venus-1-machine", test.GenDefaultCluster().Name),
				test.GenTestEvent("event-2", corev1.EventTypeWarning, "Killed", "message killed", "Cluster", "venus-1-machine", test.GenDefaultCluster().Name),
			},
			ExpectedResult: `{"error":{"code":403,"message":"forbidden: \"john@acme.com\" doesn't belong to project my-first-project-ID"}}`,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v2/projects/%s/clusters/%s/events%s", tc.ProjectIDToSync, tc.ClusterIDToSync, tc.QueryParams), strings.NewReader(""))
			res := httptest.NewRecorder()
			kubermaticObj := make([]ctrlruntimeclient.Object, 0)
			machineObj := make([]ctrlruntimeclient.Object, 0)
			kubernetesObj := make([]ctrlruntimeclient.Object, 0)
			for _, existingEvents := range tc.ExistingEvents {
				kubernetesObj = append(kubernetesObj, existingEvents)
			}
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)

			ep, _, err := test.CreateTestEndpointAndGetClients(*tc.ExistingAPIUser, nil, kubernetesObj, machineObj, kubermaticObj, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			ep.ServeHTTP(res, req)

			if res.Code != tc.HTTPStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.HTTPStatus, res.Code, res.Body.String())
			}

			test.CompareWithResult(t, res, tc.ExpectedResult)
		})
	}
}

func TestGetClusterHealth(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		Name                   string
		Body                   string
		ExpectedResponse       string
		HTTPStatus             int
		ClusterToGet           string
		ProjectToSync          string
		ExistingAPIUser        *apiv1.User
		ExistingKubermaticObjs []ctrlruntimeclient.Object
	}{
		// scenario 1
		{
			Name:             "scenario 1: get existing cluster health status",
			Body:             ``,
			ExpectedResponse: `{"apiserver":"HealthStatusUp","applicationController":"HealthStatusUp","scheduler":"HealthStatusDown","controller":"HealthStatusUp","machineController":"HealthStatusDown","etcd":"HealthStatusUp","cloudProviderInfrastructure":"HealthStatusUp","userClusterControllerManager":"HealthStatusUp"}`,
			HTTPStatus:       http.StatusOK,
			ClusterToGet:     "keen-snyder",
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add a cluster
				test.GenCluster("clusterDefID", "clusterDef", test.GenDefaultProject().Name, time.Date(2013, 02, 04, 01, 54, 0, 0, time.UTC)),
				// add another cluster
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Status.ExtendedHealth = kubermaticv1.ExtendedClusterHealth{
						Apiserver:                    kubermaticv1.HealthStatusUp,
						ApplicationController:        kubermaticv1.HealthStatusUp,
						Scheduler:                    kubermaticv1.HealthStatusDown,
						Controller:                   kubermaticv1.HealthStatusUp,
						MachineController:            kubermaticv1.HealthStatusDown,
						Etcd:                         kubermaticv1.HealthStatusUp,
						CloudProviderInfrastructure:  kubermaticv1.HealthStatusUp,
						UserClusterControllerManager: kubermaticv1.HealthStatusUp,
					}
					return cluster
				}(),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 2
		{
			Name:             "scenario 2: the admin Bob can get John's cluster health status",
			Body:             ``,
			ExpectedResponse: `{"apiserver":"HealthStatusUp","applicationController":"HealthStatusUp","scheduler":"HealthStatusDown","controller":"HealthStatusUp","machineController":"HealthStatusDown","etcd":"HealthStatusUp","cloudProviderInfrastructure":"HealthStatusUp","userClusterControllerManager":"HealthStatusUp"}`,
			HTTPStatus:       http.StatusOK,
			ClusterToGet:     "keen-snyder",
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add admin user
				genUser("John", "john@acme.com", true),
				// add a cluster
				test.GenCluster("clusterDefID", "clusterDef", test.GenDefaultProject().Name, time.Date(2013, 02, 04, 01, 54, 0, 0, time.UTC)),
				// add another cluster
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Status.ExtendedHealth = kubermaticv1.ExtendedClusterHealth{
						Apiserver:                    kubermaticv1.HealthStatusUp,
						ApplicationController:        kubermaticv1.HealthStatusUp,
						Scheduler:                    kubermaticv1.HealthStatusDown,
						Controller:                   kubermaticv1.HealthStatusUp,
						MachineController:            kubermaticv1.HealthStatusDown,
						Etcd:                         kubermaticv1.HealthStatusUp,
						CloudProviderInfrastructure:  kubermaticv1.HealthStatusUp,
						UserClusterControllerManager: kubermaticv1.HealthStatusUp,
					}
					return cluster
				}(),
			),
			ExistingAPIUser: test.GenAPIUser("John", "john@acme.com"),
		},
		// scenario 3
		{
			Name:             "scenario 3: the user Bob can not get John's cluster health status",
			Body:             ``,
			ExpectedResponse: `{"error":{"code":403,"message":"forbidden: \"john@acme.com\" doesn't belong to project my-first-project-ID"}}`,
			HTTPStatus:       http.StatusForbidden,
			ClusterToGet:     "keen-snyder",
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add regular user John
				genUser("John", "john@acme.com", false),
				// add a cluster
				test.GenCluster("clusterDefID", "clusterDef", test.GenDefaultProject().Name, time.Date(2013, 02, 04, 01, 54, 0, 0, time.UTC)),
				// add another cluster
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Status.ExtendedHealth = kubermaticv1.ExtendedClusterHealth{
						Apiserver:                    kubermaticv1.HealthStatusUp,
						Scheduler:                    kubermaticv1.HealthStatusDown,
						Controller:                   kubermaticv1.HealthStatusUp,
						MachineController:            kubermaticv1.HealthStatusDown,
						Etcd:                         kubermaticv1.HealthStatusUp,
						CloudProviderInfrastructure:  kubermaticv1.HealthStatusUp,
						UserClusterControllerManager: kubermaticv1.HealthStatusUp,
					}
					return cluster
				}(),
			),
			ExistingAPIUser: test.GenAPIUser("John", "john@acme.com"),
		},
		// scenario 4
		{
			Name:             "scenario 4: get existing cluster health status with opa integration enabled",
			Body:             ``,
			ExpectedResponse: `{"apiserver":"HealthStatusUp","applicationController":"HealthStatusUp","scheduler":"HealthStatusDown","controller":"HealthStatusUp","machineController":"HealthStatusDown","etcd":"HealthStatusUp","cloudProviderInfrastructure":"HealthStatusUp","userClusterControllerManager":"HealthStatusUp","gatekeeperController":"HealthStatusUp","gatekeeperAudit":"HealthStatusUp"}`,
			HTTPStatus:       http.StatusOK,
			ClusterToGet:     "keen-snyder",
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add a cluster
				test.GenCluster("clusterDefID", "clusterDef", test.GenDefaultProject().Name, time.Date(2013, 02, 04, 01, 54, 0, 0, time.UTC)),
				// add another cluster
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Status.ExtendedHealth = kubermaticv1.ExtendedClusterHealth{
						Apiserver:                    kubermaticv1.HealthStatusUp,
						ApplicationController:        kubermaticv1.HealthStatusUp,
						Scheduler:                    kubermaticv1.HealthStatusDown,
						Controller:                   kubermaticv1.HealthStatusUp,
						MachineController:            kubermaticv1.HealthStatusDown,
						Etcd:                         kubermaticv1.HealthStatusUp,
						CloudProviderInfrastructure:  kubermaticv1.HealthStatusUp,
						UserClusterControllerManager: kubermaticv1.HealthStatusUp,
						GatekeeperAudit:              &healthUp,
						GatekeeperController:         &healthUp,
					}
					cluster.Spec.OPAIntegration = &kubermaticv1.OPAIntegrationSettings{Enabled: true}
					return cluster
				}(),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 5
		{
			Name:             "scenario 5: get existing cluster health status with MLA Monitoring enabled",
			Body:             ``,
			ExpectedResponse: `{"apiserver":"HealthStatusUp","applicationController":"HealthStatusUp","scheduler":"HealthStatusDown","controller":"HealthStatusUp","machineController":"HealthStatusDown","etcd":"HealthStatusUp","cloudProviderInfrastructure":"HealthStatusUp","userClusterControllerManager":"HealthStatusUp","monitoring":"HealthStatusUp"}`,
			HTTPStatus:       http.StatusOK,
			ClusterToGet:     "keen-snyder",
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add a cluster
				test.GenCluster("clusterDefID", "clusterDef", test.GenDefaultProject().Name, time.Date(2013, 02, 04, 01, 54, 0, 0, time.UTC)),
				// add another cluster
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Status.ExtendedHealth = kubermaticv1.ExtendedClusterHealth{
						Apiserver:                    kubermaticv1.HealthStatusUp,
						ApplicationController:        kubermaticv1.HealthStatusUp,
						Scheduler:                    kubermaticv1.HealthStatusDown,
						Controller:                   kubermaticv1.HealthStatusUp,
						MachineController:            kubermaticv1.HealthStatusDown,
						Etcd:                         kubermaticv1.HealthStatusUp,
						CloudProviderInfrastructure:  kubermaticv1.HealthStatusUp,
						UserClusterControllerManager: kubermaticv1.HealthStatusUp,
						Monitoring:                   &healthUp,
					}
					cluster.Spec.MLA = &kubermaticv1.MLASettings{MonitoringEnabled: true}
					return cluster
				}(),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 6
		{
			Name:             "scenario 6: get existing cluster health status with MLA Logging enabled",
			Body:             ``,
			ExpectedResponse: `{"apiserver":"HealthStatusUp","applicationController":"HealthStatusUp","scheduler":"HealthStatusDown","controller":"HealthStatusUp","machineController":"HealthStatusDown","etcd":"HealthStatusUp","cloudProviderInfrastructure":"HealthStatusUp","userClusterControllerManager":"HealthStatusUp","logging":"HealthStatusUp"}`,
			HTTPStatus:       http.StatusOK,
			ClusterToGet:     "keen-snyder",
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add a cluster
				test.GenCluster("clusterDefID", "clusterDef", test.GenDefaultProject().Name, time.Date(2013, 02, 04, 01, 54, 0, 0, time.UTC)),
				// add another cluster
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Status.ExtendedHealth = kubermaticv1.ExtendedClusterHealth{
						Apiserver:                    kubermaticv1.HealthStatusUp,
						ApplicationController:        kubermaticv1.HealthStatusUp,
						Scheduler:                    kubermaticv1.HealthStatusDown,
						Controller:                   kubermaticv1.HealthStatusUp,
						MachineController:            kubermaticv1.HealthStatusDown,
						Etcd:                         kubermaticv1.HealthStatusUp,
						CloudProviderInfrastructure:  kubermaticv1.HealthStatusUp,
						UserClusterControllerManager: kubermaticv1.HealthStatusUp,
						Logging:                      &healthUp,
					}
					cluster.Spec.MLA = &kubermaticv1.MLASettings{LoggingEnabled: true}
					return cluster
				}(),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 7
		{
			Name:             "scenario 7: get existing cluster health status with MLA Logging enabled and alertmanager config",
			Body:             ``,
			ExpectedResponse: `{"apiserver":"HealthStatusUp","applicationController":"HealthStatusUp","scheduler":"HealthStatusDown","controller":"HealthStatusUp","machineController":"HealthStatusDown","etcd":"HealthStatusUp","cloudProviderInfrastructure":"HealthStatusUp","userClusterControllerManager":"HealthStatusUp","logging":"HealthStatusUp","alertmanagerConfig":"HealthStatusUp"}`,
			HTTPStatus:       http.StatusOK,
			ClusterToGet:     "keen-snyder",
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add a cluster
				test.GenCluster("clusterDefID", "clusterDef", test.GenDefaultProject().Name, time.Date(2013, 02, 04, 01, 54, 0, 0, time.UTC)),
				// add another cluster
				func() *kubermaticv1.Cluster {
					cluster := test.GenCluster("keen-snyder", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					cluster.Status.ExtendedHealth = kubermaticv1.ExtendedClusterHealth{
						Apiserver:                    kubermaticv1.HealthStatusUp,
						ApplicationController:        kubermaticv1.HealthStatusUp,
						Scheduler:                    kubermaticv1.HealthStatusDown,
						Controller:                   kubermaticv1.HealthStatusUp,
						MachineController:            kubermaticv1.HealthStatusDown,
						Etcd:                         kubermaticv1.HealthStatusUp,
						CloudProviderInfrastructure:  kubermaticv1.HealthStatusUp,
						UserClusterControllerManager: kubermaticv1.HealthStatusUp,
						Logging:                      &healthUp,
						AlertmanagerConfig:           &healthUp,
					}
					cluster.Spec.MLA = &kubermaticv1.MLASettings{LoggingEnabled: true}
					return cluster
				}(),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v2/projects/%s/clusters/%s/health", tc.ProjectToSync, tc.ClusterToGet), strings.NewReader(tc.Body))
			res := httptest.NewRecorder()
			var kubermaticObj []ctrlruntimeclient.Object
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
			ep, err := test.CreateTestEndpoint(*tc.ExistingAPIUser, []ctrlruntimeclient.Object{}, kubermaticObj, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			ep.ServeHTTP(res, req)

			if res.Code != tc.HTTPStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.HTTPStatus, res.Code, res.Body.String())
			}

			test.CompareWithResult(t, res, tc.ExpectedResponse)
		})
	}
}

func TestGetClusterMetrics(t *testing.T) {
	t.Parallel()
	cpuQuantity, err := resource.ParseQuantity("290")
	if err != nil {
		t.Fatal(err)
	}
	memoryQuantity, err := resource.ParseQuantity("687202304")
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		Name                   string
		ExpectedResponse       string
		HTTPStatus             int
		ClusterToGet           string
		ExistingAPIUser        *apiv1.User
		ExistingKubermaticObjs []ctrlruntimeclient.Object
		ExistingNodes          []*corev1.Node
		ExistingPodMetrics     []*v1beta1.PodMetrics
		ExistingNodeMetrics    []*v1beta1.NodeMetrics
	}{
		// scenario 1
		{
			Name:             "scenario 1: gets cluster metrics",
			ExpectedResponse: `{"name":"defClusterID","controlPlane":{"memoryTotalBytes":1310,"cpuTotalMillicores":580000},"nodes":{"memoryTotalBytes":1310,"memoryAvailableBytes":1310,"memoryUsedPercentage":100,"cpuTotalMillicores":580000,"cpuAvailableMillicores":580000,"cpuUsedPercentage":100}}`,
			ClusterToGet:     test.GenDefaultCluster().Name,
			HTTPStatus:       http.StatusOK,
			ExistingNodes: []*corev1.Node{
				{ObjectMeta: metav1.ObjectMeta{Name: "venus"}, Status: corev1.NodeStatus{Allocatable: map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity}}},
				{ObjectMeta: metav1.ObjectMeta{Name: "mars"}, Status: corev1.NodeStatus{Allocatable: map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity}}},
			},
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
			ExistingPodMetrics: []*v1beta1.PodMetrics{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "pod1", Namespace: "cluster-defClusterID"},
					Containers: []v1beta1.ContainerMetrics{
						{
							Name:  "c1-pod1",
							Usage: map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity},
						},
						{
							Name:  "c2-pod1",
							Usage: map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity},
						},
					},
				},
			},
			ExistingNodeMetrics: []*v1beta1.NodeMetrics{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "venus"},
					Usage:      map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "mars"},
					Usage:      map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity},
				},
			},
		},
		// scenario 2
		{
			Name:             "scenario 2: the admin John can get any cluster metrics",
			ExpectedResponse: `{"name":"defClusterID","controlPlane":{"memoryTotalBytes":1310,"cpuTotalMillicores":580000},"nodes":{"memoryTotalBytes":1310,"memoryAvailableBytes":1310,"memoryUsedPercentage":100,"cpuTotalMillicores":580000,"cpuAvailableMillicores":580000,"cpuUsedPercentage":100}}`,
			ClusterToGet:     test.GenDefaultCluster().Name,
			HTTPStatus:       http.StatusOK,
			ExistingNodes: []*corev1.Node{
				{ObjectMeta: metav1.ObjectMeta{Name: "venus"}, Status: corev1.NodeStatus{Allocatable: map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity}}},
				{ObjectMeta: metav1.ObjectMeta{Name: "mars"}, Status: corev1.NodeStatus{Allocatable: map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity}}},
			},
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				genUser("John", "john@acme.com", true),
				test.GenDefaultCluster(),
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
			),
			ExistingAPIUser: test.GenAPIUser("John", "john@acme.com"),
			ExistingPodMetrics: []*v1beta1.PodMetrics{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "pod1", Namespace: "cluster-defClusterID"},
					Containers: []v1beta1.ContainerMetrics{
						{
							Name:  "c1-pod1",
							Usage: map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity},
						},
						{
							Name:  "c2-pod1",
							Usage: map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity},
						},
					},
				},
			},
			ExistingNodeMetrics: []*v1beta1.NodeMetrics{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "venus"},
					Usage:      map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "mars"},
					Usage:      map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity},
				},
			},
		},
		// scenario 3
		{
			Name:             "scenario 3: the user John can not get Bob's cluster metrics",
			ExpectedResponse: `{"error":{"code":403,"message":"forbidden: \"john@acme.com\" doesn't belong to project my-first-project-ID"}}`,
			ClusterToGet:     test.GenDefaultCluster().Name,
			HTTPStatus:       http.StatusForbidden,
			ExistingNodes: []*corev1.Node{
				{ObjectMeta: metav1.ObjectMeta{Name: "venus"}, Status: corev1.NodeStatus{Allocatable: map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity}}},
				{ObjectMeta: metav1.ObjectMeta{Name: "mars"}, Status: corev1.NodeStatus{Allocatable: map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity}}},
			},
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				genUser("John", "john@acme.com", false),
				test.GenDefaultCluster(),
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
			),
			ExistingAPIUser: test.GenAPIUser("John", "john@acme.com"),
			ExistingPodMetrics: []*v1beta1.PodMetrics{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "pod1", Namespace: "cluster-defClusterID"},
					Containers: []v1beta1.ContainerMetrics{
						{
							Name:  "c1-pod1",
							Usage: map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity},
						},
						{
							Name:  "c2-pod1",
							Usage: map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity},
						},
					},
				},
			},
			ExistingNodeMetrics: []*v1beta1.NodeMetrics{
				{
					ObjectMeta: metav1.ObjectMeta{Name: "venus"},
					Usage:      map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity},
				},
				{
					ObjectMeta: metav1.ObjectMeta{Name: "mars"},
					Usage:      map[corev1.ResourceName]resource.Quantity{"cpu": cpuQuantity, "memory": memoryQuantity},
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			var kubernetesObj []ctrlruntimeclient.Object
			var kubeObj []ctrlruntimeclient.Object
			var kubermaticObj []ctrlruntimeclient.Object
			for _, existingMetric := range tc.ExistingPodMetrics {
				kubernetesObj = append(kubernetesObj, existingMetric)
			}
			for _, existingMetric := range tc.ExistingNodeMetrics {
				kubernetesObj = append(kubernetesObj, existingMetric)
			}
			for _, node := range tc.ExistingNodes {
				kubeObj = append(kubeObj, node)
			}
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v2/projects/%s/clusters/%s/metrics", test.ProjectName, tc.ClusterToGet), strings.NewReader(""))
			res := httptest.NewRecorder()

			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
			ep, _, err := test.CreateTestEndpointAndGetClients(*tc.ExistingAPIUser, nil, kubeObj, kubernetesObj, kubermaticObj, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			ep.ServeHTTP(res, req)

			if res.Code != tc.HTTPStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.HTTPStatus, res.Code, res.Body.String())
			}

			test.CompareWithResult(t, res, tc.ExpectedResponse)
		})
	}
}

func TestListNamespace(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name                   string
		expectedResponse       string
		httpStatus             int
		clusterToGet           string
		existingAPIUser        *apiv1.User
		existingKubermaticObjs []ctrlruntimeclient.Object
		existingKubernetesObjs []ctrlruntimeclient.Object
	}{
		// scenario 1
		{
			name:             "scenario 1: get cluster namespaces",
			expectedResponse: `[{"name":"default"},{"name":"kube-admin"}]`,
			clusterToGet:     test.GenDefaultCluster().Name,
			httpStatus:       http.StatusOK,
			existingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
			),
			existingKubernetesObjs: []ctrlruntimeclient.Object{
				&corev1.Namespace{
					ObjectMeta: metav1.ObjectMeta{Name: "default"},
				},
				&corev1.Namespace{
					ObjectMeta: metav1.ObjectMeta{Name: "kube-admin"},
				},
			},
			existingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 2
		{
			name:             "scenario 2: the admin John can get Bob's cluster namespaces",
			expectedResponse: `[{"name":"default"},{"name":"kube-admin"}]`,
			clusterToGet:     test.GenDefaultCluster().Name,
			httpStatus:       http.StatusOK,
			existingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				genUser("John", "john@acme.com", true),
			),
			existingKubernetesObjs: []ctrlruntimeclient.Object{
				&corev1.Namespace{
					ObjectMeta: metav1.ObjectMeta{Name: "default"},
				},
				&corev1.Namespace{
					ObjectMeta: metav1.ObjectMeta{Name: "kube-admin"},
				},
			},
			existingAPIUser: test.GenAPIUser("John", "john@acme.com"),
		},
		// scenario 3
		{
			name:             "scenario 3: the user John can not get Bob's cluster namespaces",
			expectedResponse: `{"error":{"code":403,"message":"forbidden: \"john@acme.com\" doesn't belong to project my-first-project-ID"}}`,
			clusterToGet:     test.GenDefaultCluster().Name,
			httpStatus:       http.StatusForbidden,
			existingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
				genUser("John", "john@acme.com", false),
			),
			existingKubernetesObjs: []ctrlruntimeclient.Object{
				&corev1.Namespace{
					ObjectMeta: metav1.ObjectMeta{Name: "default"},
				},
				&corev1.Namespace{
					ObjectMeta: metav1.ObjectMeta{Name: "kube-admin"},
				},
			},
			existingAPIUser: test.GenAPIUser("John", "john@acme.com"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var kubernetesObj []ctrlruntimeclient.Object
			var kubeObj []ctrlruntimeclient.Object
			kubeObj = append(kubeObj, tc.existingKubernetesObjs...)
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v2/projects/%s/clusters/%s/namespaces", test.ProjectName, tc.clusterToGet), strings.NewReader(""))
			res := httptest.NewRecorder()
			var kubermaticObj []ctrlruntimeclient.Object
			kubermaticObj = append(kubermaticObj, tc.existingKubermaticObjs...)
			ep, _, err := test.CreateTestEndpointAndGetClients(*tc.existingAPIUser, nil, kubeObj, kubernetesObj, kubermaticObj, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			ep.ServeHTTP(res, req)

			if res.Code != tc.httpStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.httpStatus, res.Code, res.Body.String())
			}

			test.CompareWithResult(t, res, tc.expectedResponse)
		})
	}
}

func TestDetachSSHKeyFromClusterEndpoint(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		Name                             string
		Body                             string
		KeyToDelete                      string
		ProjectToSync                    string
		ClusterToSync                    string
		ExpectedDeleteResponse           string
		ExpectedDeleteHTTPStatus         int
		ExistingAPIUser                  *apiv1.User
		ExistingSSHKeys                  []*kubermaticv1.UserSSHKey
		ExistingKubermaticObjs           []ctrlruntimeclient.Object
		ExpectedResponseOnGetAfterDelete string
		ExpectedGetHTTPStatus            int
	}{
		// scenario 1
		{
			Name:                             "scenario 1: detaches one key from the cluster",
			Body:                             ``,
			KeyToDelete:                      "key-c08aa5c7abf34504f18552846485267d-yafn",
			ExpectedDeleteResponse:           `{}`,
			ExpectedDeleteHTTPStatus:         http.StatusOK,
			ExpectedGetHTTPStatus:            http.StatusOK,
			ExpectedResponseOnGetAfterDelete: `[{"id":"key-abc-yafn","name":"key-display-name","creationTimestamp":"0001-01-01T00:00:00Z","spec":{"fingerprint":"","publicKey":""}}]`,
			ProjectToSync:                    test.GenDefaultProject().Name,
			ExistingAPIUser:                  test.GenDefaultAPIUser(),
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add a cluster
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
				// add ssh keys
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-c08aa5c7abf34504f18552846485267d-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{"clusterAbcID"},
					},
				},
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-abc-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Name:     "key-display-name",
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{"clusterAbcID"},
					},
				},
			),
			ClusterToSync: "clusterAbcID",
		},
		// scenario 2
		{
			Name:                             "scenario 2: the admin John detaches one key from the Bob cluster",
			Body:                             ``,
			KeyToDelete:                      "key-c08aa5c7abf34504f18552846485267d-yafn",
			ExpectedDeleteResponse:           `{}`,
			ExpectedDeleteHTTPStatus:         http.StatusOK,
			ExpectedGetHTTPStatus:            http.StatusOK,
			ExpectedResponseOnGetAfterDelete: `[{"id":"key-abc-yafn","name":"key-display-name","creationTimestamp":"0001-01-01T00:00:00Z","spec":{"fingerprint":"","publicKey":""}}]`,
			ProjectToSync:                    test.GenDefaultProject().Name,
			ExistingAPIUser:                  test.GenAPIUser("John", "john@acme.com"),
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				genUser("John", "john@acme.com", true),
				// add a cluster
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
				// add ssh keys
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-c08aa5c7abf34504f18552846485267d-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{"clusterAbcID"},
					},
				},
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-abc-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Name:     "key-display-name",
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{"clusterAbcID"},
					},
				},
			),
			ClusterToSync: "clusterAbcID",
		},
		// scenario 3
		{
			Name:                             "scenario 3: the user John can not detach any key from the Bob cluster",
			Body:                             ``,
			KeyToDelete:                      "key-c08aa5c7abf34504f18552846485267d-yafn",
			ExpectedDeleteResponse:           `{"error":{"code":403,"message":"forbidden: \"john@acme.com\" doesn't belong to project my-first-project-ID"}}`,
			ExpectedDeleteHTTPStatus:         http.StatusForbidden,
			ExpectedGetHTTPStatus:            http.StatusForbidden,
			ExpectedResponseOnGetAfterDelete: `{"error":{"code":403,"message":"forbidden: \"john@acme.com\" doesn't belong to project my-first-project-ID"}}`,
			ProjectToSync:                    test.GenDefaultProject().Name,
			ExistingAPIUser:                  test.GenAPIUser("John", "john@acme.com"),
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				genUser("John", "john@acme.com", false),
				// add a cluster
				test.GenCluster("clusterAbcID", "clusterAbc", test.GenDefaultProject().Name, time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC)),
				// add ssh keys
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-c08aa5c7abf34504f18552846485267d-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{"clusterAbcID"},
					},
				},
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-abc-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Name:     "key-display-name",
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{"clusterAbcID"},
					},
				},
			),
			ClusterToSync: "clusterAbcID",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			var ep http.Handler
			{
				var err error
				req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v2/projects/%s/clusters/%s/sshkeys/%s", tc.ProjectToSync, tc.ClusterToSync, tc.KeyToDelete), strings.NewReader(tc.Body))
				res := httptest.NewRecorder()
				var kubermaticObj []ctrlruntimeclient.Object
				kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
				ep, err = test.CreateTestEndpoint(*tc.ExistingAPIUser, []ctrlruntimeclient.Object{}, kubermaticObj, nil, hack.NewTestRouting)
				if err != nil {
					t.Fatalf("failed to create test endpoint: %v", err)
				}

				ep.ServeHTTP(res, req)

				if res.Code != tc.ExpectedDeleteHTTPStatus {
					t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.ExpectedDeleteHTTPStatus, res.Code, res.Body.String())
				}
				test.CompareWithResult(t, res, tc.ExpectedDeleteResponse)
			}
		})
	}
}

func TestAssignSSHKeyToClusterEndpoint(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		Name                   string
		SSHKeyID               string
		ExpectedResponse       string
		HTTPStatus             int
		ProjectToSync          string
		ClusterToSync          string
		ExistingAPIUser        *apiv1.User
		ExistingKubermaticObjs []ctrlruntimeclient.Object
		ExpectedSSHKeys        []*kubermaticv1.UserSSHKey
	}{
		// scenario 1
		{
			Name:             "scenario 1: an ssh key that belongs to the given project is assigned to the cluster",
			SSHKeyID:         "key-c08aa5c7abf34504f18552846485267d-yafn",
			ExpectedResponse: `{"id":"key-c08aa5c7abf34504f18552846485267d-yafn","name":"","creationTimestamp":"0001-01-01T00:00:00Z","spec":{"fingerprint":"","publicKey":""}}`,
			HTTPStatus:       http.StatusCreated,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingAPIUser:  test.GenDefaultAPIUser(),
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add a cluster
				test.GenDefaultCluster(),
				// add a ssh key
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-c08aa5c7abf34504f18552846485267d-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{test.GenDefaultCluster().Name},
					},
				},
			),
			ClusterToSync: test.GenDefaultCluster().Name,
		},
		// scenario 2
		{
			Name:             "scenario 2: an ssh key that does not belong to the given project cannot be assigned to the cluster",
			SSHKeyID:         "key-c08aa5c7abf34504f18552846485267d-yafn",
			ExpectedResponse: `{"error":{"code":500,"message":"the given ssh key key-c08aa5c7abf34504f18552846485267d-yafn does not belong to the given project my-first-project (my-first-project-ID)"}}`,
			HTTPStatus:       http.StatusInternalServerError,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingAPIUser:  test.GenDefaultAPIUser(),
			ExpectedSSHKeys:  []*kubermaticv1.UserSSHKey{},
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add a cluster
				test.GenDefaultCluster(),
				// add an ssh key
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-c08aa5c7abf34504f18552846485267d-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Project: "differentProject",
					},
				},
			),
			ClusterToSync: test.GenDefaultCluster().Name,
		},
		// scenario 3
		{
			Name:             "scenario 3: the admin John can assign ssh key to the Bob's cluster",
			SSHKeyID:         "key-c08aa5c7abf34504f18552846485267d-yafn",
			ExpectedResponse: `{"id":"key-c08aa5c7abf34504f18552846485267d-yafn","name":"","creationTimestamp":"0001-01-01T00:00:00Z","spec":{"fingerprint":"","publicKey":""}}`,
			HTTPStatus:       http.StatusCreated,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingAPIUser:  test.GenAPIUser("John", "john@acme.com"),
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				genUser("John", "john@acme.com", true),
				// add a cluster
				test.GenDefaultCluster(),
				// add a ssh key
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-c08aa5c7abf34504f18552846485267d-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{test.GenDefaultCluster().Name},
					},
				},
			),
			ClusterToSync: test.GenDefaultCluster().Name,
		},
		// scenario 4
		{
			Name:             "scenario 4: the user John can not assign ssh key to the Bob's cluster",
			SSHKeyID:         "key-c08aa5c7abf34504f18552846485267d-yafn",
			ExpectedResponse: `{"error":{"code":403,"message":"forbidden: \"john@acme.com\" doesn't belong to project my-first-project-ID"}}`,
			HTTPStatus:       http.StatusForbidden,
			ProjectToSync:    test.GenDefaultProject().Name,
			ExistingAPIUser:  test.GenAPIUser("John", "john@acme.com"),
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				genUser("John", "john@acme.com", false),
				// add a cluster
				test.GenDefaultCluster(),
				// add a ssh key
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name: "key-c08aa5c7abf34504f18552846485267d-yafn",
					},
					Spec: kubermaticv1.SSHKeySpec{
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{test.GenDefaultCluster().Name},
					},
				},
			),
			ClusterToSync: test.GenDefaultCluster().Name,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v2/projects/%s/clusters/%s/sshkeys/%s", tc.ProjectToSync, tc.ClusterToSync, tc.SSHKeyID), nil)
			res := httptest.NewRecorder()
			var kubermaticObj []ctrlruntimeclient.Object
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
			ep, err := test.CreateTestEndpoint(*tc.ExistingAPIUser, []ctrlruntimeclient.Object{}, kubermaticObj, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			ep.ServeHTTP(res, req)

			if res.Code != tc.HTTPStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.HTTPStatus, res.Code, res.Body.String())
			}

			test.CompareWithResult(t, res, tc.ExpectedResponse)
		})
	}
}

func TestListSSHKeysAssignedToClusterEndpoint(t *testing.T) {
	t.Parallel()
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	creationTime, err := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		Name                   string
		Body                   string
		ProjectToSync          string
		ClusterToSync          string
		ExpectedKeys           []apiv1.SSHKey
		HTTPStatus             int
		ExistingProject        *kubermaticv1.Project
		ExistingKubermaticUser *kubermaticv1.User
		ExistingAPIUser        *apiv1.User
		ExistingCluster        *kubermaticv1.Cluster
		ExistingSSHKeys        []*kubermaticv1.UserSSHKey
		ExistingKubermaticObjs []ctrlruntimeclient.Object
	}{
		// scenario 1
		{
			Name: "scenario 1: gets a list of ssh keys assigned to cluster",
			Body: ``,
			ExpectedKeys: []apiv1.SSHKey{
				{
					ObjectMeta: apiv1.ObjectMeta{
						ID:                "key-c08aa5c7abf34504f18552846485267d-yafn",
						Name:              "yafn",
						CreationTimestamp: apiv1.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC),
					},
				},
				{
					ObjectMeta: apiv1.ObjectMeta{
						ID:                "key-abc-yafn",
						Name:              "abcd",
						CreationTimestamp: apiv1.Date(2013, 02, 03, 19, 55, 0, 0, time.UTC),
					},
				},
			},
			HTTPStatus:    http.StatusOK,
			ProjectToSync: test.GenDefaultProject().Name,
			ExistingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				// add a cluster
				test.GenDefaultCluster(),
				// add ssh keys
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name:              "key-c08aa5c7abf34504f18552846485267d-yafn",
						CreationTimestamp: metav1.NewTime(creationTime),
					},
					Spec: kubermaticv1.SSHKeySpec{
						Name:     "yafn",
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{test.GenDefaultCluster().Name},
					},
				},
				&kubermaticv1.UserSSHKey{
					ObjectMeta: metav1.ObjectMeta{
						Name:              "key-abc-yafn",
						CreationTimestamp: metav1.NewTime(creationTime.Add(time.Minute)),
					},
					Spec: kubermaticv1.SSHKeySpec{
						Name:     "abcd",
						Project:  test.GenDefaultProject().Name,
						Clusters: []string{test.GenDefaultCluster().Name},
					},
				},
			),
			ExistingAPIUser: test.GenDefaultAPIUser(),
			ClusterToSync:   test.GenDefaultCluster().Name,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v2/projects/%s/clusters/%s/sshkeys", tc.ProjectToSync, tc.ClusterToSync), strings.NewReader(tc.Body))
			res := httptest.NewRecorder()
			var kubermaticObj []ctrlruntimeclient.Object
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
			ep, err := test.CreateTestEndpoint(*tc.ExistingAPIUser, []ctrlruntimeclient.Object{}, kubermaticObj, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			ep.ServeHTTP(res, req)

			if res.Code != tc.HTTPStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.HTTPStatus, res.Code, res.Body.String())
			}

			actualKeys := test.NewSSHKeyV1SliceWrapper{}
			actualKeys.DecodeOrDie(res.Body, t).Sort()

			wrappedExpectedKeys := test.NewSSHKeyV1SliceWrapper(tc.ExpectedKeys)
			wrappedExpectedKeys.Sort()

			actualKeys.EqualOrDie(wrappedExpectedKeys, t)
		})
	}
}

func TestRevokeClusterAdminTokenEndpoint(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name                   string
		expectedResponse       string
		httpStatus             int
		clusterToGet           *kubermaticv1.Cluster
		projectToSync          string
		existingAPIUser        *apiv1.User
		existingKubermaticObjs []ctrlruntimeclient.Object
		existingKubernetesObjs []ctrlruntimeclient.Object
	}{
		// scenario 1
		{
			name:             "scenario 1: the owner user revokes admin cluster token",
			expectedResponse: `{}`,
			clusterToGet:     test.GenDefaultCluster(),
			httpStatus:       http.StatusOK,
			existingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenDefaultCluster(),
			),
			existingAPIUser: test.GenDefaultAPIUser(),
		},
		// scenario 2
		{
			name:             "scenario 2: the admin John revokes Bob's cluster token",
			expectedResponse: `{}`,
			clusterToGet:     test.GenDefaultCluster(),
			httpStatus:       http.StatusOK,
			existingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenAdminUser("John", "john@acme.com", true),
				test.GenDefaultCluster(),
			),
			existingAPIUser: test.GenAPIUser("John", "john@acme.com"),
		},
		// scenario 3
		{
			name:             "scenario 3: the user John can not revoke Bob's cluster token",
			expectedResponse: `{"error":{"code":403,"message":"forbidden: \"john@acme.com\" doesn't belong to project my-first-project-ID"}}`,
			clusterToGet:     test.GenDefaultCluster(),
			httpStatus:       http.StatusForbidden,
			existingKubermaticObjs: test.GenDefaultKubermaticObjects(
				test.GenTestSeed(),
				test.GenAdminUser("John", "john@acme.com", false),
				test.GenDefaultCluster(),
			),
			existingAPIUser: test.GenAPIUser("John", "john@acme.com"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ep, clientsSets, err := test.CreateTestEndpointAndGetClients(*tc.existingAPIUser, nil, []ctrlruntimeclient.Object{}, []ctrlruntimeclient.Object{}, tc.existingKubermaticObjs, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			// perform test
			res := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v2/projects/%s/clusters/%s/token", test.ProjectName, tc.clusterToGet.Name), nil)
			ep.ServeHTTP(res, req)

			// check assertions
			test.CheckStatusCode(tc.httpStatus, res, t)
			test.CompareWithResult(t, res, tc.expectedResponse)
			if tc.httpStatus == http.StatusOK {
				updatedCluster := &kubermaticv1.Cluster{}
				if err := clientsSets.FakeClient.Get(context.Background(), types.NamespacedName{Name: test.DefaultClusterID}, updatedCluster); err != nil {
					t.Fatalf("failed to get cluster from fake client: %v", err)
				}
				updatedToken := updatedCluster.Status.Address.AdminToken
				if err := kuberneteshelper.ValidateKubernetesToken(updatedToken); err != nil {
					t.Fatalf("generated token '%s' is malformed: %v", updatedToken, err)
				}
				if updatedToken == tc.clusterToGet.Status.Address.AdminToken {
					t.Fatalf("generated token '%s' is exactly the same as the old one : %s", updatedToken, tc.clusterToGet.Status.Address.AdminToken)
				}
			}
		})
	}
}

func genUser(name, email string, isAdmin bool) *kubermaticv1.User {
	user := test.GenUser("", name, email)
	user.Spec.IsAdmin = isAdmin
	return user
}

func decodeProjectClusterListOrDie(r io.Reader, t *testing.T) test.NewClusterV1SliceWrapper {
	t.Helper()
	clusterListResp := &apiv2.ProjectClusterList{}
	dec := json.NewDecoder(r)
	err := dec.Decode(clusterListResp)
	if err != nil {
		t.Fatal(err)
	}
	return test.NewClusterV1SliceWrapper(clusterListResp.Clusters)
}
