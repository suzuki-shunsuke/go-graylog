package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasIndexSet returns whether the user exists.
func (lgc *Logic) HasIndexSet(id string) (bool, error) {
	return lgc.store.HasIndexSet(id)
}

// GetIndexSets returns a list of all index sets.
func (lgc *Logic) GetIndexSets(skip, limit int) ([]graylog.IndexSet, int, int, error) {
	iss, total, err := lgc.store.GetIndexSets(skip, limit)
	if err != nil {
		return iss, total, 500, err
	}
	return iss, total, 200, nil
}

// GetIndexSet returns an index set.
// If an index set is not found, returns an error.
func (lgc *Logic) GetIndexSet(id string) (*graylog.IndexSet, int, error) {
	if id == "" {
		return nil, 400, fmt.Errorf("index set id is empty")
	}
	if err := ValidateObjectID(id); err != nil {
		// unfortunately graylog returns not 400 but 404.
		return nil, 404, err
	}
	is, err := lgc.store.GetIndexSet(id)
	if err != nil {
		return is, 500, err
	}
	if is == nil {
		return nil, 404, fmt.Errorf("no index set <%s> is found", id)
	}
	return is, 200, err
}

// AddIndexSet adds an index set to the Mock Server.
func (lgc *Logic) AddIndexSet(is *graylog.IndexSet) (int, error) {
	// Class org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy not subtype of [simple type, class org.graylog2.plugin.indexer.rotation.RotationStrategyConfig] (through reference chain: org.graylog2.rest.resources.system.indexer.responses.IndexSetSummary["rotation_strategy"])
	if is == nil {
		return 400, fmt.Errorf("index set is nil")
	}
	// indexPrefix unique check
	ok, err := lgc.store.IsConflictIndexPrefix(is.ID, is.IndexPrefix)
	if err != nil {
		return 500, err
	}
	if ok {
		return 400, fmt.Errorf(
			`index prefix "%s" would conflict with an existing index set`,
			is.IndexPrefix)
	}
	is.SetCreateDefaultValues()
	if err := validator.CreateValidator.Struct(is); err != nil {
		return 400, err
	}
	is.Default = false
	if err := lgc.store.AddIndexSet(is); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateIndexSet updates an index set at the Mock Server.
func (lgc *Logic) UpdateIndexSet(prms *graylog.IndexSetUpdateParams) (*graylog.IndexSet, int, error) {
	if prms == nil {
		return nil, 400, fmt.Errorf("index set is nil")
	}
	if err := validator.UpdateValidator.Struct(prms); err != nil {
		return nil, 400, err
	}
	ok, err := lgc.HasIndexSet(prms.ID)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "id": prms.ID,
		}).Error("lgc.HasIndexSet() is failure")
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("no indexSet found with id <%s>", prms.ID)
	}
	// indexPrefix unique check
	ok, err = lgc.store.IsConflictIndexPrefix(prms.ID, prms.IndexPrefix)
	if err != nil {
		return nil, 500, err
	}
	if ok {
		return nil, 400, fmt.Errorf(
			`index prefix "%s" would conflict with an existing index set`,
			prms.IndexPrefix)
	}
	defID, err := lgc.store.GetDefaultIndexSetID()
	if err != nil {
		return nil, 500, err
	}
	if defID == prms.ID && prms.Writable != nil && !(*prms.Writable) {
		return nil, 409, fmt.Errorf("default index set must be writable")
	}

	is, err := lgc.store.UpdateIndexSet(prms)
	if err != nil {
		return nil, 500, err
	}
	return is, 200, nil
}

// DeleteIndexSet removes a index set from the Mock Server.
func (lgc *Logic) DeleteIndexSet(id string) (int, error) {
	ok, err := lgc.HasIndexSet(id)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("lgc.HasIndexSet() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no indexSet with id <%s> is not found", id)
	}
	defID, err := lgc.store.GetDefaultIndexSetID()
	if err != nil {
		return 500, err
	}
	if id == defID {
		return 400, fmt.Errorf("default index set <%s> cannot be deleted", id)
	}
	if err := lgc.store.DeleteIndexSet(id); err != nil {
		return 500, err
	}
	return 204, nil
}

// SetDefaultIndexSet sets a default index set
func (lgc *Logic) SetDefaultIndexSet(id string) (*graylog.IndexSet, int, error) {
	is, sc, err := lgc.GetIndexSet(id)
	if err != nil {
		return nil, sc, err
	}
	if is == nil {
		return nil, 404, fmt.Errorf("no indexSet found with id <%s>", id)
	}
	if !is.Writable {
		return nil, 409, fmt.Errorf("default index set must be writable")
	}
	if err := lgc.store.SetDefaultIndexSetID(id); err != nil {
		return nil, 500, err
	}
	is.Default = true
	return is, 200, nil
}
