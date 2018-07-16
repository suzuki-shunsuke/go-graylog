package plain

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

// HasAlert returns whether the alert exists.
func (store *Store) HasAlert(id string) (bool, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	_, ok := store.alerts[id]
	return ok, nil
}

// GetAlert returns an alert.
func (store *Store) GetAlert(id string) (*graylog.Alert, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	s, ok := store.alerts[id]
	if ok {
		return &s, nil
	}
	return nil, nil
}

// GetAlerts returns Alerts.
func (store *Store) GetAlerts(since, limit int) ([]graylog.Alert, int, error) {
	// TODO treat since parameter
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	size := len(store.alerts)
	if size == 0 {
		return nil, 0, nil
	}
	if limit > size {
		limit = size
	}

	arr := make([]graylog.Alert, limit)
	i := 0
	for _, a := range store.alerts {
		arr[i] = a
		i++
		if i == limit {
			break
		}
	}
	return arr, limit, nil
}
