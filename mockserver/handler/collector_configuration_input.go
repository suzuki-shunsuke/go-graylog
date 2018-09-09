package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/util"
)

// HandleCreateCollectorConfigurationInput is the handler of Create a CollectorConfiguration Input API.
func HandleCreateCollectorConfigurationInput(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// TODO authorize
	cfgID := ps.PathParam("collectorConfigurationInputID")
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("tags", "inputs", "outputs", "snippets"),
			Optional:     nil,
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}
	input := &graylog.CollectorConfigurationInput{}
	if err := util.MSDecode(body, input); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as CollectorConfigurationInput")
		return nil, 400, err
	}
	sc, err = lgc.AddCollectorConfigurationInput(cfgID, input)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	// 202 no content
	return nil, sc, nil
}

// HandleUpdateCollectorConfigurationInput is the handler of Update a CollectorConfiguration Input API.
func HandleUpdateCollectorConfigurationInput(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	cfgID := ps.PathParam("collectorConfigurationID")
	inputID := ps.PathParam("collectorConfigurationInputID")
	// TODO authorize
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("backend", "type", "name", "forward_to"),
			Optional:     set.NewStrSet("properties"),
			Ignored:      set.NewStrSet("input_id"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}
	input := &graylog.CollectorConfigurationInput{}
	if err := util.MSDecode(body, input); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as CollectorConfiguration Input")
		return nil, 400, err
	}

	lgc.Logger().WithFields(log.Fields{
		"body": body, "input": input,
		"collector_configuration_id":       cfgID,
		"collector_configuration_input_id": inputID,
	}).Debug("request body")

	sc, err = lgc.UpdateCollectorConfigurationInput(cfgID, inputID, input)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	// 202 no content
	return nil, sc, nil
}

// HandleDeleteCollectorConfigurationInput is the handler of Delete an CollectorConfiguration Input API.
func HandleDeleteCollectorConfigurationInput(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	id := ps.PathParam("collectorConfigurationID")
	inputID := ps.PathParam("collectorConfigurationInputID")
	// TODO authorize
	sc, err := lgc.DeleteCollectorConfigurationInput(id, inputID)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}
