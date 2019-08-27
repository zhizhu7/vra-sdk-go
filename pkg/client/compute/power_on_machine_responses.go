// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// PowerOnMachineReader is a Reader for the PowerOnMachine structure.
type PowerOnMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PowerOnMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPowerOnMachineAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPowerOnMachineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPowerOnMachineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPowerOnMachineAccepted creates a PowerOnMachineAccepted with default headers values
func NewPowerOnMachineAccepted() *PowerOnMachineAccepted {
	return &PowerOnMachineAccepted{}
}

/*PowerOnMachineAccepted handles this case with default header values.

successful operation
*/
type PowerOnMachineAccepted struct {
	Payload *models.RequestTracker
}

func (o *PowerOnMachineAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/power-on][%d] powerOnMachineAccepted  %+v", 202, o.Payload)
}

func (o *PowerOnMachineAccepted) GetPayload() *models.RequestTracker {
	return o.Payload
}

func (o *PowerOnMachineAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPowerOnMachineForbidden creates a PowerOnMachineForbidden with default headers values
func NewPowerOnMachineForbidden() *PowerOnMachineForbidden {
	return &PowerOnMachineForbidden{}
}

/*PowerOnMachineForbidden handles this case with default header values.

Forbidden
*/
type PowerOnMachineForbidden struct {
}

func (o *PowerOnMachineForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/power-on][%d] powerOnMachineForbidden ", 403)
}

func (o *PowerOnMachineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPowerOnMachineNotFound creates a PowerOnMachineNotFound with default headers values
func NewPowerOnMachineNotFound() *PowerOnMachineNotFound {
	return &PowerOnMachineNotFound{}
}

/*PowerOnMachineNotFound handles this case with default header values.

Not Found
*/
type PowerOnMachineNotFound struct {
}

func (o *PowerOnMachineNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/power-on][%d] powerOnMachineNotFound ", 404)
}

func (o *PowerOnMachineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
