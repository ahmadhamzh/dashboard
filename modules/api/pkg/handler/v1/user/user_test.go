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

package user_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	apiv1 "k8c.io/dashboard/v2/pkg/api/v1"
	"k8c.io/dashboard/v2/pkg/handler/test"
	"k8c.io/dashboard/v2/pkg/handler/test/hack"
	kubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/resources"
	"k8c.io/machine-controller/sdk/providerconfig"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	utilruntime.Must(kubermaticv1.AddToScheme(scheme.Scheme))
}

func TestGetUsersForProject(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		Name                        string
		ExpectedResponse            []apiv1.User
		ExpectedResponseString      string
		ExpectedActions             int
		ExpectedUserAfterInvitation *kubermaticv1.User
		ProjectToGet                string
		HTTPStatus                  int
		ExistingAPIUser             apiv1.User
		ExistingKubermaticObjs      []ctrlruntimeclient.Object
	}{
		{
			Name:         "scenario 1: get a list of user for a project 'foo'",
			HTTPStatus:   http.StatusOK,
			ProjectToGet: "foo-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("foo", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("bar", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("zorg", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("foo-ID", "john@acme.com", "owners"),
				test.GenBinding("bar-ID", "john@acme.com", "editors"),
				test.GenBinding("foo-ID", "alice@acme.com", "viewers"),
				test.GenBinding("foo-ID", "bob@acme.com", "editors"),
				test.GenBinding("bar-ID", "bob@acme.com", "editors"),
				test.GenBinding("zorg-ID", "bob@acme.com", "editors"),
				/*add users*/
				func() *kubermaticv1.User {
					user := genUser("", "john", "john@acme.com")
					user.CreationTimestamp = metav1.NewTime(time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					return user
				}(),
				func() *kubermaticv1.User {
					user := genUser("", "alice", "alice@acme.com")
					user.CreationTimestamp = metav1.NewTime(time.Date(2013, 02, 03, 19, 55, 0, 0, time.UTC))
					return user
				}(),
				func() *kubermaticv1.User {
					user := genUser("", "bob", "bob@acme.com")
					user.CreationTimestamp = metav1.NewTime(time.Date(2013, 02, 03, 19, 56, 0, 0, time.UTC))
					return user
				}(),
			},
			ExistingAPIUser: *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: []apiv1.User{
				{
					ObjectMeta: apiv1.ObjectMeta{
						ID:                "4b2d8785b49bad23638b17d8db76857a79bf79441241a78a97d88cc64bbf766e",
						Name:              "john",
						CreationTimestamp: apiv1.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC),
					},
					Email:    "john@acme.com",
					LastSeen: &[]apiv1.Time{apiv1.NewTime(test.UserLastSeen)}[0],
					Projects: []apiv1.ProjectGroup{
						{
							GroupPrefix: "owners",
							ID:          "foo-ID",
						},
					},
				},

				{
					ObjectMeta: apiv1.ObjectMeta{
						ID:                "0a0a58273565a8f3dcf779375d9debd0f685d94dc56651a16bff3bf901c0b127",
						Name:              "alice",
						CreationTimestamp: apiv1.Date(2013, 02, 03, 19, 55, 0, 0, time.UTC),
					},
					Email:    "alice@acme.com",
					LastSeen: &[]apiv1.Time{apiv1.NewTime(test.UserLastSeen)}[0],
					Projects: []apiv1.ProjectGroup{
						{
							GroupPrefix: "viewers",
							ID:          "foo-ID",
						},
					},
				},

				{
					ObjectMeta: apiv1.ObjectMeta{
						ID:                "405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6",
						Name:              "bob",
						CreationTimestamp: apiv1.Date(2013, 02, 03, 19, 56, 0, 0, time.UTC),
					},
					Email:    "bob@acme.com",
					LastSeen: &[]apiv1.Time{apiv1.NewTime(test.UserLastSeen)}[0],
					Projects: []apiv1.ProjectGroup{
						{
							GroupPrefix: "editors",
							ID:          "foo-ID",
						},
					},
				},
			},
		},
		{
			Name:         "scenario 2: get a list of user for a project 'foo' for external user",
			HTTPStatus:   http.StatusForbidden,
			ProjectToGet: "foo2-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("foo2", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("bar2", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("bar-ID", "alice2@acme.com", "viewers"),
				test.GenBinding("foo2-ID", "bob@acme.com", "editors"),
				test.GenBinding("bar2-ID", "bob@acme.com", "editors"),
				/*add users*/
				genUser("", "alice2", "alice2@acme.com"),
				genDefaultUser(), /*bob*/
			},
			ExistingAPIUser:        *genAPIUser("alice2", "alice2@acme.com"),
			ExpectedResponseString: `{"error":{"code":403,"message":"forbidden: \"alice2@acme.com\" doesn't belong to project foo2-ID"}}`,
		},
		{
			Name:         "scenario 3: the admin can get a list of user for any project",
			HTTPStatus:   http.StatusOK,
			ProjectToGet: "foo-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("foo", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("bar", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("zorg", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("foo-ID", "john@acme.com", "owners"),
				test.GenBinding("bar-ID", "john@acme.com", "editors"),
				test.GenBinding("foo-ID", "alice@acme.com", "viewers"),
				test.GenBinding("foo-ID", "bob@acme.com", "editors"),
				test.GenBinding("bar-ID", "bob@acme.com", "editors"),
				test.GenBinding("zorg-ID", "bob@acme.com", "editors"),
				/*add users*/
				func() *kubermaticv1.User {
					user := genUser("", "john", "john@acme.com")
					user.CreationTimestamp = metav1.NewTime(time.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC))
					return user
				}(),
				func() *kubermaticv1.User {
					user := genUser("", "alice", "alice@acme.com")
					user.CreationTimestamp = metav1.NewTime(time.Date(2013, 02, 03, 19, 55, 0, 0, time.UTC))
					return user
				}(),
				func() *kubermaticv1.User {
					user := genUser("", "bob", "bob@acme.com")
					user.CreationTimestamp = metav1.NewTime(time.Date(2013, 02, 03, 19, 56, 0, 0, time.UTC))
					return user
				}(),
				genDefaultAdminUser(),
			},
			ExistingAPIUser: *genAPIUser("admin", "admin@acme.com"),
			ExpectedResponse: []apiv1.User{
				{
					ObjectMeta: apiv1.ObjectMeta{
						ID:                "4b2d8785b49bad23638b17d8db76857a79bf79441241a78a97d88cc64bbf766e",
						Name:              "john",
						CreationTimestamp: apiv1.Date(2013, 02, 03, 19, 54, 0, 0, time.UTC),
					},
					Email:    "john@acme.com",
					LastSeen: &[]apiv1.Time{apiv1.NewTime(test.UserLastSeen)}[0],
					Projects: []apiv1.ProjectGroup{
						{
							GroupPrefix: "owners",
							ID:          "foo-ID",
						},
					},
				},

				{
					ObjectMeta: apiv1.ObjectMeta{
						ID:                "0a0a58273565a8f3dcf779375d9debd0f685d94dc56651a16bff3bf901c0b127",
						Name:              "alice",
						CreationTimestamp: apiv1.Date(2013, 02, 03, 19, 55, 0, 0, time.UTC),
					},
					Email:    "alice@acme.com",
					LastSeen: &[]apiv1.Time{apiv1.NewTime(test.UserLastSeen)}[0],
					Projects: []apiv1.ProjectGroup{
						{
							GroupPrefix: "viewers",
							ID:          "foo-ID",
						},
					},
				},

				{
					ObjectMeta: apiv1.ObjectMeta{
						ID:                "405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6",
						Name:              "bob",
						CreationTimestamp: apiv1.Date(2013, 02, 03, 19, 56, 0, 0, time.UTC),
					},
					Email:    "bob@acme.com",
					LastSeen: &[]apiv1.Time{apiv1.NewTime(test.UserLastSeen)}[0],
					Projects: []apiv1.ProjectGroup{
						{
							GroupPrefix: "editors",
							ID:          "foo-ID",
						},
					},
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/projects/%s/users", tc.ProjectToGet), nil)
			res := httptest.NewRecorder()
			kubermaticObj := []ctrlruntimeclient.Object{}
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
			ep, err := test.CreateTestEndpoint(tc.ExistingAPIUser, []ctrlruntimeclient.Object{}, kubermaticObj, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			ep.ServeHTTP(res, req)

			if res.Code != tc.HTTPStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.HTTPStatus, res.Code, res.Body.String())
			}

			if len(tc.ExpectedResponse) > 0 {
				actualUsers := test.NewUserV1SliceWrapper{}
				actualUsers.DecodeOrDie(res.Body, t).Sort()

				wrappedExpectedUsers := test.NewUserV1SliceWrapper(tc.ExpectedResponse)
				wrappedExpectedUsers.Sort()

				actualUsers.EqualOrDie(wrappedExpectedUsers, t)
			}

			if len(tc.ExpectedResponseString) > 0 {
				test.CompareWithResult(t, res, tc.ExpectedResponseString)
			}
		})
	}
}

