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

// GetNsxTCloudAccountsReader is a Reader for the GetNsxTCloudAccounts structure.
type GetNsxTCloudAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNsxTCloudAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNsxTCloudAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetNsxTCloudAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNsxTCloudAccountsOK creates a GetNsxTCloudAccountsOK with default headers values
func NewGetNsxTCloudAccountsOK() *GetNsxTCloudAccountsOK {
	return &GetNsxTCloudAccountsOK{}
}

/*GetNsxTCloudAccountsOK handles this case with default header values.

successful operation
*/
type GetNsxTCloudAccountsOK struct {
	Payload *models.CloudAccountNsxTResult
}

func (o *GetNsxTCloudAccountsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-nsx-t][%d] getNsxTCloudAccountsOK  %+v", 200, o.Payload)
}

func (o *GetNsxTCloudAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountNsxTResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNsxTCloudAccountsForbidden creates a GetNsxTCloudAccountsForbidden with default headers values
func NewGetNsxTCloudAccountsForbidden() *GetNsxTCloudAccountsForbidden {
	return &GetNsxTCloudAccountsForbidden{}
}

/*GetNsxTCloudAccountsForbidden handles this case with default header values.

Forbidden
*/
type GetNsxTCloudAccountsForbidden struct {
}

func (o *GetNsxTCloudAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-nsx-t][%d] getNsxTCloudAccountsForbidden ", 403)
}

func (o *GetNsxTCloudAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
