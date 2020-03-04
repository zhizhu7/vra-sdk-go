// Code generated by go-swagger; DO NOT EDIT.

package marketplace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// FilterItemsUsingGETReader is a Reader for the FilterItemsUsingGET structure.
type FilterItemsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilterItemsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFilterItemsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewFilterItemsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFilterItemsUsingGETOK creates a FilterItemsUsingGETOK with default headers values
func NewFilterItemsUsingGETOK() *FilterItemsUsingGETOK {
	return &FilterItemsUsingGETOK{}
}

/*FilterItemsUsingGETOK handles this case with default header values.

Filter entries
*/
type FilterItemsUsingGETOK struct {
	Payload *models.MarketplaceFilterEntries
}

func (o *FilterItemsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /content/api/marketplace/sources/{sourceId}/filters/{filterId}][%d] filterItemsUsingGETOK  %+v", 200, o.Payload)
}

func (o *FilterItemsUsingGETOK) GetPayload() *models.MarketplaceFilterEntries {
	return o.Payload
}

func (o *FilterItemsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MarketplaceFilterEntries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilterItemsUsingGETNotFound creates a FilterItemsUsingGETNotFound with default headers values
func NewFilterItemsUsingGETNotFound() *FilterItemsUsingGETNotFound {
	return &FilterItemsUsingGETNotFound{}
}

/*FilterItemsUsingGETNotFound handles this case with default header values.

Source or Filter not found
*/
type FilterItemsUsingGETNotFound struct {
}

func (o *FilterItemsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /content/api/marketplace/sources/{sourceId}/filters/{filterId}][%d] filterItemsUsingGETNotFound ", 404)
}

func (o *FilterItemsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}