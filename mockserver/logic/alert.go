package logic

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetAlert returns an alert.
// If an alert is not found, returns an error.
func (lgc *Logic) GetAlert(id string) (*graylog.Alert, int, error) {
	if id == "" {
		return nil, 400, fmt.Errorf("alert id is empty")
	}
	if err := ValidateObjectID(id); err != nil {
		// unfortunately graylog returns not 400 but 404.
		return nil, 404, err
	}
	alert, err := lgc.store.GetAlert(id)
	if err != nil {
		return alert, 500, err
	}
	if alert == nil {
		return nil, 404, fmt.Errorf("no alert with id <%s> is found", id)
	}
	return alert, 200, nil
}

// GetAlerts returns a list of alerts.
func (lgc *Logic) GetAlerts(since, limit int) ([]graylog.Alert, int, int, error) {
	arr, total, err := lgc.store.GetAlerts(since, limit)
	if err != nil {
		return nil, 0, 500, err
	}
	return arr, total, 200, nil
}
