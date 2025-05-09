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

package v1

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	apiv1 "k8c.io/dashboard/v2/pkg/api/v1"
	"k8c.io/kubermatic/v2/pkg/version/kubermatic"
)

// GetKubermaticVersion returns the versions of running Kubermatic components.
//
// At this time we're only interested in the version of the API server,
// since it knows its own version and can answer the endpoint promptly.
//
// This version string is constructed with `git describe` and embedded in the binary during build.
func GetKubermaticVersion(versions kubermatic.Versions) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		versions := apiv1.KubermaticVersions{API: versions.GitVersion}
		return versions, nil
	}
}
