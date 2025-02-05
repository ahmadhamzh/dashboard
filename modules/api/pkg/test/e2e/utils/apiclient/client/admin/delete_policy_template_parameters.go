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

// NewDeletePolicyTemplateParams creates a new DeletePolicyTemplateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePolicyTemplateParams() *DeletePolicyTemplateParams {
	return &DeletePolicyTemplateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePolicyTemplateParamsWithTimeout creates a new DeletePolicyTemplateParams object
// with the ability to set a timeout on a request.
func NewDeletePolicyTemplateParamsWithTimeout(timeout time.Duration) *DeletePolicyTemplateParams {
	return &DeletePolicyTemplateParams{
		timeout: timeout,
	}
}

// NewDeletePolicyTemplateParamsWithContext creates a new DeletePolicyTemplateParams object
// with the ability to set a context for a request.
func NewDeletePolicyTemplateParamsWithContext(ctx context.Context) *DeletePolicyTemplateParams {
	return &DeletePolicyTemplateParams{
		Context: ctx,
	}
}

// NewDeletePolicyTemplateParamsWithHTTPClient creates a new DeletePolicyTemplateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePolicyTemplateParamsWithHTTPClient(client *http.Client) *DeletePolicyTemplateParams {
	return &DeletePolicyTemplateParams{
		HTTPClient: client,
	}
}

/*
DeletePolicyTemplateParams contains all the parameters to send to the API endpoint

	for the delete policy template operation.

	Typically these are written to a http.Request.
*/
type DeletePolicyTemplateParams struct {

	// ProjectID.
	ProjectID *string

	// TemplateName.
	PolicyTemplateName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete policy template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePolicyTemplateParams) WithDefaults() *DeletePolicyTemplateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete policy template params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePolicyTemplateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete policy template params
func (o *DeletePolicyTemplateParams) WithTimeout(timeout time.Duration) *DeletePolicyTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete policy template params
func (o *DeletePolicyTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete policy template params
func (o *DeletePolicyTemplateParams) WithContext(ctx context.Context) *DeletePolicyTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete policy template params
func (o *DeletePolicyTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete policy template params
func (o *DeletePolicyTemplateParams) WithHTTPClient(client *http.Client) *DeletePolicyTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete policy template params
func (o *DeletePolicyTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the delete policy template params
func (o *DeletePolicyTemplateParams) WithProjectID(projectID *string) *DeletePolicyTemplateParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete policy template params
func (o *DeletePolicyTemplateParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithPolicyTemplateName adds the templateName to the delete policy template params
func (o *DeletePolicyTemplateParams) WithPolicyTemplateName(templateName string) *DeletePolicyTemplateParams {
	o.SetPolicyTemplateName(templateName)
	return o
}

// SetPolicyTemplateName adds the templateName to the delete policy template params
func (o *DeletePolicyTemplateParams) SetPolicyTemplateName(templateName string) {
	o.PolicyTemplateName = templateName
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePolicyTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param template_name
	if err := r.SetPathParam("template_name", o.PolicyTemplateName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
