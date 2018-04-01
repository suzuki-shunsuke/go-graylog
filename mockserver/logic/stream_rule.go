package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasStreamRule
func (ms *Logic) HasStreamRule(streamID, streamRuleID string) (bool, error) {
	return ms.store.HasStreamRule(streamID, streamRuleID)
}

// AddStreamRule adds a stream rule to the Server.
func (ms *Logic) AddStreamRule(rule *graylog.StreamRule) (int, error) {
	if err := validator.CreateValidator.Struct(rule); err != nil {
		return 400, err
	}
	ok, err := ms.HasStream(rule.StreamID)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream is not found: <%s>", rule.StreamID)
	}
	if err := ms.store.AddStreamRule(rule); err != nil {
		return 500, err
	}
	return 201, nil
}

// UpdateStreamRule updates a stream rule of the Server.
func (ms *Logic) UpdateStreamRule(rule *graylog.StreamRule) (int, error) {
	// PUT /streams/{streamid}/rules/{streamRuleID} Update a stream rule
	if err := validator.UpdateValidator.Struct(rule); err != nil {
		return 400, err
	}
	ok, err := ms.HasStreamRule(rule.StreamID, rule.ID)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream rule is not found: <%s>", rule.StreamID)
	}
	if err := ms.store.UpdateStreamRule(rule); err != nil {
		return 500, err
	}
	return 204, nil
}

// DeleteStreamRule deletes a stream rule from the Server.
func (ms *Logic) DeleteStreamRule(streamID, streamRuleID string) (int, error) {
	ok, err := ms.HasStream(streamID)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": streamID,
		}).Error("ms.HasStream() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream found with id <%s>", streamID)
	}
	ok, err = ms.HasStreamRule(streamID, streamRuleID)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "streamID": streamID, "streamRuleID": streamRuleID,
		}).Error("ms.HasStreamRule() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream rule found with id <%s>", streamRuleID)
	}

	if err := ms.store.DeleteStreamRule(streamID, streamRuleID); err != nil {
		return 500, err
	}
	return 200, nil
}

// GetStreamRules returns a list of all stream rules of a given stream.
func (ms *Logic) GetStreamRules(streamID string) ([]graylog.StreamRule, int, error) {
	ok, err := ms.HasStream(streamID)
	if err != nil {
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("no stream is not found: <%s>", streamID)
	}
	rules, err := ms.store.GetStreamRules(streamID)
	if err != nil {
		return nil, 500, err
	}
	return rules, 200, nil
}

// GetStreamRule returns a stream rule.
func (ms *Logic) GetStreamRule(streamID, streamRuleID string) (*graylog.StreamRule, int, error) {
	ok, err := ms.HasStream(streamID)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": streamID,
		}).Error("ms.HasStream() is failure")
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("no stream found with id <%s>", streamID)
	}
	ok, err = ms.HasStreamRule(streamID, streamRuleID)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "streamID": streamID, "streamRuleID": streamRuleID,
		}).Error("ms.HasStreamRule() is failure")
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("no stream rule found with id <%s>", streamRuleID)
	}

	rule, err := ms.store.GetStreamRule(streamID, streamRuleID)
	if err != nil {
		return rule, 500, err
	}
	return rule, 200, nil
}
