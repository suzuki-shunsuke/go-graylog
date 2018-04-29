package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/util"
	"github.com/suzuki-shunsuke/go-set"
)

// HandleGetIndexSets is the handler of Get Index Sets API.
func HandleGetIndexSets(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// GET /system/indices/index_sets Get a list of all index sets
	skip := 0
	limit := 0
	stats := false
	query := r.URL.Query()
	s, ok := query["skip"]
	var err error
	if ok && len(s) > 0 {
		skip, err = strconv.Atoi(s[0])
		if err != nil {
			lgc.Logger().WithFields(log.Fields{
				"error": err, "param_name": "skip", "value": s[0],
			}).Warn("failed to convert string to integer")
			// Unfortunately, graylog returns 404
			// https://github.com/Graylog2/graylog2-server/issues/4721
			return nil, 404, fmt.Errorf("HTTP 404 Not Found")
		}
	}
	l, ok := query["limit"]
	if ok && len(l) > 0 {
		limit, err = strconv.Atoi(l[0])
		if err != nil {
			lgc.Logger().WithFields(log.Fields{
				"error": err, "param_name": "limit", "value": l[0],
			}).Warn("failed to convert string to integer")
			// Unfortunately, graylog returns 404
			// https://github.com/Graylog2/graylog2-server/issues/4721
			return nil, 404, fmt.Errorf("HTTP 404 Not Found")
		}
	}
	st, ok := query["stats"]
	if ok && len(st) > 0 {
		stats, err = strconv.ParseBool(st[0])
		if err != nil {
			lgc.Logger().WithFields(log.Fields{
				"error": err, "param_name": "stats", "value": st[0],
			}).Warn("failed to convert string to bool")
			// Unfortunately, graylog ignores invalid stats parameter
			// TODO send issue
			stats = false
		}
	}

	arr, total, sc, err := lgc.GetIndexSets(skip, limit)
	if err != nil {
		logic.LogWE(sc, lgc.Logger().WithFields(log.Fields{
			"error": err, "skip": skip, "limit": limit, "status_code": sc,
		}), "failed to get index sets")
		return arr, sc, err
	}
	if stats {
		stats, sc, err := lgc.GetIndexSetStatsMap()
		if err != nil {
			return nil, sc, err
		}
		return &graylog.IndexSetsBody{
			IndexSets: arr, Total: total, Stats: stats}, sc, nil
	}
	return &graylog.IndexSetsBody{
		IndexSets: arr, Total: total,
		Stats: map[string]graylog.IndexSetStats{},
	}, sc, nil
}

// HandleGetIndexSet is the handler of Get an Index Set API.
func HandleGetIndexSet(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /system/indices/index_sets/{id} Get index set
	id := ps.ByName("indexSetID")
	if id == "stats" {
		return HandleGetTotalIndexSetStats(user, lgc, w, r, ps)
	}
	if sc, err := lgc.Authorize(user, "indexsets:read", id); err != nil {
		return nil, sc, err
	}
	return lgc.GetIndexSet(id)
}

// HandleCreateIndexSet is the handler of Create an Index Set API.
func HandleCreateIndexSet(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// POST /system/indices/index_sets Create index set
	if sc, err := lgc.Authorize(user, "indexsets:create"); err != nil {
		return nil, sc, err
	}
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required: set.NewStrSet(
				"title", "index_prefix", "rotation_strategy_class", "rotation_strategy",
				"retention_strategy_class", "retention_strategy", "creation_date",
				"index_analyzer", "shards", "index_optimization_max_num_segments"),
			Optional:     set.NewStrSet("description", "replicas", "index_optimization_disabled", "writable"),
			Ignored:      set.NewStrSet("default"),
			ExtForbidden: true,
		})
	if err != nil {
		return body, sc, err
	}

	is := &graylog.IndexSet{}
	if err := util.MSDecode(body, is); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as indexSet")
		return nil, 400, err
	}

	lgc.Logger().WithFields(log.Fields{
		"body": body, "index_set": is,
	}).Debug("request body")
	if is.ID == "" {
		sc, err = lgc.AddIndexSet(is)
		if err != nil {
			return nil, sc, err
		}
		if err := lgc.Save(); err != nil {
			return nil, 500, err
		}
		return is, sc, nil
	}
	is, sc, err = lgc.UpdateIndexSet(is.NewUpdateParams())
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return is, sc, nil
}

// HandleUpdateIndexSet is the handler of Update an Index Set API.
func HandleUpdateIndexSet(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /system/indices/index_sets/{id} Update index set
	id := ps.ByName("indexSetID")
	prms := &graylog.IndexSetUpdateParams{}
	if sc, err := lgc.Authorize(user, "indexsets:edit", id); err != nil {
		return nil, sc, err
	}

	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required: set.NewStrSet(
				"title", "index_prefix", "rotation_strategy_class", "rotation_strategy",
				"retention_strategy_class", "retention_strategy",
				"index_analyzer", "shards", "index_optimization_max_num_segments"),
			Optional: set.NewStrSet("description", "replicas", "index_optimization_disabled", "writable"),
			Ignored:  set.NewStrSet("default", "creation_date"),
		})
	if err != nil {
		return nil, sc, err
	}

	if err := util.MSDecode(body, prms); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Warn("Failed to parse request body as indexSetUpdateParams")
		return nil, 400, err
	}
	prms.ID = id
	lgc.Logger().WithFields(log.Fields{
		"body": body, "index_set": prms,
	}).Debug("request body")
	is, sc, err := lgc.UpdateIndexSet(prms)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return is, sc, nil
}

// HandleDeleteIndexSet is the handler of Delete an Index Set API.
func HandleDeleteIndexSet(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// DELETE /system/indices/index_sets/{id} Delete index set
	id := ps.ByName("indexSetID")
	if sc, err := lgc.Authorize(user, "indexsets:delete", id); err != nil {
		return nil, sc, err
	}
	sc, err := lgc.DeleteIndexSet(id)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}

// HandleSetDefaultIndexSet is the handler of Set the default Index Set API.
func HandleSetDefaultIndexSet(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /system/indices/index_sets/{id}/default Set default index set
	id := ps.ByName("indexSetID")
	if sc, err := lgc.Authorize(user, "indexsets:edit", id); err != nil {
		return nil, sc, err
	}
	is, sc, err := lgc.SetDefaultIndexSet(id)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return is, 200, nil
}
