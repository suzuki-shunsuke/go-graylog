package mockserver

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasInput
func (ms *MockServer) HasInput(id string) (bool, error) {
	return ms.store.HasInput(id)
}

// GetInput returns an input.
func (ms *MockServer) GetInput(id string) (*graylog.Input, error) {
	return ms.store.GetInput(id)
}

// AddInput adds an input to the mock server.
func (ms *MockServer) AddInput(input *graylog.Input) (*graylog.Input, int, error) {
	if err := validator.CreateValidator.Struct(input); err != nil {
		return nil, 400, err
	}
	s := *input
	s.ID = randStringBytesMaskImprSrc(24)
	i, err := ms.store.AddInput(&s)
	if err != nil {
		return nil, 500, err
	}
	return i, 200, nil
}

// UpdateInput updates an input at the MockServer.
// Required: Title, Type, Configuration
// Allowed: Global, Node
func (ms *MockServer) UpdateInput(input *graylog.Input) (int, error) {
	if err := validator.UpdateValidator.Struct(input); err != nil {
		return 400, err
	}
	if err := ms.store.UpdateInput(input); err != nil {
		return 500, err
	}
	return 200, nil
}

// DeleteInput deletes a input from the mock server.
func (ms *MockServer) DeleteInput(id string) (int, error) {
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

func (ms *MockServer) InputList() ([]graylog.Input, error) {
	return ms.store.GetInputs()
}

// GET /system/inputs/{inputID} Get information of a single input on this node
func (ms *MockServer) handleGetInput(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	ms.handleInit(w, r, false)
	id := ps.ByName("inputID")
	input, err := ms.GetInput(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.GetInput() is failure")
		return 500, nil, err
	}
	if input == nil {
		return 404, nil, fmt.Errorf("No input found with id %s", id)
	}
	return 200, input, nil
}

// PUT /system/inputs/{inputID} Update input on this node
func (ms *MockServer) handleUpdateInput(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err,
		}).Error("ms.handleInit() is failure")
		return 500, nil, err
	}
	id := ps.ByName("inputID")
	requiredFields := []string{"title", "type", "configuration"}
	allowedFields := []string{"global", "node"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		return sc, nil, fmt.Errorf(msg)
	}

	input := &graylog.Input{}
	if err := msDecode(body, input); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as Input")
		return 400, nil, err
	}

	ms.Logger().WithFields(log.Fields{
		"body": string(b), "input": input, "id": id,
	}).Debug("request body")

	input.ID = id
	if sc, err := ms.UpdateInput(input); err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 200, input, nil
}

// DELETE /system/inputs/{inputID} Terminate input on this node
func (ms *MockServer) handleDeleteInput(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	ms.handleInit(w, r, false)
	id := ps.ByName("inputID")
	if sc, err := ms.DeleteInput(id); err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 200, nil, nil
}

// POST /system/inputs Launch input on this node
func (ms *MockServer) handleCreateInput(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		return 500, nil, err
	}

	requiredFields := []string{"title", "type", "configuration"}
	allowedFields := []string{"global", "node"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		return sc, nil, fmt.Errorf(msg)
	}

	input := &graylog.Input{}
	if err := msDecode(body, input); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as Input")
		return 400, nil, err
	}

	input, sc, err = ms.AddInput(input)
	if err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	d := map[string]string{"id": input.ID}
	return 200, &d, nil
}

// GET /system/inputs Get all inputs
func (ms *MockServer) handleGetInputs(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	ms.handleInit(w, r, false)
	arr, err := ms.InputList()
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err,
		}).Error("ms.InputList() is failure")
		return 500, nil, err
	}
	inputs := &graylog.InputsBody{Inputs: arr, Total: len(arr)}
	return 200, inputs, nil
}
