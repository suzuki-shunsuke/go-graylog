package graylog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (ms *MockServer) IndexSetList() []IndexSet {
	if ms.IndexSets == nil {
		return []IndexSet{}
	}
	arr := make([]IndexSet, len(ms.IndexSets))
	i := 0
	for _, index := range ms.IndexSets {
		arr[i] = index
		i++
	}
	return arr
}

func validateIndexSet(indexSet *IndexSet) (int, []byte) {
	return 200, []byte("")
}

// GET /system/indices/index_sets Get a list of all index sets
func (ms *MockServer) handleGetIndexSets(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	w.Header().Set("Content-Type", "application/json")
	arr := ms.IndexSetList()
	indexSets := indexSetsBody{
		IndexSets: arr, Total: len(arr), Stats: &IndexSetStats{}}
	b, err := json.Marshal(&indexSets)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

// GET /system/indices/index_sets/{id} Get index set
func (ms *MockServer) handleGetIndexSet(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("indexSetId")
	indexSet, ok := ms.IndexSets[id]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(`{"type": "ApiError", "message": "No indexSet found with id %s"}`, id)))
		return
	}
	b, err := json.Marshal(&indexSet)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

// POST /system/indices/index_sets Create index set
func (ms *MockServer) handleCreateIndexSet(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	indexSet := IndexSet{Id: randStringBytesMaskImprSrc(24)}
	err = json.Unmarshal(b, &indexSet)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	sc, msg := validateIndexSet(&indexSet)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	ms.IndexSets[indexSet.Id] = indexSet
	b, err = json.Marshal(&indexSet)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

// PUT /system/indices/index_sets/{id} Update index set
func (ms *MockServer) handleUpdateIndexSet(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	id := ps.ByName("indexSetId")
	if _, ok := ms.IndexSets[id]; !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(`{"type": "ApiError", "message": "No indexSet found with id %s"}`, id)))
		return
	}
	indexSet := IndexSet{}
	err = json.Unmarshal(b, &indexSet)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	sc, msg := validateIndexSet(&indexSet)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	delete(ms.IndexSets, id)
	ms.IndexSets[indexSet.Id] = indexSet
	b, err = json.Marshal(&indexSet)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

// DELETE /system/indices/index_sets/{id} Delete index set
func (ms *MockServer) handleDeleteIndexSet(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("indexSetId")
	_, ok := ms.IndexSets[id]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(`{"type": "ApiError", "message": "No indexSet found with id %s"}`, id)))
		return
	}
	delete(ms.IndexSets, id)
}
