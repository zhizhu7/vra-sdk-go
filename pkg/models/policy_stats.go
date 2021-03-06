// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PolicyStats PolicyStats
// swagger:model PolicyStats
type PolicyStats struct {

	// conflict count
	ConflictCount int64 `json:"conflictCount,omitempty"`

	// enforced count
	EnforcedCount int64 `json:"enforcedCount,omitempty"`

	// not enforced count
	NotEnforcedCount int64 `json:"notEnforcedCount,omitempty"`
}

// Validate validates this policy stats
func (m *PolicyStats) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyStats) UnmarshalBinary(b []byte) error {
	var res PolicyStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
