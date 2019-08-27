// Code generated by go-swagger; DO NOT EDIT.

package load_balancer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetLoadBalancersReader is a Reader for the GetLoadBalancers structure.
type GetLoadBalancersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLoadBalancersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLoadBalancersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetLoadBalancersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLoadBalancersOK creates a GetLoadBalancersOK with default headers values
func NewGetLoadBalancersOK() *GetLoadBalancersOK {
	return &GetLoadBalancersOK{}
}

/*GetLoadBalancersOK handles this case with default header values.

successful operation
*/
type GetLoadBalancersOK struct {
	Payload *models.LoadBalancerResult
}

func (o *GetLoadBalancersOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/load-balancers][%d] getLoadBalancersOK  %+v", 200, o.Payload)
}

func (o *GetLoadBalancersOK) GetPayload() *models.LoadBalancerResult {
	return o.Payload
}

func (o *GetLoadBalancersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoadBalancerResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLoadBalancersForbidden creates a GetLoadBalancersForbidden with default headers values
func NewGetLoadBalancersForbidden() *GetLoadBalancersForbidden {
	return &GetLoadBalancersForbidden{}
}

/*GetLoadBalancersForbidden handles this case with default header values.

Forbidden
*/
type GetLoadBalancersForbidden struct {
}

func (o *GetLoadBalancersForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/load-balancers][%d] getLoadBalancersForbidden ", 403)
}

func (o *GetLoadBalancersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
