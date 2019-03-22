// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// DeleteMachineDiskReader is a Reader for the DeleteMachineDisk structure.
type DeleteMachineDiskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMachineDiskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMachineDiskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteMachineDiskForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteMachineDiskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteMachineDiskOK creates a DeleteMachineDiskOK with default headers values
func NewDeleteMachineDiskOK() *DeleteMachineDiskOK {
	return &DeleteMachineDiskOK{}
}

/*DeleteMachineDiskOK handles this case with default header values.

successful operation
*/
type DeleteMachineDiskOK struct {
	Payload *models.BlockDevice
}

func (o *DeleteMachineDiskOK) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/disks/{id1}][%d] deleteMachineDiskOK  %+v", 200, o.Payload)
}

func (o *DeleteMachineDiskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlockDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMachineDiskForbidden creates a DeleteMachineDiskForbidden with default headers values
func NewDeleteMachineDiskForbidden() *DeleteMachineDiskForbidden {
	return &DeleteMachineDiskForbidden{}
}

/*DeleteMachineDiskForbidden handles this case with default header values.

Forbidden
*/
type DeleteMachineDiskForbidden struct {
}

func (o *DeleteMachineDiskForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/disks/{id1}][%d] deleteMachineDiskForbidden ", 403)
}

func (o *DeleteMachineDiskForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMachineDiskNotFound creates a DeleteMachineDiskNotFound with default headers values
func NewDeleteMachineDiskNotFound() *DeleteMachineDiskNotFound {
	return &DeleteMachineDiskNotFound{}
}

/*DeleteMachineDiskNotFound handles this case with default header values.

Not Found
*/
type DeleteMachineDiskNotFound struct {
}

func (o *DeleteMachineDiskNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/disks/{id1}][%d] deleteMachineDiskNotFound ", 404)
}

func (o *DeleteMachineDiskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
