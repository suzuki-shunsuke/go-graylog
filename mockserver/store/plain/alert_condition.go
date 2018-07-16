package plain

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

// GetAlertConditions returns Alert Conditions.
func (store *Store) GetAlertConditions() ([]graylog.AlertCondition, int, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	size := len(store.alertConditions)
	if size == 0 {
		return nil, 0, nil
	}
	arr := make([]graylog.AlertCondition, size)
	i := 0
	for _, cond := range store.alertConditions {
		arr[i] = cond
		i++
	}
	return arr, size, nil
}
