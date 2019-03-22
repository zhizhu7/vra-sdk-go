// Code generated by go-swagger; DO NOT EDIT.

package catalog_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new catalog items API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog items API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCatalogItemUsingGET returns the catalog item with the specified ID
*/
func (a *Client) GetCatalogItemUsingGET(params *GetCatalogItemUsingGETParams) (*GetCatalogItemUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogItemUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCatalogItemUsingGET",
		Method:             "GET",
		PathPattern:        "/catalog/api/items/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCatalogItemUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCatalogItemUsingGETOK), nil

}

/*
GetCatalogItemsUsingGET returns a paginated list of catalog items
*/
func (a *Client) GetCatalogItemsUsingGET(params *GetCatalogItemsUsingGETParams) (*GetCatalogItemsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogItemsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCatalogItemsUsingGET",
		Method:             "GET",
		PathPattern:        "/catalog/api/items",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCatalogItemsUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCatalogItemsUsingGETOK), nil

}

/*
RequestCatalogItemUsingPOST creates a deployment from a catalog item
*/
func (a *Client) RequestCatalogItemUsingPOST(params *RequestCatalogItemUsingPOSTParams) (*RequestCatalogItemUsingPOSTOK, *RequestCatalogItemUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRequestCatalogItemUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "requestCatalogItemUsingPOST",
		Method:             "POST",
		PathPattern:        "/catalog/api/items/{id}/request",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RequestCatalogItemUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *RequestCatalogItemUsingPOSTOK:
		return value, nil, nil
	case *RequestCatalogItemUsingPOSTCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
