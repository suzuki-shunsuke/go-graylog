package graylog

// GET /streams/{streamid}/rules Get a list of all stream rules

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// AddStreamRule adds a stream rule to the MockServer.
func (ms *MockServer) AddStreamRule(rule *StreamRule) error {
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
		rules = map[string]StreamRule{}
	}
	rules[rule.ID] = *rule
	ms.streamRules[rule.StreamID] = rules
	return nil
}

// StreamRuleList returns a list of all stream rules of a given stream.
func (ms *MockServer) StreamRuleList(streamID string) []StreamRule {
	rules, ok := ms.streamRules[streamID]
	if !ok || rules == nil {
		return []StreamRule{}
	}
	arr := make([]StreamRule, len(rules))
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
) {
	ms.handleInit(w, r, false)
	streamID := ps.ByName("streamID")
	arr := ms.StreamRuleList(streamID)
	body := &streamRulesBody{StreamRules: arr, Total: len(arr)}
	writeOr500Error(w, body)
}

func (mc *MockServer) validateCreateStreamRule(rule *StreamRule) (int, []byte) {
	if err := CreateValidator.Struct(rule); err != nil {
		return 400, []byte(fmt.Sprintf(`{"type": "ApiError", "message": "%s"}`, err.Error()))
	}
	return 200, nil
}

// POST /streams/{streamid}/rules Create a stream rule
func (ms *MockServer) handleCreateStreamRule(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		write500Error(w)
		return
	}

	streamID := ps.ByName("streamID")
	ok, err := ms.HasStream(streamID)
	if err != nil {
		writeApiError(w, 500, err.Error())
		return
	}
	if !ok {
		writeApiError(w, 404, "Stream <%s> not found!", streamID)
		return
	}

	requiredFields := []string{"value", "field"}
	allowedFields := []string{
		"value", "type", "description", "inverted", "field"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	rule := &StreamRule{}
	if err := msDecode(body, rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as StreamRule")
		writeApiError(w, 400, "400 Bad Request")
		return
	}
	ms.Logger().WithFields(log.Fields{
		"body": string(b), "stream_rule": rule,
	}).Debug("request body")

	rule.StreamID = streamID
	if sc, msg := ms.validateCreateStreamRule(rule); sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	if err := ms.AddStreamRule(rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "rule": rule,
		}).Error("Faield to add rule to mock server")
		write500Error(w)
		return
	}
	ms.safeSave()
	ret := map[string]string{"streamrule_id": rule.ID}
	writeOr500Error(w, ret)
}

// null body 415 {"type": "ApiError", "message": "HTTP 415 Unsupported Media Type"}
// {} value field 400 {"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.resources.streams.rules.requests.CreateStreamRuleRequest, problem: Null value\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@162d5cc5; line: 1, column: 2]"}
// type 400 {"type": "ApiError", "message": "Unknown stream rule type 0"}
// value, type, description, inverted, field

func (ms *MockServer) handleUpdateStreamRule(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		write500Error(w)
		return
	}
	streamID := ps.ByName("streamID")

	ok, err := ms.HasStream(streamID)
	if err != nil {
		writeApiError(w, 500, err.Error())
		return
	}
	if !ok {
		writeApiError(w, 404, "Stream <%s> not found!", streamID)
		return
	}

	ruleID := ps.ByName("streamRuleID")
	rules, ok := ms.streamRules[streamID]
	if !ok || rules == nil {
		writeApiError(w, 404, "No StreamRule found with id %s", ruleID)
		return
	}
	rule, ok := rules[ruleID]
	if !ok {
		writeApiError(w, 404, "No StreamRule found with id %s", ruleID)
		return
	}

	requiredFields := []string{"value", "field"}
	allowedFields := []string{
		"value", "type", "description", "inverted", "field"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	rule = StreamRule{}
	if err := msDecode(body, &rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as StreamRule")
		writeApiError(w, 400, "400 Bad Request")
		return
	}
	ms.Logger().WithFields(log.Fields{
		"body": string(b), "stream_rule": rule,
	}).Debug("request body")

	rule.StreamID = streamID
	rule.ID = ruleID
	if err := UpdateValidator.Struct(&rule); err != nil {
		writeApiError(w, 400, err.Error())
		return
	}
	if err := ms.AddStreamRule(&rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "rule": &rule,
		}).Error("Faield to add rule to mock server")
		write500Error(w)
	}
	ms.safeSave()
	ret := map[string]string{"streamrule_id": rule.ID}
	writeOr500Error(w, ret)
}
