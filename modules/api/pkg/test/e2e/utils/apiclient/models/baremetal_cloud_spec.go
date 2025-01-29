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

// BaremetalCloudSpec BaremetalCloudSpec specifies access data for a baremetal cluster.
//
// swagger:model BaremetalCloudSpec
type BaremetalCloudSpec struct {

	// credentials reference
	CredentialsReference *GlobalSecretKeySelector `json:"credentialsReference,omitempty"`

	// tinkerbell
	Tinkerbell *TinkerbellCloudSpec `json:"tinkerbell,omitempty"`
}

// Validate validates this baremetal cloud spec
func (m *BaremetalCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialsReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTinkerbell(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaremetalCloudSpec) validateCredentialsReference(formats strfmt.Registry) error {
	if swag.IsZero(m.CredentialsReference) { // not required
		return nil
	}

	if m.CredentialsReference != nil {
		if err := m.CredentialsReference.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentialsReference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentialsReference")
			}
			return err
		}
	}

	return nil
}

func (m *BaremetalCloudSpec) validateTinkerbell(formats strfmt.Registry) error {
	if swag.IsZero(m.Tinkerbell) { // not required
		return nil
	}

	if m.Tinkerbell != nil {
		if err := m.Tinkerbell.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tinkerbell")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tinkerbell")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this baremetal cloud spec based on the context it is used
func (m *BaremetalCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentialsReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTinkerbell(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaremetalCloudSpec) contextValidateCredentialsReference(ctx context.Context, formats strfmt.Registry) error {

	if m.CredentialsReference != nil {

		if swag.IsZero(m.CredentialsReference) { // not required
			return nil
		}

		if err := m.CredentialsReference.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credentialsReference")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("credentialsReference")
			}
			return err
		}
	}

	return nil
}

func (m *BaremetalCloudSpec) contextValidateTinkerbell(ctx context.Context, formats strfmt.Registry) error {

	if m.Tinkerbell != nil {

		if swag.IsZero(m.Tinkerbell) { // not required
			return nil
		}

		if err := m.Tinkerbell.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tinkerbell")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tinkerbell")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BaremetalCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaremetalCloudSpec) UnmarshalBinary(b []byte) error {
	var res BaremetalCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
