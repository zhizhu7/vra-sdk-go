// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteAwsStorageProfileReader is a Reader for the DeleteAwsStorageProfile structure.
type DeleteAwsStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAwsStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAwsStorageProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteAwsStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAwsStorageProfileNoContent creates a DeleteAwsStorageProfileNoContent with default headers values
func NewDeleteAwsStorageProfileNoContent() *DeleteAwsStorageProfileNoContent {
	return &DeleteAwsStorageProfileNoContent{}
}

/*DeleteAwsStorageProfileNoContent handles this case with default header values.

No Content
*/
type DeleteAwsStorageProfileNoContent struct {
}

func (o *DeleteAwsStorageProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/storage-profiles-aws/{id}][%d] deleteAwsStorageProfileNoContent ", 204)
}

func (o *DeleteAwsStorageProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAwsStorageProfileForbidden creates a DeleteAwsStorageProfileForbidden with default headers values
func NewDeleteAwsStorageProfileForbidden() *DeleteAwsStorageProfileForbidden {
	return &DeleteAwsStorageProfileForbidden{}
}

/*DeleteAwsStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type DeleteAwsStorageProfileForbidden struct {
}

func (o *DeleteAwsStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/storage-profiles-aws/{id}][%d] deleteAwsStorageProfileForbidden ", 403)
}

func (o *DeleteAwsStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
