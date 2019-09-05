package endpoint

// Dashboards returns a Dashboard API's endpoint url.
func (ep *Endpoints) Dashboards() string {
	return ep.dashboards
}

// Dashboard returns a Dashboard API's endpoint url.
func (ep *Endpoints) Dashboard(id string) string {
	return ep.dashboards + "/" + id
}

// DashboardWidgetsPosition returns a Dashboard widgets position API's endpoint url.
func (ep *Endpoints) DashboardWidgetsPosition(dashboardID string) string {
	// /dashboards/{dashboardId}/positions
	return ep.dashboards + "/" + dashboardID + "/positions"
}
