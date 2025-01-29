// Code generated by go-swagger; DO NOT EDIT.

package datacenter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new datacenter API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new datacenter API client with basic auth credentials.
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

// New creates a new datacenter API client with a bearer token for authentication.
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
Client for datacenter API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDC(params *CreateDCParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDCCreated, error)

	DeleteDC(params *DeleteDCParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDCOK, error)

	GetDCForProvider(params *GetDCForProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDCForProviderOK, error)

	GetDCForSeed(params *GetDCForSeedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDCForSeedOK, error)

	GetDatacenter(params *GetDatacenterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDatacenterOK, error)

	ListDCForProvider(params *ListDCForProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListDCForProviderOK, error)

	ListDCForSeed(params *ListDCForSeedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListDCForSeedOK, error)

	ListDatacenters(params *ListDatacentersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListDatacentersOK, error)

	PatchDC(params *PatchDCParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchDCOK, error)

	UpdateDC(params *UpdateDCParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDCOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateDC creates the datacenter for a specified seed
*/
func (a *Client) CreateDC(params *CreateDCParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDCCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDCParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDC",
		Method:             "POST",
		PathPattern:        "/api/v1/seed/{seed_name}/dc",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDCReader{formats: a.formats},
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
	success, ok := result.(*CreateDCCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateDCDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteDC deletes the datacenter
*/
func (a *Client) DeleteDC(params *DeleteDCParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteDCOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDCParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteDC",
		Method:             "DELETE",
		PathPattern:        "/api/v1/seed/{seed_name}/dc/{dc}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDCReader{formats: a.formats},
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
	success, ok := result.(*DeleteDCOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteDCDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetDCForProvider gets the datacenter for the specified provider
*/
func (a *Client) GetDCForProvider(params *GetDCForProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDCForProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDCForProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDCForProvider",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/{provider_name}/dc/{dc}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDCForProviderReader{formats: a.formats},
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
	success, ok := result.(*GetDCForProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDCForProviderDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetDCForSeed returns the specified datacenter for the specified seed
*/
func (a *Client) GetDCForSeed(params *GetDCForSeedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDCForSeedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDCForSeedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDCForSeed",
		Method:             "GET",
		PathPattern:        "/api/v1/seed/{seed_name}/dc/{dc}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDCForSeedReader{formats: a.formats},
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
	success, ok := result.(*GetDCForSeedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDCForSeedDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetDatacenter get datacenter API
*/
func (a *Client) GetDatacenter(params *GetDatacenterParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDatacenterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatacenterParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDatacenter",
		Method:             "GET",
		PathPattern:        "/api/v1/dc/{dc}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDatacenterReader{formats: a.formats},
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
	success, ok := result.(*GetDatacenterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDatacenterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListDCForProvider returns all datacenters for the specified provider
*/
func (a *Client) ListDCForProvider(params *ListDCForProviderParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListDCForProviderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDCForProviderParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listDCForProvider",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/{provider_name}/dc",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDCForProviderReader{formats: a.formats},
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
	success, ok := result.(*ListDCForProviderOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDCForProviderDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListDCForSeed returns all datacenters for the specified seed
*/
func (a *Client) ListDCForSeed(params *ListDCForSeedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListDCForSeedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDCForSeedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listDCForSeed",
		Method:             "GET",
		PathPattern:        "/api/v1/seed/{seed_name}/dc",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDCForSeedReader{formats: a.formats},
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
	success, ok := result.(*ListDCForSeedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDCForSeedDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListDatacenters list datacenters API
*/
func (a *Client) ListDatacenters(params *ListDatacentersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListDatacentersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDatacentersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listDatacenters",
		Method:             "GET",
		PathPattern:        "/api/v1/dc",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDatacentersReader{formats: a.formats},
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
	success, ok := result.(*ListDatacentersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDatacentersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PatchDC patches the datacenter
*/
func (a *Client) PatchDC(params *PatchDCParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchDCOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchDCParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patchDC",
		Method:             "PATCH",
		PathPattern:        "/api/v1/seed/{seed_name}/dc/{dc}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchDCReader{formats: a.formats},
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
	success, ok := result.(*PatchDCOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchDCDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateDC updates the datacenter the datacenter spec will be overwritten with the one provided in the request
*/
func (a *Client) UpdateDC(params *UpdateDCParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDCOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDCParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDC",
		Method:             "PUT",
		PathPattern:        "/api/v1/seed/{seed_name}/dc/{dc}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDCReader{formats: a.formats},
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
	success, ok := result.(*UpdateDCOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateDCDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
