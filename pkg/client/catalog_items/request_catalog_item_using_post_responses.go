// Code generated by go-swagger; DO NOT EDIT.

package catalog_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// RequestCatalogItemUsingPOSTReader is a Reader for the RequestCatalogItemUsingPOST structure.
type RequestCatalogItemUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestCatalogItemUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRequestCatalogItemUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewRequestCatalogItemUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRequestCatalogItemUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRequestCatalogItemUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRequestCatalogItemUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRequestCatalogItemUsingPOSTOK creates a RequestCatalogItemUsingPOSTOK with default headers values
func NewRequestCatalogItemUsingPOSTOK() *RequestCatalogItemUsingPOSTOK {
	return &RequestCatalogItemUsingPOSTOK{}
}

/*RequestCatalogItemUsingPOSTOK handles this case with default header values.

OK
*/
type RequestCatalogItemUsingPOSTOK struct {
	Payload *models.CatalogItemRequestResponse
}

func (o *RequestCatalogItemUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /catalog/api/items/{id}/request][%d] requestCatalogItemUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *RequestCatalogItemUsingPOSTOK) GetPayload() *models.CatalogItemRequestResponse {
	return o.Payload
}

func (o *RequestCatalogItemUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogItemRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestCatalogItemUsingPOSTCreated creates a RequestCatalogItemUsingPOSTCreated with default headers values
func NewRequestCatalogItemUsingPOSTCreated() *RequestCatalogItemUsingPOSTCreated {
	return &RequestCatalogItemUsingPOSTCreated{}
}

/*RequestCatalogItemUsingPOSTCreated handles this case with default header values.

Created
*/
type RequestCatalogItemUsingPOSTCreated struct {
}

func (o *RequestCatalogItemUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /catalog/api/items/{id}/request][%d] requestCatalogItemUsingPOSTCreated ", 201)
}

func (o *RequestCatalogItemUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestCatalogItemUsingPOSTUnauthorized creates a RequestCatalogItemUsingPOSTUnauthorized with default headers values
func NewRequestCatalogItemUsingPOSTUnauthorized() *RequestCatalogItemUsingPOSTUnauthorized {
	return &RequestCatalogItemUsingPOSTUnauthorized{}
}

/*RequestCatalogItemUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type RequestCatalogItemUsingPOSTUnauthorized struct {
}

func (o *RequestCatalogItemUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /catalog/api/items/{id}/request][%d] requestCatalogItemUsingPOSTUnauthorized ", 401)
}

func (o *RequestCatalogItemUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestCatalogItemUsingPOSTForbidden creates a RequestCatalogItemUsingPOSTForbidden with default headers values
func NewRequestCatalogItemUsingPOSTForbidden() *RequestCatalogItemUsingPOSTForbidden {
	return &RequestCatalogItemUsingPOSTForbidden{}
}

/*RequestCatalogItemUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type RequestCatalogItemUsingPOSTForbidden struct {
}

func (o *RequestCatalogItemUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /catalog/api/items/{id}/request][%d] requestCatalogItemUsingPOSTForbidden ", 403)
}

func (o *RequestCatalogItemUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestCatalogItemUsingPOSTNotFound creates a RequestCatalogItemUsingPOSTNotFound with default headers values
func NewRequestCatalogItemUsingPOSTNotFound() *RequestCatalogItemUsingPOSTNotFound {
	return &RequestCatalogItemUsingPOSTNotFound{}
}

/*RequestCatalogItemUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type RequestCatalogItemUsingPOSTNotFound struct {
}

func (o *RequestCatalogItemUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /catalog/api/items/{id}/request][%d] requestCatalogItemUsingPOSTNotFound ", 404)
}

func (o *RequestCatalogItemUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
