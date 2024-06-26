// Code generated by go-swagger; DO NOT EDIT.

package preset

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
	"github.com/go-openapi/swag"
)

// NewListProjectPresetsParams creates a new ListProjectPresetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectPresetsParams() *ListProjectPresetsParams {
	return &ListProjectPresetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectPresetsParamsWithTimeout creates a new ListProjectPresetsParams object
// with the ability to set a timeout on a request.
func NewListProjectPresetsParamsWithTimeout(timeout time.Duration) *ListProjectPresetsParams {
	return &ListProjectPresetsParams{
		timeout: timeout,
	}
}

// NewListProjectPresetsParamsWithContext creates a new ListProjectPresetsParams object
// with the ability to set a context for a request.
func NewListProjectPresetsParamsWithContext(ctx context.Context) *ListProjectPresetsParams {
	return &ListProjectPresetsParams{
		Context: ctx,
	}
}

// NewListProjectPresetsParamsWithHTTPClient creates a new ListProjectPresetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectPresetsParamsWithHTTPClient(client *http.Client) *ListProjectPresetsParams {
	return &ListProjectPresetsParams{
		HTTPClient: client,
	}
}

/*
ListProjectPresetsParams contains all the parameters to send to the API endpoint

	for the list project presets operation.

	Typically these are written to a http.Request.
*/
type ListProjectPresetsParams struct {

	// Disabled.
	Disabled *bool

	// Name.
	Name *string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project presets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectPresetsParams) WithDefaults() *ListProjectPresetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project presets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectPresetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project presets params
func (o *ListProjectPresetsParams) WithTimeout(timeout time.Duration) *ListProjectPresetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project presets params
func (o *ListProjectPresetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project presets params
func (o *ListProjectPresetsParams) WithContext(ctx context.Context) *ListProjectPresetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project presets params
func (o *ListProjectPresetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project presets params
func (o *ListProjectPresetsParams) WithHTTPClient(client *http.Client) *ListProjectPresetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project presets params
func (o *ListProjectPresetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisabled adds the disabled to the list project presets params
func (o *ListProjectPresetsParams) WithDisabled(disabled *bool) *ListProjectPresetsParams {
	o.SetDisabled(disabled)
	return o
}

// SetDisabled adds the disabled to the list project presets params
func (o *ListProjectPresetsParams) SetDisabled(disabled *bool) {
	o.Disabled = disabled
}

// WithName adds the name to the list project presets params
func (o *ListProjectPresetsParams) WithName(name *string) *ListProjectPresetsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the list project presets params
func (o *ListProjectPresetsParams) SetName(name *string) {
	o.Name = name
}

// WithProjectID adds the projectID to the list project presets params
func (o *ListProjectPresetsParams) WithProjectID(projectID string) *ListProjectPresetsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project presets params
func (o *ListProjectPresetsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectPresetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Disabled != nil {

		// query param disabled
		var qrDisabled bool

		if o.Disabled != nil {
			qrDisabled = *o.Disabled
		}
		qDisabled := swag.FormatBool(qrDisabled)
		if qDisabled != "" {

			if err := r.SetQueryParam("disabled", qDisabled); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
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
