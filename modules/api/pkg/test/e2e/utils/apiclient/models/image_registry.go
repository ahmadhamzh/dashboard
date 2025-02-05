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

// ImageRegistry ImageRegistry defines requests to an OCI/Docker V2 registry to fetch image
// details.
//
// swagger:model ImageRegistry
type ImageRegistry struct {

	// JMESPath is an optional JSON Match Expression that can be used to
	// transform the ImageData struct returned as a result of processing
	// the image reference.
	// +optional
	JMESPath string `json:"jmesPath,omitempty"`

	// Reference is image reference to a container image in the registry.
	// Example: ghcr.io/kyverno/kyverno:latest
	Reference string `json:"reference,omitempty"`

	// image registry credentials
	ImageRegistryCredentials *ImageRegistryCredentials `json:"imageRegistryCredentials,omitempty"`
}

// Validate validates this image registry
func (m *ImageRegistry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImageRegistryCredentials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageRegistry) validateImageRegistryCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageRegistryCredentials) { // not required
		return nil
	}

	if m.ImageRegistryCredentials != nil {
		if err := m.ImageRegistryCredentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageRegistryCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageRegistryCredentials")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this image registry based on the context it is used
func (m *ImageRegistry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImageRegistryCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageRegistry) contextValidateImageRegistryCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.ImageRegistryCredentials != nil {

		if swag.IsZero(m.ImageRegistryCredentials) { // not required
			return nil
		}

		if err := m.ImageRegistryCredentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageRegistryCredentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("imageRegistryCredentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageRegistry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageRegistry) UnmarshalBinary(b []byte) error {
	var res ImageRegistry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
