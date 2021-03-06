// Code generated by go-swagger; DO NOT EDIT.

package request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// CancelRequestReader is a Reader for the CancelRequest structure.
type CancelRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 403:
		result := NewCancelRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelRequestForbidden creates a CancelRequestForbidden with default headers values
func NewCancelRequestForbidden() *CancelRequestForbidden {
	return &CancelRequestForbidden{}
}

/*CancelRequestForbidden handles this case with default header values.

Forbidden
*/
type CancelRequestForbidden struct {
}

func (o *CancelRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/request-tracker/{id}/operations/cancel][%d] cancelRequestForbidden ", 403)
}

func (o *CancelRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelRequestNotFound creates a CancelRequestNotFound with default headers values
func NewCancelRequestNotFound() *CancelRequestNotFound {
	return &CancelRequestNotFound{}
}

/*CancelRequestNotFound handles this case with default header values.

Not Found
*/
type CancelRequestNotFound struct {
	Payload *models.Error
}

func (o *CancelRequestNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/request-tracker/{id}/operations/cancel][%d] cancelRequestNotFound  %+v", 404, o.Payload)
}

func (o *CancelRequestNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CancelRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
