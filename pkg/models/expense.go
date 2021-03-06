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

// Expense Expense
// swagger:model Expense
type Expense struct {

	// Additional expense incurred for the resource.
	// Read Only: true
	AdditionalExpense float64 `json:"additionalExpense,omitempty"`

	// Expense sync message code if any.
	// Read Only: true
	Code string `json:"code,omitempty"`

	// Compute expense of the resource.
	// Read Only: true
	ComputeExpense float64 `json:"computeExpense,omitempty"`

	// Last expense sync time.
	// Read Only: true
	// Format: date-time
	LastUpdatedTime strfmt.DateTime `json:"lastUpdatedTime,omitempty"`

	// Expense sync message if any.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Network expense of the resource.
	// Read Only: true
	NetworkExpense float64 `json:"networkExpense,omitempty"`

	// Storage expense of the resource.
	// Read Only: true
	StorageExpense float64 `json:"storageExpense,omitempty"`

	// Total expense of the resource.
	// Read Only: true
	TotalExpense float64 `json:"totalExpense,omitempty"`

	// Monetary unit.
	// Read Only: true
	Unit string `json:"unit,omitempty"`
}

// Validate validates this expense
func (m *Expense) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastUpdatedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Expense) validateLastUpdatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedTime", "body", "date-time", m.LastUpdatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Expense) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Expense) UnmarshalBinary(b []byte) error {
	var res Expense
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
