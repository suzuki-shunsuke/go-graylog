package plain

import (
	"fmt"
	"time"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// HasInput returns whether the input exists.
func (store *Store) HasInput(id string) (bool, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	_, ok := store.inputs[id]
	return ok, nil
}

// GetInput returns an input.
func (store *Store) GetInput(id string) (*graylog.Input, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	s, ok := store.inputs[id]
	if ok {
		return &s, nil
	}
	return nil, nil
}

// AddInput adds an input to the store.
func (store *Store) AddInput(input *graylog.Input) error {
	if input == nil {
		return fmt.Errorf("input is nil")
	}
	if input.ID == "" {
		input.ID = st.NewObjectID()
	}
	input.CreatedAt = time.Now().Format("2006-01-02T15:04:05.000Z")

	store.imutex.Lock()
	defer store.imutex.Unlock()
	store.inputs[input.ID] = *input
	return nil
}

// UpdateInput updates an input at the Store.
// Required: Title, Type, Attributes
// Allowed: Global, Node
func (store *Store) UpdateInput(prms *graylog.InputUpdateParams) (*graylog.Input, error) {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	input, ok := store.inputs[prms.ID]
	if !ok {
		return nil, fmt.Errorf("the input <%s> is not found", prms.ID)
	}
	input.Title = prms.Title
	input.Attributes = prms.Attributes
	if prms.Global == nil {
		input.Global = *prms.Global
	}
	if prms.Node == "" {
		input.Node = prms.Node
	}
	store.inputs[input.ID] = input
	return &input, nil
}

// DeleteInput deletes an input from the store.
func (store *Store) DeleteInput(id string) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	delete(store.inputs, id)
	return nil
}

// GetInputs returns inputs.
func (store *Store) GetInputs() ([]graylog.Input, int, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	size := len(store.inputs)
	arr := make([]graylog.Input, size)
	i := 0
	for _, input := range store.inputs {
		arr[i] = input
		i++
	}
	return arr, size, nil
}
