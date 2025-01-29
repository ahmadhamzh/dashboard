// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyTemplate PolicyTemplate defines a reusable blueprint of a Kyverno policy.
//
// swagger:model PolicyTemplate
type PolicyTemplate struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this policy template
func (m *PolicyTemplate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this policy template based on context it is used
func (m *PolicyTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyTemplate) UnmarshalBinary(b []byte) error {
	var res PolicyTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
