package client

import (
	"context"
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// CreateDashboard creates a new dashboard account.
func (client *Client) CreateDashboard(dashboard *graylog.Dashboard) (*ErrorInfo, error) {
	return client.CreateDashboardContext(context.Background(), dashboard)
}

// CreateDashboardContext creates a new dashboard account with a context.
func (client *Client) CreateDashboardContext(
	ctx context.Context, dashboard *graylog.Dashboard,
) (*ErrorInfo, error) {
	if dashboard == nil {
		return nil, fmt.Errorf("dashboard is nil")
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
	return ei, fmt.Errorf(`response doesn't have the field "dashboard_id"`)
}

// DeleteDashboard deletes a given dashboard.
func (client *Client) DeleteDashboard(id string) (*ErrorInfo, error) {
	return client.DeleteDashboardContext(context.Background(), id)
}

// DeleteDashboardContext deletes a given dashboard with a context.
func (client *Client) DeleteDashboardContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, fmt.Errorf("id is empty")
	}
	u, err := client.Endpoints().Dashboard(id)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}

// GetDashboard returns a given dashboard.
func (client *Client) GetDashboard(id string) (*graylog.Dashboard, *ErrorInfo, error) {
	return client.GetDashboardContext(context.Background(), id)
}

// GetDashboardContext returns a given dashboard with a context.
func (client *Client) GetDashboardContext(
	ctx context.Context, id string,
) (*graylog.Dashboard, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is empty")
	}
	u, err := client.Endpoints().Dashboard(id)
	if err != nil {
		return nil, nil, err
	}
	dashboard := &graylog.Dashboard{}
	ei, err := client.callGet(ctx, u.String(), nil, dashboard)
	return dashboard, ei, err
}

// GetDashboards returns all dashboards.
func (client *Client) GetDashboards() ([]graylog.Dashboard, int, *ErrorInfo, error) {
	return client.GetDashboardsContext(context.Background())
}

// GetDashboardsContext returns all dashboards with a context.
func (client *Client) GetDashboardsContext(ctx context.Context) ([]graylog.Dashboard, int, *ErrorInfo, error) {
	dashboards := &graylog.DashboardsBody{}
	ei, err := client.callGet(
		ctx, client.Endpoints().Dashboards(), nil, dashboards)
	return dashboards.Dashboards, dashboards.Total, ei, err
}

// UpdateDashboard updates a given dashboard.
func (client *Client) UpdateDashboard(prms *graylog.Dashboard) (*ErrorInfo, error) {
	return client.UpdateDashboardContext(context.Background(), prms)
}

// UpdateDashboardContext updates a given dashboard with a context.
func (client *Client) UpdateDashboardContext(
	ctx context.Context, dashboard *graylog.Dashboard,
) (*ErrorInfo, error) {
	if dashboard == nil {
		return nil, fmt.Errorf("dashboard is nil")
	}
	if dashboard.ID == "" {
		return nil, fmt.Errorf("id is empty")
	}
	u, err := client.Endpoints().Dashboard(dashboard.ID)
	if err != nil {
		return nil, err
	}
	return client.callPut(ctx, u.String(), map[string]interface{}{
		"title":       dashboard.Title,
		"description": dashboard.Description,
	}, nil)
}

// UpdateDashboardWidgetPositions updates the positions of dashboard widgets.
func (client *Client) UpdateDashboardWidgetPositions(
	dashboardID string, positions []graylog.DashboardWidgetPosition,
) (*ErrorInfo, error) {
	return client.UpdateDashboardWidgetPositionsContext(
		context.Background(), dashboardID, positions)
}

// UpdateDashboardWidgetPositionsContext updates the positions of dashboard widgets.
func (client *Client) UpdateDashboardWidgetPositionsContext(
	ctx context.Context, dashboardID string,
	positions []graylog.DashboardWidgetPosition,
) (*ErrorInfo, error) {
	if dashboardID == "" {
		return nil, fmt.Errorf("id is empty")
	}
	u, err := client.Endpoints().DashboardWidgetsPosition(dashboardID)
	if err != nil {
		return nil, err
	}
	return client.callPut(ctx, u.String(), map[string]interface{}{
		"positions": positions,
	}, nil)
}
