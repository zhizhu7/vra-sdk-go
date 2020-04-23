// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateVmcCloudAccountReader is a Reader for the CreateVmcCloudAccount structure.
type CreateVmcCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVmcCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateVmcCloudAccountCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVmcCloudAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateVmcCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateVmcCloudAccountCreated creates a CreateVmcCloudAccountCreated with default headers values
func NewCreateVmcCloudAccountCreated() *CreateVmcCloudAccountCreated {
	return &CreateVmcCloudAccountCreated{}
}

/*CreateVmcCloudAccountCreated handles this case with default header values.

successful operation
*/
type CreateVmcCloudAccountCreated struct {
	Payload *models.CloudAccountVmc
}

func (o *CreateVmcCloudAccountCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc][%d] createVmcCloudAccountCreated  %+v", 201, o.Payload)
}

func (o *CreateVmcCloudAccountCreated) GetPayload() *models.CloudAccountVmc {
	return o.Payload
}

func (o *CreateVmcCloudAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountVmc)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVmcCloudAccountBadRequest creates a CreateVmcCloudAccountBadRequest with default headers values
func NewCreateVmcCloudAccountBadRequest() *CreateVmcCloudAccountBadRequest {
	return &CreateVmcCloudAccountBadRequest{}
}

/*CreateVmcCloudAccountBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type CreateVmcCloudAccountBadRequest struct {
}

func (o *CreateVmcCloudAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc][%d] createVmcCloudAccountBadRequest ", 400)
}

func (o *CreateVmcCloudAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateVmcCloudAccountForbidden creates a CreateVmcCloudAccountForbidden with default headers values
func NewCreateVmcCloudAccountForbidden() *CreateVmcCloudAccountForbidden {
	return &CreateVmcCloudAccountForbidden{}
}

/*CreateVmcCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type CreateVmcCloudAccountForbidden struct {
}

func (o *CreateVmcCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-vmc][%d] createVmcCloudAccountForbidden ", 403)
}

func (o *CreateVmcCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}