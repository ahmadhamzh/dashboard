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

package hack

import (
	"net/http"

	"github.com/gorilla/mux"
	prometheusapi "github.com/prometheus/client_golang/api"
	"github.com/prometheus/client_golang/prometheus"

	"k8c.io/dashboard/v2/pkg/handler"
	"k8c.io/dashboard/v2/pkg/handler/v1/common"
	v2 "k8c.io/dashboard/v2/pkg/handler/v2"
	"k8c.io/dashboard/v2/pkg/provider"
	authtypes "k8c.io/dashboard/v2/pkg/provider/auth/types"
	"k8c.io/dashboard/v2/pkg/provider/kubernetes"
	"k8c.io/dashboard/v2/pkg/serviceaccount"
	"k8c.io/dashboard/v2/pkg/watcher"
	kubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/features"
	kubermaticlog "k8c.io/kubermatic/v2/pkg/log"
	"k8c.io/kubermatic/v2/pkg/resources/certificates"
	"k8c.io/kubermatic/v2/pkg/version/kubermatic"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

// NewTestRouting is a hack that helps us avoid circular imports
// for example handler package uses v1/dc and v1/dc needs handler for testing.
func NewTestRouting(
	adminProvider provider.AdminProvider,
	settingsProvider provider.SettingsProvider,
	userInfoGetter provider.UserInfoGetter,
	seedsGetter provider.SeedsGetter,
	seedClientGetter provider.SeedClientGetter,
	configGetter provider.KubermaticConfigurationGetter,
	clusterProvidersGetter provider.ClusterProviderGetter,
	addonProviderGetter provider.AddonProviderGetter,
	addonConfigProvider provider.AddonConfigProvider,
	sshKeyProvider provider.SSHKeyProvider,
	privilegedSSHKeyProvider provider.PrivilegedSSHKeyProvider,
	userProvider provider.UserProvider,
	serviceAccountProvider provider.ServiceAccountProvider,
	privilegedServiceAccountProvider provider.PrivilegedServiceAccountProvider,
	serviceAccountTokenProvider provider.ServiceAccountTokenProvider,
	privilegedServiceAccountTokenProvider provider.PrivilegedServiceAccountTokenProvider,
	projectProvider provider.ProjectProvider,
	privilegedProjectProvider provider.PrivilegedProjectProvider,
	issuerVerifier authtypes.OIDCIssuerVerifier,
	tokenVerifiers authtypes.TokenVerifier,
	tokenExtractors authtypes.TokenExtractor,
	prometheusClient prometheusapi.Client,
	projectMemberProvider *kubernetes.ProjectMemberProvider,
	privilegedProjectMemberProvider provider.PrivilegedProjectMemberProvider,
	saTokenAuthenticator serviceaccount.TokenAuthenticator,
	saTokenGenerator serviceaccount.TokenGenerator,
	eventRecorderProvider provider.EventRecorderProvider,
	presetProvider provider.PresetProvider,
	admissionPluginProvider provider.AdmissionPluginsProvider,
	settingsWatcher watcher.SettingsWatcher,
	userWatcher watcher.UserWatcher,
	externalClusterProvider provider.ExternalClusterProvider,
	privilegedExternalClusterProvider provider.PrivilegedExternalClusterProvider,
	constraintTemplateProvider provider.ConstraintTemplateProvider,
	constraintProviderGetter provider.ConstraintProviderGetter,
	alertmanagerProviderGetter provider.AlertmanagerProviderGetter,
	clusterTemplateProvider provider.ClusterTemplateProvider,
	clusterTemplateInstanceProviderGetter provider.ClusterTemplateInstanceProviderGetter,
	ruleGroupProviderGetter provider.RuleGroupProviderGetter,
	kubermaticVersions kubermatic.Versions,
	defaultConstraintProvider provider.DefaultConstraintProvider,
	privilegedAllowedRegistryProvider provider.PrivilegedAllowedRegistryProvider,
	etcdBackupConfigProviderGetter provider.EtcdBackupConfigProviderGetter,
	etcdRestoreProviderGetter provider.EtcdRestoreProviderGetter,
	etcdBackupConfigProjectProviderGetter provider.EtcdBackupConfigProjectProviderGetter,
	etcdRestoreProjectProviderGetter provider.EtcdRestoreProjectProviderGetter,
	backupCredentialsProviderGetter provider.BackupCredentialsProviderGetter,
	privilegedMLAAdminSettingProviderGetter provider.PrivilegedMLAAdminSettingProviderGetter,
	masterClient ctrlruntimeclient.Client,
	featureGatesProvider provider.FeatureGatesProvider,
	seedProvider provider.SeedProvider,
	resourceQuotaProvider provider.ResourceQuotaProvider,
	groupProjectBindingProvider provider.GroupProjectBindingProvider,
	applicationDefinitionProvider provider.ApplicationDefinitionProvider,
	privilegedIPAMPoolProviderGetter provider.PrivilegedIPAMPoolProviderGetter,
	privilegedOperatingSystemProfileProviderGetter provider.PrivilegedOperatingSystemProfileProviderGetter,
	fakeOIDCVerifierIssuerGetter provider.OIDCIssuerVerifierGetter,
	features features.FeatureGate) http.Handler {
	routingParams := handler.RoutingParams{
		Log:                                            kubermaticlog.Logger,
		PresetProvider:                                 presetProvider,
		SeedsGetter:                                    seedsGetter,
		SeedsClientGetter:                              seedClientGetter,
		KubermaticConfigurationGetter:                  configGetter,
		SSHKeyProvider:                                 sshKeyProvider,
		PrivilegedSSHKeyProvider:                       privilegedSSHKeyProvider,
		UserProvider:                                   userProvider,
		ServiceAccountProvider:                         serviceAccountProvider,
		PrivilegedServiceAccountProvider:               privilegedServiceAccountProvider,
		ServiceAccountTokenProvider:                    serviceAccountTokenProvider,
		PrivilegedServiceAccountTokenProvider:          privilegedServiceAccountTokenProvider,
		ProjectProvider:                                projectProvider,
		PrivilegedProjectProvider:                      privilegedProjectProvider,
		TokenVerifiers:                                 tokenVerifiers,
		TokenExtractors:                                tokenExtractors,
		ClusterProviderGetter:                          clusterProvidersGetter,
		AddonProviderGetter:                            addonProviderGetter,
		AddonConfigProvider:                            addonConfigProvider,
		PrometheusClient:                               prometheusClient,
		ProjectMemberProvider:                          projectMemberProvider,
		PrivilegedProjectMemberProvider:                privilegedProjectMemberProvider,
		UserProjectMapper:                              projectMemberProvider, /*satisfies also a different interface*/
		SATokenAuthenticator:                           saTokenAuthenticator,
		SATokenGenerator:                               saTokenGenerator,
		EventRecorderProvider:                          eventRecorderProvider,
		ExposeStrategy:                                 kubermaticv1.ExposeStrategyNodePort,
		UserInfoGetter:                                 userInfoGetter,
		SettingsProvider:                               settingsProvider,
		AdminProvider:                                  adminProvider,
		AdmissionPluginProvider:                        admissionPluginProvider,
		SettingsWatcher:                                settingsWatcher,
		UserWatcher:                                    userWatcher,
		ExternalClusterProvider:                        externalClusterProvider,
		PrivilegedExternalClusterProvider:              privilegedExternalClusterProvider,
		FeatureGatesProvider:                           featureGatesProvider,
		DefaultConstraintProvider:                      defaultConstraintProvider,
		ConstraintTemplateProvider:                     constraintTemplateProvider,
		ConstraintProviderGetter:                       constraintProviderGetter,
		AlertmanagerProviderGetter:                     alertmanagerProviderGetter,
		ClusterTemplateProvider:                        clusterTemplateProvider,
		ClusterTemplateInstanceProviderGetter:          clusterTemplateInstanceProviderGetter,
		RuleGroupProviderGetter:                        ruleGroupProviderGetter,
		PrivilegedAllowedRegistryProvider:              privilegedAllowedRegistryProvider,
		EtcdBackupConfigProviderGetter:                 etcdBackupConfigProviderGetter,
		EtcdRestoreProviderGetter:                      etcdRestoreProviderGetter,
		EtcdBackupConfigProjectProviderGetter:          etcdBackupConfigProjectProviderGetter,
		EtcdRestoreProjectProviderGetter:               etcdRestoreProjectProviderGetter,
		BackupCredentialsProviderGetter:                backupCredentialsProviderGetter,
		PrivilegedMLAAdminSettingProviderGetter:        privilegedMLAAdminSettingProviderGetter,
		SeedProvider:                                   seedProvider,
		ResourceQuotaProvider:                          resourceQuotaProvider,
		GroupProjectBindingProvider:                    groupProjectBindingProvider,
		ApplicationDefinitionProvider:                  applicationDefinitionProvider,
		Versions:                                       kubermaticVersions,
		CABundle:                                       certificates.NewFakeCABundle().CertPool(),
		Features:                                       features,
		PrivilegedIPAMPoolProviderGetter:               privilegedIPAMPoolProviderGetter,
		PrivilegedOperatingSystemProfileProviderGetter: privilegedOperatingSystemProfileProviderGetter,
		OIDCIssuerVerifierProviderGetter:               fakeOIDCVerifierIssuerGetter,
	}

	r := handler.NewRouting(routingParams, masterClient)
	rv2 := v2.NewV2Routing(routingParams)

	mainRouter := mux.NewRouter()
	v1Router := mainRouter.PathPrefix("/api/v1").Subrouter()
	v2Router := mainRouter.PathPrefix("/api/v2").Subrouter()
	r.RegisterV1(v1Router, generateDefaultMetrics())
	r.RegisterV1Optional(v1Router, true)
	r.RegisterV1Admin(v1Router)
	r.RegisterV1Websocket(v1Router)
	rv2.RegisterV2(v2Router, true)
	return mainRouter
}

func generateDefaultMetrics() common.ServerMetrics {
	return common.ServerMetrics{
		InitNodeDeploymentFailures: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "kubermatic_api_failed_init_node_deployment_total",
				Help: "The number of times initial node deployment couldn't be created within the timeout",
			},
			[]string{"cluster", "datacenter"},
		),
	}
}
