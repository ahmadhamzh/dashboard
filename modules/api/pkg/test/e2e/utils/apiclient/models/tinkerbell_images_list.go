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

// TinkerbellImagesList TinkerbellImagesList represents list of available Tinkerbell images with their categories.
//
// swagger:model TinkerbellImagesList
type TinkerbellImagesList struct {

	// standard
	Standard *TinkerbellImages `json:"standard,omitempty"`
}

// Validate validates this tinkerbell images list
func (m *TinkerbellImagesList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStandard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TinkerbellImagesList) validateStandard(formats strfmt.Registry) error {
	if swag.IsZero(m.Standard) { // not required
		return nil
	}

	if m.Standard != nil {
		if err := m.Standard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standard")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tinkerbell images list based on the context it is used
func (m *TinkerbellImagesList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStandard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TinkerbellImagesList) contextValidateStandard(ctx context.Context, formats strfmt.Registry) error {

	if m.Standard != nil {

		if swag.IsZero(m.Standard) { // not required
			return nil
		}

		if err := m.Standard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("standard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("standard")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TinkerbellImagesList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TinkerbellImagesList) UnmarshalBinary(b []byte) error {
	var res TinkerbellImagesList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
