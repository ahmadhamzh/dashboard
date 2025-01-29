// Code generated by go-swagger; DO NOT EDIT.

package gke

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new gke API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new gke API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new gke API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for gke API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ListGKEClusterDiskTypes(params *ListGKEClusterDiskTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEClusterDiskTypesOK, error)

	ListGKEClusterImages(params *ListGKEClusterImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEClusterImagesOK, error)

	ListGKEClusterSizes(params *ListGKEClusterSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEClusterSizesOK, error)

	ListGKEClusterZones(params *ListGKEClusterZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEClusterZonesOK, error)

	ListGKEDiskTypes(params *ListGKEDiskTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEDiskTypesOK, error)

	ListGKEImages(params *ListGKEImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEImagesOK, error)

	ListGKEVMSizes(params *ListGKEVMSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEVMSizesOK, error)

	ListGKEVersions(params *ListGKEVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEVersionsOK, error)

	ListGKEZones(params *ListGKEZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEZonesOK, error)

	ListProjectGCPVMSizes(params *ListProjectGCPVMSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGCPVMSizesOK, error)

	ListProjectGCPZones(params *ListProjectGCPZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGCPZonesOK, error)

	ListProjectGKEDiskTypes(params *ListProjectGKEDiskTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGKEDiskTypesOK, error)

	ListProjectGKEImages(params *ListProjectGKEImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGKEImagesOK, error)

	ListProjectGKEVMSizes(params *ListProjectGKEVMSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGKEVMSizesOK, error)

	ListProjectGKEVersions(params *ListProjectGKEVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGKEVersionsOK, error)

	ListProjectGKEZones(params *ListProjectGKEZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGKEZonesOK, error)

	ValidateGKECredentials(params *ValidateGKECredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ValidateGKECredentialsOK, error)

	ValidateProjectGKECredentials(params *ValidateProjectGKECredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ValidateProjectGKECredentialsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ListGKEClusterDiskTypes gets g k e cluster machine disk types
*/
func (a *Client) ListGKEClusterDiskTypes(params *ListGKEClusterDiskTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEClusterDiskTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGKEClusterDiskTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGKEClusterDiskTypes",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/gke/disktypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGKEClusterDiskTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGKEClusterDiskTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGKEClusterDiskTypesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGKEClusterImages gets g k e cluster images
*/
func (a *Client) ListGKEClusterImages(params *ListGKEClusterImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEClusterImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGKEClusterImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGKEClusterImages",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/gke/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGKEClusterImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGKEClusterImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGKEClusterImagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGKEClusterSizes gets g k e cluster machine sizes
*/
func (a *Client) ListGKEClusterSizes(params *ListGKEClusterSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEClusterSizesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGKEClusterSizesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGKEClusterSizes",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/gke/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGKEClusterSizesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGKEClusterSizesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGKEClusterSizesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGKEClusterZones gets g k e cluster zones
*/
func (a *Client) ListGKEClusterZones(params *ListGKEClusterZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEClusterZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGKEClusterZonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGKEClusterZones",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/gke/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGKEClusterZonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGKEClusterZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGKEClusterZonesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGKEDiskTypes gets g k e machine disk types
*/
func (a *Client) ListGKEDiskTypes(params *ListGKEDiskTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEDiskTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGKEDiskTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGKEDiskTypes",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/gke/disktypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGKEDiskTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGKEDiskTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGKEDiskTypesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGKEImages Lists GKE image types
*/
func (a *Client) ListGKEImages(params *ListGKEImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGKEImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGKEImages",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/gke/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGKEImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGKEImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGKEImagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGKEVMSizes Lists GKE vmsizes
*/
func (a *Client) ListGKEVMSizes(params *ListGKEVMSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEVMSizesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGKEVMSizesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGKEVMSizes",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/gke/vmsizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGKEVMSizesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGKEVMSizesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGKEVMSizesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGKEVersions Lists GKE versions
*/
func (a *Client) ListGKEVersions(params *ListGKEVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGKEVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGKEVersions",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/gke/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGKEVersionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGKEVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGKEVersionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListGKEZones Lists GKE zones
*/
func (a *Client) ListGKEZones(params *ListGKEZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListGKEZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListGKEZonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listGKEZones",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/gke/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListGKEZonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListGKEZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListGKEZonesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectGCPVMSizes lists g c p VM sizes
*/
func (a *Client) ListProjectGCPVMSizes(params *ListProjectGCPVMSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGCPVMSizesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectGCPVMSizesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectGCPVMSizes",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/gcp/vmsizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectGCPVMSizesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectGCPVMSizesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectGCPVMSizesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	ListProjectGCPZones lists g c p zones

	Produces

application/json
*/
func (a *Client) ListProjectGCPZones(params *ListProjectGCPZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGCPZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectGCPZonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectGCPZones",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/gcp/{dc}/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectGCPZonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectGCPZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectGCPZonesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectGKEDiskTypes lists g k e machine disk types
*/
func (a *Client) ListProjectGKEDiskTypes(params *ListProjectGKEDiskTypesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGKEDiskTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectGKEDiskTypesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectGKEDiskTypes",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/gke/disktypes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectGKEDiskTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectGKEDiskTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectGKEDiskTypesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectGKEImages lists g k e image types
*/
func (a *Client) ListProjectGKEImages(params *ListProjectGKEImagesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGKEImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectGKEImagesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectGKEImages",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/gke/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectGKEImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectGKEImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectGKEImagesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectGKEVMSizes lists g k e VM sizes
*/
func (a *Client) ListProjectGKEVMSizes(params *ListProjectGKEVMSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGKEVMSizesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectGKEVMSizesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectGKEVMSizes",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/gke/vmsizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectGKEVMSizesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectGKEVMSizesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectGKEVMSizesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectGKEVersions lists g k e versions
*/
func (a *Client) ListProjectGKEVersions(params *ListProjectGKEVersionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGKEVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectGKEVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectGKEVersions",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/gke/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectGKEVersionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectGKEVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectGKEVersionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListProjectGKEZones lists g k e zones
*/
func (a *Client) ListProjectGKEZones(params *ListProjectGKEZonesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListProjectGKEZonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectGKEZonesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listProjectGKEZones",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/gke/zones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectGKEZonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListProjectGKEZonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListProjectGKEZonesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ValidateGKECredentials Validates GKE credentials
*/
func (a *Client) ValidateGKECredentials(params *ValidateGKECredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ValidateGKECredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateGKECredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateGKECredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/providers/gke/validatecredentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateGKECredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateGKECredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ValidateGKECredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ValidateProjectGKECredentials validates g k e credentials
*/
func (a *Client) ValidateProjectGKECredentials(params *ValidateProjectGKECredentialsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ValidateProjectGKECredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateProjectGKECredentialsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateProjectGKECredentials",
		Method:             "GET",
		PathPattern:        "/api/v2/projects/{project_id}/providers/gke/validatecredentials",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateProjectGKECredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ValidateProjectGKECredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ValidateProjectGKECredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
