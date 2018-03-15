package mockserver

// GET /streams/{streamid}/rules Get a list of all stream rules

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
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

func (ms *MockServer) handleGetStreamRules(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	// GET /streams/{streamid}/rules Get a list of all stream rules
	ms.handleInit(w, r, false)
	streamID := ps.ByName("streamID")
	arr, sc, err := ms.StreamRuleList(streamID)
	if err != nil {
		return sc, nil, err
	}
	body := &graylog.StreamRulesBody{StreamRules: arr, Total: len(arr)}
	return 200, body, nil
}

func (ms *MockServer) handleCreateStreamRule(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	// POST /streams/{streamid}/rules Create a stream rule
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		return 500, nil, err
	}

	streamID := ps.ByName("streamID")
	ok, err := ms.HasStream(streamID)
	if err != nil {
		return 500, nil, err
	}
	if !ok {
		return 404, nil, fmt.Errorf("Stream <%s> not found!", streamID)
	}

	requiredFields := []string{"value", "field"}
	allowedFields := []string{
		"value", "type", "description", "inverted", "field"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		return sc, nil, fmt.Errorf(msg)
	}

	rule := &graylog.StreamRule{}
	if err := msDecode(body, rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as StreamRule")
		return 400, nil, err
	}
	ms.Logger().WithFields(log.Fields{
		"body": string(b), "stream_rule": rule,
	}).Debug("request body")

	rule.StreamID = streamID
	sc, err = ms.AddStreamRule(rule)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "rule": rule,
		}).Error("Faield to add rule to mock server")
		return sc, nil, err
	}
	ms.safeSave()
	ret := map[string]string{"streamrule_id": rule.ID}
	return 201, ret, nil
}

// null body 415 {"type": "ApiError", "message": "HTTP 415 Unsupported Media Type"}
// {} value field 400 {"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.resources.streams.rules.requests.CreateStreamRuleRequest, problem: Null value\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@162d5cc5; line: 1, column: 2]"}
// type 400 {"type": "ApiError", "message": "Unknown stream rule type 0"}
// value, type, description, inverted, field

func (ms *MockServer) handleUpdateStreamRule(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	// PUT /streams/{streamid}/rules/{streamRuleID} Update a stream rule
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		return 500, nil, err
	}
	streamID := ps.ByName("streamID")

	ok, err := ms.HasStream(streamID)
	if err != nil {
		return 500, nil, err
	}
	if !ok {
		return 404, nil, fmt.Errorf("Stream <%s> not found!", streamID)
	}

	ruleID := ps.ByName("streamRuleID")
	rules, ok := ms.streamRules[streamID]
	if !ok || rules == nil {
		return 404, nil, fmt.Errorf("No StreamRule found with id %s", ruleID)
	}
	rule, ok := rules[ruleID]
	if !ok {
		return 404, nil, fmt.Errorf("No StreamRule found with id %s", ruleID)
	}

	requiredFields := []string{"value", "field"}
	allowedFields := []string{
		"value", "type", "description", "inverted", "field"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		return sc, nil, fmt.Errorf(msg)
	}

	rule = graylog.StreamRule{}
	if err := msDecode(body, &rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as StreamRule")
		return 400, nil, err
	}
	ms.Logger().WithFields(log.Fields{
		"body": string(b), "stream_rule": rule,
	}).Debug("request body")

	rule.StreamID = streamID
	rule.ID = ruleID
	if err := validator.UpdateValidator.Struct(&rule); err != nil {
		return 400, nil, err
	}
	if sc, err := ms.UpdateStreamRule(&rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "rule": &rule,
		}).Error("Faield to add rule to mock server")
		return sc, nil, err
	}
	ms.safeSave()
	ret := map[string]string{"streamrule_id": rule.ID}
	return 200, ret, nil
}

func (ms *MockServer) handleDeleteStreamRule(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	// DELETE /streams/{streamid}/rules/{streamRuleId} Delete a stream rule
	ms.handleInit(w, r, false)
	streamID := ps.ByName("streamID")
	id := ps.ByName("streamRuleID")
	sc, err := ms.DeleteStreamRule(streamID, id)
	return sc, nil, err
	ok, err := ms.HasStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.HasStream() is failure")
		return 500, nil, err
	}

	if !ok {
		return 404, nil, fmt.Errorf("No stream found with id %s", id)
	}
	ms.DeleteStream(id)
	ms.safeSave()
	return 204, nil, nil
}
