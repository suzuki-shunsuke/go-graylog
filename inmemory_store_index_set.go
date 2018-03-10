package graylog

// HasIndexSet
func (store *InMemoryStore) HasIndexSet(id string) (bool, error) {
	_, ok := store.indexSets[id]
	return ok, nil
}

// GetIndexSet returns an index set.
func (store *InMemoryStore) GetIndexSet(id string) (IndexSet, bool, error) {
	is, ok := store.indexSets[id]
	return is, ok, nil
}

// AddIndexSet adds an index set to the store.
func (store *InMemoryStore) AddIndexSet(indexSet *IndexSet) (*IndexSet, int, error) {
	store.indexSets[indexSet.Id] = *indexSet
	return indexSet, 200, nil
}

// UpdateIndexSet updates an index set at the Mock Server.
func (store *InMemoryStore) UpdateIndexSet(
	indexSet *IndexSet,
) (int, error) {
	store.indexSets[indexSet.Id] = *indexSet
	return 200, nil
}

// DeleteIndexSet removes a index set from the Mock Server.
func (store *InMemoryStore) DeleteIndexSet(id string) (int, error) {
	delete(store.indexSets, id)
	return 200, nil
}

// GetIndexSets returns a list of all index sets.
func (store *InMemoryStore) GetIndexSets() ([]IndexSet, error) {
	arr := make([]IndexSet, len(store.indexSets))
	i := 0
	for _, index := range store.indexSets {
		arr[i] = index
		i++
	}
	return arr, nil
}