func TestDeleteUserFromProject(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		Name                   string
		Body                   string
		ExpectedResponse       string
		ProjectToSync          string
		UserIDToDelete         string
		HTTPStatus             int
		ExistingAPIUser        apiv1.User
		ExistingKubermaticObjs []ctrlruntimeclient.Object
	}{
		// scenario 1
		{
			Name:          "scenario 1: john the owner of the plan9 project removes bob from the project",
			Body:          `{"id":"bobID", "email":"bob@acme.com", "projects":[{"id":"plan9", "group":"editors"}]}`,
			HTTPStatus:    http.StatusOK,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("plan9-ID", "bob@acme.com", "viewers"),
				test.GenBinding("planX-ID", "bob@acme.com", "viewers"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			UserIDToDelete:   genDefaultUser().Name,
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{}`,
		},

		// scenario 2
		{
			Name:          "scenario 2: john the owner of the plan9 project removes bob, but bob is not a member of the project",
			Body:          `{"id":"bobID", "email":"bob@acme.com", "projects":[{"id":"plan9", "group":"editors"}]}`,
			HTTPStatus:    http.StatusBadRequest,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("planX-ID", "bob@acme.com", "viewers"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			UserIDToDelete:   genDefaultUser().Name,
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"error":{"code":400,"message":"cannot delete the user bob@acme.com from the project plan9-ID because the user is not a member of the project"}}`,
		},

		// scenario 3
		{
			Name:          "scenario 3: john the owner of the plan9 project removes himself from the project",
			Body:          fmt.Sprintf(`{"id":"%s", "email":"%s", "projects":[{"id":"plan9", "group":"owners"}]}`, test.UserID, test.UserEmail),
			HTTPStatus:    http.StatusForbidden,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
			},
			UserIDToDelete:   genUser("", "john", "john@acme.com").Name,
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"error":{"code":403,"message":"you cannot delete yourself from the project"}}`,
		},

		// scenario 4
		{
			Name:          "scenario 4: email case insensitive. Remove bob from the project where email is Bob@acme.com instead bob@acme.com",
			Body:          `{"id":"bobID", "email":"Bob@acme.com", "projects":[{"id":"plan9", "group":"editors"}]}`,
			HTTPStatus:    http.StatusOK,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("plan9-ID", "bob@acme.com", "viewers"),
				test.GenBinding("planX-ID", "bob@acme.com", "viewers"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			UserIDToDelete:   genDefaultUser().Name,
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{}`,
		},

		// scenario 5
		{
			Name:          "scenario 5: the admin can remove any member from the project",
			Body:          `{"id":"bobID", "email":"bob@acme.com", "projects":[{"id":"plan9", "group":"editors"}]}`,
			HTTPStatus:    http.StatusOK,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("plan9-ID", "bob@acme.com", "viewers"),
				test.GenBinding("planX-ID", "bob@acme.com", "viewers"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
				genDefaultAdminUser(),
			},
			UserIDToDelete:   genDefaultUser().Name,
			ExistingAPIUser:  *genAPIUser("admin", "admin@acme.com"),
			ExpectedResponse: `{}`,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/projects/%s/users/%s", tc.ProjectToSync, tc.UserIDToDelete), strings.NewReader(tc.Body))
			res := httptest.NewRecorder()
			kubermaticObj := []ctrlruntimeclient.Object{}
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
			ep, err := test.CreateTestEndpoint(tc.ExistingAPIUser, nil, kubermaticObj, nil, hack.NewTestRouting)
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

func TestEditUserInProject(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		Name                   string
		Body                   string
		ExpectedResponse       string
		ProjectToSync          string
		UserIDToUpdate         string
		HTTPStatus             int
		ExistingAPIUser        apiv1.User
		ExistingKubermaticObjs []ctrlruntimeclient.Object
	}{
		// scenario 1
		{
			Name:          "scenario 1: john the owner of the plan9 project changes the group for bob from viewers to editors",
			Body:          `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6", "email":"bob@acme.com", "projects":[{"id":"plan9-ID", "group":"editors"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
			HTTPStatus:    http.StatusOK,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("plan9-ID", "bob@acme.com", "viewers"),
				test.GenBinding("my-third-project-ID", "bob@acme.com", "viewers"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			UserIDToUpdate:   genDefaultUser().Name,
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6","name":"Bob","creationTimestamp":"0001-01-01T00:00:00Z","email":"bob@acme.com","projects":[{"id":"plan9-ID","group":"editors"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
		},

		// scenario 2
		{
			Name:          "scenario 2: john the owner of the plan9 project changes the group for bob, but bob is not a member of the project",
			Body:          `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6", "email":"bob@acme.com", "projects":[{"id":"plan9-ID", "group":"editors"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
			HTTPStatus:    http.StatusBadRequest,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("my-third-project-ID", "bob@acme.com", "viewers"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			UserIDToUpdate:   genDefaultUser().Name,
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"error":{"code":400,"message":"cannot change the membership of the user bob@acme.com for the project plan9-ID because the user is not a member of the project"}}`,
		},

		// scenario 3
		{
			Name:          "scenario 3: john the owner of the plan9 project changes the group for bob from viewers to owners",
			Body:          `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6", "email":"bob@acme.com", "projects":[{"id":"plan9-ID", "group":"owners"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
			HTTPStatus:    http.StatusOK,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("plan9-ID", "bob@acme.com", "viewers"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			UserIDToUpdate:   genDefaultUser().Name,
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6","name":"Bob","creationTimestamp":"0001-01-01T00:00:00Z","email":"bob@acme.com","projects":[{"id":"plan9-ID","group":"owners"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
		},

		// scenario 4
		{
			Name:          "scenario 4: john the owner of the plan9 project changes the group for bob from viewers to admins(wrong name)",
			Body:          `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6", "email":"bob@acme.com", "projects":[{"id":"plan9-ID", "group":"admins"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
			HTTPStatus:    http.StatusBadRequest,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("plan9-ID", "bob@acme.com", "viewers"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			UserIDToUpdate:   genDefaultUser().Name,
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"error":{"code":400,"message":"invalid group name admins"}}`,
		},

		// scenario 5
		{
			Name:          "scenario 5: email case insensitive. Changes the group for bob from viewers to editors where email is BOB@ACME.COM instead bob@acme.com",
			Body:          `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6", "email":"BOB@ACME.COM", "projects":[{"id":"plan9-ID", "group":"editors"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
			HTTPStatus:    http.StatusOK,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("plan9-ID", "bob@acme.com", "viewers"),
				test.GenBinding("my-third-project-ID", "bob@acme.com", "viewers"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			UserIDToUpdate:   genDefaultUser().Name,
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6","name":"Bob","creationTimestamp":"0001-01-01T00:00:00Z","email":"bob@acme.com","projects":[{"id":"plan9-ID","group":"editors"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
		},
		// scenario 6
		{
			Name:          "scenario 6: the admin changes the group for bob from viewers to editors",
			Body:          `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6", "email":"bob@acme.com", "projects":[{"id":"plan9-ID", "group":"editors"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
			HTTPStatus:    http.StatusOK,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("plan9-ID", "bob@acme.com", "viewers"),
				test.GenBinding("my-third-project-ID", "bob@acme.com", "viewers"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
				genDefaultAdminUser(),
			},
			UserIDToUpdate:   genDefaultUser().Name,
			ExistingAPIUser:  *genAPIUser("admin", "admin@acme.com"),
			ExpectedResponse: `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6","name":"Bob","creationTimestamp":"0001-01-01T00:00:00Z","email":"bob@acme.com","projects":[{"id":"plan9-ID","group":"editors"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/v1/projects/%s/users/%s", tc.ProjectToSync, tc.UserIDToUpdate), strings.NewReader(tc.Body))
			res := httptest.NewRecorder()
			kubermaticObj := []ctrlruntimeclient.Object{}
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
			ep, err := test.CreateTestEndpoint(tc.ExistingAPIUser, []ctrlruntimeclient.Object{}, kubermaticObj, nil, hack.NewTestRouting)
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

func TestAddUserToProject(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		Name                   string
		Body                   string
		ExpectedResponse       string
		ProjectToSync          string
		HTTPStatus             int
		ExistingAPIUser        apiv1.User
		ExistingKubermaticObjs []ctrlruntimeclient.Object
	}{
		{
			Name:          "scenario 1: john the owner of the plan9 project invites bob to the project as an editor",
			Body:          `{"email":"bob@acme.com", "projects":[{"id":"plan9-ID", "group":"editors"}]}`,
			HTTPStatus:    http.StatusCreated,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("my-third-project-ID", "john@acme.com", "editors"),
				test.GenBinding("placeX-ID", "bob@acme.com", "editors"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6","name":"Bob","creationTimestamp":"0001-01-01T00:00:00Z","email":"bob@acme.com","projects":[{"id":"plan9-ID","group":"editors"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
		},

		{
			Name:          "scenario 2: john the owner of the plan9 project tries to invite bob to another project",
			Body:          `{"email":"bob@acme.com", "projects":[{"id":"moby", "group":"editors"}]}`,
			HTTPStatus:    http.StatusForbidden,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/* add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("my-third-project-ID", "john@acme.com", "editors"),
				test.GenBinding("placeX-ID", "bob@acme.com", "editors"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"error":{"code":403,"message":"you can only assign the user to plan9-ID project"}}`,
		},

		{
			Name:          "scenario 3: john the owner of the plan9 project tries to invite  himself to another group",
			Body:          fmt.Sprintf(`{"email":"%s", "projects":[{"id":"plan9-ID", "group":"editors"}]}`, test.UserEmail),
			HTTPStatus:    http.StatusForbidden,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("my-third-project-ID", "john@acme.com", "editors"),
				test.GenBinding("placeX-ID", "bob@acme.com", "editors"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"error":{"code":403,"message":"you cannot assign yourself to a different group"}}`,
		},

		{
			Name:          "scenario 4: john the owner of the plan9 project invites bob to the project as an owner",
			Body:          `{"email":"bob@acme.com", "projects":[{"id":"plan9-ID", "group":"owners"}]}`,
			HTTPStatus:    http.StatusCreated,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("my-third-project-ID", "john@acme.com", "editors"),
				test.GenBinding("placeX-ID", "bob@acme.com", "editors"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6","name":"Bob","creationTimestamp":"0001-01-01T00:00:00Z","email":"bob@acme.com","projects":[{"id":"plan9-ID","group":"owners"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
		},

		{
			Name:          "scenario 5: john the owner of the plan9 project invites bob to the project one more time",
			Body:          `{"email":"bob@acme.com", "projects":[{"id":"plan9-ID", "group":"editors"}]}`,
			HTTPStatus:    http.StatusBadRequest,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("my-third-project-ID", "john@acme.com", "editors"),
				test.GenBinding("plan9-ID", "bob@acme.com", "editors"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"error":{"code":400,"message":"cannot add the user bob@acme.com to the project plan9-ID because user is already in the project"}}`,
		},

		{
			Name:          "scenario 6: email case insensitive. Bob is invited to the project one more time and emil starts with capital letter",
			Body:          `{"email":"Bob@acme.com", "projects":[{"id":"plan9-ID", "group":"editors"}]}`,
			HTTPStatus:    http.StatusBadRequest,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("my-third-project-ID", "john@acme.com", "editors"),
				test.GenBinding("plan9-ID", "bob@acme.com", "editors"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"error":{"code":400,"message":"cannot add the user Bob@acme.com to the project plan9-ID because user is already in the project"}}`,
		},

		{
			Name:          "scenario 7: email case insensitive. Bob is invited to the project as an editor with capital letter email",
			Body:          `{"email":"BOB@ACME.COM", "projects":[{"id":"plan9-ID", "group":"editors"}]}`,
			HTTPStatus:    http.StatusCreated,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("my-third-project-ID", "john@acme.com", "editors"),
				test.GenBinding("placeX-ID", "bob@acme.com", "editors"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
			},
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6","name":"Bob","creationTimestamp":"0001-01-01T00:00:00Z","email":"bob@acme.com","projects":[{"id":"plan9-ID","group":"editors"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
		},

		{
			Name: "scenario 8: john tries to add a service account to a project",
			Body: func() string {
				sa := test.GenProjectServiceAccount("1", "test-1", "editors", "plan9-ID")
				return fmt.Sprintf(`{"email":"%s", "projects":[{"id":"plan9-ID", "group":"editors"}]}`, sa.Spec.Email)
			}(),
			HTTPStatus:    http.StatusBadRequest,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/* add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("my-third-project-ID", "john@acme.com", "editors"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				test.GenProjectServiceAccount("1", "test-1", "editors", "plan9-ID"),
			},
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedResponse: `{"error":{"code":400,"message":"cannot add the given member serviceaccount-1@sa.kubermatic.io to the project plan9 because the email indicates a service account"}}`,
		},
		{
			Name:          "scenario 9: the admin invites bob to the project as an editor",
			Body:          `{"email":"bob@acme.com", "projects":[{"id":"plan9-ID", "group":"editors"}]}`,
			HTTPStatus:    http.StatusCreated,
			ProjectToSync: "plan9-ID",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("my-first-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("my-third-project", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				test.GenBinding("my-third-project-ID", "john@acme.com", "editors"),
				test.GenBinding("placeX-ID", "bob@acme.com", "editors"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
				genDefaultUser(), /*bob*/
				genDefaultAdminUser(),
			},
			ExistingAPIUser:  *genAPIUser("admin", "admin@acme.com"),
			ExpectedResponse: `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6","name":"Bob","creationTimestamp":"0001-01-01T00:00:00Z","email":"bob@acme.com","projects":[{"id":"plan9-ID","group":"editors"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/v1/projects/%s/users", tc.ProjectToSync), strings.NewReader(tc.Body))
			res := httptest.NewRecorder()
			kubermaticObj := []ctrlruntimeclient.Object{}
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
			ep, err := test.CreateTestEndpoint(tc.ExistingAPIUser, []ctrlruntimeclient.Object{}, kubermaticObj, nil, hack.NewTestRouting)
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

func TestGetCurrentUser(t *testing.T) {
	testcases := []struct {
		Name                   string
		ExpectedResponse       string
		ExpectedStatus         int
		ExistingKubermaticObjs []ctrlruntimeclient.Object
		ExistingAPIUser        apiv1.User
	}{
		{
			Name: "scenario 1: get john's profile (no projects assigned)",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add users*/
				genUser("", "john", "john@acme.com"),
			},
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedStatus:   http.StatusOK,
			ExpectedResponse: `{"id":"4b2d8785b49bad23638b17d8db76857a79bf79441241a78a97d88cc64bbf766e","name":"john","creationTimestamp":"0001-01-01T00:00:00Z","email":"john@acme.com","lastSeen":"2020-12-31T23:00:00Z"}`,
		},

		{
			Name: "scenario 2: get john's profile (one project assigned)",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add projects*/
				test.GenProject("moby", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				test.GenProject("plan9", kubermaticv1.ProjectActive, test.DefaultCreationTimestamp()),
				/*add bindings*/
				test.GenBinding("plan9-ID", "john@acme.com", "owners"),
				/*add users*/
				genUser("", "john", "john@acme.com"),
			},
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedStatus:   http.StatusOK,
			ExpectedResponse: `{"id":"4b2d8785b49bad23638b17d8db76857a79bf79441241a78a97d88cc64bbf766e","name":"john","creationTimestamp":"0001-01-01T00:00:00Z","email":"john@acme.com","projects":[{"id":"plan9-ID","group":"owners"}],"lastSeen":"2020-12-31T23:00:00Z"}`,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			kubermaticObj := []ctrlruntimeclient.Object{}
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
			ep, err := test.CreateTestEndpoint(tc.ExistingAPIUser, []ctrlruntimeclient.Object{}, kubermaticObj, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			req := httptest.NewRequest(http.MethodGet, "/api/v1/me", nil)
			res := httptest.NewRecorder()
			ep.ServeHTTP(res, req)

			if res.Code != tc.ExpectedStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.ExpectedStatus, res.Code, res.Body.String())
			}
			test.CompareWithResult(t, res, tc.ExpectedResponse)
		})
	}
}

