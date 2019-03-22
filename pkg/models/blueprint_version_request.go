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

// BlueprintVersionRequest BlueprintVersionRequest
// swagger:model BlueprintVersionRequest
type BlueprintVersionRequest struct {

	// Blueprint version change log
	ChangeLog string `json:"changeLog,omitempty"`

	// Blueprint version description
	Description string `json:"description,omitempty"`

	// Blueprint version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this blueprint version request
func (m *BlueprintVersionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlueprintVersionRequest) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlueprintVersionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlueprintVersionRequest) UnmarshalBinary(b []byte) error {
	var res BlueprintVersionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
