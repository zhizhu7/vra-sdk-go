// Code generated by go-swagger; DO NOT EDIT.

package fabric_flavors

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

// NewGetFabricFlavorsParams creates a new GetFabricFlavorsParams object
// with the default values initialized.
func NewGetFabricFlavorsParams() *GetFabricFlavorsParams {
	var ()
	return &GetFabricFlavorsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFabricFlavorsParamsWithTimeout creates a new GetFabricFlavorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFabricFlavorsParamsWithTimeout(timeout time.Duration) *GetFabricFlavorsParams {
	var ()
	return &GetFabricFlavorsParams{

		timeout: timeout,
	}
}

// NewGetFabricFlavorsParamsWithContext creates a new GetFabricFlavorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFabricFlavorsParamsWithContext(ctx context.Context) *GetFabricFlavorsParams {
	var ()
	return &GetFabricFlavorsParams{

		Context: ctx,
	}
}

// NewGetFabricFlavorsParamsWithHTTPClient creates a new GetFabricFlavorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFabricFlavorsParamsWithHTTPClient(client *http.Client) *GetFabricFlavorsParams {
	var ()
	return &GetFabricFlavorsParams{
		HTTPClient: client,
	}
}

/*GetFabricFlavorsParams contains all the parameters to send to the API endpoint
for the get fabric flavors operation typically these are written to a http.Request
*/
type GetFabricFlavorsParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fabric flavors params
func (o *GetFabricFlavorsParams) WithTimeout(timeout time.Duration) *GetFabricFlavorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fabric flavors params
func (o *GetFabricFlavorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fabric flavors params
func (o *GetFabricFlavorsParams) WithContext(ctx context.Context) *GetFabricFlavorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fabric flavors params
func (o *GetFabricFlavorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fabric flavors params
func (o *GetFabricFlavorsParams) WithHTTPClient(client *http.Client) *GetFabricFlavorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fabric flavors params
func (o *GetFabricFlavorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get fabric flavors params
func (o *GetFabricFlavorsParams) WithAPIVersion(aPIVersion string) *GetFabricFlavorsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get fabric flavors params
func (o *GetFabricFlavorsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetFabricFlavorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
