package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/util"
)

// HandleCreateCollectorConfigurationOutput is the handler of Create a CollectorConfiguration Output API.
func HandleCreateCollectorConfigurationOutput(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// TODO authorize
	cfgID := ps.PathParam("collectorConfigurationID")
	// Known properties include: output_id, type, name, backend, properties
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("type", "name", "backend"),
			Optional:     set.NewStrSet("properties"),
			Ignored:      set.NewStrSet("output_id"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}
	output := &graylog.CollectorConfigurationOutput{}
	if err := util.MSDecode(body, output); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as CollectorConfigurationOutput")
		return nil, 400, err
	}
	sc, err = lgc.AddCollectorConfigurationOutput(cfgID, output)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	// 202 no content
	return nil, sc, nil
}

// HandleUpdateCollectorConfigurationOutput is the handler of Update a CollectorConfiguration Output API.
func HandleUpdateCollectorConfigurationOutput(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	cfgID := ps.PathParam("collectorConfigurationID")
	outputID := ps.PathParam("collectorConfigurationOutputID")
	// TODO authorize
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("backend", "type", "name"),
			Optional:     set.NewStrSet("properties"),
			Ignored:      set.NewStrSet("output_id"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}
	output := &graylog.CollectorConfigurationOutput{}
	if err := util.MSDecode(body, output); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as CollectorConfiguration Output")
		return nil, 400, err
	}

	lgc.Logger().WithFields(log.Fields{
		"body": body, "output": output,
		"collector_configuration_id":        cfgID,
		"collector_configuration_output_id": outputID,
	}).Debug("request body")

	sc, err = lgc.UpdateCollectorConfigurationOutput(cfgID, outputID, output)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	// 202 no content
	return nil, sc, nil
}

// HandleDeleteCollectorConfigurationOutput is the handler of Delete an CollectorConfiguration Output API.
func HandleDeleteCollectorConfigurationOutput(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	id := ps.PathParam("collectorConfigurationID")
	outputID := ps.PathParam("collectorConfigurationOutputID")
	// TODO authorize
	sc, err := lgc.DeleteCollectorConfigurationOutput(id, outputID)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}
