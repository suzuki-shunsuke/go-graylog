package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

// CreateCollectorConfiguration creates a collector configuration.
func (client *Client) CreateCollectorConfiguration(
	ctx context.Context, cfg *graylog.CollectorConfiguration,
) (*ErrorInfo, error) {
	// POST /plugins/org.graylog.plugins.collector/configurations Create new collector configuration
	if cfg == nil {
		return nil, errors.New("collector configuration is nil")
	}
	if cfg.Inputs == nil {
		cfg.Inputs = []graylog.CollectorConfigurationInput{}
	}
	if cfg.Outputs == nil {
		cfg.Outputs = []graylog.CollectorConfigurationOutput{}
	}
	if cfg.Snippets == nil {
		cfg.Snippets = []graylog.CollectorConfigurationSnippet{}
	}
	if cfg.Tags == nil {
		cfg.Tags = set.StrSet{}
	}
	return client.callPost(ctx, client.Endpoints().CollectorConfigurations(), cfg, cfg)
}

// GetCollectorConfigurations returns all collector configurations.
func (client *Client) GetCollectorConfigurations(ctx context.Context) ([]graylog.CollectorConfiguration, int, *ErrorInfo, error) {
	cfgs := &graylog.CollectorConfigurationsBody{}
	ei, err := client.callGet(
		ctx, client.Endpoints().CollectorConfigurations(), nil, cfgs)
	return cfgs.Configurations, cfgs.Total, ei, err
}

// GetCollectorConfiguration returns a given user.
func (client *Client) GetCollectorConfiguration(
	ctx context.Context, id string,
) (*graylog.CollectorConfiguration, *ErrorInfo, error) {
	// GET /api/plugins/org.graylog.plugins.collector/configurations/:id
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	cfg := &graylog.CollectorConfiguration{}
	ei, err := client.callGet(
		ctx, client.Endpoints().CollectorConfiguration(id), nil, cfg)
	return cfg, ei, err
}

// RenameCollectorConfiguration renames a collector configuration.
func (client *Client) RenameCollectorConfiguration(
	ctx context.Context, id, name string,
) (*graylog.CollectorConfiguration, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is nil")
	}
	if name == "" {
		return nil, nil, errors.New("name is nil")
	}
	input := graylog.CollectorConfiguration{
		Name:     name,
		Inputs:   []graylog.CollectorConfigurationInput{},
		Outputs:  []graylog.CollectorConfigurationOutput{},
		Snippets: []graylog.CollectorConfigurationSnippet{},
	}
	cfg := graylog.CollectorConfiguration{Name: name}
	ei, err := client.callPut(
		ctx, client.Endpoints().CollectorConfigurationName(id), &input, &cfg)
	return &cfg, ei, err
}

// DeleteCollectorConfiguration deletes a collector configuration.
func (client *Client) DeleteCollectorConfiguration(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return client.callDelete(
		ctx, client.Endpoints().CollectorConfiguration(id), nil, nil)
}
