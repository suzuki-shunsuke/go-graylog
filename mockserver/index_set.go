package mockserver

import (
	"fmt"

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
