// Code generated by go-swagger; DO NOT EDIT.

package blueprint

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

// NewListBlueprintVersionsUsingGETParams creates a new ListBlueprintVersionsUsingGETParams object
// with the default values initialized.
func NewListBlueprintVersionsUsingGETParams() *ListBlueprintVersionsUsingGETParams {
	var ()
	return &ListBlueprintVersionsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListBlueprintVersionsUsingGETParamsWithTimeout creates a new ListBlueprintVersionsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListBlueprintVersionsUsingGETParamsWithTimeout(timeout time.Duration) *ListBlueprintVersionsUsingGETParams {
	var ()
	return &ListBlueprintVersionsUsingGETParams{

		timeout: timeout,
	}
}

// NewListBlueprintVersionsUsingGETParamsWithContext creates a new ListBlueprintVersionsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewListBlueprintVersionsUsingGETParamsWithContext(ctx context.Context) *ListBlueprintVersionsUsingGETParams {
	var ()
	return &ListBlueprintVersionsUsingGETParams{

		Context: ctx,
	}
}

// NewListBlueprintVersionsUsingGETParamsWithHTTPClient creates a new ListBlueprintVersionsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListBlueprintVersionsUsingGETParamsWithHTTPClient(client *http.Client) *ListBlueprintVersionsUsingGETParams {
	var ()
	return &ListBlueprintVersionsUsingGETParams{
		HTTPClient: client,
	}
}

/*ListBlueprintVersionsUsingGETParams contains all the parameters to send to the API endpoint
for the list blueprint versions using g e t operation typically these are written to a http.Request
*/
type ListBlueprintVersionsUsingGETParams struct {

	/*DollarOrderby
	  Sorting criteria in the format: property (asc|desc). Default sort order is descending on updatedAt. Sorting is supported on fields createdAt, updatedAt, createdBy, updatedBy, name, version.

	*/
	DollarOrderby []string
	/*DollarSelect
	  Fields to include in content. Use * to get all fields.

	*/
	DollarSelect []string
	/*DollarSkip
	  Number of records you want to skip

	*/
	DollarSkip *int32
	/*DollarTop
	  Number of records you want

	*/
	DollarTop *int32
	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about

	*/
	APIVersion *string
	/*BlueprintID
	  blueprintId

	*/
	BlueprintID strfmt.UUID
	/*Status
	  Filter by blueprint status: versioned / released

	*/
	Status *string
	/*Version
	  Filter by version

	*/
	Version *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) WithTimeout(timeout time.Duration) *ListBlueprintVersionsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) WithContext(ctx context.Context) *ListBlueprintVersionsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) WithHTTPClient(client *http.Client) *ListBlueprintVersionsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarOrderby adds the dollarOrderby to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) WithDollarOrderby(dollarOrderby []string) *ListBlueprintVersionsUsingGETParams {
	o.SetDollarOrderby(dollarOrderby)
	return o
}

// SetDollarOrderby adds the dollarOrderby to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) SetDollarOrderby(dollarOrderby []string) {
	o.DollarOrderby = dollarOrderby
}

// WithDollarSelect adds the dollarSelect to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) WithDollarSelect(dollarSelect []string) *ListBlueprintVersionsUsingGETParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) SetDollarSelect(dollarSelect []string) {
	o.DollarSelect = dollarSelect
}

// WithDollarSkip adds the dollarSkip to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) WithDollarSkip(dollarSkip *int32) *ListBlueprintVersionsUsingGETParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) SetDollarSkip(dollarSkip *int32) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) WithDollarTop(dollarTop *int32) *ListBlueprintVersionsUsingGETParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) SetDollarTop(dollarTop *int32) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) WithAPIVersion(aPIVersion *string) *ListBlueprintVersionsUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBlueprintID adds the blueprintID to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) WithBlueprintID(blueprintID strfmt.UUID) *ListBlueprintVersionsUsingGETParams {
	o.SetBlueprintID(blueprintID)
	return o
}

// SetBlueprintID adds the blueprintId to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) SetBlueprintID(blueprintID strfmt.UUID) {
	o.BlueprintID = blueprintID
}

// WithStatus adds the status to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) WithStatus(status *string) *ListBlueprintVersionsUsingGETParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) SetStatus(status *string) {
	o.Status = status
}

// WithVersion adds the version to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) WithVersion(version *string) *ListBlueprintVersionsUsingGETParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the list blueprint versions using get params
func (o *ListBlueprintVersionsUsingGETParams) SetVersion(version *string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ListBlueprintVersionsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDollarOrderby := o.DollarOrderby

	joinedDollarOrderby := swag.JoinByFormat(valuesDollarOrderby, "multi")
	// query array param $orderby
	if err := r.SetQueryParam("$orderby", joinedDollarOrderby...); err != nil {
		return err
	}

	valuesDollarSelect := o.DollarSelect

	joinedDollarSelect := swag.JoinByFormat(valuesDollarSelect, "multi")
	// query array param $select
	if err := r.SetQueryParam("$select", joinedDollarSelect...); err != nil {
		return err
	}

	if o.DollarSkip != nil {

		// query param $skip
		var qrDollarSkip int32
		if o.DollarSkip != nil {
			qrDollarSkip = *o.DollarSkip
		}
		qDollarSkip := swag.FormatInt32(qrDollarSkip)
		if qDollarSkip != "" {
			if err := r.SetQueryParam("$skip", qDollarSkip); err != nil {
				return err
			}
		}

	}

	if o.DollarTop != nil {

		// query param $top
		var qrDollarTop int32
		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := swag.FormatInt32(qrDollarTop)
		if qDollarTop != "" {
			if err := r.SetQueryParam("$top", qDollarTop); err != nil {
				return err
			}
		}

	}

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

	// path param blueprintId
	if err := r.SetPathParam("blueprintId", o.BlueprintID.String()); err != nil {
		return err
	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if o.Version != nil {

		// query param version
		var qrVersion string
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := qrVersion
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}