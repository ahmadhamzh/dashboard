// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HelmSource helm source
//
// swagger:model HelmSource
type HelmSource struct {

	// Name of the Chart.
	// +kubebuilder:validation:MinLength=1
	ChartName string `json:"chartName,omitempty"`

	// Version of the Chart.
	// +kubebuilder:validation:MinLength=1
	ChartVersion string `json:"chartVersion,omitempty"`

	// Insecure disables certificate validation when using an HTTPS registry. This setting has no
	// effect when using a plaintext connection.
	Insecure bool `json:"insecure,omitempty"`

	// PlainHTTP will enable HTTP-only (i.e. unencrypted) traffic for oci:// URLs. By default HTTPS
	// is used when communicating with an oci:// URL.
	PlainHTTP bool `json:"plainHTTP,omitempty"`

	// URL of the Helm repository the following schemes are supported:
	//
	// http://example.com/myrepo (HTTP)
	// https://example.com/myrepo (HTTPS)
	// oci://example.com:5000/myrepo (OCI, HTTPS by default, use plainHTTP to enable unencrypted HTTP)
	URL string `json:"url,omitempty"`

	// credentials
	Credentials *HelmCredentials `json:"credentials,omitempty"`
}

// Validate validates this helm source
func (m *HelmSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HelmSource) validateCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.Credentials) { // not required
		return nil
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this helm source based on the context it is used
func (m *HelmSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HelmSource) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {
		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HelmSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HelmSource) UnmarshalBinary(b []byte) error {
	var res HelmSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
