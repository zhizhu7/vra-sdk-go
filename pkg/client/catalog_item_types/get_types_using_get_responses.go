// Code generated by go-swagger; DO NOT EDIT.

package catalog_item_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetTypesUsingGETReader is a Reader for the GetTypesUsingGET structure.
type GetTypesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTypesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTypesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetTypesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetTypesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetTypesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTypesUsingGETOK creates a GetTypesUsingGETOK with default headers values
func NewGetTypesUsingGETOK() *GetTypesUsingGETOK {
	return &GetTypesUsingGETOK{}
}

/*GetTypesUsingGETOK handles this case with default header values.

OK
*/
type GetTypesUsingGETOK struct {
	Payload *models.CatalogItemType
}

func (o *GetTypesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /catalog/api/types][%d] getTypesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetTypesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogItemType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTypesUsingGETUnauthorized creates a GetTypesUsingGETUnauthorized with default headers values
func NewGetTypesUsingGETUnauthorized() *GetTypesUsingGETUnauthorized {
	return &GetTypesUsingGETUnauthorized{}
}

/*GetTypesUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetTypesUsingGETUnauthorized struct {
}

func (o *GetTypesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /catalog/api/types][%d] getTypesUsingGETUnauthorized ", 401)
}

func (o *GetTypesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTypesUsingGETForbidden creates a GetTypesUsingGETForbidden with default headers values
func NewGetTypesUsingGETForbidden() *GetTypesUsingGETForbidden {
	return &GetTypesUsingGETForbidden{}
}

/*GetTypesUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetTypesUsingGETForbidden struct {
}

func (o *GetTypesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /catalog/api/types][%d] getTypesUsingGETForbidden ", 403)
}

func (o *GetTypesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTypesUsingGETNotFound creates a GetTypesUsingGETNotFound with default headers values
func NewGetTypesUsingGETNotFound() *GetTypesUsingGETNotFound {
	return &GetTypesUsingGETNotFound{}
}

/*GetTypesUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetTypesUsingGETNotFound struct {
}

func (o *GetTypesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /catalog/api/types][%d] getTypesUsingGETNotFound ", 404)
}

func (o *GetTypesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
