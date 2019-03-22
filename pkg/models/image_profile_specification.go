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

// ImageProfileSpecification Specification for image profile.
// swagger:model ImageProfileSpecification
type ImageProfileSpecification struct {

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// Image mapping defined for the corresponding region.
	// Required: true
	ImageMapping map[string]FabricImageDescription `json:"imageMapping"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// The id of the region for which this profile is created
	// Required: true
	RegionID *string `json:"regionId"`
}

// Validate validates this image profile specification
func (m *ImageProfileSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImageMapping(formats); err != nil {
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

func (m *ImageProfileSpecification) validateImageMapping(formats strfmt.Registry) error {

	for k := range m.ImageMapping {

		if err := validate.Required("imageMapping"+"."+k, "body", m.ImageMapping[k]); err != nil {
			return err
		}
		if val, ok := m.ImageMapping[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ImageProfileSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ImageProfileSpecification) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("regionId", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageProfileSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageProfileSpecification) UnmarshalBinary(b []byte) error {
	var res ImageProfileSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
