// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewListPolicyTemplateParams creates a new ListPolicyTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListPolicyTemplateParams() *ListPolicyTemplateParams {
	return &ListPolicyTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListPolicyTemplateParamsWithTimeout creates a new ListPolicyTemplateParams object
// with the ability to set a timeout on a request.
func NewListPolicyTemplateParamsWithTimeout(timeout time.Duration) *ListPolicyTemplateParams {
	return &ListPolicyTemplateParams{
		timeout: timeout,
	}
}

// NewListPolicyTemplateParamsWithContext creates a new ListPolicyTemplateParams object
// with the ability to set a context for a request.
func NewListPolicyTemplateParamsWithContext(ctx context.Context) *ListPolicyTemplateParams {
	return &ListPolicyTemplateParams{
		Context: ctx,
	}
}

// NewListPolicyTemplateParamsWithHTTPClient creates a new ListPolicyTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewListPolicyTemplateParamsWithHTTPClient(client *http.Client) *ListPolicyTemplateParams {
	return &ListPolicyTemplateParams{
		HTTPClient: client,
	}
}

/*
ListPolicyTemplateParams contains all the parameters to send to the API endpoint

	for the list policy template operation.

	Typically these are written to a http.Request.
*/
type ListPolicyTemplateParams struct {

	// ProjectID.
	ProjectID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list policy template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPolicyTemplateParams) WithDefaults() *ListPolicyTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list policy template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListPolicyTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list policy template params
func (o *ListPolicyTemplateParams) WithTimeout(timeout time.Duration) *ListPolicyTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list policy template params
func (o *ListPolicyTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list policy template params
func (o *ListPolicyTemplateParams) WithContext(ctx context.Context) *ListPolicyTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list policy template params
func (o *ListPolicyTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list policy template params
func (o *ListPolicyTemplateParams) WithHTTPClient(client *http.Client) *ListPolicyTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list policy template params
func (o *ListPolicyTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the list policy template params
func (o *ListPolicyTemplateParams) WithProjectID(projectID *string) *ListPolicyTemplateParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list policy template params
func (o *ListPolicyTemplateParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListPolicyTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ProjectID != nil {

		// query param project_id
		var qrProjectID string

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {

			if err := r.SetQueryParam("project_id", qProjectID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
