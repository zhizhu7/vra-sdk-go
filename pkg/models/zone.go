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

// Zone Description of a compute placement zone.  This can be used to specify a subset of compute resources within a region where machines can be placed.**region** - Region - Region for the zone.<br>**self** - Zone - Self link to this zone
// swagger:model Zone
type Zone struct {

	// HATEOAS of the entity
	// Required: true
	Links map[string]Href `json:"_links"`

	// Date when the entity was created. The date is in ISO 6801 and UTC.
	CreatedAt string `json:"createdAt,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// The id of this resource instance
	// Required: true
	ID *string `json:"id"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// The id of the organization this entity belongs to.
	OrganizationID string `json:"organizationId,omitempty"`

	// Email of the user that owns the entity.
	Owner string `json:"owner,omitempty"`

	// The placement policy for the zone.
	PlacementPolicy string `json:"placementPolicy,omitempty"`

	// A set of tag keys and optional values that were set on this placement.
	Tags []*Tag `json:"tags"`

	// A set of tag keys and optional values for compute resource filtering.
	TagsToMatch []*Tag `json:"tagsToMatch"`

	// Date when the entity was last updated. The date is ISO 8601 and UTC.
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this zone
func (m *Zone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagsToMatch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Zone) validateLinks(formats strfmt.Registry) error {

	for k := range m.Links {

		if err := validate.Required("_links"+"."+k, "body", m.Links[k]); err != nil {
			return err
		}
		if val, ok := m.Links[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Zone) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Zone) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Zone) validateTagsToMatch(formats strfmt.Registry) error {

	if swag.IsZero(m.TagsToMatch) { // not required
		return nil
	}

	for i := 0; i < len(m.TagsToMatch); i++ {
		if swag.IsZero(m.TagsToMatch[i]) { // not required
			continue
		}

		if m.TagsToMatch[i] != nil {
			if err := m.TagsToMatch[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagsToMatch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Zone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Zone) UnmarshalBinary(b []byte) error {
	var res Zone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
