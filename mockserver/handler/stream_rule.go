package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-set"
)

// HandleGetStreamRules
func HandleGetStreamRules(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /streams/{streamid}/rules Get a list of all stream rules
	streamID := ps.ByName("streamID")
	arr, total, sc, err := ms.GetStreamRules(streamID)
	if err != nil {
		return nil, sc, err
	}
	return &graylog.StreamRulesBody{StreamRules: arr, Total: total}, 200, nil
}

// HandleCreateStreamRule
func HandleCreateStreamRule(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// POST /streams/{streamid}/rules Create a stream rule
	streamID := ps.ByName("streamID")
	ok, err := ms.HasStream(streamID)
	if err != nil {
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("stream <%s> not found!", streamID)
	}

	requiredFields := set.NewStrSet("value", "field")
	allowedFields := set.NewStrSet(
		"value", "type", "description", "inverted", "field")
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if sc != 200 {
		return nil, sc, err
	}

	rule := &graylog.StreamRule{}
	if err := msDecode(body, rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as StreamRule")
		return nil, 400, err
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
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return map[string]string{"streamrule_id": rule.ID}, 201, nil
}

// null body 415 {"type": "ApiError", "message": "HTTP 415 Unsupported Media Type"}
// {} value field 400 {"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.resources.streams.rules.requests.CreateStreamRuleRequest, problem: Null value\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@162d5cc5; line: 1, column: 2]"}
// type 400 {"type": "ApiError", "message": "Unknown stream rule type 0"}
// value, type, description, inverted, field

// HandleUpdateStreamRule
func HandleUpdateStreamRule(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /streams/{streamid}/rules/{streamRuleID} Update a stream rule
	streamID := ps.ByName("streamID")
	ruleID := ps.ByName("streamRuleID")

	requiredFields := set.NewStrSet("value", "field")
	allowedFields := set.NewStrSet(
		"value", "type", "description", "inverted", "field")
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if sc != 200 {
		return nil, sc, err
	}
	rule := &graylog.StreamRule{}
	if err := msDecode(body, rule); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as StreamRule")
		return nil, 400, err
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
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return map[string]string{"streamrule_id": rule.ID}, 200, nil
}

// HandleDeleteStreamRule
func HandleDeleteStreamRule(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// DELETE /streams/{streamid}/rules/{streamRuleId} Delete a stream rule
	streamID := ps.ByName("streamID")
	id := ps.ByName("streamRuleID")
	// TODO authorization
	sc, err := ms.DeleteStreamRule(streamID, id)
	return nil, sc, err
}

// HandleGetStreamRule
func HandleGetStreamRule(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /streams/{streamid}/rules/{streamRuleId} Get a single stream rules
	// TODO authorization
	return ms.GetStreamRule(
		ps.ByName("streamID"), ps.ByName("streamRuleID"))
}
