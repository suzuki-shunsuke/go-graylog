package plain

import (
	"fmt"
	"time"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// AddDashboard adds an dashboard to the store.
func (store *Store) AddDashboard(dashboard *graylog.Dashboard) error {
	if dashboard == nil {
		return fmt.Errorf("dashboard is nil")
	}
	if dashboard.ID == "" {
		dashboard.ID = st.NewObjectID()
	}
	dashboard.CreatedAt = time.Now().Format("2006-01-02T15:04:05.000Z")

	store.imutex.Lock()
	defer store.imutex.Unlock()
	store.dashboards[dashboard.ID] = *dashboard
	return nil
}

// DeleteDashboard deletes an dashboard from the store.
func (store *Store) DeleteDashboard(id string) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	delete(store.dashboards, id)
	return nil
}

// GetDashboard returns an dashboard.
func (store *Store) GetDashboard(id string) (*graylog.Dashboard, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	s, ok := store.dashboards[id]
	if ok {
		return &s, nil
	}
	return nil, nil
}

// GetDashboards returns dashboards.
func (store *Store) GetDashboards() ([]graylog.Dashboard, int, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	size := len(store.dashboards)
	arr := make([]graylog.Dashboard, size)
	i := 0
	for _, dashboard := range store.dashboards {
		arr[i] = dashboard
		i++
	}
	return arr, size, nil
}

// HasDashboard returns whether the dashboard exists.
func (store *Store) HasDashboard(id string) (bool, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	_, ok := store.dashboards[id]
	return ok, nil
}

// UpdateDashboard updates an dashboard at the Store.
// Allowed: title, description
func (store *Store) UpdateDashboard(dashboard *graylog.Dashboard) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	id := dashboard.ID
	d, ok := store.dashboards[id]
	if !ok {
		return fmt.Errorf("the dashboard <%s> is not found", id)
	}
	if dashboard.Title != "" {
		d.Title = dashboard.Title
	}
	if dashboard.Description != "" {
		d.Description = dashboard.Description
	}
	store.dashboards[id] = d
	return nil
}
