package mockserver

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
)

func (ms *MockServer) handleGetStreamRules(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	// GET /streams/{streamid}/rules Get a list of all stream rules
	streamID := ps.ByName("streamID")
	arr, sc, err := ms.GetStreamRules(streamID)
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
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if sc != 200 {
		return sc, nil, err
	}

	rule := &graylog.StreamRule{}
	if err := msDecode(body, rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as StreamRule")
		return 400, nil, err
	}
	ms.Logger().WithFields(log.Fields{
		"body": body, "stream_rule": rule,
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
	streamID := ps.ByName("streamID")
	ruleID := ps.ByName("streamRuleID")

	requiredFields := []string{"value", "field"}
	allowedFields := []string{
		"value", "type", "description", "inverted", "field"}
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if sc != 200 {
		return sc, nil, err
	}
	rule := &graylog.StreamRule{}
	if err := msDecode(body, rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as StreamRule")
		return 400, nil, err
	}
	ms.Logger().WithFields(log.Fields{
		"body": body, "stream_rule": rule,
	}).Debug("request body")

	rule.StreamID = streamID
	rule.ID = ruleID
	if sc, err := ms.UpdateStreamRule(rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "rule": &rule,
		}).Error("Faield to update stream rule")
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

func (ms *MockServer) handleGetStreamRule(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	// GET /streams/{streamid}/rules/{streamRuleId} Get a single stream rules
	rule, sc, err := ms.GetStreamRule(
		ps.ByName("streamID"), ps.ByName("streamRuleID"))
	return sc, rule, err
}
