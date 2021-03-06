// Code generated by go-swagger; DO NOT EDIT.

package fabric_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetFabricNetworksReader is a Reader for the GetFabricNetworks structure.
type GetFabricNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFabricNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFabricNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFabricNetworksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFabricNetworksOK creates a GetFabricNetworksOK with default headers values
func NewGetFabricNetworksOK() *GetFabricNetworksOK {
	return &GetFabricNetworksOK{}
}

/*GetFabricNetworksOK handles this case with default header values.

successful operation
*/
type GetFabricNetworksOK struct {
	Payload *models.FabricNetworkResult
}

func (o *GetFabricNetworksOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-networks][%d] getFabricNetworksOK  %+v", 200, o.Payload)
}

func (o *GetFabricNetworksOK) GetPayload() *models.FabricNetworkResult {
	return o.Payload
}

func (o *GetFabricNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricNetworkResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricNetworksForbidden creates a GetFabricNetworksForbidden with default headers values
func NewGetFabricNetworksForbidden() *GetFabricNetworksForbidden {
	return &GetFabricNetworksForbidden{}
}

/*GetFabricNetworksForbidden handles this case with default header values.

Forbidden
*/
type GetFabricNetworksForbidden struct {
}

func (o *GetFabricNetworksForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-networks][%d] getFabricNetworksForbidden ", 403)
}

func (o *GetFabricNetworksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
