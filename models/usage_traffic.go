// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UsageTraffic usage traffic
//
// swagger:model UsageTraffic
type UsageTraffic struct {

	// consumed
	Consumed int64 `json:"consumed,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// limit
	Limit int64 `json:"limit,omitempty"`

	// receive
	Receive int64 `json:"receive,omitempty"`

	// transmit
	Transmit int64 `json:"transmit,omitempty"`
}

// Validate validates this usage traffic
func (m *UsageTraffic) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this usage traffic based on context it is used
func (m *UsageTraffic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UsageTraffic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsageTraffic) UnmarshalBinary(b []byte) error {
	var res UsageTraffic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
