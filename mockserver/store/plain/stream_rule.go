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
	stream, ok := store.streams[streamID]
	if !ok {
		return false, nil
	}
	for _, rule := range stream.Rules {
		if rule.ID == streamRuleID {
			return true, nil
		}
	}
	return false, nil
}

// GetStreamRule returns a stream rule.
func (store *Store) GetStreamRule(streamID, streamRuleID string) (*graylog.StreamRule, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	stream, ok := store.streams[streamID]
	if !ok {
		return nil, nil
	}
	for _, rule := range stream.Rules {
		if rule.ID == streamRuleID {
			return &rule, nil
		}
	}
	return nil, nil
}

// GetStreamRules returns stream rules of the given stream.
func (store *Store) GetStreamRules(id string) ([]graylog.StreamRule, int, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	stream, ok := store.streams[id]
	if !ok {
		return nil, 0, nil
	}
	return stream.Rules, len(stream.Rules), nil
}

// AddStreamRule adds a stream rule.
func (store *Store) AddStreamRule(rule *graylog.StreamRule) error {
	if rule == nil {
		return fmt.Errorf("rule is nil")
	}
	streamID := rule.StreamID
	if streamID == "" {
		return fmt.Errorf("stream id is empty")
	}

	store.imutex.Lock()
	defer store.imutex.Unlock()

	stream, ok := store.streams[streamID]
	if !ok {
		return fmt.Errorf(`stream "%s" is not found`, streamID)
	}

	if rule.ID == "" {
		rule.ID = st.NewObjectID()
	}
	stream.Rules = append(stream.Rules, *rule)
	store.streams[streamID] = stream
	return nil
}

// UpdateStreamRule updates a stream rule.
func (store *Store) UpdateStreamRule(prms *graylog.StreamRuleUpdateParams) error {
	if prms == nil {
		return fmt.Errorf("rule is nil")
	}
	streamID := prms.StreamID
	if streamID == "" {
		return fmt.Errorf("stream id is empty")
	}
	id := prms.ID
	if id == "" {
		return fmt.Errorf("stream rule id is empty")
	}

	store.imutex.Lock()
	defer store.imutex.Unlock()

	stream, ok := store.streams[streamID]
	if !ok {
		return fmt.Errorf(`stream "%s" is not found`, streamID)
	}

	for i, rule := range stream.Rules {
		if rule.ID != id {
			continue
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
		stream.Rules[i] = rule
		return nil
	}
	return fmt.Errorf("no stream rule with id <%s> is found", id)
}

// DeleteStreamRule deletes a stream rule.
func (store *Store) DeleteStreamRule(streamID, streamRuleID string) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()

	if streamID == "" {
		return fmt.Errorf("failed to delete a stream rule: stream id is empty")
	}
	if streamRuleID == "" {
		return fmt.Errorf("failed to delete a stream rule: stream rule id is empty")
	}

	stream, ok := store.streams[streamID]
	if !ok {
		return fmt.Errorf(`stream "%s" is not found`, streamID)
	}

	rules := []graylog.StreamRule{}
	removed := false
	for _, rule := range stream.Rules {
		if rule.ID == streamRuleID {
			removed = true
			continue
		}
		rules = append(rules, rule)
	}
	if !removed {
		return fmt.Errorf("no stream rule with id <%s> is found", streamRuleID)
	}
	stream.Rules = rules
	store.streams[streamID] = stream
	return nil
}
