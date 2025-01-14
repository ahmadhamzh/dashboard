/*
Copyright 2024 The Kubermatic Kubernetes Platform contributors.

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

package kubelb

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"k8c.io/dashboard/v2/pkg/provider"
)

func ListKubeLBTenants(seedsGetter provider.SeedsGetter) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return listKubeLBTenants(ctx, request, seedsGetter)
	}
}

func DecodeListKubeLBTenantsReq(c context.Context, r *http.Request) (interface{}, error) {
	return decodeListKubeLBTenantsReq(c, r)
}

func GetKubeLBTenants(seedsGetter provider.SeedsGetter) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return getKubeLBTenants(ctx, request, seedsGetter)
	}
}

func DecodeGetKubeLBTenantsReq(c context.Context, r *http.Request) (interface{}, error) {
	return decodeGetKubeLBTenantsReq(c, r)
}
