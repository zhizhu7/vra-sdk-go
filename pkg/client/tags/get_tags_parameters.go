// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// NewGetTagsParams creates a new GetTagsParams object
// with the default values initialized.
func NewGetTagsParams() *GetTagsParams {
	var ()
	return &GetTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTagsParamsWithTimeout creates a new GetTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTagsParamsWithTimeout(timeout time.Duration) *GetTagsParams {
	var ()
	return &GetTagsParams{

		timeout: timeout,
	}
}

// NewGetTagsParamsWithContext creates a new GetTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTagsParamsWithContext(ctx context.Context) *GetTagsParams {
	var ()
	return &GetTagsParams{

		Context: ctx,
	}
}

// NewGetTagsParamsWithHTTPClient creates a new GetTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTagsParamsWithHTTPClient(client *http.Client) *GetTagsParams {
	var ()
	return &GetTagsParams{
		HTTPClient: client,
	}
}

/*GetTagsParams contains all the parameters to send to the API endpoint
for the get tags operation typically these are written to a http.Request
*/
type GetTagsParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tags params
func (o *GetTagsParams) WithTimeout(timeout time.Duration) *GetTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tags params
func (o *GetTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tags params
func (o *GetTagsParams) WithContext(ctx context.Context) *GetTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tags params
func (o *GetTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tags params
func (o *GetTagsParams) WithHTTPClient(client *http.Client) *GetTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tags params
func (o *GetTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get tags params
func (o *GetTagsParams) WithAPIVersion(aPIVersion string) *GetTagsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get tags params
func (o *GetTagsParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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