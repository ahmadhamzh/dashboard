//go:build ee

/*
                  Kubermatic Enterprise Read-Only License
                         Version 1.0 ("KERO-1.0”)
                     Copyright © 2024 Kubermatic GmbH

   1.	You may only view, read and display for studying purposes the source
      code of the software licensed under this license, and, to the extent
      explicitly provided under this license, the binary code.
   2.	Any use of the software which exceeds the foregoing right, including,
      without limitation, its execution, compilation, copying, modification
      and distribution, is expressly prohibited.
   3.	THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
      EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
      MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
      IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
      CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
      TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
      SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

   END OF TERMS AND CONDITIONS
*/

package kubelb

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	apiv1 "k8c.io/dashboard/v2/pkg/api/v1"
	"k8c.io/dashboard/v2/pkg/handler/middleware"
	"k8c.io/dashboard/v2/pkg/handler/v1/common"
	"k8c.io/dashboard/v2/pkg/provider"
	kubelbv1alpha1 "k8c.io/kubelb/api/kubelb.k8c.io/v1alpha1"
	v1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	utilerrors "k8c.io/kubermatic/v2/pkg/util/errors"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/clientcmd"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// swagger:parameters listSeedKubeLBTenants
type listSeedKubeLBTenantsReq struct {
	// in: path
	// required: true
	SeedName string `json:"seed_name"`
}

type getKubeLBTenantsReq struct {
	// in: path
	// required: true
	SeedName string `json:"seed_name"`
	// in: path
	// required: true
	DC string `json:"dc"`
	// in: path
	// required: true
	TenantName string `json:"tenant_name"`
}

func (req listSeedKubeLBTenantsReq) GetSeedCluster() apiv1.SeedCluster {
	return apiv1.SeedCluster{
		SeedName: req.SeedName,
	}
}

func (req getKubeLBTenantsReq) GetSeedCluster() apiv1.SeedCluster {
	return apiv1.SeedCluster{
		SeedName: req.SeedName,
	}
}

func ListKubeLBTenants(ctx context.Context, request interface{}, seedsGetter provider.SeedsGetter) (interface{}, error) {
	req, ok := request.(listSeedKubeLBTenantsReq)
	if !ok {
		return nil, utilerrors.NewBadRequest("invalid request")
	}
	privilegedClusterProvider := ctx.Value(middleware.PrivilegedClusterProviderContextKey).(provider.PrivilegedClusterProvider)
	seedClient := privilegedClusterProvider.GetSeedClusterAdminRuntimeClient()
	seeds, err := seedsGetter()
	if err != nil {
		return nil, common.KubernetesErrorToHTTPError(err)
	}
	tenantsMap := make(map[string]kubelbv1alpha1.Tenant)
	seed := seeds[req.SeedName]
	if seed.Spec.KubeLB != nil && seed.Spec.KubeLB.Kubeconfig.Name != "" {
		seedTenants, err := getListOfTenantsForSeed(ctx, seedClient, seed)
		if err != nil {
			return nil, err
		}
		for _, tenant := range seedTenants.Items {
			tenantsMap[tenant.Name] = tenant
		}
	}
	dataCenters := seed.Spec.Datacenters
	for _, dataCenter := range dataCenters {

		dcTenants, err := getListOfTenantsForDatacenter(ctx, seedClient, seed, dataCenter)
		if err != nil {
			continue
		}
		for _, tenant := range dcTenants.Items {
			if _, exists := tenantsMap[tenant.Name]; !exists {
				tenantsMap[tenant.Name] = tenant
			}
		}
	}

	return tenantsMap, nil
}

func DecodeListKubeLBTenantsReq(c context.Context, r *http.Request) (interface{}, error) {
	var req listSeedKubeLBTenantsReq
	Seedname := mux.Vars(r)["seed_name"]
	if Seedname == "" {
		return nil, utilerrors.NewBadRequest("'seed_name' parameter is required but was not provided")
	}
	req.SeedName = Seedname
	return req, nil
}

func GetKubeLBTenants(ctx context.Context, request interface{}, seedsGetter provider.SeedsGetter) (interface{}, error) {
	req, ok := request.(getKubeLBTenantsReq)
	if !ok {
		return nil, utilerrors.NewBadRequest("invalid request")
	}
	privilegedClusterProvider := ctx.Value(middleware.PrivilegedClusterProviderContextKey).(provider.PrivilegedClusterProvider)
	seedClient := privilegedClusterProvider.GetSeedClusterAdminRuntimeClient()
	seeds, err := seedsGetter()
	if err != nil {
		return nil, common.KubernetesErrorToHTTPError(err)
	}
	seed := seeds[req.SeedName]
	dc := seed.Spec.Datacenters[req.DC]
	secret, err := getDataCenterKubeLBKubeconfigSecret(ctx, seedClient, seed, dc)

	if err != nil {
		return nil, fmt.Errorf("failed to get kubeLB secret: %w", err)
	}

	kubeLBManagerKubeconfig := secret.Data["kubeconfig"]
	if len(kubeLBManagerKubeconfig) == 0 {
		return nil, fmt.Errorf("no kubeconfig found")
	}

	kubeLBKubeconfig, err := clientcmd.Load(kubeLBManagerKubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to parse kubeconfig: %w", err)
	}
	cfg, err := clientcmd.NewInteractiveClientConfig(*kubeLBKubeconfig, "", nil, nil, nil).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load kubeconfig: %w", err)
	}
	client, err := ctrlruntimeclient.New(cfg, ctrlruntimeclient.Options{})
	tenant := &kubelbv1alpha1.Tenant{}
	namespacedName := types.NamespacedName{
		Namespace: seed.Namespace,
		Name:      req.TenantName,
	}
	if err := client.Get(ctx, namespacedName, tenant); err != nil {
		return nil, fmt.Errorf("failed to get tenant%s: %w", err)
	}

	return tenant, nil
}

