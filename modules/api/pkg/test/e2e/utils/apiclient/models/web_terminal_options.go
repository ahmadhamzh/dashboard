// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebTerminalOptions web terminal options
//
// swagger:model WebTerminalOptions
type WebTerminalOptions struct {

	// AdditionalEnvironmentVariables are the additional environment variables that can be set for the Web Terminal.
	AdditionalEnvironmentVariables []*EnvVar `json:"additionalEnvironmentVariables"`

	// EnableInternetAccess enables the Web Terminal feature to access the internet.
	EnableInternetAccess bool `json:"enableInternetAccess,omitempty"`

	// Enabled enables the Web Terminal feature for the user clusters.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this web terminal options
func (m *WebTerminalOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalEnvironmentVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebTerminalOptions) validateAdditionalEnvironmentVariables(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalEnvironmentVariables) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalEnvironmentVariables); i++ {
		if swag.IsZero(m.AdditionalEnvironmentVariables[i]) { // not required
			continue
		}

		if m.AdditionalEnvironmentVariables[i] != nil {
			if err := m.AdditionalEnvironmentVariables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalEnvironmentVariables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalEnvironmentVariables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this web terminal options based on the context it is used
func (m *WebTerminalOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalEnvironmentVariables(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebTerminalOptions) contextValidateAdditionalEnvironmentVariables(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalEnvironmentVariables); i++ {

		if m.AdditionalEnvironmentVariables[i] != nil {

			if swag.IsZero(m.AdditionalEnvironmentVariables[i]) { // not required
				return nil
			}

			if err := m.AdditionalEnvironmentVariables[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additionalEnvironmentVariables" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additionalEnvironmentVariables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebTerminalOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebTerminalOptions) UnmarshalBinary(b []byte) error {
	var res WebTerminalOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