func TestNewUser(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		Name                      string
		HTTPStatus                int
		ExpectedResponse          string
		ExpectedKubermaticUser    *kubermaticv1.User
		ExistingKubermaticObjects []ctrlruntimeclient.Object
		ExistingAPIUser           *apiv1.User
	}{
		{
			Name:             "scenario 1: successfully creates the initial new user resource",
			ExpectedResponse: `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6","name":"Bob","creationTimestamp":"0001-01-01T00:00:00Z","email":"bob@acme.com","isAdmin":true,"lastSeen":"2020-12-31T23:00:00Z"}`,
			HTTPStatus:       http.StatusOK,
			ExpectedKubermaticUser: func() *kubermaticv1.User {
				expectedKubermaticUser := test.GenDefaultUser()
				expectedKubermaticUser.UID = ""
				return expectedKubermaticUser
			}(),
			ExistingAPIUser: genDefaultAPIUser(),
		},

		{
			Name:             "scenario 2: fails when creating a user without an email address",
			ExpectedResponse: `{"error":{"code":401,"message":"not authorized"}}`,
			HTTPStatus:       http.StatusUnauthorized,
			ExistingAPIUser: func() *apiv1.User {
				apiUser := genDefaultAPIUser()
				apiUser.Email = ""
				return apiUser
			}(),
		},

		{
			Name:             "scenario 3: creating a user if already exists doesn't have effect",
			ExpectedResponse: `{"id":"405ac8384fa984f787f9486daf34d84d98f20c4d6a12e2cc4ed89be3bcb06ad6","name":"Bob","creationTimestamp":"0001-01-01T00:00:00Z","email":"bob@acme.com","lastSeen":"2020-12-31T23:00:00Z"}`,
			HTTPStatus:       http.StatusOK,
			ExistingKubermaticObjects: []ctrlruntimeclient.Object{
				genDefaultUser(),
			},
			ExistingAPIUser: genDefaultAPIUser(),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			ep, err := test.CreateTestEndpoint(*tc.ExistingAPIUser, []ctrlruntimeclient.Object{}, tc.ExistingKubermaticObjects, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			// act
			req := httptest.NewRequest(http.MethodGet, "/api/v1/me", nil)
			res := httptest.NewRecorder()
			ep.ServeHTTP(res, req)

			// validate
			if res.Code != tc.HTTPStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.HTTPStatus, res.Code, res.Body.String())
			}
			test.CompareWithResult(t, res, tc.ExpectedResponse)
		})
	}
}

