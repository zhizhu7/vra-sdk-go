// Code generated by go-swagger; DO NOT EDIT.

package blueprint_validation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new blueprint validation API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for blueprint validation API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ValidateBlueprintUsingPOST validates a blueprint
*/
func (a *Client) ValidateBlueprintUsingPOST(params *ValidateBlueprintUsingPOSTParams) (*ValidateBlueprintUsingPOSTOK, *ValidateBlueprintUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateBlueprintUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateBlueprintUsingPOST",
		Method:             "POST",
		PathPattern:        "/blueprint/api/blueprint-validation",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateBlueprintUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ValidateBlueprintUsingPOSTOK:
		return value, nil, nil
	case *ValidateBlueprintUsingPOSTCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
