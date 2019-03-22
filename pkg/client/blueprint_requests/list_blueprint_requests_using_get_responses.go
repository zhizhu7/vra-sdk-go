// Code generated by go-swagger; DO NOT EDIT.

package blueprint_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// ListBlueprintRequestsUsingGETReader is a Reader for the ListBlueprintRequestsUsingGET structure.
type ListBlueprintRequestsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBlueprintRequestsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListBlueprintRequestsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListBlueprintRequestsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListBlueprintRequestsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListBlueprintRequestsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListBlueprintRequestsUsingGETOK creates a ListBlueprintRequestsUsingGETOK with default headers values
func NewListBlueprintRequestsUsingGETOK() *ListBlueprintRequestsUsingGETOK {
	return &ListBlueprintRequestsUsingGETOK{}
}

/*ListBlueprintRequestsUsingGETOK handles this case with default header values.

OK
*/
type ListBlueprintRequestsUsingGETOK struct {
	Payload *models.QueryResultBlueprintRequest
}

func (o *ListBlueprintRequestsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-requests][%d] listBlueprintRequestsUsingGETOK  %+v", 200, o.Payload)
}

func (o *ListBlueprintRequestsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryResultBlueprintRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBlueprintRequestsUsingGETUnauthorized creates a ListBlueprintRequestsUsingGETUnauthorized with default headers values
func NewListBlueprintRequestsUsingGETUnauthorized() *ListBlueprintRequestsUsingGETUnauthorized {
	return &ListBlueprintRequestsUsingGETUnauthorized{}
}

/*ListBlueprintRequestsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type ListBlueprintRequestsUsingGETUnauthorized struct {
}

func (o *ListBlueprintRequestsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-requests][%d] listBlueprintRequestsUsingGETUnauthorized ", 401)
}

func (o *ListBlueprintRequestsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListBlueprintRequestsUsingGETForbidden creates a ListBlueprintRequestsUsingGETForbidden with default headers values
func NewListBlueprintRequestsUsingGETForbidden() *ListBlueprintRequestsUsingGETForbidden {
	return &ListBlueprintRequestsUsingGETForbidden{}
}

/*ListBlueprintRequestsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type ListBlueprintRequestsUsingGETForbidden struct {
}

func (o *ListBlueprintRequestsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-requests][%d] listBlueprintRequestsUsingGETForbidden ", 403)
}

func (o *ListBlueprintRequestsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListBlueprintRequestsUsingGETNotFound creates a ListBlueprintRequestsUsingGETNotFound with default headers values
func NewListBlueprintRequestsUsingGETNotFound() *ListBlueprintRequestsUsingGETNotFound {
	return &ListBlueprintRequestsUsingGETNotFound{}
}

/*ListBlueprintRequestsUsingGETNotFound handles this case with default header values.

Not Found
*/
type ListBlueprintRequestsUsingGETNotFound struct {
}

func (o *ListBlueprintRequestsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-requests][%d] listBlueprintRequestsUsingGETNotFound ", 404)
}

func (o *ListBlueprintRequestsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
