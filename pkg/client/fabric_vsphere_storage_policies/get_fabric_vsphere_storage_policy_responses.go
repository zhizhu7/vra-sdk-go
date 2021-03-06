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

// GetFabricVSphereStoragePolicyReader is a Reader for the GetFabricVSphereStoragePolicy structure.
type GetFabricVSphereStoragePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFabricVSphereStoragePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFabricVSphereStoragePolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFabricVSphereStoragePolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFabricVSphereStoragePolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFabricVSphereStoragePolicyOK creates a GetFabricVSphereStoragePolicyOK with default headers values
func NewGetFabricVSphereStoragePolicyOK() *GetFabricVSphereStoragePolicyOK {
	return &GetFabricVSphereStoragePolicyOK{}
}

/*GetFabricVSphereStoragePolicyOK handles this case with default header values.

successful operation
*/
type GetFabricVSphereStoragePolicyOK struct {
	Payload *models.FabricVsphereStoragePolicy
}

func (o *GetFabricVSphereStoragePolicyOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-vsphere-storage-policies/{id}][%d] getFabricVSphereStoragePolicyOK  %+v", 200, o.Payload)
}

func (o *GetFabricVSphereStoragePolicyOK) GetPayload() *models.FabricVsphereStoragePolicy {
	return o.Payload
}

func (o *GetFabricVSphereStoragePolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricVsphereStoragePolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricVSphereStoragePolicyForbidden creates a GetFabricVSphereStoragePolicyForbidden with default headers values
func NewGetFabricVSphereStoragePolicyForbidden() *GetFabricVSphereStoragePolicyForbidden {
	return &GetFabricVSphereStoragePolicyForbidden{}
}

/*GetFabricVSphereStoragePolicyForbidden handles this case with default header values.

Forbidden
*/
type GetFabricVSphereStoragePolicyForbidden struct {
}

func (o *GetFabricVSphereStoragePolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-vsphere-storage-policies/{id}][%d] getFabricVSphereStoragePolicyForbidden ", 403)
}

func (o *GetFabricVSphereStoragePolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFabricVSphereStoragePolicyNotFound creates a GetFabricVSphereStoragePolicyNotFound with default headers values
func NewGetFabricVSphereStoragePolicyNotFound() *GetFabricVSphereStoragePolicyNotFound {
	return &GetFabricVSphereStoragePolicyNotFound{}
}

/*GetFabricVSphereStoragePolicyNotFound handles this case with default header values.

Not Found
*/
type GetFabricVSphereStoragePolicyNotFound struct {
	Payload *models.Error
}

func (o *GetFabricVSphereStoragePolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-vsphere-storage-policies/{id}][%d] getFabricVSphereStoragePolicyNotFound  %+v", 404, o.Payload)
}

func (o *GetFabricVSphereStoragePolicyNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFabricVSphereStoragePolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
