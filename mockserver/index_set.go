package mockserver

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

func (ms *MockServer) HasIndexSet(id string) (bool, error) {
	return ms.store.HasIndexSet(id)
}

func (ms *MockServer) GetIndexSet(id string) (*graylog.IndexSet, error) {
	return ms.store.GetIndexSet(id)
}

// AddIndexSet adds an index set to the Mock Server.
func (ms *MockServer) AddIndexSet(indexSet *graylog.IndexSet) (int, error) {
	if err := validator.CreateValidator.Struct(indexSet); err != nil {
		return 400, err
	}
	// indexPrefix unique check
	ok, err := ms.store.IsConflictIndexPrefix("", indexSet.IndexPrefix)
	if err != nil {
		return 500, err
	}
	if ok {
		return 400, fmt.Errorf(
			`Index prefix "%s" would conflict with an existing index set!`,
			indexSet.IndexPrefix)
	}
	indexSet.ID = randStringBytesMaskImprSrc(24)
	indexSet.Default = false
	if err := ms.store.AddIndexSet(indexSet); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateIndexSet updates an index set at the Mock Server.
func (ms *MockServer) UpdateIndexSet(
	indexSet *graylog.IndexSet,
) (int, error) {
	ok, err := ms.HasIndexSet(indexSet.ID)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": indexSet.ID,
		}).Error("ms.HasIndexSet() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("No indexSet found with id %s", indexSet.ID)
	}
	if err := validator.UpdateValidator.Struct(indexSet); err != nil {
		return 400, err
	}
	// indexPrefix unique check
	ok, err = ms.store.IsConflictIndexPrefix(indexSet.ID, indexSet.IndexPrefix)
	if err != nil {
		return 500, err
	}
	if ok {
		return 400, fmt.Errorf(
			`Index prefix "%s" would conflict with an existing index set!`,
			indexSet.IndexPrefix)
	}

	if err := ms.store.UpdateIndexSet(indexSet); err != nil {
		return 500, err
	}
	return 200, nil
}

// DeleteIndexSet removes a index set from the Mock Server.
func (ms *MockServer) DeleteIndexSet(id string) (int, error) {
	ok, err := ms.HasIndexSet(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.HasIndexSet() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("No indexSet with id %s is not found", id)
	}
	defID, err := ms.store.GetDefaultIndexSetID()
	if err != nil {
		return 500, err
	}
	if id == defID {
		return 400, fmt.Errorf("default index set <%s> cannot be deleted", id)
	}
	if err := ms.store.DeleteIndexSet(id); err != nil {
		return 500, err
	}
	return 200, nil
}

// IndexSetList returns a list of all index sets.
func (ms *MockServer) IndexSetList() ([]graylog.IndexSet, error) {
	return ms.store.GetIndexSets()
}

// GET /system/indices/index_sets Get a list of all index sets
func (ms *MockServer) handleGetIndexSets(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	arr, err := ms.IndexSetList()
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
func (ms *MockServer) handleGetIndexSet(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("indexSetID")
	if id == "stats" {
		return ms.handleGetAllIndexSetsStats(w, r, ps)
	}
	indexSet, err := ms.GetIndexSet(id)
	if err != nil {
		ms.logger.WithFields(log.Fields{
			"error": err, "id": id,
		}).Info("ms.GetIndexSet() is failure")
		return 500, nil, err
	}
	if indexSet == nil {
		return 404, nil, fmt.Errorf("No indexSet found with id %s", id)
	}
	return 200, indexSet, nil
}

// POST /system/indices/index_sets Create index set
func (ms *MockServer) handleCreateIndexSet(
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
	sc, msg, body := validateRequestBody(
		r.Body, requiredFields, allowedFields, acceptedFields)
	if sc != 200 {
		return sc, nil, fmt.Errorf(msg)
	}

	indexSet := &graylog.IndexSet{}
	if err := msDecode(body, indexSet); err != nil {
		ms.logger.WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as indexSet")
		return 400, nil, err
	}

	ms.Logger().WithFields(log.Fields{
		"body": body, "index_set": indexSet,
	}).Debug("request body")
	sc, err := ms.AddIndexSet(indexSet)
	if err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 201, indexSet, nil
}

// PUT /system/indices/index_sets/{id} Update index set
func (ms *MockServer) handleUpdateIndexSet(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("indexSetID")
	indexSet, err := ms.GetIndexSet(id)
	if err != nil {
		ms.logger.WithFields(log.Fields{
			"error": err, "id": id,
		}).Info("ms.GetIndexSet() is failure")
		return 500, nil, err
	}
	if indexSet == nil {
		return 404, nil, fmt.Errorf("No indexSet found with id %s", id)
	}

	// default can't change (ignored)
	requiredFields := []string{
		"title", "index_prefix", "rotation_strategy_class", "rotation_strategy",
		"retention_strategy_class", "retention_strategy", "creation_date",
		"index_analyzer", "shards", "index_optimization_max_num_segments"}
	acceptedFields := []string{
		"description", "replicas", "index_optimization_disabled", "writable"}
	sc, msg, body := validateRequestBody(r.Body, requiredFields, nil, acceptedFields)
	if sc != 200 {
		return sc, nil, fmt.Errorf(msg)
	}

	if err := msDecode(body, indexSet); err != nil {
		ms.logger.WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as indexSet")
		return 400, nil, err
	}

	ms.Logger().WithFields(log.Fields{
		"body": body, "index_set": indexSet,
	}).Debug("request body")
	indexSet.ID = id
	if sc, err := ms.UpdateIndexSet(indexSet); err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 200, indexSet, nil
}

// DELETE /system/indices/index_sets/{id} Delete index set
func (ms *MockServer) handleDeleteIndexSet(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("indexSetID")
	if sc, err := ms.DeleteIndexSet(id); err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 204, nil, nil
}

// PUT /system/indices/index_sets/{id}/default Set default index set
func (ms *MockServer) handleSetDefaultIndexSet(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("indexSetID")
	indexSet, err := ms.GetIndexSet(id)
	if err != nil {
		ms.logger.WithFields(log.Fields{
			"error": err, "id": id,
		}).Info("ms.GetIndexSet() is failure")
		return 500, nil, err
	}
	if indexSet == nil {
		return 404, nil, fmt.Errorf("No indexSet found with id %s", id)
	}
	if !indexSet.Writable {
		return 409, nil, fmt.Errorf("Default index set must be writable.")
	}
	if err := ms.store.SetDefaultIndexSetID(id); err != nil {
		return 500, nil, err
	}
	ms.safeSave()
	indexSet.Default = true
	return 200, indexSet, nil
}
