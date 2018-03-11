package graylog

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (ms *MockServer) HasIndexSet(id string) (bool, error) {
	return ms.store.HasIndexSet(id)
}

func (ms *MockServer) GetIndexSet(id string) (*IndexSet, error) {
	return ms.store.GetIndexSet(id)
}

// AddIndexSet adds an index set to the Mock Server.
func (ms *MockServer) AddIndexSet(indexSet *IndexSet) (*IndexSet, int, error) {
	if err := CreateValidator.Struct(indexSet); err != nil {
		return nil, 400, err
	}
	// indexPrefix unique check
	for _, is := range ms.indexSets {
		if is.IndexPrefix == indexSet.IndexPrefix {
			return nil, 400, fmt.Errorf(
				`Index prefix "%s" would conflict with an existing index set!`,
				indexSet.IndexPrefix)
		}
	}
	s := *indexSet
	s.ID = randStringBytesMaskImprSrc(24)
	i, err := ms.store.AddIndexSet(&s)
	if err != nil {
		return nil, 500, err
	}
	return i, 200, nil
}

// UpdateIndexSet updates an index set at the Mock Server.
func (ms *MockServer) UpdateIndexSet(
	indexSet *IndexSet,
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
	if err := UpdateValidator.Struct(indexSet); err != nil {
		return 400, err
	}
	// indexPrefix unique check
	for _, is := range ms.indexSets {
		if is.IndexPrefix == indexSet.IndexPrefix && is.ID != indexSet.ID {
			return 400, fmt.Errorf(
				`Index prefix "%s" would conflict with an existing index set!`,
				indexSet.IndexPrefix)
		}
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
func (ms *MockServer) IndexSetList() ([]IndexSet, error) {
	return ms.store.GetIndexSets()
}

// GET /system/indices/index_sets Get a list of all index sets
func (ms *MockServer) handleGetIndexSets(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	ms.handleInit(w, r, false)
	arr, err := ms.IndexSetList()
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err,
		}).Error("ms.HasIndexList() is failure")
		write500Error(w)
		return
	}
	indexSets := &indexSetsBody{
		IndexSets: arr, Total: len(arr), Stats: &IndexSetStats{}}
	writeOr500Error(w, indexSets)
}

// GET /system/indices/index_sets/{id} Get index set
func (ms *MockServer) handleGetIndexSet(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	id := ps.ByName("indexSetID")
	if id == "stats" {
		ms.handleGetAllIndexSetsStats(w, r, ps)
		return
	}
	ms.handleInit(w, r, false)
	indexSet, err := ms.GetIndexSet(id)
	if err != nil {
		ms.logger.WithFields(log.Fields{
			"error": err, "id": id,
		}).Info("ms.GetIndexSet() is failure")
		writeApiError(w, 500, err.Error())
		return
	}
	if indexSet == nil {
		writeApiError(w, 404, "No indexSet found with id %s", id)
		return
	}
	writeOr500Error(w, indexSet)
}

// POST /system/indices/index_sets Create index set
func (ms *MockServer) handleCreateIndexSet(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		write500Error(w)
		return
	}

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
		b, requiredFields, allowedFields, acceptedFields)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	indexSet := &IndexSet{}
	if err := msDecode(body, indexSet); err != nil {
		ms.logger.WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as indexSet")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	ms.Logger().WithFields(log.Fields{
		"body": string(b), "index_set": indexSet,
	}).Debug("request body")
	if is, sc, err := ms.AddIndexSet(indexSet); err != nil {
		writeApiError(w, sc, err.Error())
		return
	} else {
		ms.safeSave()
		writeOr500Error(w, is)
	}
}

// PUT /system/indices/index_sets/{id} Update index set
func (ms *MockServer) handleUpdateIndexSet(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		write500Error(w)
		return
	}
	id := ps.ByName("indexSetID")
	indexSet, err := ms.GetIndexSet(id)
	if err != nil {
		ms.logger.WithFields(log.Fields{
			"error": err, "id": id,
		}).Info("ms.GetIndexSet() is failure")
		writeApiError(w, 500, err.Error())
		return
	}
	if indexSet == nil {
		writeApiError(w, 404, "No indexSet found with id %s", id)
		return
	}

	// default can't change (ignored)
	requiredFields := []string{
		"title", "index_prefix", "rotation_strategy_class", "rotation_strategy",
		"retention_strategy_class", "retention_strategy", "creation_date",
		"index_analyzer", "shards", "index_optimization_max_num_segments"}
	acceptedFields := []string{
		"description", "replicas", "index_optimization_disabled", "writable"}
	sc, msg, body := validateRequestBody(b, requiredFields, nil, acceptedFields)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	if err := msDecode(body, indexSet); err != nil {
		ms.logger.WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as indexSet")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	ms.Logger().WithFields(log.Fields{
		"body": string(b), "index_set": indexSet,
	}).Debug("request body")
	indexSet.ID = id
	if sc, err := ms.UpdateIndexSet(indexSet); err != nil {
		writeApiError(w, sc, err.Error())
		return
	}
	ms.safeSave()
	writeOr500Error(w, indexSet)
}

// DELETE /system/indices/index_sets/{id} Delete index set
func (ms *MockServer) handleDeleteIndexSet(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	id := ps.ByName("indexSetID")
	if sc, err := ms.DeleteIndexSet(id); err != nil {
		writeApiError(w, sc, err.Error())
		return
	}
	ms.safeSave()
}

// PUT /system/indices/index_sets/{id}/default Set default index set
func (ms *MockServer) handleSetDefaultIndexSet(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	id := ps.ByName("indexSetID")
	indexSet, err := ms.GetIndexSet(id)
	if err != nil {
		ms.logger.WithFields(log.Fields{
			"error": err, "id": id,
		}).Info("ms.GetIndexSet() is failure")
		writeApiError(w, 500, err.Error())
		return
	}
	if indexSet == nil {
		writeApiError(w, 404, "No indexSet found with id %s", id)
		return
	}
	if !indexSet.Writable {
		writeApiError(w, 409, "Default index set must be writable.")
		return
	}
	if err := ms.store.SetDefaultIndexSetID(id); err != nil {
		writeApiError(w, 500, err.Error())
	}
	ms.safeSave()
	indexSet.Default = true
	writeOr500Error(w, indexSet)
}

// GET /system/indices/index_sets/{id}/stats Get index set statistics
func (ms *MockServer) handleGetIndexSetStats(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	id := ps.ByName("indexSetID")
	indexSetStats, ok := ms.indexSetStats[id]
	if !ok {
		writeApiError(w, 404, "No indexSet found with id %s", id)
		return
	}
	writeOr500Error(w, &indexSetStats)
}

// GET /system/indices/index_sets/stats Get stats of all index sets
func (ms *MockServer) handleGetAllIndexSetsStats(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	writeOr500Error(w, ms.AllIndexSetsStats())
}
