// Code generated by go-swagger; DO NOT EDIT.

package blueprint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetBlueprintInputsSchemaUsingGETReader is a Reader for the GetBlueprintInputsSchemaUsingGET structure.
type GetBlueprintInputsSchemaUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlueprintInputsSchemaUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBlueprintInputsSchemaUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetBlueprintInputsSchemaUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetBlueprintInputsSchemaUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetBlueprintInputsSchemaUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlueprintInputsSchemaUsingGETOK creates a GetBlueprintInputsSchemaUsingGETOK with default headers values
func NewGetBlueprintInputsSchemaUsingGETOK() *GetBlueprintInputsSchemaUsingGETOK {
	return &GetBlueprintInputsSchemaUsingGETOK{}
}

/*GetBlueprintInputsSchemaUsingGETOK handles this case with default header values.

OK
*/
type GetBlueprintInputsSchemaUsingGETOK struct {
	Payload *models.PropertyDefinition
}

func (o *GetBlueprintInputsSchemaUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/inputs-schema][%d] getBlueprintInputsSchemaUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetBlueprintInputsSchemaUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PropertyDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlueprintInputsSchemaUsingGETUnauthorized creates a GetBlueprintInputsSchemaUsingGETUnauthorized with default headers values
func NewGetBlueprintInputsSchemaUsingGETUnauthorized() *GetBlueprintInputsSchemaUsingGETUnauthorized {
	return &GetBlueprintInputsSchemaUsingGETUnauthorized{}
}

/*GetBlueprintInputsSchemaUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetBlueprintInputsSchemaUsingGETUnauthorized struct {
}

func (o *GetBlueprintInputsSchemaUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/inputs-schema][%d] getBlueprintInputsSchemaUsingGETUnauthorized ", 401)
}

func (o *GetBlueprintInputsSchemaUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintInputsSchemaUsingGETForbidden creates a GetBlueprintInputsSchemaUsingGETForbidden with default headers values
func NewGetBlueprintInputsSchemaUsingGETForbidden() *GetBlueprintInputsSchemaUsingGETForbidden {
	return &GetBlueprintInputsSchemaUsingGETForbidden{}
}

/*GetBlueprintInputsSchemaUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetBlueprintInputsSchemaUsingGETForbidden struct {
}

func (o *GetBlueprintInputsSchemaUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/inputs-schema][%d] getBlueprintInputsSchemaUsingGETForbidden ", 403)
}

func (o *GetBlueprintInputsSchemaUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintInputsSchemaUsingGETNotFound creates a GetBlueprintInputsSchemaUsingGETNotFound with default headers values
func NewGetBlueprintInputsSchemaUsingGETNotFound() *GetBlueprintInputsSchemaUsingGETNotFound {
	return &GetBlueprintInputsSchemaUsingGETNotFound{}
}

/*GetBlueprintInputsSchemaUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetBlueprintInputsSchemaUsingGETNotFound struct {
}

func (o *GetBlueprintInputsSchemaUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/inputs-schema][%d] getBlueprintInputsSchemaUsingGETNotFound ", 404)
}

func (o *GetBlueprintInputsSchemaUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
