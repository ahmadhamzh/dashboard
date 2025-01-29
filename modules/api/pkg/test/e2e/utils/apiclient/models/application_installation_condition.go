// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ApplicationInstallationCondition application installation condition
//
// swagger:model ApplicationInstallationCondition
type ApplicationInstallationCondition struct {

	// Last time we got an update on a given condition.
	// +optional
	// Format: date-time
	LastHeartbeatTime strfmt.DateTime `json:"lastHeartbeatTime,omitempty"`

	// Last time the condition transit from one status to another.
	// +optional
	// Format: date-time
	LastTransitionTime strfmt.DateTime `json:"lastTransitionTime,omitempty"`

	// Human readable message indicating details about last transition.
	Message string `json:"message,omitempty"`

	// (brief) reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// Type of ApplicationInstallation condition.
	// ManifestsRetrieved ManifestsRetrieved  ManifestsRetrieved indicates all necessary manifests have been fetched from the external source.
	// Ready Ready  Ready describes all components have been successfully rolled out and are ready.
	// Enum: ["ManifestsRetrieved","Ready"]
	Type string `json:"type,omitempty"`

	// status
	Status ConditionStatus `json:"status,omitempty"`
}

// Validate validates this application installation condition
func (m *ApplicationInstallationCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastHeartbeatTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastTransitionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationInstallationCondition) validateLastHeartbeatTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastHeartbeatTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastHeartbeatTime", "body", "date-time", m.LastHeartbeatTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationInstallationCondition) validateLastTransitionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.LastTransitionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastTransitionTime", "body", "date-time", m.LastTransitionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var applicationInstallationConditionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ManifestsRetrieved","Ready"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		applicationInstallationConditionTypeTypePropEnum = append(applicationInstallationConditionTypeTypePropEnum, v)
	}
}

const (

	// ApplicationInstallationConditionTypeManifestsRetrieved captures enum value "ManifestsRetrieved"
	ApplicationInstallationConditionTypeManifestsRetrieved string = "ManifestsRetrieved"

	// ApplicationInstallationConditionTypeReady captures enum value "Ready"
	ApplicationInstallationConditionTypeReady string = "Ready"
)

// prop value enum
func (m *ApplicationInstallationCondition) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, applicationInstallationConditionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApplicationInstallationCondition) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ApplicationInstallationCondition) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this application installation condition based on the context it is used
func (m *ApplicationInstallationCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationInstallationCondition) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationInstallationCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationInstallationCondition) UnmarshalBinary(b []byte) error {
	var res ApplicationInstallationCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
