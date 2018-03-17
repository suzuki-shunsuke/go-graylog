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
func (ms *MockServer) UpdateIndexSet(is *graylog.IndexSet) (int, error) {
	if is == nil {
		return 400, fmt.Errorf("index set is nil")
	}
	if err := validator.UpdateValidator.Struct(is); err != nil {
		return 400, err
	}
	ok, err := ms.HasIndexSet(is.ID)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": is.ID,
		}).Error("ms.HasIndexSet() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("No indexSet found with id %s", is.ID)
	}
	// indexPrefix unique check
	ok, err = ms.store.IsConflictIndexPrefix(is.ID, is.IndexPrefix)
	if err != nil {
		return 500, err
	}
	if ok {
		return 400, fmt.Errorf(
			`Index prefix "%s" would conflict with an existing index set!`,
			is.IndexPrefix)
	}

	if err := ms.store.UpdateIndexSet(is); err != nil {
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

// GetIndexSets returns a list of all index sets.
func (ms *MockServer) GetIndexSets() ([]graylog.IndexSet, error) {
	return ms.store.GetIndexSets()
}

// SetDefaultIndexSet sets a default index set
func (ms *MockServer) SetDefaultIndexSet(id string) (*graylog.IndexSet, int, error) {
	is, err := ms.GetIndexSet(id)
	if err != nil {
		ms.logger.WithFields(log.Fields{
			"error": err, "id": id,
		}).Info("ms.GetIndexSet() is failure")
		return nil, 500, err
	}
	if is == nil {
		return nil, 404, fmt.Errorf("No indexSet found with id %s", id)
	}
	if !is.Writable {
		return nil, 409, fmt.Errorf("Default index set must be writable.")
	}
	if err := ms.store.SetDefaultIndexSetID(id); err != nil {
		return nil, 500, err
	}
	is.Default = true
	return is, 200, nil
}
