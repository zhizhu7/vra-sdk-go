// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewAddMachineDiskParams creates a new AddMachineDiskParams object
// with the default values initialized.
func NewAddMachineDiskParams() *AddMachineDiskParams {
	var ()
	return &AddMachineDiskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddMachineDiskParamsWithTimeout creates a new AddMachineDiskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddMachineDiskParamsWithTimeout(timeout time.Duration) *AddMachineDiskParams {
	var ()
	return &AddMachineDiskParams{

		timeout: timeout,
	}
}

// NewAddMachineDiskParamsWithContext creates a new AddMachineDiskParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddMachineDiskParamsWithContext(ctx context.Context) *AddMachineDiskParams {
	var ()
	return &AddMachineDiskParams{

		Context: ctx,
	}
}

// NewAddMachineDiskParamsWithHTTPClient creates a new AddMachineDiskParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddMachineDiskParamsWithHTTPClient(client *http.Client) *AddMachineDiskParams {
	var ()
	return &AddMachineDiskParams{
		HTTPClient: client,
	}
}

/*AddMachineDiskParams contains all the parameters to send to the API endpoint
for the add machine disk operation typically these are written to a http.Request
*/
type AddMachineDiskParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion string
	/*Body
	  Disk Specification instance

	*/
	Body *models.DiskAttachmentSpecification
	/*ID
	  The ID of the machine.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add machine disk params
func (o *AddMachineDiskParams) WithTimeout(timeout time.Duration) *AddMachineDiskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add machine disk params
func (o *AddMachineDiskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add machine disk params
func (o *AddMachineDiskParams) WithContext(ctx context.Context) *AddMachineDiskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add machine disk params
func (o *AddMachineDiskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add machine disk params
func (o *AddMachineDiskParams) WithHTTPClient(client *http.Client) *AddMachineDiskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add machine disk params
func (o *AddMachineDiskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the add machine disk params
func (o *AddMachineDiskParams) WithAPIVersion(aPIVersion string) *AddMachineDiskParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the add machine disk params
func (o *AddMachineDiskParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the add machine disk params
func (o *AddMachineDiskParams) WithBody(body *models.DiskAttachmentSpecification) *AddMachineDiskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add machine disk params
func (o *AddMachineDiskParams) SetBody(body *models.DiskAttachmentSpecification) {
	o.Body = body
}

// WithID adds the id to the add machine disk params
func (o *AddMachineDiskParams) WithID(id string) *AddMachineDiskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the add machine disk params
func (o *AddMachineDiskParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AddMachineDiskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
