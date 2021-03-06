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

// Claims claims
// swagger:model Claims
type Claims struct {

	// account
	Account string `json:"account,omitempty"`

	// audience
	// Unique: true
	Audience []string `json:"audience"`

	// client Id
	ClientID string `json:"clientId,omitempty"`

	// expiration time
	ExpirationTime int64 `json:"expirationTime,omitempty"`

	// issued at
	IssuedAt int64 `json:"issuedAt,omitempty"`

	// issuer
	Issuer string `json:"issuer,omitempty"`

	// jwt Id
	JwtID string `json:"jwtId,omitempty"`

	// not before
	NotBefore int64 `json:"notBefore,omitempty"`

	// properties
	Properties map[string]string `json:"properties,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`
}

// Validate validates this claims
func (m *Claims) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAudience(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Claims) validateAudience(formats strfmt.Registry) error {

	if swag.IsZero(m.Audience) { // not required
		return nil
	}

	if err := validate.UniqueItems("audience", "body", m.Audience); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Claims) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Claims) UnmarshalBinary(b []byte) error {
	var res Claims
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
