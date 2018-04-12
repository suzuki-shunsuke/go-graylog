package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-set"
)

// HandleGetIndexSets
func HandleGetIndexSets(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// GET /system/indices/index_sets Get a list of all index sets
	// TODO skip limit stats
	skip := 0
	limit := 0
	stats := false
	arr, total, sc, err := ms.GetIndexSets(skip, limit)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err,
		}).Error("ms.HasIndexList() is failure")
		return arr, sc, err
	}
	if stats {
		// TODO
		stats, sc, err := ms.GetIndexSetStatsMap()
		if err != nil {
			return nil, sc, err
		}
		return &graylog.IndexSetsBody{
			IndexSets: arr, Total: total, Stats: stats}, 200, nil
	}
	return &graylog.IndexSetsBody{
		IndexSets: arr, Total: total,
		Stats: map[string]graylog.IndexSetStats{},
	}, sc, nil
}

// HandleGetIndexSet
func HandleGetIndexSet(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /system/indices/index_sets/{id} Get index set
	id := ps.ByName("indexSetID")
	if id == "stats" {
		return HandleGetTotalIndexSetStats(user, ms, w, r, ps)
	}
	if sc, err := ms.Authorize(user, "indexsets:read", id); err != nil {
		return nil, sc, err
	}
	return ms.GetIndexSet(id)
}

// HandleCreateIndexSet
func HandleCreateIndexSet(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// POST /system/indices/index_sets Create index set
	if sc, err := ms.Authorize(user, "indexsets:create"); err != nil {
		return nil, sc, err
	}
	requiredFields := set.NewStrSet(
		"title", "index_prefix", "rotation_strategy_class", "rotation_strategy",
		"retention_strategy_class", "retention_strategy", "creation_date",
		"index_analyzer", "shards", "index_optimization_max_num_segments")
	allowedFields := set.NewStrSet(
		"description", "replicas", "index_optimization_disabled",
		"writable", "default")
	acceptedFields := set.NewStrSet(
		"description", "replicas", "index_optimization_disabled", "writable")
	body, sc, err := validateRequestBody(
		r.Body, requiredFields, allowedFields, acceptedFields)
	if err != nil {
		return body, sc, err
	}

	indexSet := &graylog.IndexSet{}
	if err := msDecode(body, indexSet); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as indexSet")
		return nil, 400, err
	}

	ms.Logger().WithFields(log.Fields{
		"body": body, "index_set": indexSet,
	}).Debug("request body")
	sc, err = ms.AddIndexSet(indexSet)
	if err != nil {
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return indexSet, 201, nil
}

// HandleUpdateIndexSet
func HandleUpdateIndexSet(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /system/indices/index_sets/{id} Update index set
	is := &graylog.IndexSet{ID: ps.ByName("indexSetID")}
	if sc, err := ms.Authorize(user, "indexsets:edit", is.ID); err != nil {
		return nil, sc, err
	}

	// default can't change (ignored)
	requiredFields := set.NewStrSet(
		"title", "index_prefix", "rotation_strategy_class", "rotation_strategy",
		"retention_strategy_class", "retention_strategy",
		"index_analyzer", "shards", "index_optimization_max_num_segments")
	acceptedFields := set.NewStrSet(
		"description", "replicas", "index_optimization_disabled", "writable")
	body, sc, err := validateRequestBody(r.Body, requiredFields, nil, acceptedFields)
	if err != nil {
		return nil, sc, err
	}

	if err := msDecode(body, is); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as indexSet")
		return nil, 400, err
	}

	ms.Logger().WithFields(log.Fields{
		"body": body, "index_set": is,
	}).Debug("request body")
	if sc, err := ms.UpdateIndexSet(is); err != nil {
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return is, 200, nil
}

// HandleDeleteIndexSet
func HandleDeleteIndexSet(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// DELETE /system/indices/index_sets/{id} Delete index set
	id := ps.ByName("indexSetID")
	if sc, err := ms.Authorize(user, "indexsets:delete", id); err != nil {
		return nil, sc, err
	}
	if sc, err := ms.DeleteIndexSet(id); err != nil {
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return nil, 204, nil
}

// HandleSetDefaultIndexSet
func HandleSetDefaultIndexSet(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /system/indices/index_sets/{id}/default Set default index set
	id := ps.ByName("indexSetID")
	if sc, err := ms.Authorize(user, "indexsets:edit", id); err != nil {
		return nil, sc, err
	}
	is, sc, err := ms.SetDefaultIndexSet(id)
	if err != nil {
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return is, 200, nil
}
