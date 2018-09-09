package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/util"
)

// HandleGetCollectorConfiguration is the handler of Get an CollectorConfiguration API.
func HandleGetCollectorConfiguration(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /api/plugins/org.graylog.plugins.collector/configurations/:id
	id := ps.PathParam("collectorConfigurationID")
	// TODO authorize
	return lgc.GetCollectorConfiguration(id)
}

// HandleGetCollectorConfigurations is the handler of Get Collector Configurations API.
func HandleGetCollectorConfigurations(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// GET /plugins/org.graylog.plugins.collector/configurations List all collector configurations
	arr, total, sc, err := lgc.GetCollectorConfigurations()
	if err != nil {
		return arr, sc, err
	}
	cfgs := &graylog.CollectorConfigurationsBody{Configurations: arr, Total: total}
	return cfgs, sc, nil
}

// HandleCreateCollectorConfiguration is the handler of Create an CollectorConfiguration API.
func HandleCreateCollectorConfiguration(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// POST /system/inputs Launch input on this node
	// TODO authorize
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("tags", "inputs", "outputs", "snippets"),
			Optional:     set.NewStrSet("name", "id"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}
	cfg := &graylog.CollectorConfiguration{}
	if err := util.MSDecode(body, cfg); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as CollectorConfiguration")
		return nil, 400, err
	}
	sc, err = lgc.AddCollectorConfiguration(cfg)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return cfg, sc, nil
}

// HandleRenameCollectorConfiguration is the handler of Rename a CollectorConfiguration API.
func HandleRenameCollectorConfiguration(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// PUT /plugins/org.graylog.plugins.collector/configurations/{id}/name Updates a collector configuration name
	id := ps.PathParam("collectorConfigurationID")
	// TODO authorize
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("tags", "inputs", "outputs", "snippets"),
			Optional:     set.NewStrSet("name", "id"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}
	prms := &graylog.CollectorConfiguration{}
	if err := util.MSDecode(body, prms); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as CollectorConfiguration")
		return nil, 400, err
	}

	lgc.Logger().WithFields(log.Fields{
		"body": body, "input": prms, "id": id,
	}).Debug("request body")

	cfg, sc, err := lgc.RenameCollectorConfiguration(id, prms.Name)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return cfg, sc, nil
}

// HandleDeleteCollectorConfiguration is the handler of Delete an CollectorConfiguration API.
func HandleDeleteCollectorConfiguration(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// DELETE /plugins/org.graylog.plugins.collector/configurations/{id} Delete a collector configuration
	id := ps.PathParam("collectorConfigurationID")
	// TODO authorize
	sc, err := lgc.DeleteCollectorConfiguration(id)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}
