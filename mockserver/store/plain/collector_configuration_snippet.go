package plain

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// HasCollectorConfigurationSnippet returns whether the collector configuration snippet exists.
func (store *Store) HasCollectorConfigurationSnippet(cfgID, snippetID string) (bool, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	cfg, ok := store.collectorConfigurations[cfgID]
	if !ok {
		return false, nil
	}
	for _, snippet := range cfg.Snippets {
		if snippet.SnippetID == snippetID {
			return true, nil
		}
	}
	return false, nil
}

// AddCollectorConfigurationSnippet adds a collector configuration snippet to the store.
func (store *Store) AddCollectorConfigurationSnippet(cfgID string, snippet *graylog.CollectorConfigurationSnippet) error {
	if cfgID == "" {
		return fmt.Errorf("id is required")
	}
	if snippet == nil {
		return fmt.Errorf("collector configuration snippet is nil")
	}
	if snippet.SnippetID == "" {
		snippet.SnippetID = st.NewObjectID()
	}

	store.imutex.Lock()
	defer store.imutex.Unlock()
	cfg, ok := store.collectorConfigurations[cfgID]
	if !ok {
		return fmt.Errorf("collector configuration <%s> is not found", cfgID)
	}
	cfg.Snippets = append(cfg.Snippets, *snippet)
	store.collectorConfigurations[cfgID] = cfg
	return nil
}

// UpdateCollectorConfigurationSnippet updates a collector configuration snippet.
func (store *Store) UpdateCollectorConfigurationSnippet(cfgID, snippetID string, snippet *graylog.CollectorConfigurationSnippet) error {
	if cfgID == "" {
		return fmt.Errorf("collector configuration id is required")
	}
	if snippetID == "" {
		return fmt.Errorf("collector configuration snippet id is required")
	}
	if snippet == nil {
		return fmt.Errorf("collector configuration snippet is nil")
	}
	store.imutex.Lock()
	defer store.imutex.Unlock()
	cfg, ok := store.collectorConfigurations[cfgID]
	if !ok {
		return fmt.Errorf("the collector configuration <%s> is not found", cfgID)
	}
	for i, a := range cfg.Snippets {
		if a.SnippetID == snippetID {
			cfg.Snippets[i] = *snippet
			store.collectorConfigurations[cfgID] = cfg
			return nil
		}
	}
	return fmt.Errorf("collector configuration snippet is not found")
}

// DeleteCollectorConfigurationSnippet deletes a collector configuration snippet from the store.
func (store *Store) DeleteCollectorConfigurationSnippet(cfgID, snippetID string) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	cfg, ok := store.collectorConfigurations[cfgID]
	if !ok {
		return fmt.Errorf("the collector configuration <%s> is not found", cfgID)
	}
	snippets := []graylog.CollectorConfigurationSnippet{}
	removed := false
	for _, a := range cfg.Snippets {
		if a.SnippetID == snippetID {
			removed = true
			continue
		}
		snippets = append(snippets, a)
	}
	if !removed {
		return fmt.Errorf("the collector configuration snippet is not found")
	}
	cfg.Snippets = snippets
	store.collectorConfigurations[cfgID] = cfg
	return nil
}
