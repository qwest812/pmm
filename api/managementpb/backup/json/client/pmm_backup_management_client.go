// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/percona/pmm/api/managementpb/backup/json/client/artifacts"
	"github.com/percona/pmm/api/managementpb/backup/json/client/backups"
	"github.com/percona/pmm/api/managementpb/backup/json/client/locations"
	"github.com/percona/pmm/api/managementpb/backup/json/client/restore_history"
)

// Default PMM backup management HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new PMM backup management HTTP client.
func NewHTTPClient(formats strfmt.Registry) *PMMBackupManagement {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new PMM backup management HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *PMMBackupManagement {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new PMM backup management client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *PMMBackupManagement {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(PMMBackupManagement)
	cli.Transport = transport
	cli.Artifacts = artifacts.New(transport, formats)
	cli.Backups = backups.New(transport, formats)
	cli.Locations = locations.New(transport, formats)
	cli.RestoreHistory = restore_history.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// PMMBackupManagement is a client for PMM backup management
type PMMBackupManagement struct {
	Artifacts artifacts.ClientService

	Backups backups.ClientService

	Locations locations.ClientService

	RestoreHistory restore_history.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *PMMBackupManagement) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Artifacts.SetTransport(transport)
	c.Backups.SetTransport(transport)
	c.Locations.SetTransport(transport)
	c.RestoreHistory.SetTransport(transport)
}
