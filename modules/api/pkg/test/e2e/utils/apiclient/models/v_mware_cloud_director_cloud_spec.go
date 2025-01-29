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

// VMwareCloudDirectorCloudSpec VMwareCloudDirectorCloudSpec specifies access data to VMware Cloud Director cloud.
//
// swagger:model VMwareCloudDirectorCloudSpec
type VMwareCloudDirectorCloudSpec struct {

	// The VMware Cloud Director API token.
	// +optional
	APIToken string `json:"apiToken,omitempty"`

	// The name of organizational virtual data center network that will be associated with the VMs and vApp.
	// Deprecated: OVDCNetwork has been deprecated starting with KKP 2.25 and will be removed in KKP 2.27+. It is recommended to use OVDCNetworks instead.
	OVDCNetwork string `json:"ovdcNetwork,omitempty"`

	// OVDCNetworks is the list of organizational virtual data center networks that will be attached to the vApp and can be consumed the VMs.
	OVDCNetworks []string `json:"ovdcNetworks"`

	// The name of organization to use.
	// +optional
	Organization string `json:"organization,omitempty"`

	// The VMware Cloud Director user password.
	// +optional
	Password string `json:"password,omitempty"`

	// The VMware Cloud Director user name.
	// +optional
	Username string `json:"username,omitempty"`

	// VApp used for isolation of VMs and their associated network
	// +optional
	VApp string `json:"vapp,omitempty"`

	// The organizational virtual data center.
	// +optional
	VDC string `json:"vdc,omitempty"`

	// credentials reference
	CredentialsReference *GlobalSecretKeySelector `json:"credentialsReference,omitempty"`

	// csi
	Csi *VMwareCloudDirectorCSIConfig `json:"csi,omitempty"`
}

// Validate validates this v mware cloud director cloud spec
func (m *VMwareCloudDirectorCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialsReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsi(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMwareCloudDirectorCloudSpec) validateCredentialsReference(formats strfmt.Registry) error {
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

func (m *VMwareCloudDirectorCloudSpec) validateCsi(formats strfmt.Registry) error {
	if swag.IsZero(m.Csi) { // not required
		return nil
	}

	if m.Csi != nil {
		if err := m.Csi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csi")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v mware cloud director cloud spec based on the context it is used
func (m *VMwareCloudDirectorCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentialsReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCsi(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMwareCloudDirectorCloudSpec) contextValidateCredentialsReference(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMwareCloudDirectorCloudSpec) contextValidateCsi(ctx context.Context, formats strfmt.Registry) error {

	if m.Csi != nil {

		if swag.IsZero(m.Csi) { // not required
			return nil
		}

		if err := m.Csi.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csi")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMwareCloudDirectorCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMwareCloudDirectorCloudSpec) UnmarshalBinary(b []byte) error {
	var res VMwareCloudDirectorCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
