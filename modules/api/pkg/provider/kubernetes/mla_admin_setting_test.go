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

package kubernetes_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"k8c.io/dashboard/v2/pkg/handler/test"
	"k8c.io/dashboard/v2/pkg/provider"
	"k8c.io/dashboard/v2/pkg/provider/kubernetes"
	kubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/resources"
	"k8c.io/kubermatic/v2/pkg/test/fake"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	testMLAAdminSettingName        = resources.MLAAdminSettingsName
	testMLAAdminSettingClusterName = "test-mla-admin-setting"
)

func TestGetMLAAdminSetting(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name                    string
		existingObjects         []ctrlruntimeclient.Object
		cluster                 *kubermaticv1.Cluster
		expectedMLAAdminSetting *kubermaticv1.MLAAdminSetting
		expectedError           string
	}{
		{
			name: "get mlaAdminSetting",
			existingObjects: []ctrlruntimeclient.Object{
				test.GenMLAAdminSetting(testMLAAdminSettingName, testMLAAdminSettingClusterName, 1),
			},
			cluster:                 genCluster(testMLAAdminSettingClusterName, "kubernetes", "my-first-project-ID", "test-mla-admin-setting", "john@acme.com"),
			expectedMLAAdminSetting: test.GenMLAAdminSetting(testMLAAdminSettingName, testMLAAdminSettingClusterName, 1),
		},
		{
			name:          "mlaAdminSetting is not found",
			cluster:       genCluster(testMLAAdminSettingClusterName, "kubernetes", "my-first-project-ID", "test-mla-admin-setting", "john@acme.com"),
			expectedError: "mlaadminsettings.kubermatic.k8c.io \"mla-admin-settings\" not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := fake.NewClientBuilder().
				WithObjects(tc.existingObjects...).
				Build()
			mlaAdminSettingProvider := kubernetes.NewPrivilegedMLAAdminSettingProvider(client)
			mlaAdminSetting, err := mlaAdminSettingProvider.GetUnsecured(context.Background(), tc.cluster)
			if len(tc.expectedError) == 0 {
				if err != nil {
					t.Fatal(err)
				}
				tc.expectedMLAAdminSetting.ResourceVersion = mlaAdminSetting.ResourceVersion
				assert.Equal(t, tc.expectedMLAAdminSetting, mlaAdminSetting)
			} else {
				if err == nil {
					t.Fatalf("expected error message")
				}
				assert.Equal(t, tc.expectedError, err.Error())
			}
		})
	}
}

func TestCreateMLAAdminSetting(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name                    string
		existingObjects         []ctrlruntimeclient.Object
		cluster                 *kubermaticv1.Cluster
		expectedMLAAdminSetting *kubermaticv1.MLAAdminSetting
		expectedError           string
	}{
		{
			name:                    "create mlaAdminSetting",
			cluster:                 genCluster(testMLAAdminSettingClusterName, "kubernetes", "my-first-project-ID", "test-mla-admin-setting", "john@acme.com"),
			expectedMLAAdminSetting: test.GenMLAAdminSetting(testMLAAdminSettingName, testMLAAdminSettingClusterName, 1),
		},
		{
			name: "create mlaAdminSetting which already exists",
			existingObjects: []ctrlruntimeclient.Object{
				test.GenMLAAdminSetting(testMLAAdminSettingName, testMLAAdminSettingClusterName, 1),
			},
			cluster:                 genCluster(testMLAAdminSettingClusterName, "kubernetes", "my-first-project-ID", "test-mla-admin-setting", "john@acme.com"),
			expectedMLAAdminSetting: test.GenMLAAdminSetting(testMLAAdminSettingName, testMLAAdminSettingClusterName, 1),
			expectedError:           "mlaadminsettings.kubermatic.k8c.io \"mla-admin-settings\" already exists",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := fake.NewClientBuilder().
				WithObjects(tc.existingObjects...).
				Build()
			mlaAdminSettingProvider := kubernetes.NewPrivilegedMLAAdminSettingProvider(client)
			_, err := mlaAdminSettingProvider.CreateUnsecured(context.Background(), tc.expectedMLAAdminSetting)
			if len(tc.expectedError) == 0 {
				if err != nil {
					t.Fatal(err)
				}
				mlaAdminSetting, err := mlaAdminSettingProvider.GetUnsecured(context.Background(), tc.cluster)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.expectedMLAAdminSetting, mlaAdminSetting)
			} else {
				if err == nil {
					t.Fatalf("expected error message")
				}
				assert.Equal(t, tc.expectedError, err.Error())
			}
		})
	}
}

