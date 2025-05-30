// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubeLBSeedSettingsAPI kube l b seed settings API
//
// swagger:model KubeLBSeedSettingsAPI
type KubeLBSeedSettingsAPI struct {

	// enable for all datacenters
	EnableForAllDatacenters bool `json:"enableForAllDatacenters,omitempty"`
}

// Validate validates this kube l b seed settings API
func (m *KubeLBSeedSettingsAPI) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this kube l b seed settings API based on context it is used
func (m *KubeLBSeedSettingsAPI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubeLBSeedSettingsAPI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeLBSeedSettingsAPI) UnmarshalBinary(b []byte) error {
	var res KubeLBSeedSettingsAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
