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

// RoleBinding RoleBinding references a role, but does not contain it.
//
// swagger:model RoleBinding
type RoleBinding struct {

	// Indicates the scope of this binding.
	Namespace string `json:"namespace,omitempty"`

	// role ref name
	RoleRefName string `json:"roleRefName,omitempty"`

	// Subjects holds references to the objects the role applies to.
	Subjects []*Subject `json:"subjects"`
}

// Validate validates this role binding
func (m *RoleBinding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleBinding) validateSubjects(formats strfmt.Registry) error {
	if swag.IsZero(m.Subjects) { // not required
		return nil
	}

	for i := 0; i < len(m.Subjects); i++ {
		if swag.IsZero(m.Subjects[i]) { // not required
			continue
		}

		if m.Subjects[i] != nil {
			if err := m.Subjects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this role binding based on the context it is used
func (m *RoleBinding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoleBinding) contextValidateSubjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Subjects); i++ {

		if m.Subjects[i] != nil {

			if swag.IsZero(m.Subjects[i]) { // not required
				return nil
			}

			if err := m.Subjects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subjects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subjects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoleBinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoleBinding) UnmarshalBinary(b []byte) error {
	var res RoleBinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
