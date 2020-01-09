package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v9"
)

// CreateEventNotification creates a new event notification.
func (client *Client) CreateEventNotification(
	ctx context.Context, notif *graylog.EventNotification,
) (*ErrorInfo, error) {
	// required: title, type, configuration
	if notif == nil {
		return nil, errors.New("event notification is nil")
	}
	return client.callPost(
		ctx, client.Endpoints().EventNotifications(),
		map[string]interface{}{
			"title":       notif.Title,
			"description": notif.Description,
			"config":      notif.Config,
		}, notif)
}

// GetEventNotifications returns all event notifications.
func (client *Client) GetEventNotifications(ctx context.Context) (
	*graylog.EventNotificationsBody, *ErrorInfo, error,
) {
	notifs := &graylog.EventNotificationsBody{}
	ei, err := client.callGet(ctx, client.Endpoints().EventNotifications(), nil, notifs)
	return notifs, ei, err
}

// GetEventNotification returns a given event notification.
func (client *Client) GetEventNotification(
	ctx context.Context, id string,
) (*graylog.EventNotification, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	notif := &graylog.EventNotification{}
	ei, err := client.callGet(ctx, client.Endpoints().EventNotification(id), nil, notif)
	return notif, ei, err
}

// UpdateEventNotification updates a given event notification.
func (client *Client) UpdateEventNotification(
	ctx context.Context, notif *graylog.EventNotification,
) (*ErrorInfo, error) {
	if notif == nil {
		return nil, errors.New("event notification is nil")
	}
	if notif.ID == "" {
		return nil, errors.New("id is empty")
	}
	ei, err := client.callPut(
		ctx, client.Endpoints().EventNotification(notif.ID), map[string]interface{}{
			"id":          notif.ID,
			"title":       notif.Title,
			"description": notif.Description,
			"config":      notif.Config,
		}, notif)
	return ei, err
}

// DeleteEventNotification deletes a given event notification.
func (client *Client) DeleteEventNotification(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return client.callDelete(ctx, client.Endpoints().EventNotification(id), nil, nil)
}
