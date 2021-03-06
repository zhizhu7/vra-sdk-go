// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewDeletePolicyUsingDELETE1Params creates a new DeletePolicyUsingDELETE1Params object
// with the default values initialized.
func NewDeletePolicyUsingDELETE1Params() *DeletePolicyUsingDELETE1Params {
	var ()
	return &DeletePolicyUsingDELETE1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePolicyUsingDELETE1ParamsWithTimeout creates a new DeletePolicyUsingDELETE1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePolicyUsingDELETE1ParamsWithTimeout(timeout time.Duration) *DeletePolicyUsingDELETE1Params {
	var ()
	return &DeletePolicyUsingDELETE1Params{

		timeout: timeout,
	}
}

// NewDeletePolicyUsingDELETE1ParamsWithContext creates a new DeletePolicyUsingDELETE1Params object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePolicyUsingDELETE1ParamsWithContext(ctx context.Context) *DeletePolicyUsingDELETE1Params {
	var ()
	return &DeletePolicyUsingDELETE1Params{

		Context: ctx,
	}
}

// NewDeletePolicyUsingDELETE1ParamsWithHTTPClient creates a new DeletePolicyUsingDELETE1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePolicyUsingDELETE1ParamsWithHTTPClient(client *http.Client) *DeletePolicyUsingDELETE1Params {
	var ()
	return &DeletePolicyUsingDELETE1Params{
		HTTPClient: client,
	}
}

/*DeletePolicyUsingDELETE1Params contains all the parameters to send to the API endpoint
for the delete policy using d e l e t e 1 operation typically these are written to a http.Request
*/
type DeletePolicyUsingDELETE1Params struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /catalog/api/about

	*/
	APIVersion *string
	/*ID
	  Policy ID

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete policy using d e l e t e 1 params
func (o *DeletePolicyUsingDELETE1Params) WithTimeout(timeout time.Duration) *DeletePolicyUsingDELETE1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete policy using d e l e t e 1 params
func (o *DeletePolicyUsingDELETE1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete policy using d e l e t e 1 params
func (o *DeletePolicyUsingDELETE1Params) WithContext(ctx context.Context) *DeletePolicyUsingDELETE1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete policy using d e l e t e 1 params
func (o *DeletePolicyUsingDELETE1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete policy using d e l e t e 1 params
func (o *DeletePolicyUsingDELETE1Params) WithHTTPClient(client *http.Client) *DeletePolicyUsingDELETE1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete policy using d e l e t e 1 params
func (o *DeletePolicyUsingDELETE1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete policy using d e l e t e 1 params
func (o *DeletePolicyUsingDELETE1Params) WithAPIVersion(aPIVersion *string) *DeletePolicyUsingDELETE1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete policy using d e l e t e 1 params
func (o *DeletePolicyUsingDELETE1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete policy using d e l e t e 1 params
func (o *DeletePolicyUsingDELETE1Params) WithID(id strfmt.UUID) *DeletePolicyUsingDELETE1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete policy using d e l e t e 1 params
func (o *DeletePolicyUsingDELETE1Params) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePolicyUsingDELETE1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
