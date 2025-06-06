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

package kubernetes

import (
	"context"

	"k8c.io/dashboard/v2/pkg/provider"
	kubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/resources"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// PrivilegedMLAAdminSettingProvider struct that holds required components in order to manage MLAAdminSetting objects.
type PrivilegedMLAAdminSettingProvider struct {
	// privilegedClient is used for admins to interact with MLAAdminSetting objects.
	privilegedClient ctrlruntimeclient.Client
}

var _ provider.PrivilegedMLAAdminSettingProvider = &PrivilegedMLAAdminSettingProvider{}

// NewPrivilegedMLAAdminSettingProvider returns a MLAAdminSetting provider.
func NewPrivilegedMLAAdminSettingProvider(privilegedClient ctrlruntimeclient.Client) *PrivilegedMLAAdminSettingProvider {
	return &PrivilegedMLAAdminSettingProvider{
		privilegedClient: privilegedClient,
	}
}

func (p *PrivilegedMLAAdminSettingProvider) GetUnsecured(ctx context.Context, cluster *kubermaticv1.Cluster) (*kubermaticv1.MLAAdminSetting, error) {
	mlaAdminSetting := &kubermaticv1.MLAAdminSetting{}
	if err := p.privilegedClient.Get(ctx, types.NamespacedName{
		Name:      resources.MLAAdminSettingsName,
		Namespace: cluster.Status.NamespaceName,
	}, mlaAdminSetting); err != nil {
		return nil, err
	}
	return mlaAdminSetting, nil
}

func (p *PrivilegedMLAAdminSettingProvider) CreateUnsecured(ctx context.Context, mlaAdminSetting *kubermaticv1.MLAAdminSetting) (*kubermaticv1.MLAAdminSetting, error) {
	err := p.privilegedClient.Create(ctx, mlaAdminSetting)
	return mlaAdminSetting, err
}

func (p *PrivilegedMLAAdminSettingProvider) UpdateUnsecured(ctx context.Context, newMLAAdminSetting *kubermaticv1.MLAAdminSetting) (*kubermaticv1.MLAAdminSetting, error) {
	err := p.privilegedClient.Update(ctx, newMLAAdminSetting)
	return newMLAAdminSetting, err
}

func (p *PrivilegedMLAAdminSettingProvider) DeleteUnsecured(ctx context.Context, cluster *kubermaticv1.Cluster) error {
	return p.privilegedClient.Delete(ctx, &kubermaticv1.MLAAdminSetting{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.MLAAdminSettingsName,
			Namespace: cluster.Status.NamespaceName,
		},
	})
}

func PrivilegedMLAAdminSettingProviderFactory(mapper meta.RESTMapper, seedKubeconfigGetter provider.SeedKubeconfigGetter) provider.PrivilegedMLAAdminSettingProviderGetter {
	return func(seed *kubermaticv1.Seed) (provider.PrivilegedMLAAdminSettingProvider, error) {
		cfg, err := seedKubeconfigGetter(seed)
		if err != nil {
			return nil, err
		}
		privilegedClient, err := ctrlruntimeclient.New(cfg, ctrlruntimeclient.Options{Mapper: mapper})
		if err != nil {
			return nil, err
		}
		return NewPrivilegedMLAAdminSettingProvider(
			privilegedClient,
		), nil
	}
}
