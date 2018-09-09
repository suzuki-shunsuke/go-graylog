package plain

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// HasCollectorConfigurationInput returns whether the collector configuration input exists.
func (store *Store) HasCollectorConfigurationInput(cfgID, inputID string) (bool, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	cfg, ok := store.collectorConfigurations[cfgID]
	if !ok {
		return false, nil
	}
	for _, input := range cfg.Inputs {
		if input.InputID == inputID {
			return true, nil
		}
	}
	return false, nil
}

// AddCollectorConfigurationInput adds a collector configuration input to the store.
func (store *Store) AddCollectorConfigurationInput(cfgID string, input *graylog.CollectorConfigurationInput) error {
	if cfgID == "" {
		return fmt.Errorf("id is required")
	}
	if input == nil {
		return fmt.Errorf("collector configuration input is nil")
	}
	if input.InputID == "" {
		input.InputID = st.NewObjectID()
	}

	store.imutex.Lock()
	defer store.imutex.Unlock()
	cfg, ok := store.collectorConfigurations[cfgID]
	if !ok {
		return fmt.Errorf("collector configuration <%s> is not found", cfgID)
	}
	cfg.Inputs = append(cfg.Inputs, *input)
	store.collectorConfigurations[cfgID] = cfg
	return nil
}

// UpdateCollectorConfigurationInput updates a collector configuration input.
func (store *Store) UpdateCollectorConfigurationInput(cfgID, inputID string, input *graylog.CollectorConfigurationInput) error {
	if cfgID == "" {
		return fmt.Errorf("collector configuration id is required")
	}
	if inputID == "" {
		return fmt.Errorf("collector configuration input id is required")
	}
	if input == nil {
		return fmt.Errorf("collector configuration input is nil")
	}
	store.imutex.Lock()
	defer store.imutex.Unlock()
	cfg, ok := store.collectorConfigurations[cfgID]
	if !ok {
		return fmt.Errorf("the collector configuration <%s> is not found", cfgID)
	}
	for i, a := range cfg.Inputs {
		if a.InputID == inputID {
			cfg.Inputs[i] = *input
			store.collectorConfigurations[cfgID] = cfg
			return nil
		}
	}
	return fmt.Errorf("collector configuration input is not found")
}

// DeleteCollectorConfigurationInput deletes a collector configuration input from the store.
func (store *Store) DeleteCollectorConfigurationInput(cfgID, inputID string) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	cfg, ok := store.collectorConfigurations[cfgID]
	if !ok {
		return fmt.Errorf("the collector configuration <%s> is not found", cfgID)
	}
	inputs := []graylog.CollectorConfigurationInput{}
	removed := false
	for _, a := range cfg.Inputs {
		if a.InputID == inputID {
			removed = true
			continue
		}
		inputs = append(inputs, a)
	}
	if !removed {
		return fmt.Errorf("the collector configuration input is not found")
	}
	cfg.Inputs = inputs
	store.collectorConfigurations[cfgID] = cfg
	return nil
}
