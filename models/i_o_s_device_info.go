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

// IOSDeviceInfo i o s device info
//
// swagger:model IOSDeviceInfo
type IOSDeviceInfo struct {

	// identifier for vendor
	IdentifierForVendor string `json:"identifier_for_vendor,omitempty"`

	// is physical device
	IsPhysicalDevice string `json:"is_physical_device,omitempty"`

	// localized model
	LocalizedModel string `json:"localized_model,omitempty"`

	// model
	Model string `json:"model,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// system name
	SystemName string `json:"system_name,omitempty"`

	// system version
	SystemVersion string `json:"system_version,omitempty"`

	// utsname
	Utsname *IOSDeviceInfoUtsname `json:"utsname,omitempty"`
}

// Validate validates this i o s device info
func (m *IOSDeviceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUtsname(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IOSDeviceInfo) validateUtsname(formats strfmt.Registry) error {
	if swag.IsZero(m.Utsname) { // not required
		return nil
	}

	if m.Utsname != nil {
		if err := m.Utsname.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("utsname")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this i o s device info based on the context it is used
func (m *IOSDeviceInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUtsname(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IOSDeviceInfo) contextValidateUtsname(ctx context.Context, formats strfmt.Registry) error {

	if m.Utsname != nil {
		if err := m.Utsname.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("utsname")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IOSDeviceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IOSDeviceInfo) UnmarshalBinary(b []byte) error {
	var res IOSDeviceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IOSDeviceInfoUtsname i o s device info utsname
//
// swagger:model IOSDeviceInfoUtsname
type IOSDeviceInfoUtsname struct {

	// machine
	Machine string `json:"machine,omitempty"`

	// nodename
	Nodename string `json:"nodename,omitempty"`

	// release
	Release string `json:"release,omitempty"`

	// sysname
	Sysname string `json:"sysname,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this i o s device info utsname
func (m *IOSDeviceInfoUtsname) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this i o s device info utsname based on context it is used
func (m *IOSDeviceInfoUtsname) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IOSDeviceInfoUtsname) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IOSDeviceInfoUtsname) UnmarshalBinary(b []byte) error {
	var res IOSDeviceInfoUtsname
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}