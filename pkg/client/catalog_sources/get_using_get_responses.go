// Code generated by go-swagger; DO NOT EDIT.

package catalog_sources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetUsingGETReader is a Reader for the GetUsingGET structure.
type GetUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsingGETOK creates a GetUsingGETOK with default headers values
func NewGetUsingGETOK() *GetUsingGETOK {
	return &GetUsingGETOK{}
}

/*GetUsingGETOK handles this case with default header values.

OK
*/
type GetUsingGETOK struct {
	Payload *models.CatalogSource
}

func (o *GetUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /catalog/api/admin/sources/{sourceId}][%d] getUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetUsingGETOK) GetPayload() *models.CatalogSource {
	return o.Payload
}

func (o *GetUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogSource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsingGETUnauthorized creates a GetUsingGETUnauthorized with default headers values
func NewGetUsingGETUnauthorized() *GetUsingGETUnauthorized {
	return &GetUsingGETUnauthorized{}
}

/*GetUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetUsingGETUnauthorized struct {
}

func (o *GetUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /catalog/api/admin/sources/{sourceId}][%d] getUsingGETUnauthorized ", 401)
}

func (o *GetUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGETForbidden creates a GetUsingGETForbidden with default headers values
func NewGetUsingGETForbidden() *GetUsingGETForbidden {
	return &GetUsingGETForbidden{}
}

/*GetUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetUsingGETForbidden struct {
}

func (o *GetUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /catalog/api/admin/sources/{sourceId}][%d] getUsingGETForbidden ", 403)
}

func (o *GetUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGETNotFound creates a GetUsingGETNotFound with default headers values
func NewGetUsingGETNotFound() *GetUsingGETNotFound {
	return &GetUsingGETNotFound{}
}

/*GetUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetUsingGETNotFound struct {
}

func (o *GetUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /catalog/api/admin/sources/{sourceId}][%d] getUsingGETNotFound ", 404)
}

func (o *GetUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
