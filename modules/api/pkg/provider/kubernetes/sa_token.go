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

package kubernetes

import (
	"context"
	"fmt"
	"strings"

	"k8c.io/dashboard/v2/pkg/provider"
	kubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/resources"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	labelTokenName = "token"
	tokenPrefix    = "sa-token-"
)

// NewServiceAccountProvider returns a service account provider.
func NewServiceAccountTokenProvider(impersonationClient ImpersonationClient, clientPrivileged ctrlruntimeclient.Client) (*ServiceAccountTokenProvider, error) {
	return &ServiceAccountTokenProvider{
		kubernetesImpersonationClient: impersonationClient,
		kubernetesClientPrivileged:    clientPrivileged,
	}, nil
}

// ServiceAccountProvider manages service account resources.
type ServiceAccountTokenProvider struct {
	// kubernetesImpersonationClient is used as a ground for impersonation
	kubernetesImpersonationClient ImpersonationClient

	kubernetesClientPrivileged ctrlruntimeclient.Client
}

var _ provider.ServiceAccountTokenProvider = &ServiceAccountTokenProvider{}
var _ provider.PrivilegedServiceAccountTokenProvider = &ServiceAccountTokenProvider{}

// Create creates a new token for service account.
func (p *ServiceAccountTokenProvider) Create(ctx context.Context, userInfo *provider.UserInfo, sa *kubermaticv1.User, projectID, tokenName, tokenID, token string) (*corev1.Secret, error) {
	if userInfo == nil {
		return nil, apierrors.NewBadRequest("userInfo cannot be nil")
	}
	if sa == nil {
		return nil, apierrors.NewBadRequest("service account cannot be nil")
	}

	secret := genToken(sa, projectID, tokenName, tokenID, token)
	secret.Namespace = resources.KubermaticNamespace
	kubernetesImpersonatedClient, err := createImpersonationClientWrapperFromUserInfo(userInfo, p.kubernetesImpersonationClient)
	if err != nil {
		return nil, apierrors.NewInternalError(err)
	}

	if err := kubernetesImpersonatedClient.Create(ctx, secret); err != nil {
		return nil, err
	}
	secret.Name = removeTokenPrefix(secret.Name)
	return secret, nil
}

// CreateUnsecured creates a new token
//
// Note that this function:
// is unsafe in a sense that it uses privileged account to create the resource.
func (p *ServiceAccountTokenProvider) CreateUnsecured(ctx context.Context, sa *kubermaticv1.User, projectID, tokenName, tokenID, token string) (*corev1.Secret, error) {
	if sa == nil {
		return nil, apierrors.NewBadRequest("service account cannot be nil")
	}

	secret := genToken(sa, projectID, tokenName, tokenID, token)
	secret.Namespace = resources.KubermaticNamespace

	if err := p.kubernetesClientPrivileged.Create(ctx, secret); err != nil {
		return nil, err
	}
	secret.Name = removeTokenPrefix(secret.Name)
	return secret, nil
}

func genToken(sa *kubermaticv1.User, projectID, tokenName, tokenID, token string) *corev1.Secret {
	secret := &corev1.Secret{}
	secret.Name = addTokenPrefix(tokenID)
	secret.OwnerReferences = []metav1.OwnerReference{
		{
			APIVersion: kubermaticv1.SchemeGroupVersion.String(),
			Kind:       kubermaticv1.UserKindName,
			UID:        sa.GetUID(),
			Name:       sa.Name,
		},
	}
	secret.Labels = map[string]string{
		kubermaticv1.ProjectIDLabelKey: projectID,
		"name":                         tokenName,
	}
	secret.Data = make(map[string][]byte)
	secret.Data[labelTokenName] = []byte(token)
	secret.Type = "Opaque"
	return secret
}

