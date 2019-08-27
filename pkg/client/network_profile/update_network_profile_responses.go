// Code generated by go-swagger; DO NOT EDIT.

package network_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateNetworkProfileReader is a Reader for the UpdateNetworkProfile structure.
type UpdateNetworkProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateNetworkProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateNetworkProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateNetworkProfileOK creates a UpdateNetworkProfileOK with default headers values
func NewUpdateNetworkProfileOK() *UpdateNetworkProfileOK {
	return &UpdateNetworkProfileOK{}
}

/*UpdateNetworkProfileOK handles this case with default header values.

successful operation
*/
type UpdateNetworkProfileOK struct {
	Payload *models.NetworkProfile
}

func (o *UpdateNetworkProfileOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/network-profiles/{id}][%d] updateNetworkProfileOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkProfileOK) GetPayload() *models.NetworkProfile {
	return o.Payload
}

func (o *UpdateNetworkProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNetworkProfileForbidden creates a UpdateNetworkProfileForbidden with default headers values
func NewUpdateNetworkProfileForbidden() *UpdateNetworkProfileForbidden {
	return &UpdateNetworkProfileForbidden{}
}

/*UpdateNetworkProfileForbidden handles this case with default header values.

Forbidden
*/
type UpdateNetworkProfileForbidden struct {
}

func (o *UpdateNetworkProfileForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/network-profiles/{id}][%d] updateNetworkProfileForbidden ", 403)
}

func (o *UpdateNetworkProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateNetworkProfileNotFound creates a UpdateNetworkProfileNotFound with default headers values
func NewUpdateNetworkProfileNotFound() *UpdateNetworkProfileNotFound {
	return &UpdateNetworkProfileNotFound{}
}

/*UpdateNetworkProfileNotFound handles this case with default header values.

Not Found
*/
type UpdateNetworkProfileNotFound struct {
}

func (o *UpdateNetworkProfileNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/network-profiles/{id}][%d] updateNetworkProfileNotFound ", 404)
}

func (o *UpdateNetworkProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
