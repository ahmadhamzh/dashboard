// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Names names
//
// swagger:model Names
type Names struct {

	// kind
	Kind string `json:"kind,omitempty"`

	// short names
	ShortNames []string `json:"shortNames"`
}

// Validate validates this names
func (m *Names) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this names based on context it is used
func (m *Names) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Names) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Names) UnmarshalBinary(b []byte) error {
	var res Names
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}