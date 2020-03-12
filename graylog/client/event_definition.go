package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

// CreateEventDefinition creates a new event definition.
func (client *Client) CreateEventDefinition(
	ctx context.Context, definition *graylog.EventDefinition,
) (*ErrorInfo, error) {
	// required: title, description, priority, alert, config, key_spec, notification_settings

	if definition == nil {
		return nil, errors.New("event definition is nil")
	}
	data := map[string]interface{}{
		"title":                 definition.Title,
		"description":           definition.Description,
		"priority":              definition.Priority,
		"alert":                 definition.Alert,
		"key_spec":              definition.KeySpec,
		"notification_settings": definition.NotificationSettings,
		"config":                definition.Config,
	}
	if len(definition.FieldSpec) != 0 {
		data["field_spec"] = definition.FieldSpec
	}
	if len(definition.Storage) != 0 {
		data["storage"] = definition.Storage
	}
	if len(definition.Notifications) != 0 {
		data["notifications"] = definition.Notifications
	}
	return client.callPost(
		ctx, client.Endpoints().EventDefinitions(),
		data, definition)
}

// GetEventDefinitions returns all event definitions.
func (client *Client) GetEventDefinitions(ctx context.Context) (
	*graylog.EventDefinitionsBody, *ErrorInfo, error,
) {
	definitions := &graylog.EventDefinitionsBody{}
	ei, err := client.callGet(ctx, client.Endpoints().EventDefinitions(), nil, definitions)
	return definitions, ei, err
}

// GetEventDefinition returns a given event definition.
func (client *Client) GetEventDefinition(
	ctx context.Context, id string,
) (*graylog.EventDefinition, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	definition := &graylog.EventDefinition{}
	ei, err := client.callGet(ctx, client.Endpoints().EventDefinition(id), nil, definition)
	return definition, ei, err
}

// UpdateEventDefinition updates a given event definition.
func (client *Client) UpdateEventDefinition(
	ctx context.Context, definition *graylog.EventDefinition,
) (*ErrorInfo, error) {
	// required title, description, priority, alert, config, key_spec notification_settings
	if definition == nil {
		return nil, errors.New("event definition is nil")
	}
	if definition.ID == "" {
		return nil, errors.New("id is empty")
	}
	data := map[string]interface{}{
		"id":                    definition.ID,
		"title":                 definition.Title,
		"description":           definition.Description,
		"priority":              definition.Priority,
		"alert":                 definition.Alert,
		"key_spec":              definition.KeySpec,
		"notification_settings": definition.NotificationSettings,
		"config":                definition.Config,
	}
	if len(definition.FieldSpec) != 0 {
		data["field_spec"] = definition.FieldSpec
	}
	if len(definition.Storage) != 0 {
		data["storage"] = definition.Storage
	}
	if len(definition.Notifications) != 0 {
		data["notifications"] = definition.Notifications
	}
	ei, err := client.callPut(
		ctx, client.Endpoints().EventDefinition(definition.ID), data, definition)
	return ei, err
}

// DeleteEventDefinition deletes a given event definition.
func (client *Client) DeleteEventDefinition(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return client.callDelete(ctx, client.Endpoints().EventDefinition(id), nil, nil)
}
