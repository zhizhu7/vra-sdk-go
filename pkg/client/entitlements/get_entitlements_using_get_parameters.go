// Code generated by go-swagger; DO NOT EDIT.

package entitlements

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

// NewGetEntitlementsUsingGETParams creates a new GetEntitlementsUsingGETParams object
// with the default values initialized.
func NewGetEntitlementsUsingGETParams() *GetEntitlementsUsingGETParams {
	var ()
	return &GetEntitlementsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEntitlementsUsingGETParamsWithTimeout creates a new GetEntitlementsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEntitlementsUsingGETParamsWithTimeout(timeout time.Duration) *GetEntitlementsUsingGETParams {
	var ()
	return &GetEntitlementsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetEntitlementsUsingGETParamsWithContext creates a new GetEntitlementsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEntitlementsUsingGETParamsWithContext(ctx context.Context) *GetEntitlementsUsingGETParams {
	var ()
	return &GetEntitlementsUsingGETParams{

		Context: ctx,
	}
}

// NewGetEntitlementsUsingGETParamsWithHTTPClient creates a new GetEntitlementsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEntitlementsUsingGETParamsWithHTTPClient(client *http.Client) *GetEntitlementsUsingGETParams {
	var ()
	return &GetEntitlementsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetEntitlementsUsingGETParams contains all the parameters to send to the API endpoint
for the get entitlements using g e t operation typically these are written to a http.Request
*/
type GetEntitlementsUsingGETParams struct {

	/*ProjectID
	  The project id for which to return .

	*/
	ProjectID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get entitlements using g e t params
func (o *GetEntitlementsUsingGETParams) WithTimeout(timeout time.Duration) *GetEntitlementsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get entitlements using g e t params
func (o *GetEntitlementsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get entitlements using g e t params
func (o *GetEntitlementsUsingGETParams) WithContext(ctx context.Context) *GetEntitlementsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get entitlements using g e t params
func (o *GetEntitlementsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get entitlements using g e t params
func (o *GetEntitlementsUsingGETParams) WithHTTPClient(client *http.Client) *GetEntitlementsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get entitlements using g e t params
func (o *GetEntitlementsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get entitlements using g e t params
func (o *GetEntitlementsUsingGETParams) WithProjectID(projectID *string) *GetEntitlementsUsingGETParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get entitlements using g e t params
func (o *GetEntitlementsUsingGETParams) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEntitlementsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID string
		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {
			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
