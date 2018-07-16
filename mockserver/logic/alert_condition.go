package logic

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

// GetAlertConditions returns a list of alert conditions.
func (lgc *Logic) GetAlertConditions() ([]graylog.AlertCondition, int, int, error) {
	conds, total, err := lgc.store.GetAlertConditions()
	if err != nil {
		return nil, 0, 500, err
	}
	return conds, total, 200, nil
}
