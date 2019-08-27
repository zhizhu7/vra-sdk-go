// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewGetProjectsParams creates a new GetProjectsParams object
// with the default values initialized.
func NewGetProjectsParams() *GetProjectsParams {
	var ()
	return &GetProjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectsParamsWithTimeout creates a new GetProjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProjectsParamsWithTimeout(timeout time.Duration) *GetProjectsParams {
	var ()
	return &GetProjectsParams{

		timeout: timeout,
	}
}

// NewGetProjectsParamsWithContext creates a new GetProjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProjectsParamsWithContext(ctx context.Context) *GetProjectsParams {
	var ()
	return &GetProjectsParams{

		Context: ctx,
	}
}

// NewGetProjectsParamsWithHTTPClient creates a new GetProjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProjectsParamsWithHTTPClient(client *http.Client) *GetProjectsParams {
	var ()
	return &GetProjectsParams{
		HTTPClient: client,
	}
}

/*GetProjectsParams contains all the parameters to send to the API endpoint
for the get projects operation typically these are written to a http.Request
*/
type GetProjectsParams struct {

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

// WithTimeout adds the timeout to the get projects params
func (o *GetProjectsParams) WithTimeout(timeout time.Duration) *GetProjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get projects params
func (o *GetProjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get projects params
func (o *GetProjectsParams) WithContext(ctx context.Context) *GetProjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get projects params
func (o *GetProjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get projects params
func (o *GetProjectsParams) WithHTTPClient(client *http.Client) *GetProjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get projects params
func (o *GetProjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the get projects params
func (o *GetProjectsParams) WithDollarFilter(dollarFilter *string) *GetProjectsParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get projects params
func (o *GetProjectsParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithAPIVersion adds the aPIVersion to the get projects params
func (o *GetProjectsParams) WithAPIVersion(aPIVersion *string) *GetProjectsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get projects params
func (o *GetProjectsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
