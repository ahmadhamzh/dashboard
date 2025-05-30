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

package seedsettings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	apiv2 "k8c.io/dashboard/v2/pkg/api/v2"
	"k8c.io/dashboard/v2/pkg/handler/v1/common"
	"k8c.io/dashboard/v2/pkg/provider"
	kubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/kubermatic/v1"
	utilerrors "k8c.io/kubermatic/v2/pkg/util/errors"
)

// getSeedSettingsReq defines HTTP request for getSeedSettings
// swagger:parameters getSeedSettings
type getSeedSettingsReq struct {
	// in: path
	// required: true
	Name string `json:"seed_name"`
}

func DecodeGetSeedSettingsReq(c context.Context, r *http.Request) (interface{}, error) {
	var req getSeedSettingsReq
	name := mux.Vars(r)["seed_name"]
	if name == "" {
		return nil, fmt.Errorf("'seed_name' parameter is required but was not provided")
	}
	req.Name = name
	return req, nil
}

// GetSeedSettingsEndpoint returns SeedSettings for a Seed cluster.
func GetSeedSettingsEndpoint(seedsGetter provider.SeedsGetter) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(getSeedSettingsReq)
		if !ok {
			return nil, utilerrors.NewBadRequest("invalid request")
		}

		seedMap, err := seedsGetter()
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		seed, ok := seedMap[req.Name]
		if !ok {
			return nil, utilerrors.NewNotFound("Seed", req.Name)
		}
		return convertSeedToSeedSettings(seed), nil
	}
}

func convertSeedToSeedSettings(seed *kubermaticv1.Seed) *apiv2.SeedSettings {
	seedSettings := &apiv2.SeedSettings{}
	if seed.Spec.MLA != nil && seed.Spec.MLA.UserClusterMLAEnabled {
		seedSettings.MLA.UserClusterMLAEnabled = true
	}

	if seed.Spec.Metering != nil {
		seedSettings.Metering.Enabled = seed.Spec.Metering.Enabled
		seedSettings.Metering.StorageClassName = seed.Spec.Metering.StorageClassName
		seedSettings.Metering.StorageSize = seed.Spec.Metering.StorageSize
	}

	if len(seed.Spec.SeedDNSOverwrite) > 0 {
		seedSettings.SeedDNSOverwrite = seed.Spec.SeedDNSOverwrite
	}

	if seed.Spec.KubeLB != nil {
		seedSettings.KubeLB = &apiv2.KubeLBSeedSettingsAPI{
			EnableForAllDatacenters: seed.Spec.KubeLB.EnableForAllDatacenters,
		}
	}

	return seedSettings
}
