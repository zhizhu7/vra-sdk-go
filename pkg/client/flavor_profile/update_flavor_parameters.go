// Code generated by go-swagger; DO NOT EDIT.

package flavor_profile

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

// NewUpdateFlavorParams creates a new UpdateFlavorParams object
// with the default values initialized.
func NewUpdateFlavorParams() *UpdateFlavorParams {
	var ()
	return &UpdateFlavorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateFlavorParamsWithTimeout creates a new UpdateFlavorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateFlavorParamsWithTimeout(timeout time.Duration) *UpdateFlavorParams {
	var ()
	return &UpdateFlavorParams{

		timeout: timeout,
	}
}

// NewUpdateFlavorParamsWithContext creates a new UpdateFlavorParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateFlavorParamsWithContext(ctx context.Context) *UpdateFlavorParams {
	var ()
	return &UpdateFlavorParams{

		Context: ctx,
	}
}

// NewUpdateFlavorParamsWithHTTPClient creates a new UpdateFlavorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateFlavorParamsWithHTTPClient(client *http.Client) *UpdateFlavorParams {
	var ()
	return &UpdateFlavorParams{
		HTTPClient: client,
	}
}

/*UpdateFlavorParams contains all the parameters to send to the API endpoint
for the update flavor operation typically these are written to a http.Request
*/
type UpdateFlavorParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion string
	/*Body
	  FlavorProfile instance

	*/
	Body *models.FlavorProfileSpecification
	/*ID
	  The ID of the flavor.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update flavor params
func (o *UpdateFlavorParams) WithTimeout(timeout time.Duration) *UpdateFlavorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update flavor params
func (o *UpdateFlavorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update flavor params
func (o *UpdateFlavorParams) WithContext(ctx context.Context) *UpdateFlavorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update flavor params
func (o *UpdateFlavorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update flavor params
func (o *UpdateFlavorParams) WithHTTPClient(client *http.Client) *UpdateFlavorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update flavor params
func (o *UpdateFlavorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update flavor params
func (o *UpdateFlavorParams) WithAPIVersion(aPIVersion string) *UpdateFlavorParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update flavor params
func (o *UpdateFlavorParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update flavor params
func (o *UpdateFlavorParams) WithBody(body *models.FlavorProfileSpecification) *UpdateFlavorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update flavor params
func (o *UpdateFlavorParams) SetBody(body *models.FlavorProfileSpecification) {
	o.Body = body
}

// WithID adds the id to the update flavor params
func (o *UpdateFlavorParams) WithID(id string) *UpdateFlavorParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update flavor params
func (o *UpdateFlavorParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateFlavorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
