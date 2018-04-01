package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasInput
func (ms *Logic) HasInput(id string) (bool, error) {
	return ms.store.HasInput(id)
}

// GetInput returns an input.
// If an input is not found, returns an error.
func (ms *Logic) GetInput(id string) (*graylog.Input, int, error) {
	if id == "" {
		return nil, 400, fmt.Errorf("input id is empty")
	}
	// TODO id validation
	input, err := ms.store.GetInput(id)
	if err != nil {
		return input, 500, err
	}
	if input == nil {
		return nil, 404, fmt.Errorf("no input with id <%s> is found", id)
	}
	return input, 200, nil
}

// AddInput adds an input to the mock server.
func (ms *Logic) AddInput(input *graylog.Input) (int, error) {
	if err := validator.CreateValidator.Struct(input); err != nil {
		return 400, err
	}
	if err := ms.store.AddInput(input); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateInput updates an input at the Server.
// Required: Title, Type, Configuration
// Allowed: Global, Node
func (ms *Logic) UpdateInput(input *graylog.Input) (int, error) {
	if err := validator.UpdateValidator.Struct(input); err != nil {
		return 400, err
	}
	if err := ms.store.UpdateInput(input); err != nil {
		return 500, err
	}
	return 200, nil
}

// DeleteInput deletes a input from the mock server.
func (ms *Logic) DeleteInput(id string) (int, error) {
	ok, err := ms.HasInput(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.HasInput() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("the input <%s> is not found", id)
	}
	if err := ms.store.DeleteInput(id); err != nil {
		return 500, err
	}
	return 200, nil
}

// GetInputs returns a list of inputs.
func (ms *Logic) GetInputs() ([]graylog.Input, int, error) {
	inputs, err := ms.store.GetInputs()
	if err != nil {
		return inputs, 500, err
	}
	return inputs, 200, nil
}
