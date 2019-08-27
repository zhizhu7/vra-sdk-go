// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewGetNetworkDomainsParams creates a new GetNetworkDomainsParams object
// with the default values initialized.
func NewGetNetworkDomainsParams() *GetNetworkDomainsParams {
	var ()
	return &GetNetworkDomainsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkDomainsParamsWithTimeout creates a new GetNetworkDomainsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkDomainsParamsWithTimeout(timeout time.Duration) *GetNetworkDomainsParams {
	var ()
	return &GetNetworkDomainsParams{

		timeout: timeout,
	}
}

// NewGetNetworkDomainsParamsWithContext creates a new GetNetworkDomainsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkDomainsParamsWithContext(ctx context.Context) *GetNetworkDomainsParams {
	var ()
	return &GetNetworkDomainsParams{

		Context: ctx,
	}
}

// NewGetNetworkDomainsParamsWithHTTPClient creates a new GetNetworkDomainsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkDomainsParamsWithHTTPClient(client *http.Client) *GetNetworkDomainsParams {
	var ()
	return &GetNetworkDomainsParams{
		HTTPClient: client,
	}
}

/*GetNetworkDomainsParams contains all the parameters to send to the API endpoint
for the get network domains operation typically these are written to a http.Request
*/
type GetNetworkDomainsParams struct {

	/*DollarFilter
	  Add a filter to return limited results

	*/
	DollarFilter *string
	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network domains params
func (o *GetNetworkDomainsParams) WithTimeout(timeout time.Duration) *GetNetworkDomainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network domains params
func (o *GetNetworkDomainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network domains params
func (o *GetNetworkDomainsParams) WithContext(ctx context.Context) *GetNetworkDomainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network domains params
func (o *GetNetworkDomainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network domains params
func (o *GetNetworkDomainsParams) WithHTTPClient(client *http.Client) *GetNetworkDomainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network domains params
func (o *GetNetworkDomainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the get network domains params
func (o *GetNetworkDomainsParams) WithDollarFilter(dollarFilter *string) *GetNetworkDomainsParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get network domains params
func (o *GetNetworkDomainsParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithAPIVersion adds the aPIVersion to the get network domains params
func (o *GetNetworkDomainsParams) WithAPIVersion(aPIVersion *string) *GetNetworkDomainsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get network domains params
func (o *GetNetworkDomainsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkDomainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarFilter != nil {

		// query param $filter
		var qrDollarFilter string
		if o.DollarFilter != nil {
			qrDollarFilter = *o.DollarFilter
		}
		qDollarFilter := qrDollarFilter
		if qDollarFilter != "" {
			if err := r.SetQueryParam("$filter", qDollarFilter); err != nil {
				return err
			}
		}

	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
