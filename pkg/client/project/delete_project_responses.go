// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteProjectReader is a Reader for the DeleteProject structure.
type DeleteProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteProjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteProjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteProjectNoContent creates a DeleteProjectNoContent with default headers values
func NewDeleteProjectNoContent() *DeleteProjectNoContent {
	return &DeleteProjectNoContent{}
}

/*DeleteProjectNoContent handles this case with default header values.

No Content
*/
type DeleteProjectNoContent struct {
}

func (o *DeleteProjectNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/projects/{id}][%d] deleteProjectNoContent ", 204)
}

func (o *DeleteProjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectForbidden creates a DeleteProjectForbidden with default headers values
func NewDeleteProjectForbidden() *DeleteProjectForbidden {
	return &DeleteProjectForbidden{}
}

/*DeleteProjectForbidden handles this case with default header values.

Forbidden
*/
type DeleteProjectForbidden struct {
}

func (o *DeleteProjectForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/projects/{id}][%d] deleteProjectForbidden ", 403)
}

func (o *DeleteProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProjectConflict creates a DeleteProjectConflict with default headers values
func NewDeleteProjectConflict() *DeleteProjectConflict {
	return &DeleteProjectConflict{}
}

/*DeleteProjectConflict handles this case with default header values.

Conflict, when the project is in use
*/
type DeleteProjectConflict struct {
}

func (o *DeleteProjectConflict) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/projects/{id}][%d] deleteProjectConflict ", 409)
}

func (o *DeleteProjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
