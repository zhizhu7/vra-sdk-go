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

// DataCollectorRegistration Data collector registration object.<br>The supplied data collector is an OVA tool that contains the credentials and protocols needed to create a connection between a data collector appliance on a host vCenter and a vCenter-based cloud account. . The process of deploying data collector is: <br> 1. Download the data collector ova from the "ovaLink".<br>2. Import the .ova file to the vCenter Server and start the installation.<br> 3. When asked for the key, copy and use the "key" provided.<br> 4. It takes a few minutes to detect your Data Collector after it is deployed and powered on in vCenter.
// swagger:model DataCollectorRegistration
type DataCollectorRegistration struct {

	// A registration key for the data collector
	// Required: true
	Key *string `json:"key"`

	// Data collector OVA Link
	// Required: true
	OvaLink *string `json:"ovaLink"`
}

// Validate validates this data collector registration
func (m *DataCollectorRegistration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOvaLink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataCollectorRegistration) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

func (m *DataCollectorRegistration) validateOvaLink(formats strfmt.Registry) error {

	if err := validate.Required("ovaLink", "body", m.OvaLink); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataCollectorRegistration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataCollectorRegistration) UnmarshalBinary(b []byte) error {
	var res DataCollectorRegistration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
