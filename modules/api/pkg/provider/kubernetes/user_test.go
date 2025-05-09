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

package kubernetes_test

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	apiv1 "k8c.io/dashboard/v2/pkg/api/v1"
	"k8c.io/dashboard/v2/pkg/handler/test"
	"k8c.io/dashboard/v2/pkg/provider/kubernetes"
	kubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/resources"
	"k8c.io/kubermatic/v2/pkg/test/fake"
	"k8c.io/machine-controller/sdk/providerconfig"

	corev1 "k8s.io/api/core/v1"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func TestAddUserTokenToBlacklist(t *testing.T) {
	// test data
	testcases := []struct {
		name           string
		existingUser   *kubermaticv1.User
		existingObjs   []ctrlruntimeclient.Object
		token          string
		expiry         apiv1.Time
		expectedTokens []string
	}{
		{
			name: "scenario 1: add token to not existing list",
			existingObjs: []ctrlruntimeclient.Object{
				genUser("", "john", "john@acme.com"),
			},
			existingUser:   genUser("", "john", "john@acme.com"),
			token:          TestFakeToken,
			expiry:         apiv1.Date(2222, 02, 03, 19, 55, 0, 0, time.UTC),
			expectedTokens: []string{TestFakeToken},
		},
		{
			name: "scenario 2: add expired token to not existing list",
			existingObjs: []ctrlruntimeclient.Object{
				genUser("", "john", "john@acme.com"),
			},
			existingUser:   genUser("", "john", "john@acme.com"),
			token:          TestFakeToken,
			expiry:         apiv1.Date(1981, 02, 03, 19, 55, 0, 0, time.UTC),
			expectedTokens: []string{},
		},
		{
			name: "scenario 3: add token to existing list",
			existingObjs: []ctrlruntimeclient.Object{
				func() *corev1.Secret {
					user := genUser("", "john", "john@acme.com")
					return test.GenBlacklistTokenSecret(user.GetInvalidTokensReferenceSecretName(), []byte(`[{"token":"fakeTokenId-1","expiry":"2222-06-20T12:04:00Z"},{"token":"fakeTokenId-2","expiry":"2000-06-20T12:04:00Z"}]`))
				}(),
				func() *kubermaticv1.User {
					user := genUser("", "john", "john@acme.com")
					user.Spec.InvalidTokensReference = &providerconfig.GlobalSecretKeySelector{
						ObjectReference: corev1.ObjectReference{
							Name:      user.GetInvalidTokensReferenceSecretName(),
							Namespace: resources.KubermaticNamespace,
						},
					}

					return user
				}(),
			},
			existingUser:   genUser("", "john", "john@acme.com"),
			token:          "fakeTokenId-3",
			expiry:         apiv1.Date(2222, 02, 03, 19, 55, 0, 0, time.UTC),
			expectedTokens: []string{"fakeTokenId-1", "fakeTokenId-3"},
		},
		{
			name: "scenario 4: add expired token to existing list",
			existingObjs: []ctrlruntimeclient.Object{
				func() *corev1.Secret {
					user := genUser("", "john", "john@acme.com")
					return test.GenBlacklistTokenSecret(user.GetInvalidTokensReferenceSecretName(), []byte(`[{"token":"fakeTokenId-1","expiry":"2222-06-20T12:04:00Z"},{"token":"fakeTokenId-2","expiry":"2000-06-20T12:04:00Z"}]`))
				}(),
				func() *kubermaticv1.User {
					user := genUser("", "john", "john@acme.com")
					user.Spec.InvalidTokensReference = &providerconfig.GlobalSecretKeySelector{
						ObjectReference: corev1.ObjectReference{
							Name:      user.GetInvalidTokensReferenceSecretName(),
							Namespace: resources.KubermaticNamespace,
						},
					}

					return user
				}(),
			},
			existingUser:   genUser("", "john", "john@acme.com"),
			token:          "fakeTokenId-3",
			expiry:         apiv1.Date(1961, 02, 03, 19, 55, 0, 0, time.UTC),
			expectedTokens: []string{"fakeTokenId-1"},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			existingObj := []ctrlruntimeclient.Object{}
			existingObj = append(existingObj, tc.existingObjs...)
			fakeClient := fake.
				NewClientBuilder().
				WithObjects(existingObj...).
				Build()

			// fetch user to get the ResourceVersion
			user := &kubermaticv1.User{}
			if err := fakeClient.Get(ctx, ctrlruntimeclient.ObjectKeyFromObject(tc.existingUser), user); err != nil {
				t.Fatal(err)
			}

			// act
			target := kubernetes.NewUserProvider(fakeClient)
			if err := target.InvalidateToken(ctx, user, tc.token, tc.expiry); err != nil {
				t.Fatal(err)
			}
			resultList, err := target.GetInvalidatedTokens(ctx, user)
			if err != nil {
				t.Fatal(err)
			}
			sort.Strings(resultList)
			sort.Strings(tc.expectedTokens)

			assert.Equal(t, resultList, tc.expectedTokens)
		})
	}
}
