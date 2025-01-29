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

// JSONSchemaPropsOrStringArray JSONSchemaPropsOrStringArray represents a JSONSchemaProps or a string array.
//
// swagger:model JSONSchemaPropsOrStringArray
type JSONSchemaPropsOrStringArray struct {

	// +listType=atomic
	Property []string `json:"Property"`

	// schema
	Schema *JSONSchemaProps `json:"Schema,omitempty"`
}

// Validate validates this JSON schema props or string array
func (m *JSONSchemaPropsOrStringArray) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JSONSchemaPropsOrStringArray) validateSchema(formats strfmt.Registry) error {
	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Schema")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this JSON schema props or string array based on the context it is used
func (m *JSONSchemaPropsOrStringArray) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JSONSchemaPropsOrStringArray) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if m.Schema != nil {

		if swag.IsZero(m.Schema) { // not required
			return nil
		}

		if err := m.Schema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Schema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Schema")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JSONSchemaPropsOrStringArray) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONSchemaPropsOrStringArray) UnmarshalBinary(b []byte) error {
	var res JSONSchemaPropsOrStringArray
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
