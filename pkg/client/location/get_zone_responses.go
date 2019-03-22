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

// GetZoneReader is a Reader for the GetZone structure.
type GetZoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetZoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetZoneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetZoneOK creates a GetZoneOK with default headers values
func NewGetZoneOK() *GetZoneOK {
	return &GetZoneOK{}
}

/*GetZoneOK handles this case with default header values.

successful operation
*/
type GetZoneOK struct {
	Payload *models.ZoneResult
}

func (o *GetZoneOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/zones][%d] getZoneOK  %+v", 200, o.Payload)
}

func (o *GetZoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZoneResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZoneForbidden creates a GetZoneForbidden with default headers values
func NewGetZoneForbidden() *GetZoneForbidden {
	return &GetZoneForbidden{}
}

/*GetZoneForbidden handles this case with default header values.

Forbidden
*/
type GetZoneForbidden struct {
}

func (o *GetZoneForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/zones][%d] getZoneForbidden ", 403)
}

func (o *GetZoneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
