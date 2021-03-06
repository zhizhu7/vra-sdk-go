// Code generated by go-swagger; DO NOT EDIT.

package fabric_compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetFabricComputesReader is a Reader for the GetFabricComputes structure.
type GetFabricComputesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFabricComputesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFabricComputesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFabricComputesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFabricComputesOK creates a GetFabricComputesOK with default headers values
func NewGetFabricComputesOK() *GetFabricComputesOK {
	return &GetFabricComputesOK{}
}

/*GetFabricComputesOK handles this case with default header values.

successful operation
*/
type GetFabricComputesOK struct {
	Payload *models.FabricComputeResult
}

func (o *GetFabricComputesOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-computes][%d] getFabricComputesOK  %+v", 200, o.Payload)
}

func (o *GetFabricComputesOK) GetPayload() *models.FabricComputeResult {
	return o.Payload
}

func (o *GetFabricComputesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricComputeResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricComputesForbidden creates a GetFabricComputesForbidden with default headers values
func NewGetFabricComputesForbidden() *GetFabricComputesForbidden {
	return &GetFabricComputesForbidden{}
}

/*GetFabricComputesForbidden handles this case with default header values.

Forbidden
*/
type GetFabricComputesForbidden struct {
}

func (o *GetFabricComputesForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-computes][%d] getFabricComputesForbidden ", 403)
}

func (o *GetFabricComputesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
