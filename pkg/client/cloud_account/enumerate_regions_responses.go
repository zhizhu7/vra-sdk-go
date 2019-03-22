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

// EnumerateRegionsReader is a Reader for the EnumerateRegions structure.
type EnumerateRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnumerateRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEnumerateRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEnumerateRegionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEnumerateRegionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEnumerateRegionsOK creates a EnumerateRegionsOK with default headers values
func NewEnumerateRegionsOK() *EnumerateRegionsOK {
	return &EnumerateRegionsOK{}
}

/*EnumerateRegionsOK handles this case with default header values.

successful operation
*/
type EnumerateRegionsOK struct {
	Payload *models.CloudAccountRegions
}

func (o *EnumerateRegionsOK) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts/region-enumeration][%d] enumerateRegionsOK  %+v", 200, o.Payload)
}

func (o *EnumerateRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountRegions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumerateRegionsBadRequest creates a EnumerateRegionsBadRequest with default headers values
func NewEnumerateRegionsBadRequest() *EnumerateRegionsBadRequest {
	return &EnumerateRegionsBadRequest{}
}

/*EnumerateRegionsBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type EnumerateRegionsBadRequest struct {
}

func (o *EnumerateRegionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts/region-enumeration][%d] enumerateRegionsBadRequest ", 400)
}

func (o *EnumerateRegionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewEnumerateRegionsForbidden creates a EnumerateRegionsForbidden with default headers values
func NewEnumerateRegionsForbidden() *EnumerateRegionsForbidden {
	return &EnumerateRegionsForbidden{}
}

/*EnumerateRegionsForbidden handles this case with default header values.

Forbidden
*/
type EnumerateRegionsForbidden struct {
}

func (o *EnumerateRegionsForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts/region-enumeration][%d] enumerateRegionsForbidden ", 403)
}

func (o *EnumerateRegionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
