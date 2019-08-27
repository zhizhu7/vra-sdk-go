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

// UpdateAzureStorageProfileReader is a Reader for the UpdateAzureStorageProfile structure.
type UpdateAzureStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAzureStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAzureStorageProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAzureStorageProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAzureStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAzureStorageProfileOK creates a UpdateAzureStorageProfileOK with default headers values
func NewUpdateAzureStorageProfileOK() *UpdateAzureStorageProfileOK {
	return &UpdateAzureStorageProfileOK{}
}

/*UpdateAzureStorageProfileOK handles this case with default header values.

successful operation
*/
type UpdateAzureStorageProfileOK struct {
	Payload *models.AzureStorageProfile
}

func (o *UpdateAzureStorageProfileOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/storage-profiles-azure/{id}][%d] updateAzureStorageProfileOK  %+v", 200, o.Payload)
}

func (o *UpdateAzureStorageProfileOK) GetPayload() *models.AzureStorageProfile {
	return o.Payload
}

func (o *UpdateAzureStorageProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureStorageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAzureStorageProfileBadRequest creates a UpdateAzureStorageProfileBadRequest with default headers values
func NewUpdateAzureStorageProfileBadRequest() *UpdateAzureStorageProfileBadRequest {
	return &UpdateAzureStorageProfileBadRequest{}
}

/*UpdateAzureStorageProfileBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type UpdateAzureStorageProfileBadRequest struct {
}

func (o *UpdateAzureStorageProfileBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/storage-profiles-azure/{id}][%d] updateAzureStorageProfileBadRequest ", 400)
}

func (o *UpdateAzureStorageProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAzureStorageProfileForbidden creates a UpdateAzureStorageProfileForbidden with default headers values
func NewUpdateAzureStorageProfileForbidden() *UpdateAzureStorageProfileForbidden {
	return &UpdateAzureStorageProfileForbidden{}
}

/*UpdateAzureStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type UpdateAzureStorageProfileForbidden struct {
}

func (o *UpdateAzureStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/storage-profiles-azure/{id}][%d] updateAzureStorageProfileForbidden ", 403)
}

func (o *UpdateAzureStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
