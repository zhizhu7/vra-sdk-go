// Code generated by go-swagger; DO NOT EDIT.

package health_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new health status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for health status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetHealthStatusUsingGET catalogs service health

Returns health of catalog service. The HTTP 200 success status response code indicates that the health is OK.
*/
func (a *Client) GetHealthStatusUsingGET(params *GetHealthStatusUsingGETParams) (*GetHealthStatusUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHealthStatusUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHealthStatusUsingGET",
		Method:             "GET",
		PathPattern:        "/catalog/api/health",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetHealthStatusUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetHealthStatusUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getHealthStatusUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
