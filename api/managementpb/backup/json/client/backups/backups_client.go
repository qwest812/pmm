// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ChangeScheduledBackup(params *ChangeScheduledBackupParams) (*ChangeScheduledBackupOK, error)

	GetLogs(params *GetLogsParams) (*GetLogsOK, error)

	ListArtifactCompatibleServices(params *ListArtifactCompatibleServicesParams) (*ListArtifactCompatibleServicesOK, error)

	ListScheduledBackups(params *ListScheduledBackupsParams) (*ListScheduledBackupsOK, error)

	RemoveScheduledBackup(params *RemoveScheduledBackupParams) (*RemoveScheduledBackupOK, error)

	RestoreBackup(params *RestoreBackupParams) (*RestoreBackupOK, error)

	ScheduleBackup(params *ScheduleBackupParams) (*ScheduleBackupOK, error)

	StartBackup(params *StartBackupParams) (*StartBackupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ChangeScheduledBackup changes scheduled backup changes existing scheduled backup
*/
func (a *Client) ChangeScheduledBackup(params *ChangeScheduledBackupParams) (*ChangeScheduledBackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeScheduledBackupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ChangeScheduledBackup",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/Backups/ChangeScheduled",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeScheduledBackupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ChangeScheduledBackupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeScheduledBackupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetLogs gets logs returns logs for provided artifact
*/
func (a *Client) GetLogs(params *GetLogsParams) (*GetLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLogs",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/Backups/GetLogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetLogsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLogsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListArtifactCompatibleServices lists artifact compatible services lists compatible services for restoring a backup
*/
func (a *Client) ListArtifactCompatibleServices(params *ListArtifactCompatibleServicesParams) (*ListArtifactCompatibleServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListArtifactCompatibleServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListArtifactCompatibleServices",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/Backups/ListArtifactCompatibleServices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListArtifactCompatibleServicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListArtifactCompatibleServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListArtifactCompatibleServicesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListScheduledBackups lists scheduled backups returns all scheduled backups
*/
func (a *Client) ListScheduledBackups(params *ListScheduledBackupsParams) (*ListScheduledBackupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListScheduledBackupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListScheduledBackups",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/Backups/ListScheduled",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListScheduledBackupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListScheduledBackupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListScheduledBackupsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RemoveScheduledBackup removes scheduled backup removes existing scheduled backup
*/
func (a *Client) RemoveScheduledBackup(params *RemoveScheduledBackupParams) (*RemoveScheduledBackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveScheduledBackupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RemoveScheduledBackup",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/Backups/RemoveScheduled",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveScheduledBackupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveScheduledBackupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RemoveScheduledBackupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  RestoreBackup restores backup requests the backup restore

  Could return the Error message in the details containing specific ErrorCode indicating failure reason:
ERROR_CODE_XTRABACKUP_NOT_INSTALLED - xtrabackup is not installed on the service
ERROR_CODE_INVALID_XTRABACKUP - different versions of xtrabackup and xbcloud
ERROR_CODE_INCOMPATIBLE_XTRABACKUP - xtrabackup is not compatible with MySQL for taking a backup
ERROR_CODE_INCOMPATIBLE_TARGET_MYSQL - target MySQL version is not compatible with the artifact for performing a restore of the backup
*/
func (a *Client) RestoreBackup(params *RestoreBackupParams) (*RestoreBackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreBackupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RestoreBackup",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/Backups/Restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestoreBackupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestoreBackupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestoreBackupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ScheduleBackup schedules backup schedules repeated backup
*/
func (a *Client) ScheduleBackup(params *ScheduleBackupParams) (*ScheduleBackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScheduleBackupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ScheduleBackup",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/Backups/Schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ScheduleBackupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ScheduleBackupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ScheduleBackupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  StartBackup starts backup request backup specified service to location

  Could return the Error message in the details containing specific ErrorCode indicating failure reason:
ERROR_CODE_XTRABACKUP_NOT_INSTALLED - xtrabackup is not installed on the service
ERROR_CODE_INVALID_XTRABACKUP - different versions of xtrabackup and xbcloud
ERROR_CODE_INCOMPATIBLE_XTRABACKUP - xtrabackup is not compatible with MySQL for taking a backup
*/
func (a *Client) StartBackup(params *StartBackupParams) (*StartBackupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartBackupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartBackup",
		Method:             "POST",
		PathPattern:        "/v1/management/backup/Backups/Start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartBackupReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartBackupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartBackupDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
