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

// BlueprintDeploymentResourcesQueryResult BlueprintDeploymentResourcesQueryResult
// swagger:model BlueprintDeploymentResourcesQueryResult
type BlueprintDeploymentResourcesQueryResult struct {

	// Count
	// Read Only: true
	Count int64 `json:"count,omitempty"`

	// List of blueprint deployment resources
	// Read Only: true
	Objects []*DeploymentResource `json:"objects"`
}

// Validate validates this blueprint deployment resources query result
func (m *BlueprintDeploymentResourcesQueryResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlueprintDeploymentResourcesQueryResult) validateObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.Objects) { // not required
		return nil
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlueprintDeploymentResourcesQueryResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlueprintDeploymentResourcesQueryResult) UnmarshalBinary(b []byte) error {
	var res BlueprintDeploymentResourcesQueryResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
