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

// ExternalClusterMachineDeploymentCloudSpec ExternalClusterMachineDeploymentCloudSpec represents an object holding machine deployment cloud details.
//
// swagger:model ExternalClusterMachineDeploymentCloudSpec
type ExternalClusterMachineDeploymentCloudSpec struct {

	// aks
	Aks *AKSMachineDeploymentCloudSpec `json:"aks,omitempty"`

	// eks
	Eks *EKSMachineDeploymentCloudSpec `json:"eks,omitempty"`

	// gke
	Gke *GKEMachineDeploymentCloudSpec `json:"gke,omitempty"`
}

// Validate validates this external cluster machine deployment cloud spec
func (m *ExternalClusterMachineDeploymentCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGke(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalClusterMachineDeploymentCloudSpec) validateAks(formats strfmt.Registry) error {
	if swag.IsZero(m.Aks) { // not required
		return nil
	}

	if m.Aks != nil {
		if err := m.Aks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aks")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalClusterMachineDeploymentCloudSpec) validateEks(formats strfmt.Registry) error {
	if swag.IsZero(m.Eks) { // not required
		return nil
	}

	if m.Eks != nil {
		if err := m.Eks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eks")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalClusterMachineDeploymentCloudSpec) validateGke(formats strfmt.Registry) error {
	if swag.IsZero(m.Gke) { // not required
		return nil
	}

	if m.Gke != nil {
		if err := m.Gke.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gke")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gke")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this external cluster machine deployment cloud spec based on the context it is used
func (m *ExternalClusterMachineDeploymentCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGke(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalClusterMachineDeploymentCloudSpec) contextValidateAks(ctx context.Context, formats strfmt.Registry) error {

	if m.Aks != nil {

		if swag.IsZero(m.Aks) { // not required
			return nil
		}

		if err := m.Aks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aks")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalClusterMachineDeploymentCloudSpec) contextValidateEks(ctx context.Context, formats strfmt.Registry) error {

	if m.Eks != nil {

		if swag.IsZero(m.Eks) { // not required
			return nil
		}

		if err := m.Eks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eks")
			}
			return err
		}
	}

	return nil
}

func (m *ExternalClusterMachineDeploymentCloudSpec) contextValidateGke(ctx context.Context, formats strfmt.Registry) error {

	if m.Gke != nil {

		if swag.IsZero(m.Gke) { // not required
			return nil
		}

		if err := m.Gke.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gke")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gke")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternalClusterMachineDeploymentCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalClusterMachineDeploymentCloudSpec) UnmarshalBinary(b []byte) error {
	var res ExternalClusterMachineDeploymentCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
