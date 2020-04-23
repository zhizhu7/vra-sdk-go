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

// NewEnumeratePrivateImagesGcpParams creates a new EnumeratePrivateImagesGcpParams object
// with the default values initialized.
func NewEnumeratePrivateImagesGcpParams() *EnumeratePrivateImagesGcpParams {
	var ()
	return &EnumeratePrivateImagesGcpParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnumeratePrivateImagesGcpParamsWithTimeout creates a new EnumeratePrivateImagesGcpParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnumeratePrivateImagesGcpParamsWithTimeout(timeout time.Duration) *EnumeratePrivateImagesGcpParams {
	var ()
	return &EnumeratePrivateImagesGcpParams{

		timeout: timeout,
	}
}

// NewEnumeratePrivateImagesGcpParamsWithContext creates a new EnumeratePrivateImagesGcpParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnumeratePrivateImagesGcpParamsWithContext(ctx context.Context) *EnumeratePrivateImagesGcpParams {
	var ()
	return &EnumeratePrivateImagesGcpParams{

		Context: ctx,
	}
}

// NewEnumeratePrivateImagesGcpParamsWithHTTPClient creates a new EnumeratePrivateImagesGcpParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnumeratePrivateImagesGcpParamsWithHTTPClient(client *http.Client) *EnumeratePrivateImagesGcpParams {
	var ()
	return &EnumeratePrivateImagesGcpParams{
		HTTPClient: client,
	}
}

/*EnumeratePrivateImagesGcpParams contains all the parameters to send to the API endpoint
for the enumerate private images gcp operation typically these are written to a http.Request
*/
type EnumeratePrivateImagesGcpParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  Id of GCP cloud account to enumerate

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enumerate private images gcp params
func (o *EnumeratePrivateImagesGcpParams) WithTimeout(timeout time.Duration) *EnumeratePrivateImagesGcpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enumerate private images gcp params
func (o *EnumeratePrivateImagesGcpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enumerate private images gcp params
func (o *EnumeratePrivateImagesGcpParams) WithContext(ctx context.Context) *EnumeratePrivateImagesGcpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enumerate private images gcp params
func (o *EnumeratePrivateImagesGcpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enumerate private images gcp params
func (o *EnumeratePrivateImagesGcpParams) WithHTTPClient(client *http.Client) *EnumeratePrivateImagesGcpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enumerate private images gcp params
func (o *EnumeratePrivateImagesGcpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the enumerate private images gcp params
func (o *EnumeratePrivateImagesGcpParams) WithAPIVersion(aPIVersion *string) *EnumeratePrivateImagesGcpParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the enumerate private images gcp params
func (o *EnumeratePrivateImagesGcpParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the enumerate private images gcp params
func (o *EnumeratePrivateImagesGcpParams) WithID(id string) *EnumeratePrivateImagesGcpParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the enumerate private images gcp params
func (o *EnumeratePrivateImagesGcpParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EnumeratePrivateImagesGcpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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