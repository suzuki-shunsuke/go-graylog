package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

// GET /system/indices/index_sets Get a list of all index sets
func HandleGetIndexSets(
	ms *logic.Server,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	arr, err := ms.GetIndexSets()
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err,
		}).Error("ms.HasIndexList() is failure")
		return 500, nil, err
	}
	indexSets := &graylog.IndexSetsBody{
		IndexSets: arr, Total: len(arr), Stats: &graylog.IndexSetStats{}}
	return 200, indexSets, nil
}

// GET /system/indices/index_sets/{id} Get index set
func HandleGetIndexSet(
	ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("indexSetID")
	if id == "stats" {
		return HandleGetAllIndexSetsStats(ms, w, r, ps)
	}
	is, sc, err := ms.GetIndexSet(id)
	return sc, is, err
}

// POST /system/indices/index_sets Create index set
func HandleCreateIndexSet(
	ms *logic.Server,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	requiredFields := []string{
		"title", "index_prefix", "rotation_strategy_class", "rotation_strategy",
		"retention_strategy_class", "retention_strategy", "creation_date",
		"index_analyzer", "shards", "index_optimization_max_num_segments"}
	allowedFields := []string{
		"description", "replicas", "index_optimization_disabled",
		"writable", "default"}
	acceptedFields := []string{
		"description", "replicas", "index_optimization_disabled", "writable"}
	body, sc, err := validateRequestBody(
		r.Body, requiredFields, allowedFields, acceptedFields)
	if err != nil {
		return sc, nil, err
	}

	indexSet := &graylog.IndexSet{}
	if err := msDecode(body, indexSet); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as indexSet")
		return 400, nil, err
	}

	ms.Logger().WithFields(log.Fields{
		"body": body, "index_set": indexSet,
	}).Debug("request body")
	sc, err = ms.AddIndexSet(indexSet)
	if err != nil {
		return sc, nil, err
	}
	ms.SafeSave()
	return 201, indexSet, nil
}

// PUT /system/indices/index_sets/{id} Update index set
func HandleUpdateIndexSet(
	ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	// default can't change (ignored)
	requiredFields := []string{
		"title", "index_prefix", "rotation_strategy_class", "rotation_strategy",
		"retention_strategy_class", "retention_strategy", "creation_date",
		"index_analyzer", "shards", "index_optimization_max_num_segments"}
	acceptedFields := []string{
		"description", "replicas", "index_optimization_disabled", "writable"}
	body, sc, err := validateRequestBody(r.Body, requiredFields, nil, acceptedFields)
	if err != nil {
		return sc, nil, err
	}

	is := &graylog.IndexSet{ID: ps.ByName("indexSetID")}
	if err := msDecode(body, is); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as indexSet")
		return 400, nil, err
	}

	ms.Logger().WithFields(log.Fields{
		"body": body, "index_set": is,
	}).Debug("request body")
	if sc, err := ms.UpdateIndexSet(is); err != nil {
		return sc, nil, err
	}
	ms.SafeSave()
	return 200, is, nil
}

// DELETE /system/indices/index_sets/{id} Delete index set
func HandleDeleteIndexSet(
	ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("indexSetID")
	if sc, err := ms.DeleteIndexSet(id); err != nil {
		return sc, nil, err
	}
	ms.SafeSave()
	return 204, nil, nil
}

// PUT /system/indices/index_sets/{id}/default Set default index set
func HandleSetDefaultIndexSet(
	ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	is, sc, err := ms.SetDefaultIndexSet(ps.ByName("indexSetID"))
	if err != nil {
		return sc, nil, err
	}
	ms.SafeSave()
	return 200, is, nil
}
