// Code generated by go-swagger; DO NOT EDIT.

package load_balancer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewLoadBalancerDeleteOperationParams creates a new LoadBalancerDeleteOperationParams object
// with the default values initialized.
func NewLoadBalancerDeleteOperationParams() *LoadBalancerDeleteOperationParams {
	var ()
	return &LoadBalancerDeleteOperationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoadBalancerDeleteOperationParamsWithTimeout creates a new LoadBalancerDeleteOperationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoadBalancerDeleteOperationParamsWithTimeout(timeout time.Duration) *LoadBalancerDeleteOperationParams {
	var ()
	return &LoadBalancerDeleteOperationParams{

		timeout: timeout,
	}
}

// NewLoadBalancerDeleteOperationParamsWithContext creates a new LoadBalancerDeleteOperationParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoadBalancerDeleteOperationParamsWithContext(ctx context.Context) *LoadBalancerDeleteOperationParams {
	var ()
	return &LoadBalancerDeleteOperationParams{

		Context: ctx,
	}
}

// NewLoadBalancerDeleteOperationParamsWithHTTPClient creates a new LoadBalancerDeleteOperationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoadBalancerDeleteOperationParamsWithHTTPClient(client *http.Client) *LoadBalancerDeleteOperationParams {
	var ()
	return &LoadBalancerDeleteOperationParams{
		HTTPClient: client,
	}
}

/*LoadBalancerDeleteOperationParams contains all the parameters to send to the API endpoint
for the load balancer delete operation operation typically these are written to a http.Request
*/
type LoadBalancerDeleteOperationParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion string
	/*ID
	  The ID of the load balancer.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the load balancer delete operation params
func (o *LoadBalancerDeleteOperationParams) WithTimeout(timeout time.Duration) *LoadBalancerDeleteOperationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the load balancer delete operation params
func (o *LoadBalancerDeleteOperationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the load balancer delete operation params
func (o *LoadBalancerDeleteOperationParams) WithContext(ctx context.Context) *LoadBalancerDeleteOperationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the load balancer delete operation params
func (o *LoadBalancerDeleteOperationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the load balancer delete operation params
func (o *LoadBalancerDeleteOperationParams) WithHTTPClient(client *http.Client) *LoadBalancerDeleteOperationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the load balancer delete operation params
func (o *LoadBalancerDeleteOperationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the load balancer delete operation params
func (o *LoadBalancerDeleteOperationParams) WithAPIVersion(aPIVersion string) *LoadBalancerDeleteOperationParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the load balancer delete operation params
func (o *LoadBalancerDeleteOperationParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the load balancer delete operation params
func (o *LoadBalancerDeleteOperationParams) WithID(id string) *LoadBalancerDeleteOperationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the load balancer delete operation params
func (o *LoadBalancerDeleteOperationParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LoadBalancerDeleteOperationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
