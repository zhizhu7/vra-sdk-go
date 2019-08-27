// Code generated by go-swagger; DO NOT EDIT.

package fabric_vsphere_storage_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetFabricVSphereStoragePoliciesReader is a Reader for the GetFabricVSphereStoragePolicies structure.
type GetFabricVSphereStoragePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFabricVSphereStoragePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFabricVSphereStoragePoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFabricVSphereStoragePoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFabricVSphereStoragePoliciesOK creates a GetFabricVSphereStoragePoliciesOK with default headers values
func NewGetFabricVSphereStoragePoliciesOK() *GetFabricVSphereStoragePoliciesOK {
	return &GetFabricVSphereStoragePoliciesOK{}
}

/*GetFabricVSphereStoragePoliciesOK handles this case with default header values.

successful operation
*/
type GetFabricVSphereStoragePoliciesOK struct {
	Payload *models.FabricVsphereStoragePolicyResult
}

func (o *GetFabricVSphereStoragePoliciesOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-vsphere-storage-policies][%d] getFabricVSphereStoragePoliciesOK  %+v", 200, o.Payload)
}

func (o *GetFabricVSphereStoragePoliciesOK) GetPayload() *models.FabricVsphereStoragePolicyResult {
	return o.Payload
}

func (o *GetFabricVSphereStoragePoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricVsphereStoragePolicyResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricVSphereStoragePoliciesForbidden creates a GetFabricVSphereStoragePoliciesForbidden with default headers values
func NewGetFabricVSphereStoragePoliciesForbidden() *GetFabricVSphereStoragePoliciesForbidden {
	return &GetFabricVSphereStoragePoliciesForbidden{}
}

/*GetFabricVSphereStoragePoliciesForbidden handles this case with default header values.

Forbidden
*/
type GetFabricVSphereStoragePoliciesForbidden struct {
}

func (o *GetFabricVSphereStoragePoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-vsphere-storage-policies][%d] getFabricVSphereStoragePoliciesForbidden ", 403)
}

func (o *GetFabricVSphereStoragePoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
