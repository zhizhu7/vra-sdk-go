// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateVSphereStorageProfileReader is a Reader for the UpdateVSphereStorageProfile structure.
type UpdateVSphereStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVSphereStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVSphereStorageProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateVSphereStorageProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateVSphereStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateVSphereStorageProfileOK creates a UpdateVSphereStorageProfileOK with default headers values
func NewUpdateVSphereStorageProfileOK() *UpdateVSphereStorageProfileOK {
	return &UpdateVSphereStorageProfileOK{}
}

/*UpdateVSphereStorageProfileOK handles this case with default header values.

successful operation
*/
type UpdateVSphereStorageProfileOK struct {
	Payload *models.VsphereStorageProfile
}

func (o *UpdateVSphereStorageProfileOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/storage-profiles-vsphere/{id}][%d] updateVSphereStorageProfileOK  %+v", 200, o.Payload)
}

func (o *UpdateVSphereStorageProfileOK) GetPayload() *models.VsphereStorageProfile {
	return o.Payload
}

func (o *UpdateVSphereStorageProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VsphereStorageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVSphereStorageProfileBadRequest creates a UpdateVSphereStorageProfileBadRequest with default headers values
func NewUpdateVSphereStorageProfileBadRequest() *UpdateVSphereStorageProfileBadRequest {
	return &UpdateVSphereStorageProfileBadRequest{}
}

/*UpdateVSphereStorageProfileBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type UpdateVSphereStorageProfileBadRequest struct {
	Payload *models.Error
}

func (o *UpdateVSphereStorageProfileBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/storage-profiles-vsphere/{id}][%d] updateVSphereStorageProfileBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateVSphereStorageProfileBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateVSphereStorageProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVSphereStorageProfileForbidden creates a UpdateVSphereStorageProfileForbidden with default headers values
func NewUpdateVSphereStorageProfileForbidden() *UpdateVSphereStorageProfileForbidden {
	return &UpdateVSphereStorageProfileForbidden{}
}

/*UpdateVSphereStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type UpdateVSphereStorageProfileForbidden struct {
}

func (o *UpdateVSphereStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/storage-profiles-vsphere/{id}][%d] updateVSphereStorageProfileForbidden ", 403)
}

func (o *UpdateVSphereStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
