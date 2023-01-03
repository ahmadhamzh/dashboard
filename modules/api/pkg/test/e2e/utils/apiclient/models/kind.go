// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Kind Kind specifies the resource Kind and APIGroup.
//
// swagger:model Kind
type Kind struct {

	// APIGroups specifies the APIGroups of the resources
	APIGroups []string `json:"apiGroups"`

	// Kinds specifies the kinds of the resources
	Kinds []string `json:"kinds"`
}

// Validate validates this kind
func (m *Kind) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kind based on context it is used
func (m *Kind) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Kind) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Kind) UnmarshalBinary(b []byte) error {
	var res Kind
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}