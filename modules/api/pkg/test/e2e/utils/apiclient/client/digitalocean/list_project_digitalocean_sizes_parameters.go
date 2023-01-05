// Code generated by go-swagger; DO NOT EDIT.

package digitalocean

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

// NewListProjectDigitaloceanSizesParams creates a new ListProjectDigitaloceanSizesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectDigitaloceanSizesParams() *ListProjectDigitaloceanSizesParams {
	return &ListProjectDigitaloceanSizesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectDigitaloceanSizesParamsWithTimeout creates a new ListProjectDigitaloceanSizesParams object
// with the ability to set a timeout on a request.
func NewListProjectDigitaloceanSizesParamsWithTimeout(timeout time.Duration) *ListProjectDigitaloceanSizesParams {
	return &ListProjectDigitaloceanSizesParams{
		timeout: timeout,
	}
}

// NewListProjectDigitaloceanSizesParamsWithContext creates a new ListProjectDigitaloceanSizesParams object
// with the ability to set a context for a request.
func NewListProjectDigitaloceanSizesParamsWithContext(ctx context.Context) *ListProjectDigitaloceanSizesParams {
	return &ListProjectDigitaloceanSizesParams{
		Context: ctx,
	}
}

// NewListProjectDigitaloceanSizesParamsWithHTTPClient creates a new ListProjectDigitaloceanSizesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectDigitaloceanSizesParamsWithHTTPClient(client *http.Client) *ListProjectDigitaloceanSizesParams {
	return &ListProjectDigitaloceanSizesParams{
		HTTPClient: client,
	}
}

/*
ListProjectDigitaloceanSizesParams contains all the parameters to send to the API endpoint

	for the list project digitalocean sizes operation.

	Typically these are written to a http.Request.
*/
type ListProjectDigitaloceanSizesParams struct {

	// Credential.
	Credential *string

	// DatacenterName.
	DatacenterName *string

	// DoToken.
	DoToken *string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project digitalocean sizes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectDigitaloceanSizesParams) WithDefaults() *ListProjectDigitaloceanSizesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project digitalocean sizes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectDigitaloceanSizesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) WithTimeout(timeout time.Duration) *ListProjectDigitaloceanSizesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) WithContext(ctx context.Context) *ListProjectDigitaloceanSizesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) WithHTTPClient(client *http.Client) *ListProjectDigitaloceanSizesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) WithCredential(credential *string) *ListProjectDigitaloceanSizesParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDatacenterName adds the datacenterName to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) WithDatacenterName(datacenterName *string) *ListProjectDigitaloceanSizesParams {
	o.SetDatacenterName(datacenterName)
	return o
}

// SetDatacenterName adds the datacenterName to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) SetDatacenterName(datacenterName *string) {
	o.DatacenterName = datacenterName
}

// WithDoToken adds the doToken to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) WithDoToken(doToken *string) *ListProjectDigitaloceanSizesParams {
	o.SetDoToken(doToken)
	return o
}

// SetDoToken adds the doToken to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) SetDoToken(doToken *string) {
	o.DoToken = doToken
}

// WithProjectID adds the projectID to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) WithProjectID(projectID string) *ListProjectDigitaloceanSizesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project digitalocean sizes params
func (o *ListProjectDigitaloceanSizesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectDigitaloceanSizesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.DatacenterName != nil {

		// header param DatacenterName
		if err := r.SetHeaderParam("DatacenterName", *o.DatacenterName); err != nil {
			return err
		}
	}

	if o.DoToken != nil {

		// header param DoToken
		if err := r.SetHeaderParam("DoToken", *o.DoToken); err != nil {
			return err
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