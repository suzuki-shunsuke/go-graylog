package mockserver

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasStreamRule
func (ms *MockServer) HasStreamRule(streamID, streamRuleID string) (bool, error) {
	return ms.HasStreamRule(streamID, streamRuleID)
}

// AddStreamRule adds a stream rule to the MockServer.
func (ms *MockServer) AddStreamRule(rule *graylog.StreamRule) (int, error) {
	if err := validator.CreateValidator.Struct(rule); err != nil {
		return 400, err
	}
	ok, err := ms.HasStream(rule.StreamID)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream is not found: %s", rule.StreamID)
	}
	rule.ID = randStringBytesMaskImprSrc(24)
	if err = ms.store.AddStreamRule(rule); err != nil {
		return 500, err
	}
	return 201, nil
}

// UpdateStreamRule updates a stream rule of the MockServer.
func (ms *MockServer) UpdateStreamRule(rule *graylog.StreamRule) (int, error) {
	// PUT /streams/{streamid}/rules/{streamRuleID} Update a stream rule
	if err := validator.UpdateValidator.Struct(rule); err != nil {
		return 400, err
	}
	ok, err := ms.HasStream(rule.StreamID)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream is not found: %s", rule.StreamID)
	}
	r, err := ms.store.GetStreamRule(rule.StreamID, rule.ID)
	if err != nil {
		return 500, err
	}
	if r == nil {
		return 404, fmt.Errorf("no stream rule is not found: %s", rule.ID)
	}
	if err := ms.store.UpdateStreamRule(rule); err != nil {
		return 500, err
	}
	return 200, nil
}

// DeleteStreamRule deletes a stream rule from the MockServer.
func (ms *MockServer) DeleteStreamRule(streamID, streamRuleID string) (int, error) {
	ok, err := ms.HasStream(streamID)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": streamID,
		}).Error("ms.HasStream() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("No stream found with id %s", streamID)
	}
	ok, err = ms.HasStreamRule(streamID, streamRuleID)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "streamID": streamID, "streamRuleID": streamRuleID,
		}).Error("ms.HasStreamRule() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("No stream rule found with id %s", streamRuleID)
	}

	if err := ms.store.DeleteStreamRule(streamID, streamRuleID); err != nil {
		return 500, err
	}
	return 200, nil
}

// StreamRuleList returns a list of all stream rules of a given stream.
func (ms *MockServer) StreamRuleList(streamID string) ([]graylog.StreamRule, int, error) {
	ok, err := ms.HasStream(streamID)
	if err != nil {
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("no stream is not found: %s", streamID)
	}
	rules, err := ms.store.GetStreamRules(streamID)
	if err != nil {
		return nil, 500, err
	}
	return rules, 200, nil
}
