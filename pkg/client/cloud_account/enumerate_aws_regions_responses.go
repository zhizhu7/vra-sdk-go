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

// EnumerateAwsRegionsReader is a Reader for the EnumerateAwsRegions structure.
type EnumerateAwsRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnumerateAwsRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEnumerateAwsRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEnumerateAwsRegionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEnumerateAwsRegionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEnumerateAwsRegionsOK creates a EnumerateAwsRegionsOK with default headers values
func NewEnumerateAwsRegionsOK() *EnumerateAwsRegionsOK {
	return &EnumerateAwsRegionsOK{}
}

/*EnumerateAwsRegionsOK handles this case with default header values.

successful operation
*/
type EnumerateAwsRegionsOK struct {
	Payload *models.CloudAccountRegions
}

func (o *EnumerateAwsRegionsOK) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-aws/region-enumeration][%d] enumerateAwsRegionsOK  %+v", 200, o.Payload)
}

func (o *EnumerateAwsRegionsOK) GetPayload() *models.CloudAccountRegions {
	return o.Payload
}

func (o *EnumerateAwsRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountRegions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumerateAwsRegionsBadRequest creates a EnumerateAwsRegionsBadRequest with default headers values
func NewEnumerateAwsRegionsBadRequest() *EnumerateAwsRegionsBadRequest {
	return &EnumerateAwsRegionsBadRequest{}
}

/*EnumerateAwsRegionsBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type EnumerateAwsRegionsBadRequest struct {
	Payload *models.Error
}

func (o *EnumerateAwsRegionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-aws/region-enumeration][%d] enumerateAwsRegionsBadRequest  %+v", 400, o.Payload)
}

func (o *EnumerateAwsRegionsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *EnumerateAwsRegionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEnumerateAwsRegionsForbidden creates a EnumerateAwsRegionsForbidden with default headers values
func NewEnumerateAwsRegionsForbidden() *EnumerateAwsRegionsForbidden {
	return &EnumerateAwsRegionsForbidden{}
}

/*EnumerateAwsRegionsForbidden handles this case with default header values.

Forbidden
*/
type EnumerateAwsRegionsForbidden struct {
}

func (o *EnumerateAwsRegionsForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts-aws/region-enumeration][%d] enumerateAwsRegionsForbidden ", 403)
}

func (o *EnumerateAwsRegionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
