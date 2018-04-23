package plain

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// HasStreamRule returns whether the stream rule exists.
func (store *Store) HasStreamRule(streamID, streamRuleID string) (bool, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	rules, ok := store.streamRules[streamID]
	if !ok {
		return false, nil
	}
	_, ok = rules[streamRuleID]
	return ok, nil
}

// GetStreamRule returns a stream rule.
func (store *Store) GetStreamRule(streamID, streamRuleID string) (*graylog.StreamRule, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	rules, ok := store.streamRules[streamID]
	if !ok {
		return nil, nil
	}
	rule, ok := rules[streamRuleID]
	if ok {
		return &rule, nil
	}
	return nil, nil
}

// GetStreamRules returns stream rules of the given stream.
func (store *Store) GetStreamRules(id string) ([]graylog.StreamRule, int, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	rules, ok := store.streamRules[id]
	if !ok {
		return nil, 0, nil
	}
	size := len(rules)
	arr := make([]graylog.StreamRule, size)
	i := 0
	for _, rule := range rules {
		arr[i] = rule
		i++
	}
	return arr, size, nil
}

// AddStreamRule adds a stream rule.
func (store *Store) AddStreamRule(rule *graylog.StreamRule) error {
	if rule == nil {
		return fmt.Errorf("rule is nil")
	}
	store.imutex.Lock()
	defer store.imutex.Unlock()
	rules, ok := store.streamRules[rule.StreamID]
	if !ok {
		rules = map[string]graylog.StreamRule{}
	}
	if rule.ID == "" {
		rule.ID = st.NewObjectID()
	}
	rules[rule.ID] = *rule
	store.streamRules[rule.StreamID] = rules
	return nil
}

// UpdateStreamRule updates a stream rule.
func (store *Store) UpdateStreamRule(prms *graylog.StreamRuleUpdateParams) error {
	if prms == nil {
		return fmt.Errorf("rule is nil")
	}
	store.imutex.Lock()
	defer store.imutex.Unlock()
	rules, ok := store.streamRules[prms.StreamID]
	if !ok {
		return fmt.Errorf("no stream with id <%s> is found", prms.StreamID)
	}
	rule, ok := rules[prms.ID]
	if !ok {
		return fmt.Errorf("no stream rule with id <%s> is found", prms.ID)
	}
	if prms.Field != "" {
		rule.Field = prms.Field
	}
	if prms.Description != "" {
		rule.Description = prms.Description
	}
	if prms.Value != "" {
		rule.Value = prms.Value
	}
	if prms.Type != nil {
		rule.Type = *prms.Type
	}
	if prms.Inverted != nil {
		rule.Inverted = *prms.Inverted
	}
	rules[rule.ID] = rule
	store.streamRules[rule.StreamID] = rules
	return nil
}

// DeleteStreamRule deletes a stream rule.
func (store *Store) DeleteStreamRule(streamID, streamRuleID string) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	rules, ok := store.streamRules[streamID]
	if !ok {
		return nil
	}
	delete(rules, streamRuleID)
	store.streamRules[streamID] = rules
	return nil
}
