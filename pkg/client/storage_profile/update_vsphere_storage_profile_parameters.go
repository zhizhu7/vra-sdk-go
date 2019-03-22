// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

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

// NewUpdateVSphereStorageProfileParams creates a new UpdateVSphereStorageProfileParams object
// with the default values initialized.
func NewUpdateVSphereStorageProfileParams() *UpdateVSphereStorageProfileParams {
	var ()
	return &UpdateVSphereStorageProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVSphereStorageProfileParamsWithTimeout creates a new UpdateVSphereStorageProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateVSphereStorageProfileParamsWithTimeout(timeout time.Duration) *UpdateVSphereStorageProfileParams {
	var ()
	return &UpdateVSphereStorageProfileParams{

		timeout: timeout,
	}
}

// NewUpdateVSphereStorageProfileParamsWithContext creates a new UpdateVSphereStorageProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateVSphereStorageProfileParamsWithContext(ctx context.Context) *UpdateVSphereStorageProfileParams {
	var ()
	return &UpdateVSphereStorageProfileParams{

		Context: ctx,
	}
}

// NewUpdateVSphereStorageProfileParamsWithHTTPClient creates a new UpdateVSphereStorageProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateVSphereStorageProfileParamsWithHTTPClient(client *http.Client) *UpdateVSphereStorageProfileParams {
	var ()
	return &UpdateVSphereStorageProfileParams{
		HTTPClient: client,
	}
}

/*UpdateVSphereStorageProfileParams contains all the parameters to send to the API endpoint
for the update v sphere storage profile operation typically these are written to a http.Request
*/
type UpdateVSphereStorageProfileParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion string
	/*Body
	  StorageProfileVsphereSpecification instance

	*/
	Body *models.StorageProfileVsphereSpecification
	/*ID
	  The ID of the storage profile.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update v sphere storage profile params
func (o *UpdateVSphereStorageProfileParams) WithTimeout(timeout time.Duration) *UpdateVSphereStorageProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update v sphere storage profile params
func (o *UpdateVSphereStorageProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update v sphere storage profile params
func (o *UpdateVSphereStorageProfileParams) WithContext(ctx context.Context) *UpdateVSphereStorageProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update v sphere storage profile params
func (o *UpdateVSphereStorageProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update v sphere storage profile params
func (o *UpdateVSphereStorageProfileParams) WithHTTPClient(client *http.Client) *UpdateVSphereStorageProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update v sphere storage profile params
func (o *UpdateVSphereStorageProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update v sphere storage profile params
func (o *UpdateVSphereStorageProfileParams) WithAPIVersion(aPIVersion string) *UpdateVSphereStorageProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update v sphere storage profile params
func (o *UpdateVSphereStorageProfileParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update v sphere storage profile params
func (o *UpdateVSphereStorageProfileParams) WithBody(body *models.StorageProfileVsphereSpecification) *UpdateVSphereStorageProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update v sphere storage profile params
func (o *UpdateVSphereStorageProfileParams) SetBody(body *models.StorageProfileVsphereSpecification) {
	o.Body = body
}

// WithID adds the id to the update v sphere storage profile params
func (o *UpdateVSphereStorageProfileParams) WithID(id string) *UpdateVSphereStorageProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update v sphere storage profile params
func (o *UpdateVSphereStorageProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVSphereStorageProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
