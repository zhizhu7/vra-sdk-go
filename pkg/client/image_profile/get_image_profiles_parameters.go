// Code generated by go-swagger; DO NOT EDIT.

package image_profile

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

// NewGetImageProfilesParams creates a new GetImageProfilesParams object
// with the default values initialized.
func NewGetImageProfilesParams() *GetImageProfilesParams {
	var ()
	return &GetImageProfilesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetImageProfilesParamsWithTimeout creates a new GetImageProfilesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetImageProfilesParamsWithTimeout(timeout time.Duration) *GetImageProfilesParams {
	var ()
	return &GetImageProfilesParams{

		timeout: timeout,
	}
}

// NewGetImageProfilesParamsWithContext creates a new GetImageProfilesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetImageProfilesParamsWithContext(ctx context.Context) *GetImageProfilesParams {
	var ()
	return &GetImageProfilesParams{

		Context: ctx,
	}
}

// NewGetImageProfilesParamsWithHTTPClient creates a new GetImageProfilesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetImageProfilesParamsWithHTTPClient(client *http.Client) *GetImageProfilesParams {
	var ()
	return &GetImageProfilesParams{
		HTTPClient: client,
	}
}

/*GetImageProfilesParams contains all the parameters to send to the API endpoint
for the get image profiles operation typically these are written to a http.Request
*/
type GetImageProfilesParams struct {

	/*DollarFilter
	  Add a filter to return limited results

	*/
	DollarFilter *string
	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get image profiles params
func (o *GetImageProfilesParams) WithTimeout(timeout time.Duration) *GetImageProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get image profiles params
func (o *GetImageProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get image profiles params
func (o *GetImageProfilesParams) WithContext(ctx context.Context) *GetImageProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get image profiles params
func (o *GetImageProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get image profiles params
func (o *GetImageProfilesParams) WithHTTPClient(client *http.Client) *GetImageProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get image profiles params
func (o *GetImageProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the get image profiles params
func (o *GetImageProfilesParams) WithDollarFilter(dollarFilter *string) *GetImageProfilesParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get image profiles params
func (o *GetImageProfilesParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithAPIVersion adds the aPIVersion to the get image profiles params
func (o *GetImageProfilesParams) WithAPIVersion(aPIVersion *string) *GetImageProfilesParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get image profiles params
func (o *GetImageProfilesParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetImageProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
