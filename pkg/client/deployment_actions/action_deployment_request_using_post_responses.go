// Code generated by go-swagger; DO NOT EDIT.

package deployment_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ActionDeploymentRequestUsingPOSTReader is a Reader for the ActionDeploymentRequestUsingPOST structure.
type ActionDeploymentRequestUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionDeploymentRequestUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActionDeploymentRequestUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewActionDeploymentRequestUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewActionDeploymentRequestUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewActionDeploymentRequestUsingPOSTOK creates a ActionDeploymentRequestUsingPOSTOK with default headers values
func NewActionDeploymentRequestUsingPOSTOK() *ActionDeploymentRequestUsingPOSTOK {
	return &ActionDeploymentRequestUsingPOSTOK{}
}

/*ActionDeploymentRequestUsingPOSTOK handles this case with default header values.

OK
*/
type ActionDeploymentRequestUsingPOSTOK struct {
}

func (o *ActionDeploymentRequestUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /deployment/api/requests/{requestId}][%d] actionDeploymentRequestUsingPOSTOK ", 200)
}

func (o *ActionDeploymentRequestUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionDeploymentRequestUsingPOSTUnauthorized creates a ActionDeploymentRequestUsingPOSTUnauthorized with default headers values
func NewActionDeploymentRequestUsingPOSTUnauthorized() *ActionDeploymentRequestUsingPOSTUnauthorized {
	return &ActionDeploymentRequestUsingPOSTUnauthorized{}
}

/*ActionDeploymentRequestUsingPOSTUnauthorized handles this case with default header values.

Unauthorized
*/
type ActionDeploymentRequestUsingPOSTUnauthorized struct {
}

func (o *ActionDeploymentRequestUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /deployment/api/requests/{requestId}][%d] actionDeploymentRequestUsingPOSTUnauthorized ", 401)
}

func (o *ActionDeploymentRequestUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewActionDeploymentRequestUsingPOSTNotFound creates a ActionDeploymentRequestUsingPOSTNotFound with default headers values
func NewActionDeploymentRequestUsingPOSTNotFound() *ActionDeploymentRequestUsingPOSTNotFound {
	return &ActionDeploymentRequestUsingPOSTNotFound{}
}

/*ActionDeploymentRequestUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type ActionDeploymentRequestUsingPOSTNotFound struct {
}

func (o *ActionDeploymentRequestUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /deployment/api/requests/{requestId}][%d] actionDeploymentRequestUsingPOSTNotFound ", 404)
}

func (o *ActionDeploymentRequestUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}