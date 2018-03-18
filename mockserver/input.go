package mockserver

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasInput
func (ms *Server) HasInput(id string) (bool, error) {
	return ms.store.HasInput(id)
}

// GetInput returns an input.
func (ms *Server) GetInput(id string) (*graylog.Input, error) {
	return ms.store.GetInput(id)
}

// AddInput adds an input to the mock server.
func (ms *Server) AddInput(input *graylog.Input) (int, error) {
	if err := validator.CreateValidator.Struct(input); err != nil {
		return 400, err
	}
	input.ID = randStringBytesMaskImprSrc(24)
	if err := ms.store.AddInput(input); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateInput updates an input at the Server.
// Required: Title, Type, Configuration
// Allowed: Global, Node
func (ms *Server) UpdateInput(input *graylog.Input) (int, error) {
	if err := validator.UpdateValidator.Struct(input); err != nil {
		return 400, err
	}
	if err := ms.store.UpdateInput(input); err != nil {
		return 500, err
	}
	return 200, nil
}

// DeleteInput deletes a input from the mock server.
func (ms *Server) DeleteInput(id string) (int, error) {
	ok, err := ms.HasInput(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.HasInput() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("The input is not found")
	}
	if err := ms.store.DeleteInput(id); err != nil {
		return 500, err
	}
	return 200, nil
}

func (ms *Server) InputList() ([]graylog.Input, error) {
	return ms.store.GetInputs()
}
