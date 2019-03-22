// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// About State object representing an about page that includes api versioning information
// swagger:model About
type About struct {

	// The latest version of the API in yyyy-MM-dd format (UTC).
	// Required: true
	LatestAPIVersion *string `json:"latestApiVersion"`

	// A collection of all currently supported api versions.
	// Required: true
	SupportedApis []*APIDescription `json:"supportedApis"`
}

// Validate validates this about
func (m *About) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLatestAPIVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportedApis(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *About) validateLatestAPIVersion(formats strfmt.Registry) error {

	if err := validate.Required("latestApiVersion", "body", m.LatestAPIVersion); err != nil {
		return err
	}

	return nil
}

func (m *About) validateSupportedApis(formats strfmt.Registry) error {

	if err := validate.Required("supportedApis", "body", m.SupportedApis); err != nil {
		return err
	}

	for i := 0; i < len(m.SupportedApis); i++ {
		if swag.IsZero(m.SupportedApis[i]) { // not required
			continue
		}

		if m.SupportedApis[i] != nil {
			if err := m.SupportedApis[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supportedApis" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *About) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *About) UnmarshalBinary(b []byte) error {
	var res About
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
