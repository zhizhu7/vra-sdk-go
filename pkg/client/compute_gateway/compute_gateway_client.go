// Code generated by go-swagger; DO NOT EDIT.

package compute_gateway

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new compute gateway API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for compute gateway API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateComputeGateway creates a compute gateway

Create a new compute gateway.
*/
func (a *Client) CreateComputeGateway(params *CreateComputeGatewayParams) (*CreateComputeGatewayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateComputeGatewayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createComputeGateway",
		Method:             "POST",
		PathPattern:        "/iaas/api/compute-gateways",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateComputeGatewayReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateComputeGatewayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createComputeGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteComputeGateway deletes a compute gateway

Delete compute gateway with a given id
*/
func (a *Client) DeleteComputeGateway(params *DeleteComputeGatewayParams) (*DeleteComputeGatewayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteComputeGatewayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteComputeGateway",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/compute-gateways/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteComputeGatewayReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteComputeGatewayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteComputeGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetComputeGateway gets compute gateway

Get all compute gateways
*/
func (a *Client) GetComputeGateway(params *GetComputeGatewayParams) (*GetComputeGatewayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComputeGatewayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getComputeGateway",
		Method:             "GET",
		PathPattern:        "/iaas/api/compute-gateways",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComputeGatewayReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetComputeGatewayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getComputeGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetGateway gets gateway

Get gateway with a given id
*/
func (a *Client) GetGateway(params *GetGatewayParams) (*GetGatewayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGatewayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGateway",
		Method:             "GET",
		PathPattern:        "/iaas/api/compute-gateways/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetGatewayReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGatewayOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGateway: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}