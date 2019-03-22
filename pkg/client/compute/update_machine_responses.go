// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateMachineReader is a Reader for the UpdateMachine structure.
type UpdateMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 403:
		result := NewUpdateMachineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateMachineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateMachineForbidden creates a UpdateMachineForbidden with default headers values
func NewUpdateMachineForbidden() *UpdateMachineForbidden {
	return &UpdateMachineForbidden{}
}

/*UpdateMachineForbidden handles this case with default header values.

Forbidden
*/
type UpdateMachineForbidden struct {
}

func (o *UpdateMachineForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/machines/{id}][%d] updateMachineForbidden ", 403)
}

func (o *UpdateMachineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateMachineNotFound creates a UpdateMachineNotFound with default headers values
func NewUpdateMachineNotFound() *UpdateMachineNotFound {
	return &UpdateMachineNotFound{}
}

/*UpdateMachineNotFound handles this case with default header values.

Not Found
*/
type UpdateMachineNotFound struct {
}

func (o *UpdateMachineNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/machines/{id}][%d] updateMachineNotFound ", 404)
}

func (o *UpdateMachineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
