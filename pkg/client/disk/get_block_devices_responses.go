// Code generated by go-swagger; DO NOT EDIT.

package disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetBlockDevicesReader is a Reader for the GetBlockDevices structure.
type GetBlockDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlockDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBlockDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetBlockDevicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlockDevicesOK creates a GetBlockDevicesOK with default headers values
func NewGetBlockDevicesOK() *GetBlockDevicesOK {
	return &GetBlockDevicesOK{}
}

/*GetBlockDevicesOK handles this case with default header values.

successful operation
*/
type GetBlockDevicesOK struct {
	Payload *models.BlockDeviceResult
}

func (o *GetBlockDevicesOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/block-devices][%d] getBlockDevicesOK  %+v", 200, o.Payload)
}

func (o *GetBlockDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlockDeviceResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlockDevicesForbidden creates a GetBlockDevicesForbidden with default headers values
func NewGetBlockDevicesForbidden() *GetBlockDevicesForbidden {
	return &GetBlockDevicesForbidden{}
}

/*GetBlockDevicesForbidden handles this case with default header values.

Forbidden
*/
type GetBlockDevicesForbidden struct {
}

func (o *GetBlockDevicesForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/block-devices][%d] getBlockDevicesForbidden ", 403)
}

func (o *GetBlockDevicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
