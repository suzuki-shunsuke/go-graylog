package graylog

import (
	"fmt"
)

// HasInput
func (store *InMemoryStore) HasInput(id string) (bool, error) {
	_, ok := store.inputs[id]
	return ok, nil
}

// GetInput returns an input.
func (store *InMemoryStore) GetInput(id string) (Input, bool, error) {
	s, ok := store.inputs[id]
	return s, ok, nil
}

// AddInput adds an input to the store.
func (store *InMemoryStore) AddInput(input *Input) (*Input, int, error) {
	store.inputs[input.Id] = *input
	return input, 200, nil
}

// UpdateInput updates an input at the InMemoryStore.
// Required: Title, Type, Configuration
// Allowed: Global, Node
func (store *InMemoryStore) UpdateInput(input *Input) (int, error) {
	u, ok, _ := store.GetInput(input.Id)
	if !ok {
		return 404, fmt.Errorf("The input is not found")
	}
	u.Title = input.Title
	u.Type = input.Type
	u.Configuration = &(*(input.Configuration))

	u.Global = input.Global
	u.Node = input.Node

	store.inputs[u.Id] = u
	return 200, nil
}

// DeleteInput deletes an input from the store.
func (store *InMemoryStore) DeleteInput(id string) (int, error) {
	delete(store.inputs, id)
	return 200, nil
}

// GetInputs returns inputs.
func (store *InMemoryStore) GetInputs() ([]Input, error) {
	size := len(store.inputs)
	arr := make([]Input, size)
	i := 0
	for _, input := range store.inputs {
		arr[i] = input
		i++
	}
	return arr, nil
}
