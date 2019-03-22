// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CloudAccountVsphereResult State object representing a query result of vSphere cloud accounts.
// swagger:model CloudAccountVsphereResult
type CloudAccountVsphereResult struct {

	// List of content items
	// Read Only: true
	Content []*CloudAccountVsphere `json:"content"`

	// Total number of elements
	// Read Only: true
	TotalElements int64 `json:"totalElements,omitempty"`
}

// Validate validates this cloud account vsphere result
func (m *CloudAccountVsphereResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountVsphereResult) validateContent(formats strfmt.Registry) error {

	if swag.IsZero(m.Content) { // not required
		return nil
	}

	for i := 0; i < len(m.Content); i++ {
		if swag.IsZero(m.Content[i]) { // not required
			continue
		}

		if m.Content[i] != nil {
			if err := m.Content[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudAccountVsphereResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountVsphereResult) UnmarshalBinary(b []byte) error {
	var res CloudAccountVsphereResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
