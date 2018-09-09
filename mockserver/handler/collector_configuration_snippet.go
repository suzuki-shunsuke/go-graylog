package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/util"
)

// HandleCreateCollectorConfigurationSnippet is the handler of Create a CollectorConfiguration Snippet API.
func HandleCreateCollectorConfigurationSnippet(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// TODO authorize
	cfgID := ps.PathParam("collectorConfigurationSnippetID")
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("tags", "snippets", "outputs", "snippets"),
			Optional:     nil,
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}
	snippet := &graylog.CollectorConfigurationSnippet{}
	if err := util.MSDecode(body, snippet); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as CollectorConfigurationSnippet")
		return nil, 400, err
	}
	sc, err = lgc.AddCollectorConfigurationSnippet(cfgID, snippet)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	// 202 no content
	return nil, sc, nil
}

// HandleUpdateCollectorConfigurationSnippet is the handler of Update a CollectorConfiguration Snippet API.
func HandleUpdateCollectorConfigurationSnippet(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	cfgID := ps.PathParam("collectorConfigurationID")
	snippetID := ps.PathParam("collectorConfigurationSnippetID")
	// TODO authorize
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("backend", "name", "snippet"),
			Ignored:      set.NewStrSet("snippet_id"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}
	snippet := &graylog.CollectorConfigurationSnippet{}
	if err := util.MSDecode(body, snippet); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as CollectorConfiguration Snippet")
		return nil, 400, err
	}

	lgc.Logger().WithFields(log.Fields{
		"body": body, "snippet": snippet,
		"collector_configuration_id":         cfgID,
		"collector_configuration_snippet_id": snippetID,
	}).Debug("request body")

	sc, err = lgc.UpdateCollectorConfigurationSnippet(cfgID, snippetID, snippet)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	// 202 no content
	return nil, sc, nil
}

// HandleDeleteCollectorConfigurationSnippet is the handler of Delete an CollectorConfiguration Snippet API.
func HandleDeleteCollectorConfigurationSnippet(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	id := ps.PathParam("collectorConfigurationID")
	snippetID := ps.PathParam("collectorConfigurationSnippetID")
	// TODO authorize
	sc, err := lgc.DeleteCollectorConfigurationSnippet(id, snippetID)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}
