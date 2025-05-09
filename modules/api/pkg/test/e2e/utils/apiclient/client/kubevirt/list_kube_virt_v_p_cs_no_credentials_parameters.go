// Code generated by go-swagger; DO NOT EDIT.

package kubevirt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListKubeVirtVPCsNoCredentialsParams creates a new ListKubeVirtVPCsNoCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListKubeVirtVPCsNoCredentialsParams() *ListKubeVirtVPCsNoCredentialsParams {
	return &ListKubeVirtVPCsNoCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListKubeVirtVPCsNoCredentialsParamsWithTimeout creates a new ListKubeVirtVPCsNoCredentialsParams object
// with the ability to set a timeout on a request.
func NewListKubeVirtVPCsNoCredentialsParamsWithTimeout(timeout time.Duration) *ListKubeVirtVPCsNoCredentialsParams {
	return &ListKubeVirtVPCsNoCredentialsParams{
		timeout: timeout,
	}
}

// NewListKubeVirtVPCsNoCredentialsParamsWithContext creates a new ListKubeVirtVPCsNoCredentialsParams object
// with the ability to set a context for a request.
func NewListKubeVirtVPCsNoCredentialsParamsWithContext(ctx context.Context) *ListKubeVirtVPCsNoCredentialsParams {
	return &ListKubeVirtVPCsNoCredentialsParams{
		Context: ctx,
	}
}

// NewListKubeVirtVPCsNoCredentialsParamsWithHTTPClient creates a new ListKubeVirtVPCsNoCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListKubeVirtVPCsNoCredentialsParamsWithHTTPClient(client *http.Client) *ListKubeVirtVPCsNoCredentialsParams {
	return &ListKubeVirtVPCsNoCredentialsParams{
		HTTPClient: client,
	}
}

/*
ListKubeVirtVPCsNoCredentialsParams contains all the parameters to send to the API endpoint

	for the list kube virt v p cs no credentials operation.

	Typically these are written to a http.Request.
*/
type ListKubeVirtVPCsNoCredentialsParams struct {

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list kube virt v p cs no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListKubeVirtVPCsNoCredentialsParams) WithDefaults() *ListKubeVirtVPCsNoCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list kube virt v p cs no credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListKubeVirtVPCsNoCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list kube virt v p cs no credentials params
func (o *ListKubeVirtVPCsNoCredentialsParams) WithTimeout(timeout time.Duration) *ListKubeVirtVPCsNoCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list kube virt v p cs no credentials params
func (o *ListKubeVirtVPCsNoCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list kube virt v p cs no credentials params
func (o *ListKubeVirtVPCsNoCredentialsParams) WithContext(ctx context.Context) *ListKubeVirtVPCsNoCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list kube virt v p cs no credentials params
func (o *ListKubeVirtVPCsNoCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list kube virt v p cs no credentials params
func (o *ListKubeVirtVPCsNoCredentialsParams) WithHTTPClient(client *http.Client) *ListKubeVirtVPCsNoCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list kube virt v p cs no credentials params
func (o *ListKubeVirtVPCsNoCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list kube virt v p cs no credentials params
func (o *ListKubeVirtVPCsNoCredentialsParams) WithClusterID(clusterID string) *ListKubeVirtVPCsNoCredentialsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list kube virt v p cs no credentials params
func (o *ListKubeVirtVPCsNoCredentialsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list kube virt v p cs no credentials params
func (o *ListKubeVirtVPCsNoCredentialsParams) WithProjectID(projectID string) *ListKubeVirtVPCsNoCredentialsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list kube virt v p cs no credentials params
func (o *ListKubeVirtVPCsNoCredentialsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListKubeVirtVPCsNoCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
