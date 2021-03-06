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

// FlavorProfileSpecification Specification for flavor profile
// swagger:model FlavorProfileSpecification
type FlavorProfileSpecification struct {

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Map between global fabric flavor keys <String> and fabric flavor descriptions <FabricFlavorDescription>
	// Required: true
	FlavorMapping map[string]FabricFlavorDescription `json:"flavorMapping"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// The id of the region for which this profile is created
	// Required: true
	RegionID *string `json:"regionId"`
}

// Validate validates this flavor profile specification
func (m *FlavorProfileSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlavorMapping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlavorProfileSpecification) validateFlavorMapping(formats strfmt.Registry) error {

	for k := range m.FlavorMapping {

		if err := validate.Required("flavorMapping"+"."+k, "body", m.FlavorMapping[k]); err != nil {
			return err
		}
		if val, ok := m.FlavorMapping[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *FlavorProfileSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FlavorProfileSpecification) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("regionId", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlavorProfileSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlavorProfileSpecification) UnmarshalBinary(b []byte) error {
	var res FlavorProfileSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
