// Code generated by go-swagger; DO NOT EDIT.

package db_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new db clusters API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for db clusters API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteDBCluster(params *DeleteDBClusterParams) (*DeleteDBClusterOK, error)

	ListDBClusters(params *ListDBClustersParams) (*ListDBClustersOK, error)

	RestartDBCluster(params *RestartDBClusterParams) (*RestartDBClusterOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteDBCluster deletes DB cluster deletes DB cluster
*/
func (a *Client) DeleteDBCluster(params *DeleteDBClusterParams) (*DeleteDBClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDBClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteDBCluster",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/DBClusters/Delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDBClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDBClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteDBClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListDBClusters lists DB clusters returns a list of DB clusters
*/
func (a *Client) ListDBClusters(params *ListDBClustersParams) (*ListDBClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDBClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListDBClusters",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/DBClusters/List",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListDBClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListDBClustersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDBClustersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RestartDBCluster restarts DB cluster restarts DB cluster
*/
func (a *Client) RestartDBCluster(params *RestartDBClusterParams) (*RestartDBClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestartDBClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RestartDBCluster",
		Method:             "POST",
		PathPattern:        "/v1/management/DBaaS/DBClusters/Restart",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestartDBClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestartDBClusterOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestartDBClusterDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
