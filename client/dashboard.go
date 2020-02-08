package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

// CreateDashboard creates a new dashboard account.
func (client *Client) CreateDashboard(
	ctx context.Context, dashboard *graylog.Dashboard,
) (*ErrorInfo, error) {
	if dashboard == nil {
		return nil, errors.New("dashboard is nil")
	}

	ret := map[string]string{}
	ei, err := client.callPost(
		ctx, client.Endpoints().Dashboards(), map[string]interface{}{
			"title":       dashboard.Title,
			"description": dashboard.Description,
		}, &ret)
	if err != nil {
		return ei, err
	}
	if id, ok := ret["dashboard_id"]; ok {
		dashboard.ID = id
		return ei, nil
	}
	return ei, errors.New(`response doesn't have the field "dashboard_id"`)
}

// DeleteDashboard deletes a given dashboard.
func (client *Client) DeleteDashboard(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return client.callDelete(ctx, client.Endpoints().Dashboard(id), nil, nil)
}

// GetDashboard returns a given dashboard.
func (client *Client) GetDashboard(
	ctx context.Context, id string,
) (*graylog.Dashboard, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	dashboard := &graylog.Dashboard{}
	ei, err := client.callGet(ctx, client.Endpoints().Dashboard(id), nil, dashboard)
	return dashboard, ei, err
}

// GetDashboards returns all dashboards.
func (client *Client) GetDashboards(ctx context.Context) ([]graylog.Dashboard, int, *ErrorInfo, error) {
	dashboards := &graylog.DashboardsBody{}
	ei, err := client.callGet(
		ctx, client.Endpoints().Dashboards(), nil, dashboards)
	return dashboards.Dashboards, dashboards.Total, ei, err
}

// UpdateDashboard updates a given dashboard.
func (client *Client) UpdateDashboard(
	ctx context.Context, dashboard *graylog.Dashboard,
) (*ErrorInfo, error) {
	if dashboard == nil {
		return nil, errors.New("dashboard is nil")
	}
	if dashboard.ID == "" {
		return nil, errors.New("id is empty")
	}
	return client.callPut(ctx, client.Endpoints().Dashboard(dashboard.ID), map[string]interface{}{
		"title":       dashboard.Title,
		"description": dashboard.Description,
	}, nil)
}

// UpdateDashboardWidgetPositions updates the positions of dashboard widgets.
func (client *Client) UpdateDashboardWidgetPositions(
	ctx context.Context, dashboardID string,
	positions []graylog.DashboardWidgetPosition,
) (*ErrorInfo, error) {
	if dashboardID == "" {
		return nil, errors.New("id is empty")
	}
	return client.callPut(
		ctx, client.Endpoints().DashboardWidgetsPosition(dashboardID), map[string]interface{}{
			"positions": positions,
		}, nil)
}
