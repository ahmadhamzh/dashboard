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

// GCPCloudSpec GCPCloudSpec specifies access data to GCP.
//
// swagger:model GCPCloudSpec
type GCPCloudSpec struct {

	// network
	Network string `json:"network,omitempty"`

	// A CIDR range that will be used to allow access to the node port range in the firewall rules to.
	// If NodePortsAllowedIPRange nor NodePortsAllowedIPRanges is set, the node port range can be accessed from anywhere.
	NodePortsAllowedIPRange string `json:"nodePortsAllowedIPRange,omitempty"`

	// The Google Service Account (JSON format), encoded with base64.
	ServiceAccount string `json:"serviceAccount,omitempty"`

	// subnetwork
	Subnetwork string `json:"subnetwork,omitempty"`

	// credentials reference
	CredentialsReference *GlobalSecretKeySelector `json:"credentialsReference,omitempty"`

	// node ports allowed IP ranges
	NodePortsAllowedIPRanges *NetworkRanges `json:"nodePortsAllowedIPRanges,omitempty"`
}

// Validate validates this g c p cloud spec
func (m *GCPCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialsReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodePortsAllowedIPRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GCPCloudSpec) validateCredentialsReference(formats strfmt.Registry) error {
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

func (m *GCPCloudSpec) validateNodePortsAllowedIPRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.NodePortsAllowedIPRanges) { // not required
		return nil
	}

	if m.NodePortsAllowedIPRanges != nil {
		if err := m.NodePortsAllowedIPRanges.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodePortsAllowedIPRanges")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodePortsAllowedIPRanges")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this g c p cloud spec based on the context it is used
func (m *GCPCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentialsReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodePortsAllowedIPRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GCPCloudSpec) contextValidateCredentialsReference(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GCPCloudSpec) contextValidateNodePortsAllowedIPRanges(ctx context.Context, formats strfmt.Registry) error {

	if m.NodePortsAllowedIPRanges != nil {

		if swag.IsZero(m.NodePortsAllowedIPRanges) { // not required
			return nil
		}

		if err := m.NodePortsAllowedIPRanges.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodePortsAllowedIPRanges")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodePortsAllowedIPRanges")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GCPCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GCPCloudSpec) UnmarshalBinary(b []byte) error {
	var res GCPCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
