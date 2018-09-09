package plain

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// HasCollectorConfigurationOutput returns whether the collector configuration output exists.
func (store *Store) HasCollectorConfigurationOutput(cfgID, outputID string) (bool, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	cfg, ok := store.collectorConfigurations[cfgID]
	if !ok {
		return false, nil
	}
	for _, output := range cfg.Outputs {
		if output.OutputID == outputID {
			return true, nil
		}
	}
	return false, nil
}

// AddCollectorConfigurationOutput adds a collector configuration output to the store.
func (store *Store) AddCollectorConfigurationOutput(cfgID string, output *graylog.CollectorConfigurationOutput) error {
	if cfgID == "" {
		return fmt.Errorf("id is required")
	}
	if output == nil {
		return fmt.Errorf("collector configuration output is nil")
	}
	if output.OutputID == "" {
		output.OutputID = st.NewObjectID()
	}

	store.imutex.Lock()
	defer store.imutex.Unlock()
	cfg, ok := store.collectorConfigurations[cfgID]
	if !ok {
		return fmt.Errorf("collector configuration <%s> is not found", cfgID)
	}
	cfg.Outputs = append(cfg.Outputs, *output)
	store.collectorConfigurations[cfgID] = cfg
	return nil
}

// UpdateCollectorConfigurationOutput updates a collector configuration output.
func (store *Store) UpdateCollectorConfigurationOutput(cfgID, outputID string, output *graylog.CollectorConfigurationOutput) error {
	if cfgID == "" {
		return fmt.Errorf("collector configuration id is required")
	}
	if outputID == "" {
		return fmt.Errorf("collector configuration output id is required")
	}
	if output == nil {
		return fmt.Errorf("collector configuration output is nil")
	}
	store.imutex.Lock()
	defer store.imutex.Unlock()
	cfg, ok := store.collectorConfigurations[cfgID]
	if !ok {
		return fmt.Errorf("the collector configuration <%s> is not found", cfgID)
	}
	for i, a := range cfg.Outputs {
		if a.OutputID == outputID {
			cfg.Outputs[i] = *output
			store.collectorConfigurations[cfgID] = cfg
			return nil
		}
	}
	return fmt.Errorf("collector configuration output is not found")
}

// DeleteCollectorConfigurationOutput deletes a collector configuration output from the store.
func (store *Store) DeleteCollectorConfigurationOutput(cfgID, outputID string) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	cfg, ok := store.collectorConfigurations[cfgID]
	if !ok {
		return fmt.Errorf("the collector configuration <%s> is not found", cfgID)
	}
	outputs := []graylog.CollectorConfigurationOutput{}
	removed := false
	for _, a := range cfg.Outputs {
		if a.OutputID == outputID {
			removed = true
			continue
		}
		outputs = append(outputs, a)
	}
	if !removed {
		return fmt.Errorf("the collector configuration output is not found")
	}
	cfg.Outputs = outputs
	store.collectorConfigurations[cfgID] = cfg
	return nil
}