func TestUpdateMLAAdminSetting(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name                    string
		existingObjects         []ctrlruntimeclient.Object
		cluster                 *kubermaticv1.Cluster
		expectedMLAAdminSetting *kubermaticv1.MLAAdminSetting
		expectedError           string
	}{
		{
			name: "update mlaAdminSetting",
			existingObjects: []ctrlruntimeclient.Object{
				test.GenMLAAdminSetting(testMLAAdminSettingName, testMLAAdminSettingClusterName, 1),
			},
			cluster:                 genCluster(testMLAAdminSettingClusterName, "kubernetes", "my-first-project-ID", "test-mla-admin-setting", "john@acme.com"),
			expectedMLAAdminSetting: test.GenMLAAdminSetting(testMLAAdminSettingName, testMLAAdminSettingClusterName, 2),
		},
		{
			name:                    "update mlaAdminSetting which doesn't exist",
			cluster:                 genCluster(testMLAAdminSettingClusterName, "kubernetes", "my-first-project-ID", "test-mla-admin-setting", "john@acme.com"),
			expectedMLAAdminSetting: test.GenMLAAdminSetting(testMLAAdminSettingName, testMLAAdminSettingClusterName, 1),
			expectedError:           "mlaadminsettings.kubermatic.k8c.io \"mla-admin-settings\" not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := fake.NewClientBuilder().
				WithObjects(tc.existingObjects...).
				Build()
			mlaAdminSettingProvider := kubernetes.NewPrivilegedMLAAdminSettingProvider(client)
			if len(tc.expectedError) == 0 {
				currentMLAAdminSetting, err := mlaAdminSettingProvider.GetUnsecured(context.Background(), tc.cluster)
				if err != nil {
					t.Fatal(err)
				}
				tc.expectedMLAAdminSetting.ResourceVersion = currentMLAAdminSetting.ResourceVersion
				_, err = mlaAdminSettingProvider.UpdateUnsecured(context.Background(), tc.expectedMLAAdminSetting)
				if err != nil {
					t.Fatal(err)
				}
				mlaAdminSetting, err := mlaAdminSettingProvider.GetUnsecured(context.Background(), tc.cluster)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tc.expectedMLAAdminSetting, mlaAdminSetting)
			} else {
				_, err := mlaAdminSettingProvider.UpdateUnsecured(context.Background(), tc.expectedMLAAdminSetting)
				if err == nil {
					t.Fatalf("expected error message")
				}
				assert.Equal(t, tc.expectedError, err.Error())
			}
		})
	}
}

func TestDeleteMLAAdminSetting(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name                string
		existingObjects     []ctrlruntimeclient.Object
		userInfo            *provider.UserInfo
		cluster             *kubermaticv1.Cluster
		mlaAdminSettingName string
		expectedError       string
	}{
		{
			name: "delete mlaAdminSetting",
			existingObjects: []ctrlruntimeclient.Object{
				test.GenMLAAdminSetting(testMLAAdminSettingName, testMLAAdminSettingClusterName, 1),
			},
			userInfo:            &provider.UserInfo{Email: "john@acme.com", Groups: []string{"owners-abcd"}},
			cluster:             genCluster(testMLAAdminSettingClusterName, "kubernetes", "my-first-project-ID", "test-mla-admin-setting", "john@acme.com"),
			mlaAdminSettingName: testMLAAdminSettingName,
		},
		{
			name:                "delete mlaAdminSetting which doesn't exist",
			userInfo:            &provider.UserInfo{Email: "john@acme.com", Groups: []string{"owners-abcd"}},
			cluster:             genCluster(testMLAAdminSettingClusterName, "kubernetes", "my-first-project-ID", "test-mla-admin-setting", "john@acme.com"),
			mlaAdminSettingName: testMLAAdminSettingName,
			expectedError:       "mlaadminsettings.kubermatic.k8c.io \"mla-admin-settings\" not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := fake.NewClientBuilder().
				WithObjects(tc.existingObjects...).
				Build()
			mlaAdminSettingProvider := kubernetes.NewPrivilegedMLAAdminSettingProvider(client)
			err := mlaAdminSettingProvider.DeleteUnsecured(context.Background(), tc.cluster)
			if len(tc.expectedError) == 0 {
				if err != nil {
					t.Fatal(err)
				}
				_, err = mlaAdminSettingProvider.GetUnsecured(context.Background(), tc.cluster)
				assert.True(t, apierrors.IsNotFound(err))
			} else {
				if err == nil {
					t.Fatalf("expected error message")
				}
				assert.Equal(t, tc.expectedError, err.Error())
			}
		})
	}
}
