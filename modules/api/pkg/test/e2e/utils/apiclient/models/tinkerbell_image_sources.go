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

// TinkerbellImageSources TinkerbellImageSources represents Operating System image sources for Tinkerbell.
//
// swagger:model TinkerbellImageSources
type TinkerbellImageSources struct {

	// http
	HTTP *TinkerbellHTTPSource `json:"http,omitempty"`
}

// Validate validates this tinkerbell image sources
func (m *TinkerbellImageSources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTTP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TinkerbellImageSources) validateHTTP(formats strfmt.Registry) error {
	if swag.IsZero(m.HTTP) { // not required
		return nil
	}

	if m.HTTP != nil {
		if err := m.HTTP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tinkerbell image sources based on the context it is used
func (m *TinkerbellImageSources) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHTTP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TinkerbellImageSources) contextValidateHTTP(ctx context.Context, formats strfmt.Registry) error {

	if m.HTTP != nil {

		if swag.IsZero(m.HTTP) { // not required
			return nil
		}

		if err := m.HTTP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("http")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("http")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TinkerbellImageSources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TinkerbellImageSources) UnmarshalBinary(b []byte) error {
	var res TinkerbellImageSources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
