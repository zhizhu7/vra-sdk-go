// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetMachinesReader is a Reader for the GetMachines structure.
type GetMachinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMachinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetMachinesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMachinesOK creates a GetMachinesOK with default headers values
func NewGetMachinesOK() *GetMachinesOK {
	return &GetMachinesOK{}
}

/*GetMachinesOK handles this case with default header values.

successful operation
*/
type GetMachinesOK struct {
	Payload *models.MachineResult
}

func (o *GetMachinesOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines][%d] getMachinesOK  %+v", 200, o.Payload)
}

func (o *GetMachinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachinesForbidden creates a GetMachinesForbidden with default headers values
func NewGetMachinesForbidden() *GetMachinesForbidden {
	return &GetMachinesForbidden{}
}

/*GetMachinesForbidden handles this case with default header values.

Forbidden
*/
type GetMachinesForbidden struct {
}

func (o *GetMachinesForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines][%d] getMachinesForbidden ", 403)
}

func (o *GetMachinesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
