// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetProjectReader is a Reader for the GetProject structure.
type GetProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProjectOK creates a GetProjectOK with default headers values
func NewGetProjectOK() *GetProjectOK {
	return &GetProjectOK{}
}

/*GetProjectOK handles this case with default header values.

successful operation
*/
type GetProjectOK struct {
	Payload *models.Project
}

func (o *GetProjectOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/projects/{id}][%d] getProjectOK  %+v", 200, o.Payload)
}

func (o *GetProjectOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *GetProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectForbidden creates a GetProjectForbidden with default headers values
func NewGetProjectForbidden() *GetProjectForbidden {
	return &GetProjectForbidden{}
}

/*GetProjectForbidden handles this case with default header values.

Forbidden
*/
type GetProjectForbidden struct {
}

func (o *GetProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/projects/{id}][%d] getProjectForbidden ", 403)
}

func (o *GetProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectNotFound creates a GetProjectNotFound with default headers values
func NewGetProjectNotFound() *GetProjectNotFound {
	return &GetProjectNotFound{}
}

/*GetProjectNotFound handles this case with default header values.

Not Found
*/
type GetProjectNotFound struct {
	Payload *models.Error
}

func (o *GetProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/projects/{id}][%d] getProjectNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
