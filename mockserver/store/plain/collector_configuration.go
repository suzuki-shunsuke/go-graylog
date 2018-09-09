package plain

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// HasCollectorConfiguration returns whether the collector configuration exists.
func (store *Store) HasCollectorConfiguration(id string) (bool, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	_, ok := store.collectorConfigurations[id]
	return ok, nil
}

// GetCollectorConfiguration returns an input.
func (store *Store) GetCollectorConfiguration(id string) (*graylog.CollectorConfiguration, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	s, ok := store.collectorConfigurations[id]
	if ok {
		return &s, nil
	}
	return nil, nil
}

// AddCollectorConfiguration adds a collector configuration to the store.
func (store *Store) AddCollectorConfiguration(cfg *graylog.CollectorConfiguration) error {
	if cfg == nil {
		return fmt.Errorf("collector configuration is nil")
	}
	if cfg.ID == "" {
		cfg.ID = st.NewObjectID()
	}

	store.imutex.Lock()
	defer store.imutex.Unlock()
	store.collectorConfigurations[cfg.ID] = *cfg
	return nil
}

// RenameCollectorConfiguration renames a collector configuration.
func (store *Store) RenameCollectorConfiguration(id, name string) (*graylog.CollectorConfiguration, error) {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	cfg, ok := store.collectorConfigurations[id]
	if !ok {
		return nil, fmt.Errorf("the collector configuration <%s> is not found", id)
	}
	cfg.Name = name
	store.collectorConfigurations[id] = cfg
	return &cfg, nil
}

// DeleteCollectorConfiguration deletes a collector configuration from the store.
func (store *Store) DeleteCollectorConfiguration(id string) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	delete(store.collectorConfigurations, id)
	return nil
}

// GetCollectorConfigurations returns all collector configurations.
func (store *Store) GetCollectorConfigurations() ([]graylog.CollectorConfiguration, int, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	size := len(store.collectorConfigurations)
	arr := make([]graylog.CollectorConfiguration, size)
	i := 0
	for _, cfg := range store.collectorConfigurations {
		arr[i] = cfg
		i++
	}
	return arr, size, nil
}
