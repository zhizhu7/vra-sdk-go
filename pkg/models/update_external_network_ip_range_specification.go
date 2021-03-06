// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UpdateExternalNetworkIPRangeSpecification Specification for updating an ExternalNetworkIPRange
// swagger:model UpdateExternalNetworkIPRangeSpecification
type UpdateExternalNetworkIPRangeSpecification struct {

	// The id of the fabric network that this IP range should be associated with.Use "" as value of this field in order to disassociate the IP range from this network.
	FabricNetworkID string `json:"fabricNetworkId,omitempty"`
}

// Validate validates this update external network IP range specification
func (m *UpdateExternalNetworkIPRangeSpecification) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateExternalNetworkIPRangeSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateExternalNetworkIPRangeSpecification) UnmarshalBinary(b []byte) error {
	var res UpdateExternalNetworkIPRangeSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
