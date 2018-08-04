package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasStreamRule returns whether the stream sule exists.
func (lgc *Logic) HasStreamRule(streamID, streamRuleID string) (bool, error) {
	return lgc.store.HasStreamRule(streamID, streamRuleID)
}

// AddStreamRule adds a stream rule to the Server.
func (lgc *Logic) AddStreamRule(rule *graylog.StreamRule) (int, error) {
	if err := validator.CreateValidator.Struct(rule); err != nil {
		return 400, err
	}

	s, sc, err := lgc.GetStream(rule.StreamID)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "id": rule.StreamID, "sc": sc,
		}).Warn("failed to get a stream")
		return sc, err
	}
	if s.IsDefault {
		return 400, fmt.Errorf("cannot add stream rules to the default stream")
	}

	if err := lgc.store.AddStreamRule(rule); err != nil {
		return 500, err
	}
	return 201, nil
}

// UpdateStreamRule updates a stream rule of the Server.
func (lgc *Logic) UpdateStreamRule(prms *graylog.StreamRuleUpdateParams) (int, error) {
	// PUT /streams/{streamid}/rules/{streamRuleID} Update a stream rule
	if err := validator.UpdateValidator.Struct(prms); err != nil {
		return 400, err
	}
	ok, err := lgc.HasStreamRule(prms.StreamID, prms.ID)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream rule is not found: <%s>", prms.StreamID)
	}
	if err := lgc.store.UpdateStreamRule(prms); err != nil {
		return 500, err
	}
	return 204, nil
}

// DeleteStreamRule deletes a stream rule from the Server.
func (lgc *Logic) DeleteStreamRule(streamID, streamRuleID string) (int, error) {
	ok, err := lgc.HasStream(streamID)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "id": streamID,
		}).Error("lgc.HasStream() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream found with id <%s>", streamID)
	}
	ok, err = lgc.HasStreamRule(streamID, streamRuleID)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "streamID": streamID, "streamRuleID": streamRuleID,
		}).Error("lgc.HasStreamRule() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream rule found with id <%s>", streamRuleID)
	}

	if err := lgc.store.DeleteStreamRule(streamID, streamRuleID); err != nil {
		return 500, err
	}
	return 200, nil
}

// GetStreamRules returns a list of all stream rules of a given stream.
func (lgc *Logic) GetStreamRules(streamID string) ([]graylog.StreamRule, int, int, error) {
	if err := ValidateObjectID(streamID); err != nil {
		// unfortunately graylog returns not 400 but 404.
		return nil, 0, 404, err
	}
	ok, err := lgc.HasStream(streamID)
	if err != nil {
		return nil, 0, 500, err
	}
	if !ok {
		return nil, 0, 404, fmt.Errorf("no stream is not found: <%s>", streamID)
	}
	rules, total, err := lgc.store.GetStreamRules(streamID)
	if err != nil {
		return nil, 0, 500, err
	}
	return rules, total, 200, nil
}

// GetStreamRule returns a stream rule.
func (lgc *Logic) GetStreamRule(streamID, streamRuleID string) (*graylog.StreamRule, int, error) {
	ok, err := lgc.HasStream(streamID)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "id": streamID,
		}).Error("lgc.HasStream() is failure")
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("no stream found with id <%s>", streamID)
	}
	ok, err = lgc.HasStreamRule(streamID, streamRuleID)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "streamID": streamID, "streamRuleID": streamRuleID,
		}).Error("lgc.HasStreamRule() is failure")
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("no stream rule found with id <%s>", streamRuleID)
	}

	rule, err := lgc.store.GetStreamRule(streamID, streamRuleID)
	if err != nil {
		return rule, 500, err
	}
	return rule, 200, nil
}