// List  gets tokens for the given service account and project.
func (p *ServiceAccountTokenProvider) List(ctx context.Context, userInfo *provider.UserInfo, project *kubermaticv1.Project, sa *kubermaticv1.User, options *provider.ServiceAccountTokenListOptions) ([]*corev1.Secret, error) {
	if userInfo == nil {
		return nil, apierrors.NewBadRequest("userInfo cannot be nil")
	}
	if project == nil {
		return nil, apierrors.NewBadRequest("project cannot be nil")
	}
	if sa == nil {
		return nil, apierrors.NewBadRequest("sa cannot be nil")
	}
	if options == nil {
		options = &provider.ServiceAccountTokenListOptions{}
	}

	allSecrets := &corev1.SecretList{}
	if err := p.kubernetesClientPrivileged.List(ctx, allSecrets, ctrlruntimeclient.MatchingLabels{kubermaticv1.ProjectIDLabelKey: project.Name}); err != nil {
		return nil, err
	}

	resultList := make([]*corev1.Secret, 0)
	for _, secret := range allSecrets.Items {
		if isToken(&secret) {
			for _, owner := range secret.GetOwnerReferences() {
				if owner.APIVersion == kubermaticv1.SchemeGroupVersion.String() && owner.Kind == kubermaticv1.UserKindName &&
					owner.Name == sa.Name && owner.UID == sa.UID {
					resultList = append(resultList, secret.DeepCopy())
				}
			}
		}
	}

	// Note:
	// After we get the list of tokens we try to get at least one item using unprivileged account to see if the token have read access
	if len(resultList) > 0 {
		kubernetesImpersonatedClient, err := createImpersonationClientWrapperFromUserInfo(userInfo, p.kubernetesImpersonationClient)
		if err != nil {
			return nil, err
		}

		tokenToGet := resultList[0]
		if err = kubernetesImpersonatedClient.Get(ctx, ctrlruntimeclient.ObjectKey{Name: tokenToGet.Name, Namespace: resources.KubermaticNamespace}, &corev1.Secret{}); err != nil {
			return nil, err
		}
	}

	for _, token := range resultList {
		token.Name = removeTokenPrefix(token.Name)
	}

	if len(options.TokenID) == 0 {
		return resultList, nil
	}

	filteredList := make([]*corev1.Secret, 0)
	for _, token := range resultList {
		name, ok := token.Labels["name"]
		if ok {
			if name == options.TokenID {
				filteredList = append(filteredList, token)
				break
			}
		}
	}

	return filteredList, nil
}

// ListUnsecured returns all tokens in kubermatic namespace
//
// Note that this function:
// is unsafe in a sense that it uses privileged account to get the resource
// gets resources from the cache.
func (p *ServiceAccountTokenProvider) ListUnsecured(ctx context.Context, options *provider.ServiceAccountTokenListOptions) ([]*corev1.Secret, error) {
	labelSelector := labels.Everything()
	if options != nil {
		if options.LabelSelector != nil {
			labelSelector = options.LabelSelector
		}
	}
	allSecrets := &corev1.SecretList{}
	if err := p.kubernetesClientPrivileged.List(ctx, allSecrets, ctrlruntimeclient.MatchingLabelsSelector{Selector: labelSelector}); err != nil {
		return nil, err
	}
	allTokens := []*corev1.Secret{}
	for _, secret := range allSecrets.Items {
		if isToken(&secret) {
			sCpy := secret.DeepCopy()
			sCpy.Name = removeTokenPrefix(sCpy.Name)
			allTokens = append(allTokens, sCpy)
		}
	}
	if options == nil {
		return allTokens, nil
	}
	if options.TokenID != "" {
		for _, token := range allTokens {
			if token.Name == options.TokenID {
				return []*corev1.Secret{token}, nil
			}
		}
		return nil, apierrors.NewNotFound(corev1.SchemeGroupVersion.WithResource("secret").GroupResource(), options.TokenID)
	}
	if options.ServiceAccountID != "" {
		resultList := make([]*corev1.Secret, 0)
		for _, token := range allTokens {
			for _, owner := range token.GetOwnerReferences() {
				if owner.APIVersion == kubermaticv1.SchemeGroupVersion.String() && owner.Kind == kubermaticv1.UserKindName &&
					owner.Name == options.ServiceAccountID {
					resultList = append(resultList, token.DeepCopy())
				}
			}
		}
		return filterByTokenName(resultList, options.TokenName), nil
	}

	return filterByTokenName(allTokens, options.TokenName), nil
}

func filterByTokenName(allTokens []*corev1.Secret, tokenName string) []*corev1.Secret {
	if tokenName != "" {
		for _, token := range allTokens {
			name, ok := token.Labels["name"]
			if ok {
				if name == tokenName {
					return []*corev1.Secret{token}
				}
			}
		}
		return make([]*corev1.Secret, 0)
	}
	return allTokens
}

func isToken(secret *corev1.Secret) bool {
	if secret == nil {
		return false
	}
	return strings.HasPrefix(secret.Name, "sa-token")
}

// Get method returns token by name.
func (p *ServiceAccountTokenProvider) Get(ctx context.Context, userInfo *provider.UserInfo, name string) (*corev1.Secret, error) {
	if userInfo == nil {
		return nil, apierrors.NewBadRequest("userInfo cannot be nil")
	}
	if len(name) == 0 {
		return nil, apierrors.NewBadRequest("token name cannot be empty")
	}
	name = addTokenPrefix(name)

	kubernetesImpersonatedClient, err := createImpersonationClientWrapperFromUserInfo(userInfo, p.kubernetesImpersonationClient)
	if err != nil {
		return nil, apierrors.NewInternalError(err)
	}

	token := &corev1.Secret{}
	if err := kubernetesImpersonatedClient.Get(ctx, ctrlruntimeclient.ObjectKey{Name: name, Namespace: resources.KubermaticNamespace}, token); err != nil {
		return nil, err
	}
	token.Name = removeTokenPrefix(token.Name)
	return token, nil
}

