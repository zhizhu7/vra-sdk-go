// Code generated by go-swagger; DO NOT EDIT.

package disk

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

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// NewCreateBlockDeviceParams creates a new CreateBlockDeviceParams object
// with the default values initialized.
func NewCreateBlockDeviceParams() *CreateBlockDeviceParams {
	var ()
	return &CreateBlockDeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBlockDeviceParamsWithTimeout creates a new CreateBlockDeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateBlockDeviceParamsWithTimeout(timeout time.Duration) *CreateBlockDeviceParams {
	var ()
	return &CreateBlockDeviceParams{

		timeout: timeout,
	}
}

// NewCreateBlockDeviceParamsWithContext creates a new CreateBlockDeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateBlockDeviceParamsWithContext(ctx context.Context) *CreateBlockDeviceParams {
	var ()
	return &CreateBlockDeviceParams{

		Context: ctx,
	}
}

// NewCreateBlockDeviceParamsWithHTTPClient creates a new CreateBlockDeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateBlockDeviceParamsWithHTTPClient(client *http.Client) *CreateBlockDeviceParams {
	var ()
	return &CreateBlockDeviceParams{
		HTTPClient: client,
	}
}

/*CreateBlockDeviceParams contains all the parameters to send to the API endpoint
for the create block device operation typically these are written to a http.Request
*/
type CreateBlockDeviceParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion string
	/*Body
	  Disk Specification instance

	*/
	Body *models.BlockDeviceSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create block device params
func (o *CreateBlockDeviceParams) WithTimeout(timeout time.Duration) *CreateBlockDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create block device params
func (o *CreateBlockDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create block device params
func (o *CreateBlockDeviceParams) WithContext(ctx context.Context) *CreateBlockDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create block device params
func (o *CreateBlockDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create block device params
func (o *CreateBlockDeviceParams) WithHTTPClient(client *http.Client) *CreateBlockDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create block device params
func (o *CreateBlockDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create block device params
func (o *CreateBlockDeviceParams) WithAPIVersion(aPIVersion string) *CreateBlockDeviceParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create block device params
func (o *CreateBlockDeviceParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create block device params
func (o *CreateBlockDeviceParams) WithBody(body *models.BlockDeviceSpecification) *CreateBlockDeviceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create block device params
func (o *CreateBlockDeviceParams) SetBody(body *models.BlockDeviceSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBlockDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
