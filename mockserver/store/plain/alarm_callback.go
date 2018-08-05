package plain

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

// GetAlarmCallbacks returns alarm callbacks.
func (store *Store) GetAlarmCallbacks() ([]graylog.AlarmCallback, int, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	size := len(store.alarmCallbacks)
	arr := make([]graylog.AlarmCallback, size)
	i := 0
	for _, ac := range store.alarmCallbacks {
		arr[i] = ac
		i++
	}
	return arr, size, nil
}
