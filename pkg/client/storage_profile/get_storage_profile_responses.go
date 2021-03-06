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

// GetStorageProfileReader is a Reader for the GetStorageProfile structure.
type GetStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStorageProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStorageProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStorageProfileOK creates a GetStorageProfileOK with default headers values
func NewGetStorageProfileOK() *GetStorageProfileOK {
	return &GetStorageProfileOK{}
}

/*GetStorageProfileOK handles this case with default header values.

successful operation
*/
type GetStorageProfileOK struct {
	Payload *models.StorageProfile
}

func (o *GetStorageProfileOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles/{id}][%d] getStorageProfileOK  %+v", 200, o.Payload)
}

func (o *GetStorageProfileOK) GetPayload() *models.StorageProfile {
	return o.Payload
}

func (o *GetStorageProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStorageProfileForbidden creates a GetStorageProfileForbidden with default headers values
func NewGetStorageProfileForbidden() *GetStorageProfileForbidden {
	return &GetStorageProfileForbidden{}
}

/*GetStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type GetStorageProfileForbidden struct {
}

func (o *GetStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles/{id}][%d] getStorageProfileForbidden ", 403)
}

func (o *GetStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStorageProfileNotFound creates a GetStorageProfileNotFound with default headers values
func NewGetStorageProfileNotFound() *GetStorageProfileNotFound {
	return &GetStorageProfileNotFound{}
}

/*GetStorageProfileNotFound handles this case with default header values.

Not Found
*/
type GetStorageProfileNotFound struct {
	Payload *models.Error
}

func (o *GetStorageProfileNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles/{id}][%d] getStorageProfileNotFound  %+v", 404, o.Payload)
}

func (o *GetStorageProfileNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetStorageProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
