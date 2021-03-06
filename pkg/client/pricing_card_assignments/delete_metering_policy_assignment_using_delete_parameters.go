// Code generated by go-swagger; DO NOT EDIT.

package pricing_card_assignments

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

// NewDeleteMeteringPolicyAssignmentUsingDELETEParams creates a new DeleteMeteringPolicyAssignmentUsingDELETEParams object
// with the default values initialized.
func NewDeleteMeteringPolicyAssignmentUsingDELETEParams() *DeleteMeteringPolicyAssignmentUsingDELETEParams {
	var ()
	return &DeleteMeteringPolicyAssignmentUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteMeteringPolicyAssignmentUsingDELETEParamsWithTimeout creates a new DeleteMeteringPolicyAssignmentUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteMeteringPolicyAssignmentUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteMeteringPolicyAssignmentUsingDELETEParams {
	var ()
	return &DeleteMeteringPolicyAssignmentUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteMeteringPolicyAssignmentUsingDELETEParamsWithContext creates a new DeleteMeteringPolicyAssignmentUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteMeteringPolicyAssignmentUsingDELETEParamsWithContext(ctx context.Context) *DeleteMeteringPolicyAssignmentUsingDELETEParams {
	var ()
	return &DeleteMeteringPolicyAssignmentUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteMeteringPolicyAssignmentUsingDELETEParamsWithHTTPClient creates a new DeleteMeteringPolicyAssignmentUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteMeteringPolicyAssignmentUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteMeteringPolicyAssignmentUsingDELETEParams {
	var ()
	return &DeleteMeteringPolicyAssignmentUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteMeteringPolicyAssignmentUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete metering policy assignment using d e l e t e operation typically these are written to a http.Request
*/
type DeleteMeteringPolicyAssignmentUsingDELETEParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /catalog/api/about

	*/
	APIVersion *string
	/*ID
	  pricing card Assignment Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete metering policy assignment using d e l e t e params
func (o *DeleteMeteringPolicyAssignmentUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteMeteringPolicyAssignmentUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete metering policy assignment using d e l e t e params
func (o *DeleteMeteringPolicyAssignmentUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete metering policy assignment using d e l e t e params
func (o *DeleteMeteringPolicyAssignmentUsingDELETEParams) WithContext(ctx context.Context) *DeleteMeteringPolicyAssignmentUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete metering policy assignment using d e l e t e params
func (o *DeleteMeteringPolicyAssignmentUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete metering policy assignment using d e l e t e params
func (o *DeleteMeteringPolicyAssignmentUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteMeteringPolicyAssignmentUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete metering policy assignment using d e l e t e params
func (o *DeleteMeteringPolicyAssignmentUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete metering policy assignment using d e l e t e params
func (o *DeleteMeteringPolicyAssignmentUsingDELETEParams) WithAPIVersion(aPIVersion *string) *DeleteMeteringPolicyAssignmentUsingDELETEParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete metering policy assignment using d e l e t e params
func (o *DeleteMeteringPolicyAssignmentUsingDELETEParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete metering policy assignment using d e l e t e params
func (o *DeleteMeteringPolicyAssignmentUsingDELETEParams) WithID(id strfmt.UUID) *DeleteMeteringPolicyAssignmentUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete metering policy assignment using d e l e t e params
func (o *DeleteMeteringPolicyAssignmentUsingDELETEParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteMeteringPolicyAssignmentUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
