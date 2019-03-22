// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

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

// NewGetNsxVCloudAccountParams creates a new GetNsxVCloudAccountParams object
// with the default values initialized.
func NewGetNsxVCloudAccountParams() *GetNsxVCloudAccountParams {
	var ()
	return &GetNsxVCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNsxVCloudAccountParamsWithTimeout creates a new GetNsxVCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNsxVCloudAccountParamsWithTimeout(timeout time.Duration) *GetNsxVCloudAccountParams {
	var ()
	return &GetNsxVCloudAccountParams{

		timeout: timeout,
	}
}

// NewGetNsxVCloudAccountParamsWithContext creates a new GetNsxVCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNsxVCloudAccountParamsWithContext(ctx context.Context) *GetNsxVCloudAccountParams {
	var ()
	return &GetNsxVCloudAccountParams{

		Context: ctx,
	}
}

// NewGetNsxVCloudAccountParamsWithHTTPClient creates a new GetNsxVCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNsxVCloudAccountParamsWithHTTPClient(client *http.Client) *GetNsxVCloudAccountParams {
	var ()
	return &GetNsxVCloudAccountParams{
		HTTPClient: client,
	}
}

/*GetNsxVCloudAccountParams contains all the parameters to send to the API endpoint
for the get nsx v cloud account operation typically these are written to a http.Request
*/
type GetNsxVCloudAccountParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion string
	/*ID
	  The ID of the Cloud Account

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nsx v cloud account params
func (o *GetNsxVCloudAccountParams) WithTimeout(timeout time.Duration) *GetNsxVCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nsx v cloud account params
func (o *GetNsxVCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nsx v cloud account params
func (o *GetNsxVCloudAccountParams) WithContext(ctx context.Context) *GetNsxVCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nsx v cloud account params
func (o *GetNsxVCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nsx v cloud account params
func (o *GetNsxVCloudAccountParams) WithHTTPClient(client *http.Client) *GetNsxVCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nsx v cloud account params
func (o *GetNsxVCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get nsx v cloud account params
func (o *GetNsxVCloudAccountParams) WithAPIVersion(aPIVersion string) *GetNsxVCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get nsx v cloud account params
func (o *GetNsxVCloudAccountParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get nsx v cloud account params
func (o *GetNsxVCloudAccountParams) WithID(id string) *GetNsxVCloudAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get nsx v cloud account params
func (o *GetNsxVCloudAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetNsxVCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
