package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasInput returns whether the input exists.
func (lgc *Logic) HasInput(id string) (bool, error) {
	return lgc.store.HasInput(id)
}

// GetInput returns an input.
// If an input is not found, returns an error.
func (lgc *Logic) GetInput(id string) (*graylog.Input, int, error) {
	if id == "" {
		return nil, 400, fmt.Errorf("input id is empty")
	}
	if err := ValidateObjectID(id); err != nil {
		// unfortunately graylog returns not 400 but 404.
		return nil, 404, err
	}
	input, err := lgc.store.GetInput(id)
	if err != nil {
		return input, 500, err
	}
	if input == nil {
		return nil, 404, fmt.Errorf("no input with id <%s> is found", id)
	}
	return input, 200, nil
}

// AddInput adds an input to the mock server.
func (lgc *Logic) AddInput(input *graylog.Input) (int, error) {
	if err := validator.CreateValidator.Struct(input); err != nil {
		return 400, err
	}
	if err := lgc.store.AddInput(input); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateInput updates an input at the Server.
// Required: Title, Type, Attrs
// Allowed: Global, Node
func (lgc *Logic) UpdateInput(prms *graylog.InputUpdateParams) (*graylog.Input, int, error) {
	if prms == nil {
		return nil, 400, fmt.Errorf("input is nil")
	}
	if err := validator.UpdateValidator.Struct(prms); err != nil {
		return nil, 400, err
	}
	ok, err := lgc.HasInput(prms.ID)
	if err != nil {
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("the input <%s> is not found", prms.ID)
	}

	input, err := lgc.store.UpdateInput(prms)
	if err != nil {
		return nil, 500, err
	}
	return input, 200, nil
}

// DeleteInput deletes a input from the mock server.
func (lgc *Logic) DeleteInput(id string) (int, error) {
	ok, err := lgc.HasInput(id)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("lgc.HasInput() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("the input <%s> is not found", id)
	}
	if err := lgc.store.DeleteInput(id); err != nil {
		return 500, err
	}
	return 200, nil
}

// GetInputs returns a list of inputs.
func (lgc *Logic) GetInputs() ([]graylog.Input, int, int, error) {
	inputs, total, err := lgc.store.GetInputs()
	if err != nil {
		return nil, 0, 500, err
	}
	return inputs, total, 200, nil
}
