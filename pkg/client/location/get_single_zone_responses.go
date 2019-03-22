// Code generated by go-swagger; DO NOT EDIT.

package location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetSingleZoneReader is a Reader for the GetSingleZone structure.
type GetSingleZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSingleZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSingleZoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetSingleZoneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSingleZoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSingleZoneOK creates a GetSingleZoneOK with default headers values
func NewGetSingleZoneOK() *GetSingleZoneOK {
	return &GetSingleZoneOK{}
}

/*GetSingleZoneOK handles this case with default header values.

successful operation
*/
type GetSingleZoneOK struct {
	Payload *models.Zone
}

func (o *GetSingleZoneOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/zones/{id}][%d] getSingleZoneOK  %+v", 200, o.Payload)
}

func (o *GetSingleZoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Zone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSingleZoneForbidden creates a GetSingleZoneForbidden with default headers values
func NewGetSingleZoneForbidden() *GetSingleZoneForbidden {
	return &GetSingleZoneForbidden{}
}

/*GetSingleZoneForbidden handles this case with default header values.

Forbidden
*/
type GetSingleZoneForbidden struct {
}

func (o *GetSingleZoneForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/zones/{id}][%d] getSingleZoneForbidden ", 403)
}

func (o *GetSingleZoneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSingleZoneNotFound creates a GetSingleZoneNotFound with default headers values
func NewGetSingleZoneNotFound() *GetSingleZoneNotFound {
	return &GetSingleZoneNotFound{}
}

/*GetSingleZoneNotFound handles this case with default header values.

Not Found
*/
type GetSingleZoneNotFound struct {
}

func (o *GetSingleZoneNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/zones/{id}][%d] getSingleZoneNotFound ", 404)
}

func (o *GetSingleZoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
