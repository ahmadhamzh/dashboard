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

// DatacenterSpecNutanix DatacenterSpecNutanix describes a Nutanix datacenter.
//
// swagger:model DatacenterSpecNutanix
type DatacenterSpecNutanix struct {

	// Optional: AllowInsecure allows to disable the TLS certificate check against the endpoint (defaults to false)
	AllowInsecure bool `json:"allowInsecure,omitempty"`

	// Endpoint to use for accessing Nutanix Prism Central. No protocol or port should be passed,
	// for example "nutanix.example.com" or "10.0.0.1"
	Endpoint string `json:"endpoint,omitempty"`

	// Optional: Port to use when connecting to the Nutanix Prism Central endpoint (defaults to 9440)
	Port int32 `json:"port,omitempty"`

	// images
	Images ImageList `json:"images,omitempty"`
}

// Validate validates this datacenter spec nutanix
func (m *DatacenterSpecNutanix) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecNutanix) validateImages(formats strfmt.Registry) error {
	if swag.IsZero(m.Images) { // not required
		return nil
	}

	if m.Images != nil {
		if err := m.Images.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("images")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("images")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this datacenter spec nutanix based on the context it is used
func (m *DatacenterSpecNutanix) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecNutanix) contextValidateImages(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Images.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("images")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("images")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatacenterSpecNutanix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatacenterSpecNutanix) UnmarshalBinary(b []byte) error {
	var res DatacenterSpecNutanix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
