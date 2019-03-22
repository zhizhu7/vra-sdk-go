// Code generated by go-swagger; DO NOT EDIT.

package load_balancer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteLoadBalancerReader is a Reader for the DeleteLoadBalancer structure.
type DeleteLoadBalancerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLoadBalancerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 403:
		result := NewDeleteLoadBalancerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteLoadBalancerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLoadBalancerForbidden creates a DeleteLoadBalancerForbidden with default headers values
func NewDeleteLoadBalancerForbidden() *DeleteLoadBalancerForbidden {
	return &DeleteLoadBalancerForbidden{}
}

/*DeleteLoadBalancerForbidden handles this case with default header values.

Forbidden
*/
type DeleteLoadBalancerForbidden struct {
}

func (o *DeleteLoadBalancerForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/load-balancers/{id}][%d] deleteLoadBalancerForbidden ", 403)
}

func (o *DeleteLoadBalancerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteLoadBalancerNotFound creates a DeleteLoadBalancerNotFound with default headers values
func NewDeleteLoadBalancerNotFound() *DeleteLoadBalancerNotFound {
	return &DeleteLoadBalancerNotFound{}
}

/*DeleteLoadBalancerNotFound handles this case with default header values.

Not Found
*/
type DeleteLoadBalancerNotFound struct {
}

func (o *DeleteLoadBalancerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/load-balancers/{id}][%d] deleteLoadBalancerNotFound ", 404)
}

func (o *DeleteLoadBalancerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
