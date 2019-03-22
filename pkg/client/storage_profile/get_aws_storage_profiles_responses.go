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

// GetAwsStorageProfilesReader is a Reader for the GetAwsStorageProfiles structure.
type GetAwsStorageProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAwsStorageProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAwsStorageProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetAwsStorageProfilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAwsStorageProfilesOK creates a GetAwsStorageProfilesOK with default headers values
func NewGetAwsStorageProfilesOK() *GetAwsStorageProfilesOK {
	return &GetAwsStorageProfilesOK{}
}

/*GetAwsStorageProfilesOK handles this case with default header values.

successful operation
*/
type GetAwsStorageProfilesOK struct {
	Payload *models.StorageProfileAwsResult
}

func (o *GetAwsStorageProfilesOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-aws][%d] getAwsStorageProfilesOK  %+v", 200, o.Payload)
}

func (o *GetAwsStorageProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageProfileAwsResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAwsStorageProfilesForbidden creates a GetAwsStorageProfilesForbidden with default headers values
func NewGetAwsStorageProfilesForbidden() *GetAwsStorageProfilesForbidden {
	return &GetAwsStorageProfilesForbidden{}
}

/*GetAwsStorageProfilesForbidden handles this case with default header values.

Forbidden
*/
type GetAwsStorageProfilesForbidden struct {
}

func (o *GetAwsStorageProfilesForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-aws][%d] getAwsStorageProfilesForbidden ", 403)
}

func (o *GetAwsStorageProfilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
