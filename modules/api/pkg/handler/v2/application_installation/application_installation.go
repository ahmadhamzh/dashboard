/*
Copyright 2022 The Kubermatic Kubernetes Platform contributors.

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

package applicationinstallation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	apiv2 "k8c.io/dashboard/v2/pkg/api/v2"
	handlercommon "k8c.io/dashboard/v2/pkg/handler/common"
	"k8c.io/dashboard/v2/pkg/handler/middleware"
	"k8c.io/dashboard/v2/pkg/handler/v1/common"
	"k8c.io/dashboard/v2/pkg/provider"
	appskubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/apps.kubermatic/v1"
	utilerrors "k8c.io/kubermatic/v2/pkg/util/errors"
	"k8c.io/reconciler/pkg/reconciling"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/yaml"
)

func ListApplicationInstallations(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider,
	privilegedProjectProvider provider.PrivilegedProjectProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(listApplicationInstallationsReq)
		if !ok {
			return nil, utilerrors.NewBadRequest("invalid request")
		}

		clusterProvider := ctx.Value(middleware.ClusterProviderContextKey).(provider.ClusterProvider)

		cluster, err := handlercommon.GetCluster(ctx, projectProvider, privilegedProjectProvider, userInfoGetter, req.ProjectID, req.ClusterID, nil)
		if err != nil {
			return nil, err
		}

		client, err := common.GetClusterClient(ctx, userInfoGetter, clusterProvider, cluster, req.ProjectID)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		installList := &appskubermaticv1.ApplicationInstallationList{}
		if err := client.List(ctx, installList); err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		installations := make([]*apiv2.ApplicationInstallationListItem, len(installList.Items))
		for i := range installList.Items {
			installations[i] = convertInternalToAPIApplicationInstallationForList(&installList.Items[i])
		}

		return installations, nil
	}
}

func CreateApplicationInstallation(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider,
	privilegedProjectProvider provider.PrivilegedProjectProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(createApplicationInstallationReq)
		if !ok {
			return nil, utilerrors.NewBadRequest("invalid request")
		}

		clusterProvider := ctx.Value(middleware.ClusterProviderContextKey).(provider.ClusterProvider)

		cluster, err := handlercommon.GetCluster(ctx, projectProvider, privilegedProjectProvider, userInfoGetter, req.ProjectID, req.ClusterID, nil)
		if err != nil {
			return nil, err
		}

		client, err := common.GetClusterClient(ctx, userInfoGetter, clusterProvider, cluster, req.ProjectID)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		// check if namespace for CR already exists and create it if not
		reconcilers := []reconciling.NamedNamespaceReconcilerFactory{
			genericNamespaceReconciler(req.Body.Namespace),
		}
		if err := reconciling.ReconcileNamespaces(ctx, reconcilers, "", client); err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		internalAppInstall := convertAPItoInternalApplicationInstallationBody(&req.Body)

		// if the values field is set, convert all of its values into the valuesBlock field.
		// we do this as values is a deprecated field
		if err := migrateValuesToValuesBlock(&internalAppInstall.Spec); err != nil {
			return nil, utilerrors.New(http.StatusInternalServerError, fmt.Sprintf("failed to convert values into valuesBlock field: %v", err))
		}

		if err := client.Create(ctx, internalAppInstall); err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		return convertInternalToAPIApplicationInstallation(internalAppInstall), nil
	}
}

func DeleteApplicationInstallation(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider,
	privilegedProjectProvider provider.PrivilegedProjectProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(deleteApplicationInstallationReq)
		if !ok {
			return nil, utilerrors.NewBadRequest("invalid request")
		}

		clusterProvider := ctx.Value(middleware.ClusterProviderContextKey).(provider.ClusterProvider)

		cluster, err := handlercommon.GetCluster(ctx, projectProvider, privilegedProjectProvider, userInfoGetter, req.ProjectID, req.ClusterID, nil)
		if err != nil {
			return nil, err
		}

		client, err := common.GetClusterClient(ctx, userInfoGetter, clusterProvider, cluster, req.ProjectID)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		delObj := &appskubermaticv1.ApplicationInstallation{
			TypeMeta: metav1.TypeMeta{
				Kind:       appskubermaticv1.ApplicationInstallationKindName,
				APIVersion: appskubermaticv1.SchemeGroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      req.ApplicationInstallationName,
				Namespace: req.Namespace,
			},
		}
		err = client.Delete(ctx, delObj)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		return nil, nil
	}
}

func GetApplicationInstallation(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider,
	privilegedProjectProvider provider.PrivilegedProjectProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(getApplicationInstallationReq)
		if !ok {
			return nil, utilerrors.NewBadRequest("invalid request")
		}

		clusterProvider := ctx.Value(middleware.ClusterProviderContextKey).(provider.ClusterProvider)

		cluster, err := handlercommon.GetCluster(ctx, projectProvider, privilegedProjectProvider, userInfoGetter, req.ProjectID, req.ClusterID, nil)
		if err != nil {
			return nil, err
		}

		client, err := common.GetClusterClient(ctx, userInfoGetter, clusterProvider, cluster, req.ProjectID)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		applicationInstallation := &appskubermaticv1.ApplicationInstallation{}
		if err := client.Get(ctx, types.NamespacedName{Namespace: req.Namespace, Name: req.ApplicationInstallationName}, applicationInstallation); err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		return convertInternalToAPIApplicationInstallation(applicationInstallation), nil
	}
}

func UpdateApplicationInstallation(userInfoGetter provider.UserInfoGetter, projectProvider provider.ProjectProvider,
	privilegedProjectProvider provider.PrivilegedProjectProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(updateApplicationInstallationReq)
		if !ok {
			return nil, utilerrors.NewBadRequest("invalid request")
		}

		clusterProvider := ctx.Value(middleware.ClusterProviderContextKey).(provider.ClusterProvider)

		cluster, err := handlercommon.GetCluster(ctx, projectProvider, privilegedProjectProvider, userInfoGetter, req.ProjectID, req.ClusterID, nil)
		if err != nil {
			return nil, err
		}

		client, err := common.GetClusterClient(ctx, userInfoGetter, clusterProvider, cluster, req.ProjectID)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		// first fetch the currentAppInstall to make sure it exists
		currentAppInstall := &appskubermaticv1.ApplicationInstallation{}
		if err := client.Get(ctx, types.NamespacedName{Namespace: req.Namespace, Name: req.ApplicationInstallationName}, currentAppInstall); err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		newAppInstall := convertAPItoInternalApplicationInstallationBody(&req.Body)
		currentAppInstall.Spec = newAppInstall.Spec

		// if the values field is set, convert all of its values into the valuesBlock field.
		// we do this as values is a deprecated field
		if err := migrateValuesToValuesBlock(&currentAppInstall.Spec); err != nil {
			return nil, utilerrors.New(http.StatusInternalServerError, fmt.Sprintf("failed to convert values into valuesBlock field: %v", err))
		}

		err = client.Update(ctx, currentAppInstall)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		return convertInternalToAPIApplicationInstallation(currentAppInstall), nil
	}
}

func genericNamespaceReconciler(name string) reconciling.NamedNamespaceReconcilerFactory {
	return func() (string, reconciling.NamespaceReconciler) {
		return name, func(n *corev1.Namespace) (*corev1.Namespace, error) {
			return n, nil
		}
	}
}

func migrateValuesToValuesBlock(ais *appskubermaticv1.ApplicationInstallationSpec) error {
	if len(ais.Values.Raw) > 0 && string(ais.Values.Raw) != "{}" {
		y, err := yaml.JSONToYAML(ais.Values.Raw)
		if err != nil {
			return err
		}
		ais.ValuesBlock = string(y)
		ais.Values = runtime.RawExtension{}
	}
	return nil
}
