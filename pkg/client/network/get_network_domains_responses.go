// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetNetworkDomainsReader is a Reader for the GetNetworkDomains structure.
type GetNetworkDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetworkDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetNetworkDomainsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworkDomainsOK creates a GetNetworkDomainsOK with default headers values
func NewGetNetworkDomainsOK() *GetNetworkDomainsOK {
	return &GetNetworkDomainsOK{}
}

/*GetNetworkDomainsOK handles this case with default header values.

successful operation
*/
type GetNetworkDomainsOK struct {
	Payload *models.NetworkDomainResult
}

func (o *GetNetworkDomainsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/network-domains][%d] getNetworkDomainsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkDomainResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkDomainsForbidden creates a GetNetworkDomainsForbidden with default headers values
func NewGetNetworkDomainsForbidden() *GetNetworkDomainsForbidden {
	return &GetNetworkDomainsForbidden{}
}

/*GetNetworkDomainsForbidden handles this case with default header values.

Forbidden
*/
type GetNetworkDomainsForbidden struct {
}

func (o *GetNetworkDomainsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/network-domains][%d] getNetworkDomainsForbidden ", 403)
}

func (o *GetNetworkDomainsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
