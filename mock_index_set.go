package graylog

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (ms *MockServer) HasIndexSet(id string) (bool, error) {
	return ms.store.HasIndexSet(id)
}

func (ms *MockServer) GetIndexSet(id string) (IndexSet, bool, error) {
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
	s.Id = randStringBytesMaskImprSrc(24)
	return ms.store.AddIndexSet(&s)
}

// UpdateIndexSet updates an index set at the Mock Server.
func (ms *MockServer) UpdateIndexSet(
	indexSet *IndexSet,
) (int, error) {
	ok, err := ms.HasIndexSet(indexSet.Id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": indexSet.Id,
		}).Error("ms.HasIndexSet() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("No indexSet found with id %s", indexSet.Id)
	}
	if err := UpdateValidator.Struct(indexSet); err != nil {
		return 400, err
	}
	// indexPrefix unique check
	for _, is := range ms.indexSets {
		if is.IndexPrefix == indexSet.IndexPrefix && is.Id != indexSet.Id {
			return 400, fmt.Errorf(
				`Index prefix "%s" would conflict with an existing index set!`,
				indexSet.IndexPrefix)
		}
	}
	return ms.store.UpdateIndexSet(indexSet)
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
	return ms.store.DeleteIndexSet(id)
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
	id := ps.ByName("indexSetId")
	if id == "stats" {
		ms.handleGetAllIndexSetsStats(w, r, ps)
		return
	}
	ms.handleInit(w, r, false)
	indexSet, ok, err := ms.GetIndexSet(id)
	if err != nil {
		ms.logger.WithFields(log.Fields{
			"error": err, "id": id,
		}).Info("ms.GetIndexSet() is failure")
		writeApiError(w, 500, err.Error())
		return
	}
	if !ok {
		writeApiError(w, 404, "No indexSet found with id %s", id)
		return
	}
	writeOr500Error(w, &indexSet)
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
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
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
	id := ps.ByName("indexSetId")
	indexSet, ok, err := ms.GetIndexSet(id)
	if err != nil {
		ms.logger.WithFields(log.Fields{
			"error": err, "id": id,
		}).Info("ms.GetIndexSet() is failure")
		writeApiError(w, 500, err.Error())
		return
	}
	if !ok {
		writeApiError(w, 404, "No indexSet found with id %s", id)
		return
	}

	if err := json.Unmarshal(b, &indexSet); err != nil {
		writeApiError(w, 400, "400 Bad Request")
		return
	}
	indexSet.Id = id
	if sc, err := ms.UpdateIndexSet(&indexSet); err != nil {
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
	id := ps.ByName("indexSetId")
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
	id := ps.ByName("indexSetId")
	indexSet, ok, err := ms.GetIndexSet(id)
	if err != nil {
		ms.logger.WithFields(log.Fields{
			"error": err, "id": id,
		}).Info("ms.GetIndexSet() is failure")
		writeApiError(w, 500, err.Error())
		return
	}
	if !ok {
		writeApiError(w, 404, "No indexSet found with id %s", id)
		return
	}
	if !indexSet.Writable {
		writeApiError(w, 409, "Default index set must be writable.")
		return
	}
	for k, v := range ms.indexSets {
		if v.Default {
			v.Default = false
			ms.indexSets[k] = v
			break
		}
	}
	indexSet.Default = true
	ms.AddIndexSet(&indexSet)
	ms.safeSave()
	writeOr500Error(w, &indexSet)
}

// GET /system/indices/index_sets/{id}/stats Get index set statistics
func (ms *MockServer) handleGetIndexSetStats(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	id := ps.ByName("indexSetId")
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
