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

// AlibabaCloudSpec AlibabaCloudSpec specifies the access data to Alibaba.
//
// swagger:model AlibabaCloudSpec
type AlibabaCloudSpec struct {

	// The Access Key ID used to authenticate against Alibaba.
	AccessKeyID string `json:"accessKeyID,omitempty"`

	// The Access Key Secret used to authenticate against Alibaba.
	AccessKeySecret string `json:"accessKeySecret,omitempty"`

	// credentials reference
	CredentialsReference *GlobalSecretKeySelector `json:"credentialsReference,omitempty"`
}

// Validate validates this alibaba cloud spec
func (m *AlibabaCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialsReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlibabaCloudSpec) validateCredentialsReference(formats strfmt.Registry) error {
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

// ContextValidate validate this alibaba cloud spec based on the context it is used
func (m *AlibabaCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentialsReference(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlibabaCloudSpec) contextValidateCredentialsReference(ctx context.Context, formats strfmt.Registry) error {

	if m.CredentialsReference != nil {
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

// MarshalBinary interface implementation
func (m *AlibabaCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlibabaCloudSpec) UnmarshalBinary(b []byte) error {
	var res AlibabaCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
