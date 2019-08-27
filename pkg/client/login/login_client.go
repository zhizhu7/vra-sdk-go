// Code generated by go-swagger; DO NOT EDIT.

package login

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new login API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for login API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
RetrieveAuthToken retrieves auth token for local csp users

Retrieve AuthToken for local csp users.
When accessing other endpoints the `Bearer` authentication scheme and the received `token` must be provided in the `Authorization` request header field as follows:
`Authorization: Bearer {token}`
*/
func (a *Client) RetrieveAuthToken(params *RetrieveAuthTokenParams) (*RetrieveAuthTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveAuthTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveAuthToken",
		Method:             "POST",
		PathPattern:        "/iaas/api/login",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RetrieveAuthTokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RetrieveAuthTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for retrieveAuthToken: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
