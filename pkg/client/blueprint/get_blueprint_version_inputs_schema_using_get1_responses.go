// Code generated by go-swagger; DO NOT EDIT.

package blueprint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetBlueprintVersionInputsSchemaUsingGET1Reader is a Reader for the GetBlueprintVersionInputsSchemaUsingGET1 structure.
type GetBlueprintVersionInputsSchemaUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlueprintVersionInputsSchemaUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBlueprintVersionInputsSchemaUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBlueprintVersionInputsSchemaUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBlueprintVersionInputsSchemaUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBlueprintVersionInputsSchemaUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlueprintVersionInputsSchemaUsingGET1OK creates a GetBlueprintVersionInputsSchemaUsingGET1OK with default headers values
func NewGetBlueprintVersionInputsSchemaUsingGET1OK() *GetBlueprintVersionInputsSchemaUsingGET1OK {
	return &GetBlueprintVersionInputsSchemaUsingGET1OK{}
}

/*GetBlueprintVersionInputsSchemaUsingGET1OK handles this case with default header values.

OK
*/
type GetBlueprintVersionInputsSchemaUsingGET1OK struct {
	Payload *models.PropertyDefinition
}

func (o *GetBlueprintVersionInputsSchemaUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions/{version}/inputs-schema][%d] getBlueprintVersionInputsSchemaUsingGET1OK  %+v", 200, o.Payload)
}

func (o *GetBlueprintVersionInputsSchemaUsingGET1OK) GetPayload() *models.PropertyDefinition {
	return o.Payload
}

func (o *GetBlueprintVersionInputsSchemaUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PropertyDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlueprintVersionInputsSchemaUsingGET1Unauthorized creates a GetBlueprintVersionInputsSchemaUsingGET1Unauthorized with default headers values
func NewGetBlueprintVersionInputsSchemaUsingGET1Unauthorized() *GetBlueprintVersionInputsSchemaUsingGET1Unauthorized {
	return &GetBlueprintVersionInputsSchemaUsingGET1Unauthorized{}
}

/*GetBlueprintVersionInputsSchemaUsingGET1Unauthorized handles this case with default header values.

Unauthorized
*/
type GetBlueprintVersionInputsSchemaUsingGET1Unauthorized struct {
}

func (o *GetBlueprintVersionInputsSchemaUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions/{version}/inputs-schema][%d] getBlueprintVersionInputsSchemaUsingGET1Unauthorized ", 401)
}

func (o *GetBlueprintVersionInputsSchemaUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintVersionInputsSchemaUsingGET1Forbidden creates a GetBlueprintVersionInputsSchemaUsingGET1Forbidden with default headers values
func NewGetBlueprintVersionInputsSchemaUsingGET1Forbidden() *GetBlueprintVersionInputsSchemaUsingGET1Forbidden {
	return &GetBlueprintVersionInputsSchemaUsingGET1Forbidden{}
}

/*GetBlueprintVersionInputsSchemaUsingGET1Forbidden handles this case with default header values.

Forbidden
*/
type GetBlueprintVersionInputsSchemaUsingGET1Forbidden struct {
}

func (o *GetBlueprintVersionInputsSchemaUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions/{version}/inputs-schema][%d] getBlueprintVersionInputsSchemaUsingGET1Forbidden ", 403)
}

func (o *GetBlueprintVersionInputsSchemaUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintVersionInputsSchemaUsingGET1NotFound creates a GetBlueprintVersionInputsSchemaUsingGET1NotFound with default headers values
func NewGetBlueprintVersionInputsSchemaUsingGET1NotFound() *GetBlueprintVersionInputsSchemaUsingGET1NotFound {
	return &GetBlueprintVersionInputsSchemaUsingGET1NotFound{}
}

/*GetBlueprintVersionInputsSchemaUsingGET1NotFound handles this case with default header values.

Not Found
*/
type GetBlueprintVersionInputsSchemaUsingGET1NotFound struct {
}

func (o *GetBlueprintVersionInputsSchemaUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions/{version}/inputs-schema][%d] getBlueprintVersionInputsSchemaUsingGET1NotFound ", 404)
}

func (o *GetBlueprintVersionInputsSchemaUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}