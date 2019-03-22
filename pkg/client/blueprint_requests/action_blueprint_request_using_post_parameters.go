// Code generated by go-swagger; DO NOT EDIT.

package blueprint_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewActionBlueprintRequestUsingPOSTParams creates a new ActionBlueprintRequestUsingPOSTParams object
// with the default values initialized.
func NewActionBlueprintRequestUsingPOSTParams() *ActionBlueprintRequestUsingPOSTParams {
	var (
		forceDefault = bool(true)
	)
	return &ActionBlueprintRequestUsingPOSTParams{
		Force: &forceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewActionBlueprintRequestUsingPOSTParamsWithTimeout creates a new ActionBlueprintRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewActionBlueprintRequestUsingPOSTParamsWithTimeout(timeout time.Duration) *ActionBlueprintRequestUsingPOSTParams {
	var (
		forceDefault = bool(true)
	)
	return &ActionBlueprintRequestUsingPOSTParams{
		Force: &forceDefault,

		timeout: timeout,
	}
}

// NewActionBlueprintRequestUsingPOSTParamsWithContext creates a new ActionBlueprintRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewActionBlueprintRequestUsingPOSTParamsWithContext(ctx context.Context) *ActionBlueprintRequestUsingPOSTParams {
	var (
		forceDefault = bool(true)
	)
	return &ActionBlueprintRequestUsingPOSTParams{
		Force: &forceDefault,

		Context: ctx,
	}
}

// NewActionBlueprintRequestUsingPOSTParamsWithHTTPClient creates a new ActionBlueprintRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewActionBlueprintRequestUsingPOSTParamsWithHTTPClient(client *http.Client) *ActionBlueprintRequestUsingPOSTParams {
	var (
		forceDefault = bool(true)
	)
	return &ActionBlueprintRequestUsingPOSTParams{
		Force:      &forceDefault,
		HTTPClient: client,
	}
}

/*ActionBlueprintRequestUsingPOSTParams contains all the parameters to send to the API endpoint
for the action blueprint request using p o s t operation typically these are written to a http.Request
*/
type ActionBlueprintRequestUsingPOSTParams struct {

	/*Action
	  action

	*/
	Action string
	/*Force
	  Force cancellation of in progress tasks

	*/
	Force *bool
	/*RequestID
	  requestId

	*/
	RequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the action blueprint request using p o s t params
func (o *ActionBlueprintRequestUsingPOSTParams) WithTimeout(timeout time.Duration) *ActionBlueprintRequestUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the action blueprint request using p o s t params
func (o *ActionBlueprintRequestUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the action blueprint request using p o s t params
func (o *ActionBlueprintRequestUsingPOSTParams) WithContext(ctx context.Context) *ActionBlueprintRequestUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the action blueprint request using p o s t params
func (o *ActionBlueprintRequestUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the action blueprint request using p o s t params
func (o *ActionBlueprintRequestUsingPOSTParams) WithHTTPClient(client *http.Client) *ActionBlueprintRequestUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the action blueprint request using p o s t params
func (o *ActionBlueprintRequestUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAction adds the action to the action blueprint request using p o s t params
func (o *ActionBlueprintRequestUsingPOSTParams) WithAction(action string) *ActionBlueprintRequestUsingPOSTParams {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the action blueprint request using p o s t params
func (o *ActionBlueprintRequestUsingPOSTParams) SetAction(action string) {
	o.Action = action
}

// WithForce adds the force to the action blueprint request using p o s t params
func (o *ActionBlueprintRequestUsingPOSTParams) WithForce(force *bool) *ActionBlueprintRequestUsingPOSTParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the action blueprint request using p o s t params
func (o *ActionBlueprintRequestUsingPOSTParams) SetForce(force *bool) {
	o.Force = force
}

// WithRequestID adds the requestID to the action blueprint request using p o s t params
func (o *ActionBlueprintRequestUsingPOSTParams) WithRequestID(requestID string) *ActionBlueprintRequestUsingPOSTParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the action blueprint request using p o s t params
func (o *ActionBlueprintRequestUsingPOSTParams) SetRequestID(requestID string) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *ActionBlueprintRequestUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Force != nil {

		// query param force
		var qrForce bool
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {
			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}

	}

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
