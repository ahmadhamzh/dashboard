//go:build ee

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

	kubelb "k8c.io/dashboard/v2/pkg/ee/kubelb"
	"k8c.io/dashboard/v2/pkg/provider"
)

func listKubeLBTenants(ctx context.Context, request interface{}, seedsGetter provider.SeedsGetter) (interface{}, error) {
	return kubelb.ListKubeLBTenants(ctx, request, seedsGetter)
}

func decodeListKubeLBTenantsReq(c context.Context, r *http.Request) (interface{}, error) {
	return kubelb.DecodeListKubeLBTenantsReq(c, r)
}

func getKubeLBTenants(ctx context.Context, request interface{}, seedsGetter provider.SeedsGetter) (interface{}, error) {
	return kubelb.GetKubeLBTenants(ctx, request, seedsGetter)
}

func decodeGetKubeLBTenantsReq(c context.Context, r *http.Request) (interface{}, error) {
	return kubelb.DecodeGetKubeLBTenantsReq(c, r)
}
