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
	if rule.StreamId == "" {
		return errors.New("stream id is required")
	}
	if ms.Streams == nil {
		return fmt.Errorf("no stream is not found: %s", rule.StreamId)
	}
	if _, ok := ms.Streams[rule.StreamId]; !ok {
		return fmt.Errorf("no stream is not found: %s", rule.StreamId)
	}
	if rule.Id == "" {
		rule.Id = randStringBytesMaskImprSrc(24)
	}
	rules, ok := ms.StreamRules[rule.StreamId]
	if !ok || rules == nil {
		rules = map[string]StreamRule{}
	}
	rules[rule.Id] = *rule
	ms.StreamRules[rule.StreamId] = rules
	ms.safeSave()
	return nil
}

// StreamRuleList returns a list of all stream rules of a given stream.
func (ms *MockServer) StreamRuleList(streamId string) []StreamRule {
	if ms.StreamRules == nil {
		return []StreamRule{}
	}
	rules, ok := ms.StreamRules[streamId]
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
	streamId := ps.ByName("streamId")
	arr := ms.StreamRuleList(streamId)
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

	streamId := ps.ByName("streamId")
	if _, ok := ms.Streams[streamId]; !ok {
		writeApiError(w, 404, "Stream <%s> not found!", streamId)
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
		ms.Logger.WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as StreamRule")
		writeApiError(w, 400, "400 Bad Request")
		return
	}
	ms.Logger.WithFields(log.Fields{
		"body": string(b), "stream_rule": rule,
	}).Debug("request body")

	rule.StreamId = streamId
	if sc, msg := ms.validateCreateStreamRule(rule); sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	if err := ms.AddStreamRule(rule); err != nil {
		ms.Logger.WithFields(log.Fields{
			"error": err, "rule": rule,
		}).Error("Faield to add rule to mock server")
		write500Error(w)
		return
	}
	ret := map[string]string{"streamrule_id": rule.Id}
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
	streamId := ps.ByName("streamId")

	if _, ok := ms.Streams[streamId]; !ok {
		writeApiError(w, 404, "No stream found with id %s", streamId)
		return
	}
	ruleId := ps.ByName("streamRuleId")
	rules, ok := ms.StreamRules[streamId]
	if !ok || rules == nil {
		writeApiError(w, 404, "No StreamRule found with id %s", ruleId)
		return
	}
	rule, ok := rules[ruleId]
	if !ok {
		writeApiError(w, 404, "No StreamRule found with id %s", ruleId)
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
		ms.Logger.WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as StreamRule")
		writeApiError(w, 400, "400 Bad Request")
		return
	}
	ms.Logger.WithFields(log.Fields{
		"body": string(b), "stream_rule": rule,
	}).Debug("request body")

	rule.StreamId = streamId
	rule.Id = ruleId
	if err := UpdateValidator.Struct(&rule); err != nil {
		writeApiError(w, 400, err.Error())
		return
	}
	if err := ms.AddStreamRule(&rule); err != nil {
		ms.Logger.WithFields(log.Fields{
			"error": err, "rule": &rule,
		}).Error("Faield to add rule to mock server")
		write500Error(w)
	}
	ret := map[string]string{"streamrule_id": rule.Id}
	writeOr500Error(w, ret)
}
