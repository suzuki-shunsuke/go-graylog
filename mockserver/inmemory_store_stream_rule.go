package mockserver

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// HasStreamRule
func (store *InMemoryStore) HasStreamRule(streamID, streamRuleID string) (bool, error) {
	rules, ok := store.streamRules[streamID]
	if !ok {
		return false, nil
	}
	_, ok = rules[streamRuleID]
	return ok, nil
}

// GetStreamRule returns a stream rule.
func (store *InMemoryStore) GetStreamRule(streamID, streamRuleID string) (*graylog.StreamRule, error) {
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
func (store *InMemoryStore) GetStreamRules(id string) ([]graylog.StreamRule, error) {
	rules, ok := store.streamRules[id]
	if !ok {
		return nil, nil
	}
	arr := make([]graylog.StreamRule, len(rules))
	i := 0
	for _, rule := range rules {
		arr[i] = rule
		i++
	}
	return arr, nil
}

// AddStreamRule adds a stream rule.
func (store *InMemoryStore) AddStreamRule(rule *graylog.StreamRule) error {
	rules, ok := store.streamRules[rule.StreamID]
	if !ok {
		rules = map[string]graylog.StreamRule{}
	}
	rules[rule.ID] = *rule
	store.streamRules[rule.StreamID] = rules
	return nil
}

// UpdateStreamRule updates a stream rule.
func (store *InMemoryStore) UpdateStreamRule(rule *graylog.StreamRule) error {
	rules, ok := store.streamRules[rule.StreamID]
	if !ok {
		return fmt.Errorf("no stream with id <%s> is found", rule.StreamID)
	}
	if _, ok := rules[rule.ID]; !ok {
		return fmt.Errorf("no stream rule with id <%s> is found", rule.ID)
	}
	rules[rule.ID] = *rule
	store.streamRules[rule.StreamID] = rules
	return nil
}

// DeleteStreamRule deletes a stream rule.
func (store *InMemoryStore) DeleteStreamRule(streamID, streamRuleID string) error {
	rules, ok := store.streamRules[streamID]
	if !ok {
		return nil
	}
	delete(rules, streamRuleID)
	store.streamRules[streamID] = rules
	return nil
}
