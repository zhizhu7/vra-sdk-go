// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResourceActionRequest ResourceActionRequest
//
// Represents a resource day-two request
// swagger:model ResourceActionRequest
type ResourceActionRequest struct {

	// The id of the action to perform.
	// Format: uuid
	ActionID strfmt.UUID `json:"actionId,omitempty"`

	// Resource action request inputs
	Inputs interface{} `json:"inputs,omitempty"`

	// Reason for requesting a day2 operation
	Reason string `json:"reason,omitempty"`
}

// Validate validates this resource action request
func (m *ResourceActionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceActionRequest) validateActionID(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionID) { // not required
		return nil
	}

	if err := validate.FormatOf("actionId", "body", "uuid", m.ActionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceActionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceActionRequest) UnmarshalBinary(b []byte) error {
	var res ResourceActionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
