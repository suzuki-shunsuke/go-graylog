package plain

import (
	"fmt"
	"time"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// HasInput
func (store *PlainStore) HasInput(id string) (bool, error) {
	_, ok := store.inputs[id]
	return ok, nil
}

// GetInput returns an input.
func (store *PlainStore) GetInput(id string) (*graylog.Input, error) {
	s, ok := store.inputs[id]
	if ok {
		return &s, nil
	}
	return nil, nil
}

// AddInput adds an input to the store.
func (store *PlainStore) AddInput(input *graylog.Input) error {
	if input == nil {
		return fmt.Errorf("input is nil")
	}
	if input.ID == "" {
		input.ID = st.NewObjectID()
	}
	input.CreatedAt = time.Now().Format("2006-01-02T15:04:05.000Z")

	mutex.Lock()
	store.inputs[input.ID] = *input
	mutex.Unlock()
	return nil
}

// UpdateInput updates an input at the PlainStore.
// Required: Title, Type, Configuration
// Allowed: Global, Node
func (store *PlainStore) UpdateInput(input *graylog.Input) (*graylog.Input, error) {
	u, err := store.GetInput(input.ID)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, fmt.Errorf("the input <%s> is not found", input.ID)
	}
	u.Title = input.Title
	u.Type = input.Type
	u.Configuration = input.Configuration

	if input.Global != nil {
		u.Global = input.Global
	}
	if input.Node != "" {
		u.Node = input.Node
	}

	mutex.Lock()
	store.inputs[u.ID] = *u
	mutex.Unlock()
	return u, nil
}

// DeleteInput deletes an input from the store.
func (store *PlainStore) DeleteInput(id string) error {
	mutex.Lock()
	delete(store.inputs, id)
	mutex.Unlock()
	return nil
}

// GetInputs returns inputs.
func (store *PlainStore) GetInputs() ([]graylog.Input, int, error) {
	size := len(store.inputs)
	arr := make([]graylog.Input, size)
	i := 0
	for _, input := range store.inputs {
		arr[i] = input
		i++
	}
	return arr, size, nil
}
