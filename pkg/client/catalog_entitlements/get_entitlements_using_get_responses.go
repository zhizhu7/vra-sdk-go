// Code generated by go-swagger; DO NOT EDIT.

package catalog_entitlements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetEntitlementsUsingGETReader is a Reader for the GetEntitlementsUsingGET structure.
type GetEntitlementsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEntitlementsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEntitlementsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetEntitlementsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEntitlementsUsingGETOK creates a GetEntitlementsUsingGETOK with default headers values
func NewGetEntitlementsUsingGETOK() *GetEntitlementsUsingGETOK {
	return &GetEntitlementsUsingGETOK{}
}

/*GetEntitlementsUsingGETOK handles this case with default header values.

OK
*/
type GetEntitlementsUsingGETOK struct {
	Payload []*models.Entitlement
}

func (o *GetEntitlementsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /catalog/api/admin/entitlements][%d] getEntitlementsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetEntitlementsUsingGETOK) GetPayload() []*models.Entitlement {
	return o.Payload
}

func (o *GetEntitlementsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEntitlementsUsingGETUnauthorized creates a GetEntitlementsUsingGETUnauthorized with default headers values
func NewGetEntitlementsUsingGETUnauthorized() *GetEntitlementsUsingGETUnauthorized {
	return &GetEntitlementsUsingGETUnauthorized{}
}

/*GetEntitlementsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetEntitlementsUsingGETUnauthorized struct {
}

func (o *GetEntitlementsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /catalog/api/admin/entitlements][%d] getEntitlementsUsingGETUnauthorized ", 401)
}

func (o *GetEntitlementsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}