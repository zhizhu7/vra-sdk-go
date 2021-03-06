// Code generated by go-swagger; DO NOT EDIT.

package marketplace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDetailsReader is a Reader for the GetDetails structure.
type GetDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDetailsOK creates a GetDetailsOK with default headers values
func NewGetDetailsOK() *GetDetailsOK {
	return &GetDetailsOK{}
}

/*GetDetailsOK handles this case with default header values.

Content Details
*/
type GetDetailsOK struct {
	Payload *models.MarketplaceContent
}

func (o *GetDetailsOK) Error() string {
	return fmt.Sprintf("[GET /content/api/marketplace/sources/{sourceId}/contents/{contentId}][%d] getDetailsOK  %+v", 200, o.Payload)
}

func (o *GetDetailsOK) GetPayload() *models.MarketplaceContent {
	return o.Payload
}

func (o *GetDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MarketplaceContent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDetailsNotFound creates a GetDetailsNotFound with default headers values
func NewGetDetailsNotFound() *GetDetailsNotFound {
	return &GetDetailsNotFound{}
}

/*GetDetailsNotFound handles this case with default header values.

Source or Content not found
*/
type GetDetailsNotFound struct {
	Payload *models.Error
}

func (o *GetDetailsNotFound) Error() string {
	return fmt.Sprintf("[GET /content/api/marketplace/sources/{sourceId}/contents/{contentId}][%d] getDetailsNotFound  %+v", 404, o.Payload)
}

func (o *GetDetailsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetDetailsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
