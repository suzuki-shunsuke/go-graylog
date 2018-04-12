package inmemory

import (
	"fmt"
	"strings"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// HasIndexSet
func (store *InMemoryStore) HasIndexSet(id string) (bool, error) {
	for _, is := range store.indexSets {
		if is.ID == id {
			return true, nil
		}
	}
	return false, nil
}

// GetIndexSet returns an index set.
func (store *InMemoryStore) GetIndexSet(id string) (*graylog.IndexSet, error) {
	for _, is := range store.indexSets {
		if is.ID == id {
			is.Default = store.defaultIndexSetID == id
			return &is, nil
		}
	}
	return nil, nil
}

// GetDefaultIndexSetID returns a default index set id.
func (store *InMemoryStore) GetDefaultIndexSetID() (string, error) {
	return store.defaultIndexSetID, nil
}

// SetDefaultIndexSetID sets a default index set id.
func (store *InMemoryStore) SetDefaultIndexSetID(id string) error {
	is, err := store.GetIndexSet(id)
	if err != nil {
		return err
	}
	if is == nil {
		return fmt.Errorf("no index set with id <%s> is not found", id)
	}
	if !is.Writable {
		return fmt.Errorf("default index set must be writable")
	}
	store.defaultIndexSetID = id
	return nil
}

// AddIndexSet adds an index set to the store.
func (store *InMemoryStore) AddIndexSet(is *graylog.IndexSet) error {
	if is == nil {
		return fmt.Errorf("index set is nil")
	}
	if is.ID == "" {
		is.ID = st.NewObjectID()
	}
	store.indexSets = append(store.indexSets, *is)
	return nil
}

// UpdateIndexSet updates an index set at the Mock Server.
func (store *InMemoryStore) UpdateIndexSet(is *graylog.IndexSet) error {
	id := is.ID
	for i, indexSet := range store.indexSets {
		if indexSet.ID == id {
			store.indexSets[i] = *is
			return nil
		}
	}
	return fmt.Errorf("no index set with id <%s>", id)
}

// DeleteIndexSet removes a index set from the Mock Server.
func (store *InMemoryStore) DeleteIndexSet(id string) error {
	size := len(store.indexSets)
	if size == 0 {
		return nil
	}
	var arr []graylog.IndexSet
	if size == 1 {
		arr = []graylog.IndexSet{}
	} else {
		arr = make([]graylog.IndexSet, size-1)
	}
	i := 0
	for _, is := range store.indexSets {
		if is.ID == id {
			continue
		}
		arr[i] = is
		i++
	}
	store.indexSets = arr
	return nil
}

// GetIndexSets returns a list of all index sets.
func (store *InMemoryStore) GetIndexSets(skip, limit int) ([]graylog.IndexSet, int, error) {
	total := len(store.indexSets)
	size := total
	if skip < 0 {
		skip = 0
	} else {
		size -= skip
	}
	if limit > 0 && limit < size {
		size = limit
	}
	arr := make([]graylog.IndexSet, size)
	defID := store.defaultIndexSetID
	for i := 0; i < size; i++ {
		is := store.indexSets[i+skip]
		is.Default = defID == is.ID
		arr[i] = is
	}
	return arr, total, nil
}

// IsConflictIndexPrefix returns true if indexPrefix would conflict with an existing index set.
func (store *InMemoryStore) IsConflictIndexPrefix(id, prefix string) (bool, error) {
	for _, is := range store.indexSets {
		if id != is.ID && strings.HasPrefix(prefix, is.IndexPrefix) {
			return true, nil
		}
		if id != is.ID && strings.HasPrefix(is.IndexPrefix, prefix) {
			return true, nil
		}
	}
	return false, nil
}
