// Code generated by go-swagger; DO NOT EDIT.

package security_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetSecurityGroupsReader is a Reader for the GetSecurityGroups structure.
type GetSecurityGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSecurityGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetSecurityGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSecurityGroupsOK creates a GetSecurityGroupsOK with default headers values
func NewGetSecurityGroupsOK() *GetSecurityGroupsOK {
	return &GetSecurityGroupsOK{}
}

/*GetSecurityGroupsOK handles this case with default header values.

successful operation
*/
type GetSecurityGroupsOK struct {
	Payload *models.SecurityGroupResult
}

func (o *GetSecurityGroupsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/security-groups][%d] getSecurityGroupsOK  %+v", 200, o.Payload)
}

func (o *GetSecurityGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityGroupResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityGroupsForbidden creates a GetSecurityGroupsForbidden with default headers values
func NewGetSecurityGroupsForbidden() *GetSecurityGroupsForbidden {
	return &GetSecurityGroupsForbidden{}
}

/*GetSecurityGroupsForbidden handles this case with default header values.

Forbidden
*/
type GetSecurityGroupsForbidden struct {
}

func (o *GetSecurityGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/security-groups][%d] getSecurityGroupsForbidden ", 403)
}

func (o *GetSecurityGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