func DecodeGetKubeLBTenantsReq(c context.Context, r *http.Request) (interface{}, error) {
	var req getKubeLBTenantsReq
	Seedname := mux.Vars(r)["seed_name"]
	if Seedname == "" {
		return nil, utilerrors.NewBadRequest("'seed_name' parameter is required but was not provided")
	}

	dc := mux.Vars(r)["dc"]
	if dc == "" {
		return nil, utilerrors.NewBadRequest("'dc' parameter is required but was not provided")
	}

	TenantName := mux.Vars(r)["tenant_name"]
	if TenantName == "" {
		return nil, utilerrors.NewBadRequest("'tenant_name' parameter is required but was not provided")
	}

	req.SeedName = Seedname
	req.DC = dc
	req.TenantName = TenantName
	return req, nil
}

func getListOfTenantsForSeed(ctx context.Context, seedClient ctrlruntimeclient.Client, seed *v1.Seed) (*kubelbv1alpha1.TenantList, error) {
	secret, err := getSeedKubeLBKubeconfigSecret(ctx, seedClient, seed)
	tenantList := &kubelbv1alpha1.TenantList{}

	if err != nil {
		return nil, err
	}

	kubeLBManagerKubeconfig := secret.Data["kubeconfig"]
	if len(kubeLBManagerKubeconfig) == 0 {
		return nil, fmt.Errorf("no kubeconfig found")
	}
	kubeLBKubeconfig, err := clientcmd.Load(kubeLBManagerKubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to parse kubeconfig: %w", err)
	}

	cfg, err := clientcmd.NewInteractiveClientConfig(*kubeLBKubeconfig, "", nil, nil, nil).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load kubeconfig: %w", err)
	}

	client, err := ctrlruntimeclient.New(cfg, ctrlruntimeclient.Options{})

	if err := client.List(ctx, tenantList); err != nil {
		return nil, fmt.Errorf("failed to list tenants: %w", err)
	}
	return tenantList, nil
}

func getListOfTenantsForDatacenter(ctx context.Context, seedClient ctrlruntimeclient.Client, seed *v1.Seed, dc v1.Datacenter) (*kubelbv1alpha1.TenantList, error) {
	secret, err := getDataCenterKubeLBKubeconfigSecret(ctx, seedClient, seed, dc)
	tenantList := &kubelbv1alpha1.TenantList{}

	if err != nil {
		return nil, err
	}
	kubeLBManagerKubeconfig := secret.Data["kubeconfig"]
	if len(kubeLBManagerKubeconfig) == 0 {
		return nil, fmt.Errorf("no kubeconfig found")
	}
	kubeLBKubeconfig, err := clientcmd.Load(kubeLBManagerKubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to parse kubeconfig: %w", err)
	}

	cfg, err := clientcmd.NewInteractiveClientConfig(*kubeLBKubeconfig, "", nil, nil, nil).ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load kubeconfig: %w", err)
	}

	client, err := ctrlruntimeclient.New(cfg, ctrlruntimeclient.Options{})

	if err := client.List(ctx, tenantList); err != nil {
		return nil, fmt.Errorf("failed to list tenants: %w", err)

	}
	return tenantList, nil
}

func getSeedKubeLBKubeconfigSecret(ctx context.Context, client ctrlruntimeclient.Client, seed *v1.Seed) (*corev1.Secret, error) {
	var name, namespace string

	if seed.Spec.KubeLB != nil && seed.Spec.KubeLB.Kubeconfig.Name != "" {
		name = seed.Spec.KubeLB.Kubeconfig.Name
		namespace = seed.Spec.KubeLB.Kubeconfig.Namespace
	} else {
		return nil, fmt.Errorf("kubeLB management kubeconfig not found")
	}

	secret := &corev1.Secret{}
	resourceName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	if resourceName.Namespace == "" {
		resourceName.Namespace = seed.Namespace
	}
	if err := client.Get(ctx, resourceName, secret); err != nil {
		return nil, fmt.Errorf("failed to get kubeconfig secret %q: %w", resourceName.String(), err)
	}

	return secret, nil
}

func getDataCenterKubeLBKubeconfigSecret(ctx context.Context, client ctrlruntimeclient.Client, seed *v1.Seed, dc v1.Datacenter) (*corev1.Secret, error) {
	var name, namespace string

	if dc.Spec.KubeLB != nil && dc.Spec.KubeLB.Kubeconfig.Name != "" {
		name = dc.Spec.KubeLB.Kubeconfig.Name
		namespace = dc.Spec.KubeLB.Kubeconfig.Namespace
	} else {
		return nil, fmt.Errorf("kubeLB management kubeconfig not found")
	}

	secret := &corev1.Secret{}
	resourceName := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	if resourceName.Namespace == "" {
		resourceName.Namespace = seed.Namespace
	}
	if err := client.Get(ctx, resourceName, secret); err != nil {
		return nil, fmt.Errorf("failed to get kubeconfig secret %q: %w", resourceName.String(), err)
	}
	return secret, nil
}