// GetUnsecured gets the token by name
//
// Note that this function:
// is unsafe in a sense that it uses privileged account to get the resource.
func (p *ServiceAccountTokenProvider) GetUnsecured(ctx context.Context, name string) (*corev1.Secret, error) {
	if len(name) == 0 {
		return nil, apierrors.NewBadRequest("token name cannot be empty")
	}
	name = addTokenPrefix(name)

	token := &corev1.Secret{}
	if err := p.kubernetesClientPrivileged.Get(ctx, ctrlruntimeclient.ObjectKey{Namespace: resources.KubermaticNamespace, Name: name}, token); err != nil {
		return nil, err
	}
	token.Name = removeTokenPrefix(token.Name)
	return token, nil
}

// Update method updates given token.
func (p *ServiceAccountTokenProvider) Update(ctx context.Context, userInfo *provider.UserInfo, secret *corev1.Secret) (*corev1.Secret, error) {
	if userInfo == nil {
		return nil, apierrors.NewBadRequest("userInfo cannot be nil")
	}
	if secret == nil {
		return nil, apierrors.NewBadRequest("secret cannot be empty")
	}
	secretCpy := secret.DeepCopy()
	secretCpy.Name = addTokenPrefix(secretCpy.Name)
	secretCpy.Namespace = resources.KubermaticNamespace

	kubernetesImpersonatedClient, err := createImpersonationClientWrapperFromUserInfo(userInfo, p.kubernetesImpersonationClient)
	if err != nil {
		return nil, apierrors.NewInternalError(err)
	}

	if err := kubernetesImpersonatedClient.Update(ctx, secretCpy); err != nil {
		return nil, err
	}
	secretCpy.Name = removeTokenPrefix(secretCpy.Name)
	return secretCpy, nil
}

// UpdateUnsecured updates the token
//
// Note that this function:
// is unsafe in a sense that it uses privileged account to get the resource.
func (p *ServiceAccountTokenProvider) UpdateUnsecured(ctx context.Context, secret *corev1.Secret) (*corev1.Secret, error) {
	if secret == nil {
		return nil, apierrors.NewBadRequest("secret cannot be empty")
	}
	secretCpy := secret.DeepCopy()
	secretCpy.Name = addTokenPrefix(secretCpy.Name)
	secretCpy.Namespace = resources.KubermaticNamespace

	if err := p.kubernetesClientPrivileged.Update(ctx, secretCpy); err != nil {
		return nil, err
	}

	secretCpy.Name = removeTokenPrefix(secretCpy.Name)
	return secretCpy, nil
}

// Delete method deletes given token.
func (p *ServiceAccountTokenProvider) Delete(ctx context.Context, userInfo *provider.UserInfo, name string) error {
	if userInfo == nil {
		return apierrors.NewBadRequest("userInfo cannot be nil")
	}
	if len(name) == 0 {
		return apierrors.NewBadRequest("token name cannot be empty")
	}
	name = addTokenPrefix(name)

	kubernetesImpersonatedClient, err := createImpersonationClientWrapperFromUserInfo(userInfo, p.kubernetesImpersonationClient)
	if err != nil {
		return apierrors.NewInternalError(err)
	}
	return kubernetesImpersonatedClient.Delete(ctx, &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: name, Namespace: resources.KubermaticNamespace}})
}

// DeleteUnsecured deletes the token
//
// Note that this function:
// is unsafe in a sense that it uses privileged account to delete the resource.
func (p *ServiceAccountTokenProvider) DeleteUnsecured(ctx context.Context, name string) error {
	if len(name) == 0 {
		return apierrors.NewBadRequest("token name cannot be empty")
	}
	name = addTokenPrefix(name)
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: resources.KubermaticNamespace,
		},
	}
	return p.kubernetesClientPrivileged.Delete(ctx, secret)
}

// removeTokenPrefix removes "sa-token-" from a token's ID
// for example given "sa-token-gmtzqz692d" it returns "gmtzqz692d".
func removeTokenPrefix(id string) string {
	return strings.TrimPrefix(id, tokenPrefix)
}

// addTokenPrefix adds "sa-token-" prefix to a token's ID,
// for example given "gmtzqz692d" it returns "sa-token-gmtzqz692d".
func addTokenPrefix(id string) string {
	if !hasTokenPrefix(id) {
		return fmt.Sprintf("%s%s", tokenPrefix, id)
	}
	return id
}

// hasTokenPrefix checks if the given id has "sa-token-" prefix.
func hasTokenPrefix(token string) bool {
	return strings.HasPrefix(token, tokenPrefix)
}
