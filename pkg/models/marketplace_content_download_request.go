// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MarketplaceContentDownloadRequest MarketplaceContentDownloadRequest
// swagger:model MarketplaceContentDownloadRequest
type MarketplaceContentDownloadRequest struct {

	// Content Id
	// Required: true
	ContentID *string `json:"contentId"`

	// Content Name
	// Required: true
	ContentName *string `json:"contentName"`

	// Content Type
	// Required: true
	// Enum: [BLUEPRINT IMAGE ABX_SCRIPTS]
	ContentType *string `json:"contentType"`

	// Request Created Time
	// Read Only: true
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Downloaded Content Id
	// Read Only: true
	DownloadedContentID string `json:"downloadedContentId,omitempty"`

	// Execution Messages
	// Read Only: true
	ExecutionMessages []*ExecutionMessage `json:"executionMessages"`

	// Request Id
	// Read Only: true
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Request Last Updated Time
	// Read Only: true
	// Format: date-time
	LastUpdatedAt strfmt.DateTime `json:"lastUpdatedAt,omitempty"`

	// Content Source Id
	// Required: true
	// Format: uuid
	SourceID *strfmt.UUID `json:"sourceId"`

	// Request Status
	// Read Only: true
	Status string `json:"status,omitempty"`

	// Target Id
	// Required: true
	TargetID *string `json:"targetId"`

	// Target Type
	// Required: true
	// Enum: [PROJECT]
	TargetType *string `json:"targetType"`
}

// Validate validates this marketplace content download request
func (m *MarketplaceContentDownloadRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutionMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MarketplaceContentDownloadRequest) validateContentID(formats strfmt.Registry) error {

	if err := validate.Required("contentId", "body", m.ContentID); err != nil {
		return err
	}

	return nil
}

func (m *MarketplaceContentDownloadRequest) validateContentName(formats strfmt.Registry) error {

	if err := validate.Required("contentName", "body", m.ContentName); err != nil {
		return err
	}

	return nil
}

var marketplaceContentDownloadRequestTypeContentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BLUEPRINT","IMAGE","ABX_SCRIPTS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketplaceContentDownloadRequestTypeContentTypePropEnum = append(marketplaceContentDownloadRequestTypeContentTypePropEnum, v)
	}
}

const (

	// MarketplaceContentDownloadRequestContentTypeBLUEPRINT captures enum value "BLUEPRINT"
	MarketplaceContentDownloadRequestContentTypeBLUEPRINT string = "BLUEPRINT"

	// MarketplaceContentDownloadRequestContentTypeIMAGE captures enum value "IMAGE"
	MarketplaceContentDownloadRequestContentTypeIMAGE string = "IMAGE"

	// MarketplaceContentDownloadRequestContentTypeABXSCRIPTS captures enum value "ABX_SCRIPTS"
	MarketplaceContentDownloadRequestContentTypeABXSCRIPTS string = "ABX_SCRIPTS"
)

// prop value enum
func (m *MarketplaceContentDownloadRequest) validateContentTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, marketplaceContentDownloadRequestTypeContentTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MarketplaceContentDownloadRequest) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("contentType", "body", m.ContentType); err != nil {
		return err
	}

	// value enum
	if err := m.validateContentTypeEnum("contentType", "body", *m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *MarketplaceContentDownloadRequest) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MarketplaceContentDownloadRequest) validateExecutionMessages(formats strfmt.Registry) error {

	if swag.IsZero(m.ExecutionMessages) { // not required
		return nil
	}

	for i := 0; i < len(m.ExecutionMessages); i++ {
		if swag.IsZero(m.ExecutionMessages[i]) { // not required
			continue
		}

		if m.ExecutionMessages[i] != nil {
			if err := m.ExecutionMessages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("executionMessages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MarketplaceContentDownloadRequest) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MarketplaceContentDownloadRequest) validateLastUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.LastUpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastUpdatedAt", "body", "date-time", m.LastUpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MarketplaceContentDownloadRequest) validateSourceID(formats strfmt.Registry) error {

	if err := validate.Required("sourceId", "body", m.SourceID); err != nil {
		return err
	}

	if err := validate.FormatOf("sourceId", "body", "uuid", m.SourceID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MarketplaceContentDownloadRequest) validateTargetID(formats strfmt.Registry) error {

	if err := validate.Required("targetId", "body", m.TargetID); err != nil {
		return err
	}

	return nil
}

var marketplaceContentDownloadRequestTypeTargetTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROJECT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		marketplaceContentDownloadRequestTypeTargetTypePropEnum = append(marketplaceContentDownloadRequestTypeTargetTypePropEnum, v)
	}
}

const (

	// MarketplaceContentDownloadRequestTargetTypePROJECT captures enum value "PROJECT"
	MarketplaceContentDownloadRequestTargetTypePROJECT string = "PROJECT"
)

// prop value enum
func (m *MarketplaceContentDownloadRequest) validateTargetTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, marketplaceContentDownloadRequestTypeTargetTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MarketplaceContentDownloadRequest) validateTargetType(formats strfmt.Registry) error {

	if err := validate.Required("targetType", "body", m.TargetType); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetTypeEnum("targetType", "body", *m.TargetType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MarketplaceContentDownloadRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MarketplaceContentDownloadRequest) UnmarshalBinary(b []byte) error {
	var res MarketplaceContentDownloadRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}