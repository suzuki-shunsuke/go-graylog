package graylog

// GET /streams/{streamid}/rules Get a list of all stream rules

import (
	"encoding/json"
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
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	streamId := ps.ByName("streamId")
	arr := ms.StreamRuleList(streamId)
	body := &streamRulesBody{StreamRules: arr, Total: len(arr)}
	writeOr500Error(w, body)
}

func (mc *MockServer) validateCreateStreamRule(rule *StreamRule) (int, []byte) {
	key := ""
	// value, type, description, inverted, field
	switch {
	case rule.Id != "":
		key = "id"
	case rule.StreamId != "":
		key = "stream_id"
	}
	if key != "" {
		return 400, []byte(fmt.Sprintf(`{"type": "ApiError", "message": "Unable to map property %s.\nKnown properties include: value, type, description, inverted, field"}`, key))
	}
	if rule.Field == "" {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.resources.streams.rules.requests.CreateStreamRuleRequest, problem: Null field\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@2bf42a23; line: 3, column: 5]"}`)
	}
	if rule.Value == "" {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.resources.streams.rules.requests.CreateStreamRuleRequest, problem: Null value\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@a73d673; line: 3, column: 5]"}`)
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
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "Stream <%s> not found!"}`, streamId)))
		return
	}

	rule := &StreamRule{}
	if err := json.Unmarshal(b, rule); err != nil {
		ms.Logger.WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as StreamRule")
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	ms.Logger.WithFields(log.Fields{
		"body": string(b), "stream_rule": rule,
	}).Debug("request body")
	if sc, msg := ms.validateCreateStreamRule(rule); sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	rule.StreamId = streamId
	if err := ms.AddStreamRule(rule); err != nil {
		ms.Logger.WithFields(log.Fields{
			"error": err, "rule": rule,
		}).Error("Faield to add rule to mock server")
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
	}
	ret := map[string]string{"streamrule_id": rule.Id}
	writeOr500Error(w, ret)
}
