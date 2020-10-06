// Code generated by go-swagger; DO NOT EDIT.

package compute_gateway

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

// NewGetGatewayParams creates a new GetGatewayParams object
// with the default values initialized.
func NewGetGatewayParams() *GetGatewayParams {
	var ()
	return &GetGatewayParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGatewayParamsWithTimeout creates a new GetGatewayParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGatewayParamsWithTimeout(timeout time.Duration) *GetGatewayParams {
	var ()
	return &GetGatewayParams{

		timeout: timeout,
	}
}

// NewGetGatewayParamsWithContext creates a new GetGatewayParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGatewayParamsWithContext(ctx context.Context) *GetGatewayParams {
	var ()
	return &GetGatewayParams{

		Context: ctx,
	}
}

// NewGetGatewayParamsWithHTTPClient creates a new GetGatewayParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGatewayParamsWithHTTPClient(client *http.Client) *GetGatewayParams {
	var ()
	return &GetGatewayParams{
		HTTPClient: client,
	}
}

/*GetGatewayParams contains all the parameters to send to the API endpoint
for the get gateway operation typically these are written to a http.Request
*/
type GetGatewayParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the gateway.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get gateway params
func (o *GetGatewayParams) WithTimeout(timeout time.Duration) *GetGatewayParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gateway params
func (o *GetGatewayParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gateway params
func (o *GetGatewayParams) WithContext(ctx context.Context) *GetGatewayParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gateway params
func (o *GetGatewayParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gateway params
func (o *GetGatewayParams) WithHTTPClient(client *http.Client) *GetGatewayParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gateway params
func (o *GetGatewayParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get gateway params
func (o *GetGatewayParams) WithAPIVersion(aPIVersion *string) *GetGatewayParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get gateway params
func (o *GetGatewayParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get gateway params
func (o *GetGatewayParams) WithID(id string) *GetGatewayParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get gateway params
func (o *GetGatewayParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetGatewayParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// query param apiVersion
		var qrAPIVersion string
		if o.APIVersion != nil {
			qrAPIVersion = *o.APIVersion
		}
		qAPIVersion := qrAPIVersion
		if qAPIVersion != "" {
			if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
				return err
			}
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