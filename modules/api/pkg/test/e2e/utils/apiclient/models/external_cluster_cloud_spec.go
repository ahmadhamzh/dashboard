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

// ExternalClusterCloudSpec ExternalClusterCloudSpec represents an object holding cluster cloud details
//
// swagger:model ExternalClusterCloudSpec
type ExternalClusterCloudSpec struct {

	// aks
	Aks *AKSCloudSpec `json:"aks,omitempty"`

	// bring your own
	BringYourOwn BringYourOwnSpec `json:"bringYourOwn,omitempty"`

	// eks
	Eks *EKSCloudSpec `json:"eks,omitempty"`

	// gke
	Gke *GKECloudSpec `json:"gke,omitempty"`

	// kube one
	KubeOne *KubeOneSpec `json:"kubeOne,omitempty"`
}

// Validate validates this external cluster cloud spec
func (m *ExternalClusterCloudSpec) Validate(formats strfmt.Registry) error {
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

	if err := m.validateKubeOne(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalClusterCloudSpec) validateAks(formats strfmt.Registry) error {
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

func (m *ExternalClusterCloudSpec) validateEks(formats strfmt.Registry) error {
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

func (m *ExternalClusterCloudSpec) validateGke(formats strfmt.Registry) error {
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

func (m *ExternalClusterCloudSpec) validateKubeOne(formats strfmt.Registry) error {
	if swag.IsZero(m.KubeOne) { // not required
		return nil
	}

	if m.KubeOne != nil {
		if err := m.KubeOne.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubeOne")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubeOne")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this external cluster cloud spec based on the context it is used
func (m *ExternalClusterCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateKubeOne(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternalClusterCloudSpec) contextValidateAks(ctx context.Context, formats strfmt.Registry) error {

	if m.Aks != nil {
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

func (m *ExternalClusterCloudSpec) contextValidateEks(ctx context.Context, formats strfmt.Registry) error {

	if m.Eks != nil {
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

func (m *ExternalClusterCloudSpec) contextValidateGke(ctx context.Context, formats strfmt.Registry) error {

	if m.Gke != nil {
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

func (m *ExternalClusterCloudSpec) contextValidateKubeOne(ctx context.Context, formats strfmt.Registry) error {

	if m.KubeOne != nil {
		if err := m.KubeOne.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubeOne")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubeOne")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternalClusterCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalClusterCloudSpec) UnmarshalBinary(b []byte) error {
	var res ExternalClusterCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
