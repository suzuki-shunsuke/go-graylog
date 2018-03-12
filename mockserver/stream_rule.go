package mockserver

// GET /streams/{streamid}/rules Get a list of all stream rules

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// AddStreamRule adds a stream rule to the MockServer.
func (ms *MockServer) AddStreamRule(rule *graylog.StreamRule) error {
	if rule.StreamID == "" {
		return errors.New("stream id is required")
	}
	ok, err := ms.HasStream(rule.StreamID)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("no stream is not found: %s", rule.StreamID)
	}
	if rule.ID == "" {
		rule.ID = randStringBytesMaskImprSrc(24)
	}
	rules, ok := ms.streamRules[rule.StreamID]
	if !ok || rules == nil {
		rules = map[string]graylog.StreamRule{}
	}
	rules[rule.ID] = *rule
	ms.streamRules[rule.StreamID] = rules
	return nil
}

// StreamRuleList returns a list of all stream rules of a given stream.
func (ms *MockServer) StreamRuleList(streamID string) []graylog.StreamRule {
	rules, ok := ms.streamRules[streamID]
	if !ok || rules == nil {
		return []graylog.StreamRule{}
	}
	arr := make([]graylog.StreamRule, len(rules))
	i := 0
	for _, rule := range rules {
		arr[i] = rule
		i++
	}
	return arr
}

// GET /streams Get a list of all streams
func (ms *MockServer) handleGetStreamRules(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	ms.handleInit(w, r, false)
	streamID := ps.ByName("streamID")
	arr := ms.StreamRuleList(streamID)
	body := &graylog.StreamRulesBody{StreamRules: arr, Total: len(arr)}
	return 200, body, nil
}

// POST /streams/{streamid}/rules Create a stream rule
func (ms *MockServer) handleCreateStreamRule(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
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
	if err := validator.CreateValidator.Struct(rule); err != nil {
		return 400, nil, err
	}
	if err := ms.AddStreamRule(rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "rule": rule,
		}).Error("Faield to add rule to mock server")
		return 500, nil, err
	}
	ms.safeSave()
	ret := map[string]string{"streamrule_id": rule.ID}
	return 200, ret, nil
}

// null body 415 {"type": "ApiError", "message": "HTTP 415 Unsupported Media Type"}
// {} value field 400 {"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.resources.streams.rules.requests.CreateStreamRuleRequest, problem: Null value\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@162d5cc5; line: 1, column: 2]"}
// type 400 {"type": "ApiError", "message": "Unknown stream rule type 0"}
// value, type, description, inverted, field

func (ms *MockServer) handleUpdateStreamRule(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
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
	if err := ms.AddStreamRule(&rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "rule": &rule,
		}).Error("Faield to add rule to mock server")
		return 500, nil, err
	}
	ms.safeSave()
	ret := map[string]string{"streamrule_id": rule.ID}
	return 200, ret, nil
}
