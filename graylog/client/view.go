package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

// GetViews returns all views.
func (client *Client) GetViews(
	ctx context.Context,
) (*graylog.Views, *ErrorInfo, error) {
	viewsBody := &graylog.Views{}
	ei, err := client.callGet(
		ctx, client.Endpoints().Views(), nil, viewsBody)
	return viewsBody, ei, err
}

// GetView returns a given view.
func (client *Client) GetView(
	ctx context.Context, id string,
) (*graylog.View, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	view := &graylog.View{}
	ei, err := client.callGet(ctx, client.Endpoints().View(id), nil, view)
	return view, ei, err
}

// CreateView creates a view.
func (client *Client) CreateView(
	ctx context.Context, view *graylog.View,
) (*ErrorInfo, error) {
	// required: title search_id state
	// allowed: state, search_id, owner, summary, title, created_at, id, description, requires, properties, dashboard_state
	if view == nil {
		return nil, errors.New("view is nil")
	}
	if view.State == nil {
		view.State = map[string]graylog.ViewState{}
	}
	return client.callPost(ctx, client.Endpoints().Views(), view, view)
}

// UpdateView updates a view.
func (client *Client) UpdateView(
	ctx context.Context, view *graylog.View,
) (*ErrorInfo, error) {
	if view == nil {
		return nil, errors.New("view is nil")
	}
	if view.ID == "" {
		return nil, errors.New("id is empty")
	}
	if view.State == nil {
		view.State = map[string]graylog.ViewState{}
	}
	body := *view
	body.ID = ""
	return client.callPut(ctx, client.Endpoints().View(view.ID), &body, view)
}

// DeleteView deletes a view.
func (client *Client) DeleteView(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return client.callDelete(ctx, client.Endpoints().View(id), nil, nil)
}
