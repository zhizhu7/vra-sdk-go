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

// ReplaceStorageProfileReader is a Reader for the ReplaceStorageProfile structure.
type ReplaceStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReplaceStorageProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewReplaceStorageProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewReplaceStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceStorageProfileOK creates a ReplaceStorageProfileOK with default headers values
func NewReplaceStorageProfileOK() *ReplaceStorageProfileOK {
	return &ReplaceStorageProfileOK{}
}

/*ReplaceStorageProfileOK handles this case with default header values.

successful operation
*/
type ReplaceStorageProfileOK struct {
	Payload *models.StorageProfile
}

func (o *ReplaceStorageProfileOK) Error() string {
	return fmt.Sprintf("[PUT /iaas/api/storage-profiles/{id}][%d] replaceStorageProfileOK  %+v", 200, o.Payload)
}

func (o *ReplaceStorageProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceStorageProfileBadRequest creates a ReplaceStorageProfileBadRequest with default headers values
func NewReplaceStorageProfileBadRequest() *ReplaceStorageProfileBadRequest {
	return &ReplaceStorageProfileBadRequest{}
}

/*ReplaceStorageProfileBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type ReplaceStorageProfileBadRequest struct {
}

func (o *ReplaceStorageProfileBadRequest) Error() string {
	return fmt.Sprintf("[PUT /iaas/api/storage-profiles/{id}][%d] replaceStorageProfileBadRequest ", 400)
}

func (o *ReplaceStorageProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReplaceStorageProfileForbidden creates a ReplaceStorageProfileForbidden with default headers values
func NewReplaceStorageProfileForbidden() *ReplaceStorageProfileForbidden {
	return &ReplaceStorageProfileForbidden{}
}

/*ReplaceStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type ReplaceStorageProfileForbidden struct {
}

func (o *ReplaceStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[PUT /iaas/api/storage-profiles/{id}][%d] replaceStorageProfileForbidden ", 403)
}

func (o *ReplaceStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}