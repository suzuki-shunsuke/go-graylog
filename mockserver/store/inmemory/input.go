package inmemory

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// HasInput
func (store *InMemoryStore) HasInput(id string) (bool, error) {
	_, ok := store.inputs[id]
	return ok, nil
}

// GetInput returns an input.
func (store *InMemoryStore) GetInput(id string) (*graylog.Input, error) {
	s, ok := store.inputs[id]
	if ok {
		return &s, nil
	}
	return nil, nil
}

// AddInput adds an input to the store.
func (store *InMemoryStore) AddInput(input *graylog.Input) error {
	store.inputs[input.ID] = *input
	return nil
}

// UpdateInput updates an input at the InMemoryStore.
// Required: Title, Type, Configuration
// Allowed: Global, Node
func (store *InMemoryStore) UpdateInput(input *graylog.Input) error {
	u, err := store.GetInput(input.ID)
	if err != nil {
		return err
	}
	if u == nil {
		return fmt.Errorf("the input <%s> is not found", input.ID)
	}
	u.Title = input.Title
	u.Type = input.Type
	u.Configuration = &(*(input.Configuration))

	u.Global = input.Global
	u.Node = input.Node

	store.inputs[u.ID] = *u
	return nil
}

// DeleteInput deletes an input from the store.
func (store *InMemoryStore) DeleteInput(id string) error {
	delete(store.inputs, id)
	return nil
}

// GetInputs returns inputs.
func (store *InMemoryStore) GetInputs() ([]graylog.Input, error) {
	size := len(store.inputs)
	arr := make([]graylog.Input, size)
	i := 0
	for _, input := range store.inputs {
		arr[i] = input
		i++
	}
	return arr, nil
}
