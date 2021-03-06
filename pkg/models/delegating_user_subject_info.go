// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DelegatingUserSubjectInfo delegating user subject info
// swagger:model DelegatingUserSubjectInfo
type DelegatingUserSubjectInfo struct {

	// bearer token
	BearerToken string `json:"bearerToken,omitempty"`

	// subject context
	SubjectContext *SubjectContext `json:"subjectContext,omitempty"`
}

// Validate validates this delegating user subject info
func (m *DelegatingUserSubjectInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubjectContext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DelegatingUserSubjectInfo) validateSubjectContext(formats strfmt.Registry) error {

	if swag.IsZero(m.SubjectContext) { // not required
		return nil
	}

	if m.SubjectContext != nil {
		if err := m.SubjectContext.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subjectContext")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DelegatingUserSubjectInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DelegatingUserSubjectInfo) UnmarshalBinary(b []byte) error {
	var res DelegatingUserSubjectInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
