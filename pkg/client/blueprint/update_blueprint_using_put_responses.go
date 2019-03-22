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

// UpdateBlueprintUsingPUTReader is a Reader for the UpdateBlueprintUsingPUT structure.
type UpdateBlueprintUsingPUTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBlueprintUsingPUTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateBlueprintUsingPUTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewUpdateBlueprintUsingPUTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewUpdateBlueprintUsingPUTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateBlueprintUsingPUTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateBlueprintUsingPUTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateBlueprintUsingPUTOK creates a UpdateBlueprintUsingPUTOK with default headers values
func NewUpdateBlueprintUsingPUTOK() *UpdateBlueprintUsingPUTOK {
	return &UpdateBlueprintUsingPUTOK{}
}

/*UpdateBlueprintUsingPUTOK handles this case with default header values.

OK
*/
type UpdateBlueprintUsingPUTOK struct {
	Payload *models.BlueprintLinkResponse
}

func (o *UpdateBlueprintUsingPUTOK) Error() string {
	return fmt.Sprintf("[PUT /blueprint/api/blueprints/{blueprintId}][%d] updateBlueprintUsingPUTOK  %+v", 200, o.Payload)
}

func (o *UpdateBlueprintUsingPUTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlueprintLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBlueprintUsingPUTCreated creates a UpdateBlueprintUsingPUTCreated with default headers values
func NewUpdateBlueprintUsingPUTCreated() *UpdateBlueprintUsingPUTCreated {
	return &UpdateBlueprintUsingPUTCreated{}
}

/*UpdateBlueprintUsingPUTCreated handles this case with default header values.

Created
*/
type UpdateBlueprintUsingPUTCreated struct {
}

func (o *UpdateBlueprintUsingPUTCreated) Error() string {
	return fmt.Sprintf("[PUT /blueprint/api/blueprints/{blueprintId}][%d] updateBlueprintUsingPUTCreated ", 201)
}

func (o *UpdateBlueprintUsingPUTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBlueprintUsingPUTUnauthorized creates a UpdateBlueprintUsingPUTUnauthorized with default headers values
func NewUpdateBlueprintUsingPUTUnauthorized() *UpdateBlueprintUsingPUTUnauthorized {
	return &UpdateBlueprintUsingPUTUnauthorized{}
}

/*UpdateBlueprintUsingPUTUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateBlueprintUsingPUTUnauthorized struct {
}

func (o *UpdateBlueprintUsingPUTUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /blueprint/api/blueprints/{blueprintId}][%d] updateBlueprintUsingPUTUnauthorized ", 401)
}

func (o *UpdateBlueprintUsingPUTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBlueprintUsingPUTForbidden creates a UpdateBlueprintUsingPUTForbidden with default headers values
func NewUpdateBlueprintUsingPUTForbidden() *UpdateBlueprintUsingPUTForbidden {
	return &UpdateBlueprintUsingPUTForbidden{}
}

/*UpdateBlueprintUsingPUTForbidden handles this case with default header values.

Forbidden
*/
type UpdateBlueprintUsingPUTForbidden struct {
}

func (o *UpdateBlueprintUsingPUTForbidden) Error() string {
	return fmt.Sprintf("[PUT /blueprint/api/blueprints/{blueprintId}][%d] updateBlueprintUsingPUTForbidden ", 403)
}

func (o *UpdateBlueprintUsingPUTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateBlueprintUsingPUTNotFound creates a UpdateBlueprintUsingPUTNotFound with default headers values
func NewUpdateBlueprintUsingPUTNotFound() *UpdateBlueprintUsingPUTNotFound {
	return &UpdateBlueprintUsingPUTNotFound{}
}

/*UpdateBlueprintUsingPUTNotFound handles this case with default header values.

Not Found
*/
type UpdateBlueprintUsingPUTNotFound struct {
}

func (o *UpdateBlueprintUsingPUTNotFound) Error() string {
	return fmt.Sprintf("[PUT /blueprint/api/blueprints/{blueprintId}][%d] updateBlueprintUsingPUTNotFound ", 404)
}

func (o *UpdateBlueprintUsingPUTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
