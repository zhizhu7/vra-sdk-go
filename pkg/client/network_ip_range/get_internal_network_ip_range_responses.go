// Code generated by go-swagger; DO NOT EDIT.

package network_ip_range

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetInternalNetworkIPRangeReader is a Reader for the GetInternalNetworkIPRange structure.
type GetInternalNetworkIPRangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInternalNetworkIPRangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInternalNetworkIPRangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetInternalNetworkIPRangeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInternalNetworkIPRangeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInternalNetworkIPRangeOK creates a GetInternalNetworkIPRangeOK with default headers values
func NewGetInternalNetworkIPRangeOK() *GetInternalNetworkIPRangeOK {
	return &GetInternalNetworkIPRangeOK{}
}

/*GetInternalNetworkIPRangeOK handles this case with default header values.

successful operation
*/
type GetInternalNetworkIPRangeOK struct {
	Payload *models.NetworkIPRange
}

func (o *GetInternalNetworkIPRangeOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/network-ip-ranges/{id}][%d] getInternalNetworkIpRangeOK  %+v", 200, o.Payload)
}

func (o *GetInternalNetworkIPRangeOK) GetPayload() *models.NetworkIPRange {
	return o.Payload
}

func (o *GetInternalNetworkIPRangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkIPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInternalNetworkIPRangeForbidden creates a GetInternalNetworkIPRangeForbidden with default headers values
func NewGetInternalNetworkIPRangeForbidden() *GetInternalNetworkIPRangeForbidden {
	return &GetInternalNetworkIPRangeForbidden{}
}

/*GetInternalNetworkIPRangeForbidden handles this case with default header values.

Forbidden
*/
type GetInternalNetworkIPRangeForbidden struct {
}

func (o *GetInternalNetworkIPRangeForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/network-ip-ranges/{id}][%d] getInternalNetworkIpRangeForbidden ", 403)
}

func (o *GetInternalNetworkIPRangeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInternalNetworkIPRangeNotFound creates a GetInternalNetworkIPRangeNotFound with default headers values
func NewGetInternalNetworkIPRangeNotFound() *GetInternalNetworkIPRangeNotFound {
	return &GetInternalNetworkIPRangeNotFound{}
}

/*GetInternalNetworkIPRangeNotFound handles this case with default header values.

Not Found
*/
type GetInternalNetworkIPRangeNotFound struct {
	Payload *models.Error
}

func (o *GetInternalNetworkIPRangeNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/network-ip-ranges/{id}][%d] getInternalNetworkIpRangeNotFound  %+v", 404, o.Payload)
}

func (o *GetInternalNetworkIPRangeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInternalNetworkIPRangeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
