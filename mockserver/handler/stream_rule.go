package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/util"
	"github.com/suzuki-shunsuke/go-set"
)

// HandleGetStreamRules is the handler of Get Stream Rules API.
func HandleGetStreamRules(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /streams/{streamid}/rules Get a list of all stream rules
	streamID := ps.ByName("streamID")
	arr, total, sc, err := lgc.GetStreamRules(streamID)
	if err != nil {
		return nil, sc, err
	}
	return &graylog.StreamRulesBody{StreamRules: arr, Total: total}, 200, nil
}

// HandleGetStreamRule is the handler of Get a Stream Rule API.
func HandleGetStreamRule(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /streams/{streamid}/rules/{streamRuleId} Get a single stream rules
	// TODO authorization
	return lgc.GetStreamRule(
		ps.ByName("streamID"), ps.ByName("streamRuleID"))
}

// HandleCreateStreamRule is the handler of Create a Stream Rule API.
func HandleCreateStreamRule(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// POST /streams/{streamid}/rules Create a stream rule
	streamID := ps.ByName("streamID")
	ok, err := lgc.HasStream(streamID)
	if err != nil {
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("stream <%s> not found", streamID)
	}

	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("value", "field"),
			Optional:     set.NewStrSet("type", "description", "inverted"),
			ExtForbidden: true,
		})
	if sc != 200 {
		return nil, sc, err
	}

	rule := &graylog.StreamRule{}
	if err := util.MSDecode(body, rule); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as StreamRule")
		return nil, 400, err
	}
	lgc.Logger().WithFields(log.Fields{
		"body": body, "stream_rule": rule,
	}).Debug("request body")

	rule.StreamID = streamID
	sc, err = lgc.AddStreamRule(rule)
	if err != nil {
		logic.LogWE(sc, lgc.Logger().WithFields(log.Fields{
			"error": err, "rule": rule, "status_code": sc,
		}), "Faield to add rule to mock server")
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return map[string]string{"streamrule_id": rule.ID}, 201, nil
}

// type 400 {"type": "ApiError", "message": "Unknown stream rule type 0"}

// HandleUpdateStreamRule is the handler of Update a Stream Rule API.
func HandleUpdateStreamRule(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /streams/{streamid}/rules/{streamRuleID} Update a stream rule
	streamID := ps.ByName("streamID")
	ruleID := ps.ByName("streamRuleID")

	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("value", "field"),
			Optional:     set.NewStrSet("type", "description", "inverted"),
			ExtForbidden: true,
		})
	if sc != 200 {
		return nil, sc, err
	}
	prms := &graylog.StreamRuleUpdateParams{}
	if err := util.MSDecode(body, prms); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as StreamRuleUpdateParams")
		return nil, 400, err
	}
	lgc.Logger().WithFields(log.Fields{
		"body": body, "stream_rule": prms,
	}).Debug("request body")

	prms.StreamID = streamID
	prms.ID = ruleID
	if sc, err := lgc.UpdateStreamRule(prms); err != nil {
		logic.LogWE(sc, lgc.Logger().WithFields(log.Fields{
			"error": err, "rule": &prms, "status_code": sc,
		}), "faield to update stream rule")
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return map[string]string{"streamrule_id": prms.ID}, 200, nil
}

// HandleDeleteStreamRule is the handler of Delete a Stream Rule API.
func HandleDeleteStreamRule(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// DELETE /streams/{streamid}/rules/{streamRuleId} Delete a stream rule
	streamID := ps.ByName("streamID")
	id := ps.ByName("streamRuleID")
	// TODO authorization
	sc, err := lgc.DeleteStreamRule(streamID, id)
	return nil, sc, err
}
