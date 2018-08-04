package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// AddDashboard adds an dashboard to the mock server.
func (lgc *Logic) AddDashboard(dashboard *graylog.Dashboard) (int, error) {
	if err := validator.CreateValidator.Struct(dashboard); err != nil {
		return 400, err
	}
	if err := lgc.store.AddDashboard(dashboard); err != nil {
		return 500, err
	}
	return 201, nil
}

// DeleteDashboard deletes a dashboard from the mock server.
func (lgc *Logic) DeleteDashboard(id string) (int, error) {
	ok, err := lgc.HasDashboard(id)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("lgc.HasDashboard() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("the dashboard <%s> is not found", id)
	}
	if err := lgc.store.DeleteDashboard(id); err != nil {
		return 500, err
	}
	return 200, nil
}

// GetDashboard returns an dashboard.
// If an dashboard is not found, returns an error.
func (lgc *Logic) GetDashboard(id string) (*graylog.Dashboard, int, error) {
	if id == "" {
		return nil, 400, fmt.Errorf("dashboard id is empty")
	}
	if err := ValidateObjectID(id); err != nil {
		// unfortunately graylog returns not 400 but 404.
		return nil, 404, err
	}
	dashboard, err := lgc.store.GetDashboard(id)
	if err != nil {
		return dashboard, 500, err
	}
	if dashboard == nil {
		return nil, 404, fmt.Errorf("no dashboard with id <%s> is found", id)
	}
	return dashboard, 200, nil
}

// GetDashboards returns a list of dashboards.
func (lgc *Logic) GetDashboards() ([]graylog.Dashboard, int, int, error) {
	dashboards, total, err := lgc.store.GetDashboards()
	if err != nil {
		return nil, 0, 500, err
	}
	return dashboards, total, 200, nil
}

// HasDashboard returns whether the dashboard exists.
func (lgc *Logic) HasDashboard(id string) (bool, error) {
	return lgc.store.HasDashboard(id)
}

// UpdateDashboard updates an dashboard at the Server.
// Required: none
// Allowed: Title, Description
func (lgc *Logic) UpdateDashboard(dashboard *graylog.Dashboard) (int, error) {
	if dashboard == nil {
		return 400, fmt.Errorf("dashboard is nil")
	}
	if err := validator.UpdateValidator.Struct(dashboard); err != nil {
		return 400, err
	}
	ok, err := lgc.HasDashboard(dashboard.ID)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("the dashboard <%s> is not found", dashboard.ID)
	}

	if err := lgc.store.UpdateDashboard(dashboard); err != nil {
		return 500, err
	}
	return 204, nil
}
