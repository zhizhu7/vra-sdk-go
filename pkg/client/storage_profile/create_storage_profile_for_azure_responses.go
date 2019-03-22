// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// CreateStorageProfileForAzureReader is a Reader for the CreateStorageProfileForAzure structure.
type CreateStorageProfileForAzureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStorageProfileForAzureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateStorageProfileForAzureCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateStorageProfileForAzureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateStorageProfileForAzureForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateStorageProfileForAzureCreated creates a CreateStorageProfileForAzureCreated with default headers values
func NewCreateStorageProfileForAzureCreated() *CreateStorageProfileForAzureCreated {
	return &CreateStorageProfileForAzureCreated{}
}

/*CreateStorageProfileForAzureCreated handles this case with default header values.

successful operation
*/
type CreateStorageProfileForAzureCreated struct {
	Payload *models.AzureStorageProfile
}

func (o *CreateStorageProfileForAzureCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles-azure][%d] createStorageProfileForAzureCreated  %+v", 201, o.Payload)
}

func (o *CreateStorageProfileForAzureCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AzureStorageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageProfileForAzureBadRequest creates a CreateStorageProfileForAzureBadRequest with default headers values
func NewCreateStorageProfileForAzureBadRequest() *CreateStorageProfileForAzureBadRequest {
	return &CreateStorageProfileForAzureBadRequest{}
}

/*CreateStorageProfileForAzureBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type CreateStorageProfileForAzureBadRequest struct {
}

func (o *CreateStorageProfileForAzureBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles-azure][%d] createStorageProfileForAzureBadRequest ", 400)
}

func (o *CreateStorageProfileForAzureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateStorageProfileForAzureForbidden creates a CreateStorageProfileForAzureForbidden with default headers values
func NewCreateStorageProfileForAzureForbidden() *CreateStorageProfileForAzureForbidden {
	return &CreateStorageProfileForAzureForbidden{}
}

/*CreateStorageProfileForAzureForbidden handles this case with default header values.

Forbidden
*/
type CreateStorageProfileForAzureForbidden struct {
}

func (o *CreateStorageProfileForAzureForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles-azure][%d] createStorageProfileForAzureForbidden ", 403)
}

func (o *CreateStorageProfileForAzureForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
