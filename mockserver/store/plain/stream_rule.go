package plain

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// HasStreamRule
func (store *PlainStore) HasStreamRule(streamID, streamRuleID string) (bool, error) {
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
func (store *PlainStore) GetStreamRule(streamID, streamRuleID string) (*graylog.StreamRule, error) {
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
func (store *PlainStore) GetStreamRules(id string) ([]graylog.StreamRule, int, error) {
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
func (store *PlainStore) AddStreamRule(rule *graylog.StreamRule) error {
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
func (store *PlainStore) UpdateStreamRule(rule *graylog.StreamRule) error {
	if rule == nil {
		return fmt.Errorf("rule is nil")
	}
	store.imutex.Lock()
	defer store.imutex.Unlock()
	rules, ok := store.streamRules[rule.StreamID]
	if !ok {
		return fmt.Errorf("no stream with id <%s> is found", rule.StreamID)
	}
	orig, ok := rules[rule.ID]
	if !ok {
		return fmt.Errorf("no stream rule with id <%s> is found", rule.ID)
	}
	if rule.Description == "" {
		rule.Description = orig.Description
	}
	if rule.Type == 0 {
		rule.Type = orig.Type
	}
	if rule.Inverted == nil {
		rule.Inverted = orig.Inverted
	}
	rules[rule.ID] = *rule
	store.streamRules[rule.StreamID] = rules
	return nil
}

// DeleteStreamRule deletes a stream rule.
func (store *PlainStore) DeleteStreamRule(streamID, streamRuleID string) error {
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
