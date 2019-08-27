// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetNetworksReader is a Reader for the GetNetworks structure.
type GetNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetNetworksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworksOK creates a GetNetworksOK with default headers values
func NewGetNetworksOK() *GetNetworksOK {
	return &GetNetworksOK{}
}

/*GetNetworksOK handles this case with default header values.

successful operation
*/
type GetNetworksOK struct {
	Payload *models.NetworkResult
}

func (o *GetNetworksOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/networks][%d] getNetworksOK  %+v", 200, o.Payload)
}

func (o *GetNetworksOK) GetPayload() *models.NetworkResult {
	return o.Payload
}

func (o *GetNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworksForbidden creates a GetNetworksForbidden with default headers values
func NewGetNetworksForbidden() *GetNetworksForbidden {
	return &GetNetworksForbidden{}
}

/*GetNetworksForbidden handles this case with default header values.

Forbidden
*/
type GetNetworksForbidden struct {
}

func (o *GetNetworksForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/networks][%d] getNetworksForbidden ", 403)
}

func (o *GetNetworksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
