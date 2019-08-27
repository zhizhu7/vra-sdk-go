// Code generated by go-swagger; DO NOT EDIT.

package image_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetImageProfilesReader is a Reader for the GetImageProfiles structure.
type GetImageProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImageProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImageProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetImageProfilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetImageProfilesOK creates a GetImageProfilesOK with default headers values
func NewGetImageProfilesOK() *GetImageProfilesOK {
	return &GetImageProfilesOK{}
}

/*GetImageProfilesOK handles this case with default header values.

successful operation
*/
type GetImageProfilesOK struct {
	Payload *models.ImageProfileResult
}

func (o *GetImageProfilesOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/image-profiles][%d] getImageProfilesOK  %+v", 200, o.Payload)
}

func (o *GetImageProfilesOK) GetPayload() *models.ImageProfileResult {
	return o.Payload
}

func (o *GetImageProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageProfileResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageProfilesForbidden creates a GetImageProfilesForbidden with default headers values
func NewGetImageProfilesForbidden() *GetImageProfilesForbidden {
	return &GetImageProfilesForbidden{}
}

/*GetImageProfilesForbidden handles this case with default header values.

Forbidden
*/
type GetImageProfilesForbidden struct {
}

func (o *GetImageProfilesForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/image-profiles][%d] getImageProfilesForbidden ", 403)
}

func (o *GetImageProfilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
