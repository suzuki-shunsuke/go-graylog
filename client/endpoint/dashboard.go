package endpoint

import (
	"net/url"
	"path"
)

// Dashboards returns a Dashboard API's endpoint url.
func (ep *Endpoints) Dashboards() string {
	return ep.dashboards.String()
}

// Dashboard returns a Dashboard API's endpoint url.
func (ep *Endpoints) Dashboard(id string) (*url.URL, error) {
	return urlJoin(ep.dashboards, id)
}

// DashboardWidgetsPosition returns a Dashboard widgets position API's endpoint url.
func (ep *Endpoints) DashboardWidgetsPosition(dashboardID string) (*url.URL, error) {
	// /dashboards/{dashboardId}/positions
	return urlJoin(ep.dashboards, path.Join(dashboardID, "positions"))
}
