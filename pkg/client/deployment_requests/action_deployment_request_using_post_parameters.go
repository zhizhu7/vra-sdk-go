// Code generated by go-swagger; DO NOT EDIT.

package deployment_requests

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

// NewActionDeploymentRequestUsingPOSTParams creates a new ActionDeploymentRequestUsingPOSTParams object
// with the default values initialized.
func NewActionDeploymentRequestUsingPOSTParams() *ActionDeploymentRequestUsingPOSTParams {
	var ()
	return &ActionDeploymentRequestUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewActionDeploymentRequestUsingPOSTParamsWithTimeout creates a new ActionDeploymentRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewActionDeploymentRequestUsingPOSTParamsWithTimeout(timeout time.Duration) *ActionDeploymentRequestUsingPOSTParams {
	var ()
	return &ActionDeploymentRequestUsingPOSTParams{

		timeout: timeout,
	}
}

// NewActionDeploymentRequestUsingPOSTParamsWithContext creates a new ActionDeploymentRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewActionDeploymentRequestUsingPOSTParamsWithContext(ctx context.Context) *ActionDeploymentRequestUsingPOSTParams {
	var ()
	return &ActionDeploymentRequestUsingPOSTParams{

		Context: ctx,
	}
}

// NewActionDeploymentRequestUsingPOSTParamsWithHTTPClient creates a new ActionDeploymentRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewActionDeploymentRequestUsingPOSTParamsWithHTTPClient(client *http.Client) *ActionDeploymentRequestUsingPOSTParams {
	var ()
	return &ActionDeploymentRequestUsingPOSTParams{
		HTTPClient: client,
	}
}

/*ActionDeploymentRequestUsingPOSTParams contains all the parameters to send to the API endpoint
for the action deployment request using p o s t operation typically these are written to a http.Request
*/
type ActionDeploymentRequestUsingPOSTParams struct {

	/*Action
	  action

	*/
	Action string
	/*RequestID
	  requestId

	*/
	RequestID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the action deployment request using p o s t params
func (o *ActionDeploymentRequestUsingPOSTParams) WithTimeout(timeout time.Duration) *ActionDeploymentRequestUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the action deployment request using p o s t params
func (o *ActionDeploymentRequestUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the action deployment request using p o s t params
func (o *ActionDeploymentRequestUsingPOSTParams) WithContext(ctx context.Context) *ActionDeploymentRequestUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the action deployment request using p o s t params
func (o *ActionDeploymentRequestUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the action deployment request using p o s t params
func (o *ActionDeploymentRequestUsingPOSTParams) WithHTTPClient(client *http.Client) *ActionDeploymentRequestUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the action deployment request using p o s t params
func (o *ActionDeploymentRequestUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the action deployment request using p o s t params
func (o *ActionDeploymentRequestUsingPOSTParams) WithAction(action string) *ActionDeploymentRequestUsingPOSTParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the action deployment request using p o s t params
func (o *ActionDeploymentRequestUsingPOSTParams) SetAction(action string) {
	o.Action = action
}

// WithRequestID adds the requestID to the action deployment request using p o s t params
func (o *ActionDeploymentRequestUsingPOSTParams) WithRequestID(requestID strfmt.UUID) *ActionDeploymentRequestUsingPOSTParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the action deployment request using p o s t params
func (o *ActionDeploymentRequestUsingPOSTParams) SetRequestID(requestID strfmt.UUID) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *ActionDeploymentRequestUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param action
	qrAction := o.Action
	qAction := qrAction
	if qAction != "" {
		if err := r.SetQueryParam("action", qAction); err != nil {
			return err
		}
	}

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
