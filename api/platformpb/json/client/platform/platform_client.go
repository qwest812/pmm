// Code generated by go-swagger; DO NOT EDIT.

package platform

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new platform API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for platform API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	Connect(params *ConnectParams) (*ConnectOK, error)

	Disconnect(params *DisconnectParams) (*DisconnectOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  Connect connects a PMM server to the organization created on percona portal that allows the user to sign in to the PMM server with their percona account
*/
func (a *Client) Connect(params *ConnectParams) (*ConnectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConnectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Connect",
		Method:             "POST",
		PathPattern:        "/v1/Platform/Connect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ConnectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConnectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ConnectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Disconnect disconnects a PMM server from the organization created on percona portal
*/
func (a *Client) Disconnect(params *DisconnectParams) (*DisconnectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisconnectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Disconnect",
		Method:             "POST",
		PathPattern:        "/v1/Platform/Disconnect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisconnectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisconnectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DisconnectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
