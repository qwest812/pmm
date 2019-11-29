// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/percona/pmm/api/managementpb/json/client/actions"
	"github.com/percona/pmm/api/managementpb/json/client/mongo_db"
	"github.com/percona/pmm/api/managementpb/json/client/my_sql"
	"github.com/percona/pmm/api/managementpb/json/client/node"
	"github.com/percona/pmm/api/managementpb/json/client/postgre_sql"
	"github.com/percona/pmm/api/managementpb/json/client/proxy_sql"
	"github.com/percona/pmm/api/managementpb/json/client/rds"
	"github.com/percona/pmm/api/managementpb/json/client/service"
)

// Default PMM management HTTP client.
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

// NewHTTPClient creates a new PMM management HTTP client.
func NewHTTPClient(formats strfmt.Registry) *PMMManagement {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new PMM management HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *PMMManagement {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new PMM management client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *PMMManagement {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(PMMManagement)
	cli.Transport = transport

	cli.Actions = actions.New(transport, formats)

	cli.MongoDB = mongo_db.New(transport, formats)

	cli.MySQL = my_sql.New(transport, formats)

	cli.Node = node.New(transport, formats)

	cli.PostgreSQL = postgre_sql.New(transport, formats)

	cli.ProxySQL = proxy_sql.New(transport, formats)

	cli.RDS = rds.New(transport, formats)

	cli.Service = service.New(transport, formats)

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

// PMMManagement is a client for PMM management
type PMMManagement struct {
	Actions *actions.Client

	MongoDB *mongo_db.Client

	MySQL *my_sql.Client

	Node *node.Client

	PostgreSQL *postgre_sql.Client

	ProxySQL *proxy_sql.Client

	RDS *rds.Client

	Service *service.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *PMMManagement) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Actions.SetTransport(transport)

	c.MongoDB.SetTransport(transport)

	c.MySQL.SetTransport(transport)

	c.Node.SetTransport(transport)

	c.PostgreSQL.SetTransport(transport)

	c.ProxySQL.SetTransport(transport)

	c.RDS.SetTransport(transport)

	c.Service.SetTransport(transport)

}
