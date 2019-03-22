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

// NewCreateFlavorParams creates a new CreateFlavorParams object
// with the default values initialized.
func NewCreateFlavorParams() *CreateFlavorParams {
	var ()
	return &CreateFlavorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFlavorParamsWithTimeout creates a new CreateFlavorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateFlavorParamsWithTimeout(timeout time.Duration) *CreateFlavorParams {
	var ()
	return &CreateFlavorParams{

		timeout: timeout,
	}
}

// NewCreateFlavorParamsWithContext creates a new CreateFlavorParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateFlavorParamsWithContext(ctx context.Context) *CreateFlavorParams {
	var ()
	return &CreateFlavorParams{

		Context: ctx,
	}
}

// NewCreateFlavorParamsWithHTTPClient creates a new CreateFlavorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateFlavorParamsWithHTTPClient(client *http.Client) *CreateFlavorParams {
	var ()
	return &CreateFlavorParams{
		HTTPClient: client,
	}
}

/*CreateFlavorParams contains all the parameters to send to the API endpoint
for the create flavor operation typically these are written to a http.Request
*/
type CreateFlavorParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion string
	/*Body
	  FlavorProfile instance

	*/
	Body *models.FlavorProfileSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create flavor params
func (o *CreateFlavorParams) WithTimeout(timeout time.Duration) *CreateFlavorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create flavor params
func (o *CreateFlavorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create flavor params
func (o *CreateFlavorParams) WithContext(ctx context.Context) *CreateFlavorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create flavor params
func (o *CreateFlavorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create flavor params
func (o *CreateFlavorParams) WithHTTPClient(client *http.Client) *CreateFlavorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create flavor params
func (o *CreateFlavorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create flavor params
func (o *CreateFlavorParams) WithAPIVersion(aPIVersion string) *CreateFlavorParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create flavor params
func (o *CreateFlavorParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create flavor params
func (o *CreateFlavorParams) WithBody(body *models.FlavorProfileSpecification) *CreateFlavorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create flavor params
func (o *CreateFlavorParams) SetBody(body *models.FlavorProfileSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFlavorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