func TestLogoutCurrentUser(t *testing.T) {
	testcases := []struct {
		Name                   string
		ExpectedResponse       string
		ExpectedStatus         int
		ExistingKubermaticObjs []ctrlruntimeclient.Object
		ExistingKubernetesObjs []ctrlruntimeclient.Object
		ExistingAPIUser        apiv1.User
	}{
		{
			Name: "scenario 1: logout user first time",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add users*/
				genUser("", "john", "john@acme.com"),
			},
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedStatus:   http.StatusOK,
			ExpectedResponse: `{}`,
		},
		{
			Name: "scenario 2: logout user with empty blacklist token secret",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add users*/
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
			ExistingKubernetesObjs: []ctrlruntimeclient.Object{
				func() *corev1.Secret {
					user := genUser("", "john", "john@acme.com")
					return test.GenBlacklistTokenSecret(user.GetInvalidTokensReferenceSecretName(), []byte{})
				}(),
			},
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedStatus:   http.StatusOK,
			ExpectedResponse: `{}`,
		},
		{
			Name: "scenario 2: logout user when token is on blacklist",
			ExistingKubermaticObjs: []ctrlruntimeclient.Object{
				/*add users*/
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
			ExistingKubernetesObjs: []ctrlruntimeclient.Object{
				func() *corev1.Secret {
					user := genUser("", "john", "john@acme.com")
					return test.GenBlacklistTokenSecret(user.GetInvalidTokensReferenceSecretName(), []byte(`[{"token":"fakeTokenId","expiry":"2222-06-20T12:04:00Z"}]`))
				}(),
			},
			ExistingAPIUser:  *genAPIUser("john", "john@acme.com"),
			ExpectedStatus:   http.StatusUnauthorized,
			ExpectedResponse: `{"error":{"code":401,"message":"not authorized"}}`,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.Name, func(t *testing.T) {
			kubermaticObj := []ctrlruntimeclient.Object{}
			kubernetesObj := []ctrlruntimeclient.Object{}
			kubermaticObj = append(kubermaticObj, tc.ExistingKubermaticObjs...)
			kubernetesObj = append(kubernetesObj, tc.ExistingKubernetesObjs...)
			ep, err := test.CreateTestEndpoint(tc.ExistingAPIUser, kubernetesObj, kubermaticObj, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint: %v", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/v1/me/logout", nil)
			res := httptest.NewRecorder()
			ep.ServeHTTP(res, req)

			if res.Code != tc.ExpectedStatus {
				t.Fatalf("Expected HTTP status code %d, got %d: %s", tc.ExpectedStatus, res.Code, res.Body.String())
			}
			test.CompareWithResult(t, res, tc.ExpectedResponse)
		})
	}
}

// genUser generates a User resource
// note if the id is empty then it will be auto generated.
func genUser(id, name, email string) *kubermaticv1.User {
	return test.GenUser(id, name, email)
}

func genDefaultUser() *kubermaticv1.User {
	return test.GenDefaultUser()
}

func genAPIUser(name, email string) *apiv1.User {
	return test.GenAPIUser(name, email)
}

func genDefaultAPIUser() *apiv1.User {
	return test.GenDefaultAPIUser()
}

func genDefaultAdminUser() *kubermaticv1.User {
	user := test.GenUser("", "admin", "admin@acme.com")
	user.Spec.IsAdmin = true
	return user
}
