package logic

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

// GetAlarmCallbacks returns a list of inputs.
func (lgc *Logic) GetAlarmCallbacks() ([]graylog.AlarmCallback, int, int, error) {
	arr, total, err := lgc.store.GetAlarmCallbacks()
	if err != nil {
		return nil, 0, 500, err
	}
	return arr, total, 200, nil
}
