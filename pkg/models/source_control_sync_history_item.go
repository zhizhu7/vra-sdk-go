// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SourceControlSyncHistoryItem SourceControlSyncHistoryItem
// swagger:model SourceControlSyncHistoryItem
type SourceControlSyncHistoryItem struct {

	// Content Full Path
	// Read Only: true
	ContentFullPath string `json:"contentFullPath,omitempty"`

	// Content Name
	// Read Only: true
	ContentName string `json:"contentName,omitempty"`

	// Content Type
	// Read Only: true
	// Enum: [BLUEPRINT IMAGE ABX_SCRIPTS]
	ContentType string `json:"contentType,omitempty"`

	// Details
	// Read Only: true
	Details []string `json:"details"`

	// Unique Id for the Source Control History
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Integration Id
	// Read Only: true
	IntegrationID string `json:"integrationId,omitempty"`

	// Project Id
	// Read Only: true
	ProjectID string `json:"projectId,omitempty"`

	// Project Name
	// Read Only: true
	ProjectName string `json:"projectName,omitempty"`

	// Id of the sync request
	// Format: uuid
	RequestID strfmt.UUID `json:"requestId,omitempty"`

	// Content source Id
	// Read Only: true
	// Format: uuid
	SourceID strfmt.UUID `json:"sourceId,omitempty"`

	// Status
	// Read Only: true
	Status string `json:"status,omitempty"`

	// Timestamp
	// Read Only: true
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this source control sync history item
func (m *SourceControlSyncHistoryItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var sourceControlSyncHistoryItemTypeContentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BLUEPRINT","IMAGE","ABX_SCRIPTS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sourceControlSyncHistoryItemTypeContentTypePropEnum = append(sourceControlSyncHistoryItemTypeContentTypePropEnum, v)
	}
}

const (

	// SourceControlSyncHistoryItemContentTypeBLUEPRINT captures enum value "BLUEPRINT"
	SourceControlSyncHistoryItemContentTypeBLUEPRINT string = "BLUEPRINT"

	// SourceControlSyncHistoryItemContentTypeIMAGE captures enum value "IMAGE"
	SourceControlSyncHistoryItemContentTypeIMAGE string = "IMAGE"

	// SourceControlSyncHistoryItemContentTypeABXSCRIPTS captures enum value "ABX_SCRIPTS"
	SourceControlSyncHistoryItemContentTypeABXSCRIPTS string = "ABX_SCRIPTS"
)

// prop value enum
func (m *SourceControlSyncHistoryItem) validateContentTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sourceControlSyncHistoryItemTypeContentTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SourceControlSyncHistoryItem) validateContentType(formats strfmt.Registry) error {

	if swag.IsZero(m.ContentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateContentTypeEnum("contentType", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *SourceControlSyncHistoryItem) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SourceControlSyncHistoryItem) validateRequestID(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestID) { // not required
		return nil
	}

	if err := validate.FormatOf("requestId", "body", "uuid", m.RequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SourceControlSyncHistoryItem) validateSourceID(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceID) { // not required
		return nil
	}

	if err := validate.FormatOf("sourceId", "body", "uuid", m.SourceID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SourceControlSyncHistoryItem) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SourceControlSyncHistoryItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourceControlSyncHistoryItem) UnmarshalBinary(b []byte) error {
	var res SourceControlSyncHistoryItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
