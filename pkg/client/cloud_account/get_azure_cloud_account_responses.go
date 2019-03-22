// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetAzureCloudAccountReader is a Reader for the GetAzureCloudAccount structure.
type GetAzureCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAzureCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAzureCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetAzureCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAzureCloudAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAzureCloudAccountOK creates a GetAzureCloudAccountOK with default headers values
func NewGetAzureCloudAccountOK() *GetAzureCloudAccountOK {
	return &GetAzureCloudAccountOK{}
}

/*GetAzureCloudAccountOK handles this case with default header values.

successful operation
*/
type GetAzureCloudAccountOK struct {
	Payload *models.CloudAccountAzure
}

func (o *GetAzureCloudAccountOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-azure/{id}][%d] getAzureCloudAccountOK  %+v", 200, o.Payload)
}

func (o *GetAzureCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountAzure)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAzureCloudAccountForbidden creates a GetAzureCloudAccountForbidden with default headers values
func NewGetAzureCloudAccountForbidden() *GetAzureCloudAccountForbidden {
	return &GetAzureCloudAccountForbidden{}
}

/*GetAzureCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type GetAzureCloudAccountForbidden struct {
}

func (o *GetAzureCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-azure/{id}][%d] getAzureCloudAccountForbidden ", 403)
}

func (o *GetAzureCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAzureCloudAccountNotFound creates a GetAzureCloudAccountNotFound with default headers values
func NewGetAzureCloudAccountNotFound() *GetAzureCloudAccountNotFound {
	return &GetAzureCloudAccountNotFound{}
}

/*GetAzureCloudAccountNotFound handles this case with default header values.

Not Found
*/
type GetAzureCloudAccountNotFound struct {
}

func (o *GetAzureCloudAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-azure/{id}][%d] getAzureCloudAccountNotFound ", 404)
}

func (o *GetAzureCloudAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
