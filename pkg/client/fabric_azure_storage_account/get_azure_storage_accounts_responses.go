// Code generated by go-swagger; DO NOT EDIT.

package fabric_azure_storage_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetAzureStorageAccountsReader is a Reader for the GetAzureStorageAccounts structure.
type GetAzureStorageAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAzureStorageAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAzureStorageAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetAzureStorageAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAzureStorageAccountsOK creates a GetAzureStorageAccountsOK with default headers values
func NewGetAzureStorageAccountsOK() *GetAzureStorageAccountsOK {
	return &GetAzureStorageAccountsOK{}
}

/*GetAzureStorageAccountsOK handles this case with default header values.

successful operation
*/
type GetAzureStorageAccountsOK struct {
	Payload *models.FabricAzureStorageAccountResult
}

func (o *GetAzureStorageAccountsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-storage-accounts][%d] getAzureStorageAccountsOK  %+v", 200, o.Payload)
}

func (o *GetAzureStorageAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricAzureStorageAccountResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAzureStorageAccountsForbidden creates a GetAzureStorageAccountsForbidden with default headers values
func NewGetAzureStorageAccountsForbidden() *GetAzureStorageAccountsForbidden {
	return &GetAzureStorageAccountsForbidden{}
}

/*GetAzureStorageAccountsForbidden handles this case with default header values.

Forbidden
*/
type GetAzureStorageAccountsForbidden struct {
}

func (o *GetAzureStorageAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-storage-accounts][%d] getAzureStorageAccountsForbidden ", 403)
}

func (o *GetAzureStorageAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
