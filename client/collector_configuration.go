package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// CreateCollectorConfiguration creates a collector configuration.
func (client *Client) CreateCollectorConfiguration(cfg *graylog.CollectorConfiguration) (*ErrorInfo, error) {
	return client.CreateCollectorConfigurationContext(context.Background(), cfg)
}

// CreateCollectorConfigurationContext creates a collector configuration with a context.
func (client *Client) CreateCollectorConfigurationContext(
	ctx context.Context, cfg *graylog.CollectorConfiguration,
) (*ErrorInfo, error) {
	// POST /plugins/org.graylog.plugins.collector/configurations Create new collector configuration
	if cfg == nil {
		return nil, fmt.Errorf("collector configuration is nil")
	}
	if len(cfg.Inputs) == 0 {
		cfg.Inputs = []graylog.CollectorConfigurationInput{}
	}
	if len(cfg.Outputs) == 0 {
		cfg.Outputs = []graylog.CollectorConfigurationOutput{}
	}
	if len(cfg.Snippets) == 0 {
		cfg.Snippets = []graylog.CollectorConfigurationSnippet{}
	}
	return client.callPost(ctx, client.Endpoints().CollectorConfigurations(), cfg, cfg)
}

// GetCollectorConfigurations returns all collector configurations.
func (client *Client) GetCollectorConfigurations() ([]graylog.CollectorConfiguration, int, *ErrorInfo, error) {
	return client.GetCollectorConfigurationsContext(context.Background())
}

// GetCollectorConfigurationsContext returns all collector configurations with a context.
func (client *Client) GetCollectorConfigurationsContext(ctx context.Context) ([]graylog.CollectorConfiguration, int, *ErrorInfo, error) {
	cfgs := &graylog.CollectorConfigurationsBody{}
	ei, err := client.callGet(
		ctx, client.Endpoints().CollectorConfigurations(), nil, cfgs)
	return cfgs.Configurations, cfgs.Total, ei, err
}

// GetCollectorConfiguration returns a collector configuration.
func (client *Client) GetCollectorConfiguration(id string) (*graylog.CollectorConfiguration, *ErrorInfo, error) {
	return client.GetCollectorConfigurationContext(context.Background(), id)
}

// GetCollectorConfigurationContext returns a given user with a context.
func (client *Client) GetCollectorConfigurationContext(
	ctx context.Context, id string,
) (*graylog.CollectorConfiguration, *ErrorInfo, error) {
	// GET /api/plugins/org.graylog.plugins.collector/configurations/:id
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().CollectorConfiguration(id)
	if err != nil {
		return nil, nil, err
	}
	cfg := &graylog.CollectorConfiguration{}
	ei, err := client.callGet(ctx, u.String(), nil, cfg)
	return cfg, ei, err
}

// RenameCollectorConfiguration renames a collector configuration.
func (client *Client) RenameCollectorConfiguration(id, name string) (*graylog.CollectorConfiguration, *ErrorInfo, error) {
	return client.RenameCollectorConfigurationContext(context.Background(), id, name)
}

// RenameCollectorConfigurationContext renames a collector configuration with a context.
func (client *Client) RenameCollectorConfigurationContext(
	ctx context.Context, id, name string,
) (*graylog.CollectorConfiguration, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is nil")
	}
	if name == "" {
		return nil, nil, fmt.Errorf("name is nil")
	}
	input := graylog.CollectorConfiguration{
		Name:     name,
		Inputs:   []graylog.CollectorConfigurationInput{},
		Outputs:  []graylog.CollectorConfigurationOutput{},
		Snippets: []graylog.CollectorConfigurationSnippet{},
	}
	u, err := client.Endpoints().CollectorConfigurationName(id)
	if err != nil {
		return nil, nil, err
	}
	cfg := graylog.CollectorConfiguration{Name: name}
	ei, err := client.callPut(ctx, u.String(), &input, &cfg)
	return &cfg, ei, err
}

// DeleteCollectorConfiguration deletes a collector configuration.
func (client *Client) DeleteCollectorConfiguration(id string) (*ErrorInfo, error) {
	return client.DeleteCollectorConfigurationContext(context.Background(), id)
}

// DeleteCollectorConfigurationContext deletes a collector configuration with a context.
func (client *Client) DeleteCollectorConfigurationContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().CollectorConfiguration(id)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
