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

// PatchPolicyBindingBody patch policy binding body
//
// swagger:model patchPolicyBindingBody
type PatchPolicyBindingBody struct {

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// project ID
	ProjectID string `json:"projectID,omitempty"`

	// spec
	Spec *PolicyBindingSpec `json:"Spec,omitempty"`
}

// Validate validates this patch policy binding body
func (m *PatchPolicyBindingBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchPolicyBindingBody) validateSpec(formats strfmt.Registry) error {
	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Spec")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this patch policy binding body based on the context it is used
func (m *PatchPolicyBindingBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSpec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchPolicyBindingBody) contextValidateSpec(ctx context.Context, formats strfmt.Registry) error {

	if m.Spec != nil {

		if swag.IsZero(m.Spec) { // not required
			return nil
		}

		if err := m.Spec.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Spec")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Spec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchPolicyBindingBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchPolicyBindingBody) UnmarshalBinary(b []byte) error {
	var res PatchPolicyBindingBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
