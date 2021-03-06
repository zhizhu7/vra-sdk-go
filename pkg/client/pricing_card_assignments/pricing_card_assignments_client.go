// Code generated by go-swagger; DO NOT EDIT.

package pricing_card_assignments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new pricing card assignments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pricing card assignments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ChangeMeteringAssignmentStrategyUsingPATCH updates the pricing card assignment strategy for the org
*/
func (a *Client) ChangeMeteringAssignmentStrategyUsingPATCH(params *ChangeMeteringAssignmentStrategyUsingPATCHParams) (*ChangeMeteringAssignmentStrategyUsingPATCHOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeMeteringAssignmentStrategyUsingPATCHParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeMeteringAssignmentStrategyUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/price/api/private/pricing-card-assignments/strategy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangeMeteringAssignmentStrategyUsingPATCHReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeMeteringAssignmentStrategyUsingPATCHOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for changeMeteringAssignmentStrategyUsingPATCH: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateMeteringAssignmentStrategyUsingPOST selectings the new pricing card assignment strategy p r o j e c t or c l o u d z o n e are possible values can be used while creating strategy also there can be only one strategy for a given org at a given point of time

Create a new pricing card assignment strategy based on request body and validate its field according to business rules.
*/
func (a *Client) CreateMeteringAssignmentStrategyUsingPOST(params *CreateMeteringAssignmentStrategyUsingPOSTParams) (*CreateMeteringAssignmentStrategyUsingPOSTOK, *CreateMeteringAssignmentStrategyUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMeteringAssignmentStrategyUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createMeteringAssignmentStrategyUsingPOST",
		Method:             "POST",
		PathPattern:        "/price/api/private/pricing-card-assignments/strategy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMeteringAssignmentStrategyUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateMeteringAssignmentStrategyUsingPOSTOK:
		return value, nil, nil
	case *CreateMeteringAssignmentStrategyUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pricing_card_assignments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateMeteringPolicyAssignmentUsingPOST creates a new pricing card assignment

Create a new pricing card policy assignment based on request body and validate its field according to business rules. Request body with ALL entityType will delete the older assignments for the given pricingCardId
*/
func (a *Client) CreateMeteringPolicyAssignmentUsingPOST(params *CreateMeteringPolicyAssignmentUsingPOSTParams) (*CreateMeteringPolicyAssignmentUsingPOSTOK, *CreateMeteringPolicyAssignmentUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateMeteringPolicyAssignmentUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createMeteringPolicyAssignmentUsingPOST",
		Method:             "POST",
		PathPattern:        "/price/api/private/pricing-card-assignments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateMeteringPolicyAssignmentUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreateMeteringPolicyAssignmentUsingPOSTOK:
		return value, nil, nil
	case *CreateMeteringPolicyAssignmentUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pricing_card_assignments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteMeteringPolicyAssignmentUsingDELETE deletes the pricing card assignment with specified id

Deletes the pricing card assignment with the specified id
*/
func (a *Client) DeleteMeteringPolicyAssignmentUsingDELETE(params *DeleteMeteringPolicyAssignmentUsingDELETEParams) (*DeleteMeteringPolicyAssignmentUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMeteringPolicyAssignmentUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMeteringPolicyAssignmentUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/price/api/private/pricing-card-assignments/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMeteringPolicyAssignmentUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteMeteringPolicyAssignmentUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteMeteringPolicyAssignmentUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAllMeteringPolicyAssignmentsUsingGET fetches all pricing card assignment for private cloud

Returns a paginated list of pricing card assignments
*/
func (a *Client) GetAllMeteringPolicyAssignmentsUsingGET(params *GetAllMeteringPolicyAssignmentsUsingGETParams) (*GetAllMeteringPolicyAssignmentsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllMeteringPolicyAssignmentsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllMeteringPolicyAssignmentsUsingGET",
		Method:             "GET",
		PathPattern:        "/price/api/private/pricing-card-assignments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllMeteringPolicyAssignmentsUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllMeteringPolicyAssignmentsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllMeteringPolicyAssignmentsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMeteringAssignmentStrategyUsingGET fetches pricing card assignment strategy for the org

Returns a pricing card assignment strategy for the Org
*/
func (a *Client) GetMeteringAssignmentStrategyUsingGET(params *GetMeteringAssignmentStrategyUsingGETParams) (*GetMeteringAssignmentStrategyUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMeteringAssignmentStrategyUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMeteringAssignmentStrategyUsingGET",
		Method:             "GET",
		PathPattern:        "/price/api/private/pricing-card-assignments/strategy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMeteringAssignmentStrategyUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMeteringAssignmentStrategyUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMeteringAssignmentStrategyUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMeteringPolicyAssignmentUsingGET fetches pricing card assignment for private cloud by id

Returns a pricing card assignments by id
*/
func (a *Client) GetMeteringPolicyAssignmentUsingGET(params *GetMeteringPolicyAssignmentUsingGETParams) (*GetMeteringPolicyAssignmentUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMeteringPolicyAssignmentUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMeteringPolicyAssignmentUsingGET",
		Method:             "GET",
		PathPattern:        "/price/api/private/pricing-card-assignments/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMeteringPolicyAssignmentUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMeteringPolicyAssignmentUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMeteringPolicyAssignmentUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PatchMeteringPolicyAssignmentUsingPATCH updates the pricing card assignment id with the supplied id request body with a l l entity type will delete the older assignments for the given pricing card Id
*/
func (a *Client) PatchMeteringPolicyAssignmentUsingPATCH(params *PatchMeteringPolicyAssignmentUsingPATCHParams) (*PatchMeteringPolicyAssignmentUsingPATCHOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchMeteringPolicyAssignmentUsingPATCHParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchMeteringPolicyAssignmentUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/price/api/private/pricing-card-assignments/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchMeteringPolicyAssignmentUsingPATCHReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchMeteringPolicyAssignmentUsingPATCHOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patchMeteringPolicyAssignmentUsingPATCH: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
