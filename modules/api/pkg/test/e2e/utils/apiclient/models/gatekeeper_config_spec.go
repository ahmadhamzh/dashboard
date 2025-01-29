// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GatekeeperConfigSpec gatekeeper config spec
//
// swagger:model GatekeeperConfigSpec
type GatekeeperConfigSpec struct {

	// Configuration for namespace exclusion
	Match []*MatchEntry `json:"match"`

	// readiness
	Readiness *ReadinessSpec `json:"readiness,omitempty"`

	// sync
	Sync *Sync `json:"sync,omitempty"`

	// validation
	Validation *Validation `json:"validation,omitempty"`
}

// Validate validates this gatekeeper config spec
func (m *GatekeeperConfigSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadiness(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSync(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatekeeperConfigSpec) validateMatch(formats strfmt.Registry) error {
	if swag.IsZero(m.Match) { // not required
		return nil
	}

	for i := 0; i < len(m.Match); i++ {
		if swag.IsZero(m.Match[i]) { // not required
			continue
		}

		if m.Match[i] != nil {
			if err := m.Match[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("match" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("match" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GatekeeperConfigSpec) validateReadiness(formats strfmt.Registry) error {
	if swag.IsZero(m.Readiness) { // not required
		return nil
	}

	if m.Readiness != nil {
		if err := m.Readiness.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readiness")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("readiness")
			}
			return err
		}
	}

	return nil
}

func (m *GatekeeperConfigSpec) validateSync(formats strfmt.Registry) error {
	if swag.IsZero(m.Sync) { // not required
		return nil
	}

	if m.Sync != nil {
		if err := m.Sync.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync")
			}
			return err
		}
	}

	return nil
}

func (m *GatekeeperConfigSpec) validateValidation(formats strfmt.Registry) error {
	if swag.IsZero(m.Validation) { // not required
		return nil
	}

	if m.Validation != nil {
		if err := m.Validation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("validation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gatekeeper config spec based on the context it is used
func (m *GatekeeperConfigSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMatch(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReadiness(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSync(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValidation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatekeeperConfigSpec) contextValidateMatch(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Match); i++ {

		if m.Match[i] != nil {

			if swag.IsZero(m.Match[i]) { // not required
				return nil
			}

			if err := m.Match[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("match" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("match" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GatekeeperConfigSpec) contextValidateReadiness(ctx context.Context, formats strfmt.Registry) error {

	if m.Readiness != nil {

		if swag.IsZero(m.Readiness) { // not required
			return nil
		}

		if err := m.Readiness.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readiness")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("readiness")
			}
			return err
		}
	}

	return nil
}

func (m *GatekeeperConfigSpec) contextValidateSync(ctx context.Context, formats strfmt.Registry) error {

	if m.Sync != nil {

		if swag.IsZero(m.Sync) { // not required
			return nil
		}

		if err := m.Sync.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sync")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sync")
			}
			return err
		}
	}

	return nil
}

func (m *GatekeeperConfigSpec) contextValidateValidation(ctx context.Context, formats strfmt.Registry) error {

	if m.Validation != nil {

		if swag.IsZero(m.Validation) { // not required
			return nil
		}

		if err := m.Validation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("validation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatekeeperConfigSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatekeeperConfigSpec) UnmarshalBinary(b []byte) error {
	var res GatekeeperConfigSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
