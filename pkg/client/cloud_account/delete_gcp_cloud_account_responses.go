// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteGcpCloudAccountReader is a Reader for the DeleteGcpCloudAccount structure.
type DeleteGcpCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGcpCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteGcpCloudAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteGcpCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteGcpCloudAccountNoContent creates a DeleteGcpCloudAccountNoContent with default headers values
func NewDeleteGcpCloudAccountNoContent() *DeleteGcpCloudAccountNoContent {
	return &DeleteGcpCloudAccountNoContent{}
}

/*DeleteGcpCloudAccountNoContent handles this case with default header values.

No Content
*/
type DeleteGcpCloudAccountNoContent struct {
}

func (o *DeleteGcpCloudAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-gcp/{id}][%d] deleteGcpCloudAccountNoContent ", 204)
}

func (o *DeleteGcpCloudAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGcpCloudAccountForbidden creates a DeleteGcpCloudAccountForbidden with default headers values
func NewDeleteGcpCloudAccountForbidden() *DeleteGcpCloudAccountForbidden {
	return &DeleteGcpCloudAccountForbidden{}
}

/*DeleteGcpCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type DeleteGcpCloudAccountForbidden struct {
}

func (o *DeleteGcpCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/cloud-accounts-gcp/{id}][%d] deleteGcpCloudAccountForbidden ", 403)
}

func (o *DeleteGcpCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
