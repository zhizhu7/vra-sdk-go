// Code generated by go-swagger; DO NOT EDIT.

package deployments

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

// NewGetResourceByIDUsingGETParams creates a new GetResourceByIDUsingGETParams object
// with the default values initialized.
func NewGetResourceByIDUsingGETParams() *GetResourceByIDUsingGETParams {
	var ()
	return &GetResourceByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceByIDUsingGETParamsWithTimeout creates a new GetResourceByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetResourceByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetResourceByIDUsingGETParams {
	var ()
	return &GetResourceByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetResourceByIDUsingGETParamsWithContext creates a new GetResourceByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetResourceByIDUsingGETParamsWithContext(ctx context.Context) *GetResourceByIDUsingGETParams {
	var ()
	return &GetResourceByIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetResourceByIDUsingGETParamsWithHTTPClient creates a new GetResourceByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetResourceByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetResourceByIDUsingGETParams {
	var ()
	return &GetResourceByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetResourceByIDUsingGETParams contains all the parameters to send to the API endpoint
for the get resource by Id using g e t operation typically these are written to a http.Request
*/
type GetResourceByIDUsingGETParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /catalog/api/about

	*/
	APIVersion *string
	/*DepID
	  Deployment ID

	*/
	DepID strfmt.UUID
	/*ResourceID
	  Resource ID

	*/
	ResourceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get resource by Id using get params
func (o *GetResourceByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetResourceByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource by Id using get params
func (o *GetResourceByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource by Id using get params
func (o *GetResourceByIDUsingGETParams) WithContext(ctx context.Context) *GetResourceByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource by Id using get params
func (o *GetResourceByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource by Id using get params
func (o *GetResourceByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetResourceByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource by Id using get params
func (o *GetResourceByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get resource by Id using get params
func (o *GetResourceByIDUsingGETParams) WithAPIVersion(aPIVersion *string) *GetResourceByIDUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get resource by Id using get params
func (o *GetResourceByIDUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDepID adds the depID to the get resource by Id using get params
func (o *GetResourceByIDUsingGETParams) WithDepID(depID strfmt.UUID) *GetResourceByIDUsingGETParams {
	o.SetDepID(depID)
	return o
}

// SetDepID adds the depId to the get resource by Id using get params
func (o *GetResourceByIDUsingGETParams) SetDepID(depID strfmt.UUID) {
	o.DepID = depID
}

// WithResourceID adds the resourceID to the get resource by Id using get params
func (o *GetResourceByIDUsingGETParams) WithResourceID(resourceID strfmt.UUID) *GetResourceByIDUsingGETParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get resource by Id using get params
func (o *GetResourceByIDUsingGETParams) SetResourceID(resourceID strfmt.UUID) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param depId
	if err := r.SetPathParam("depId", o.DepID.String()); err != nil {
		return err
	}

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
