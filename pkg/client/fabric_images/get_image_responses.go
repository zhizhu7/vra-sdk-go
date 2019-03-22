// Code generated by go-swagger; DO NOT EDIT.

package fabric_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetImageReader is a Reader for the GetImage structure.
type GetImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetImageOK creates a GetImageOK with default headers values
func NewGetImageOK() *GetImageOK {
	return &GetImageOK{}
}

/*GetImageOK handles this case with default header values.

successful operation
*/
type GetImageOK struct {
	Payload *models.FabricImage
}

func (o *GetImageOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-images/{id}][%d] getImageOK  %+v", 200, o.Payload)
}

func (o *GetImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricImage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageForbidden creates a GetImageForbidden with default headers values
func NewGetImageForbidden() *GetImageForbidden {
	return &GetImageForbidden{}
}

/*GetImageForbidden handles this case with default header values.

Forbidden
*/
type GetImageForbidden struct {
}

func (o *GetImageForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-images/{id}][%d] getImageForbidden ", 403)
}

func (o *GetImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetImageNotFound creates a GetImageNotFound with default headers values
func NewGetImageNotFound() *GetImageNotFound {
	return &GetImageNotFound{}
}

/*GetImageNotFound handles this case with default header values.

Not Found
*/
type GetImageNotFound struct {
}

func (o *GetImageNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-images/{id}][%d] getImageNotFound ", 404)
}

func (o *GetImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
