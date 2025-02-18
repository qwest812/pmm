// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new server API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AWSInstanceCheck(params *AWSInstanceCheckParams) (*AWSInstanceCheckOK, error)

	ChangeSettings(params *ChangeSettingsParams) (*ChangeSettingsOK, error)

	CheckUpdates(params *CheckUpdatesParams) (*CheckUpdatesOK, error)

	GetSettings(params *GetSettingsParams) (*GetSettingsOK, error)

	Logs(params *LogsParams, writer io.Writer) (*LogsOK, error)

	Readiness(params *ReadinessParams) (*ReadinessOK, error)

	StartUpdate(params *StartUpdateParams) (*StartUpdateOK, error)

	TestEmailAlertingSettings(params *TestEmailAlertingSettingsParams) (*TestEmailAlertingSettingsOK, error)

	UpdateStatus(params *UpdateStatusParams) (*UpdateStatusOK, error)

	Version(params *VersionParams) (*VersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AWSInstanceCheck AWSs instance check checks AWS e c2 instance ID
*/
func (a *Client) AWSInstanceCheck(params *AWSInstanceCheckParams) (*AWSInstanceCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAWSInstanceCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AWSInstanceCheck",
		Method:             "POST",
		PathPattern:        "/v1/AWSInstanceCheck",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AWSInstanceCheckReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AWSInstanceCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AWSInstanceCheckDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ChangeSettings changes settings changes PMM server settings
*/
func (a *Client) ChangeSettings(params *ChangeSettingsParams) (*ChangeSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeSettings",
		Method:             "POST",
		PathPattern:        "/v1/Settings/Change",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  CheckUpdates checks updates checks PMM server updates availability
*/
func (a *Client) CheckUpdates(params *CheckUpdatesParams) (*CheckUpdatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckUpdatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CheckUpdates",
		Method:             "POST",
		PathPattern:        "/v1/Updates/Check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CheckUpdatesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CheckUpdatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CheckUpdatesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetSettings gets settings returns current PMM server settings
*/
func (a *Client) GetSettings(params *GetSettingsParams) (*GetSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSettings",
		Method:             "POST",
		PathPattern:        "/v1/Settings/Get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Logs logs returns logs of the PMM server
*/
func (a *Client) Logs(params *LogsParams, writer io.Writer) (*LogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Logs",
		Method:             "GET",
		PathPattern:        "/logs.zip",
		ProducesMediaTypes: []string{"application/zip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LogsReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LogsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Readiness readinesses returns an error when some PMM server component is not ready yet or is being restarted it can be used as for docker health check or kubernetes readiness probe
*/
func (a *Client) Readiness(params *ReadinessParams) (*ReadinessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadinessParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Readiness",
		Method:             "GET",
		PathPattern:        "/v1/readyz",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReadinessReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadinessOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadinessDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartUpdate starts update starts PMM server update
*/
func (a *Client) StartUpdate(params *StartUpdateParams) (*StartUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartUpdate",
		Method:             "POST",
		PathPattern:        "/v1/Updates/Start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  TestEmailAlertingSettings tests email alerting settings tests email alerting SMTP settings by sending testing email
*/
func (a *Client) TestEmailAlertingSettings(params *TestEmailAlertingSettingsParams) (*TestEmailAlertingSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestEmailAlertingSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "TestEmailAlertingSettings",
		Method:             "POST",
		PathPattern:        "/v1/Settings/TestEmailAlertingSettings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TestEmailAlertingSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TestEmailAlertingSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TestEmailAlertingSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateStatus updates status returns PMM server update status
*/
func (a *Client) UpdateStatus(params *UpdateStatusParams) (*UpdateStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateStatus",
		Method:             "POST",
		PathPattern:        "/v1/Updates/Status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateStatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Version versions returns PMM server versions
*/
func (a *Client) Version(params *VersionParams) (*VersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Version",
		Method:             "GET",
		PathPattern:        "/v1/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
