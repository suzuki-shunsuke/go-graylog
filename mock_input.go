package graylog

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// HasInput
func (ms *MockServer) HasInput(id string) (bool, error) {
	return ms.store.HasInput(id)
}

// GetInput returns an input.
func (ms *MockServer) GetInput(id string) (*Input, error) {
	return ms.store.GetInput(id)
}

// AddInput adds an input to the mock server.
func (ms *MockServer) AddInput(input *Input) (*Input, int, error) {
	if err := CreateValidator.Struct(input); err != nil {
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
func (ms *MockServer) UpdateInput(input *Input) (int, error) {
	if err := UpdateValidator.Struct(input); err != nil {
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

func (ms *MockServer) InputList() ([]Input, error) {
	return ms.store.GetInputs()
}

// GET /system/inputs/{inputID} Get information of a single input on this node
func (ms *MockServer) handleGetInput(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	id := ps.ByName("inputID")
	input, err := ms.GetInput(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.GetInput() is failure")
		write500Error(w)
		return
	}
	if input == nil {
		writeApiError(w, 404, "No input found with id %s", id)
		return
	}
	writeOr500Error(w, input)
}

// PUT /system/inputs/{inputID} Update input on this node
func (ms *MockServer) handleUpdateInput(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err,
		}).Error("ms.handleInit() is failure")
		write500Error(w)
		return
	}
	id := ps.ByName("inputID")
	requiredFields := []string{"title", "type", "configuration"}
	allowedFields := []string{"global", "node"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	input := &Input{}
	if err := msDecode(body, input); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as Input")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	ms.Logger().WithFields(log.Fields{
		"body": string(b), "input": input, "id": id,
	}).Debug("request body")

	input.ID = id
	if sc, err := ms.UpdateInput(input); err != nil {
		writeApiError(w, sc, err.Error())
		return
	}
	ms.safeSave()
	writeOr500Error(w, input)
}

// DELETE /system/inputs/{inputID} Terminate input on this node
func (ms *MockServer) handleDeleteInput(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	id := ps.ByName("inputID")
	if sc, err := ms.DeleteInput(id); err != nil {
		writeApiError(w, sc, err.Error())
		return
	}
	ms.safeSave()
}

// POST /system/inputs Launch input on this node
func (ms *MockServer) handleCreateInput(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		write500Error(w)
		return
	}

	requiredFields := []string{"title", "type", "configuration"}
	allowedFields := []string{"global", "node"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	input := &Input{}
	if err := msDecode(body, input); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as Input")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	input, sc, err = ms.AddInput(input)
	if err != nil {
		writeApiError(w, sc, err.Error())
		return
	}
	ms.safeSave()
	d := map[string]string{"id": input.ID}
	writeOr500Error(w, &d)
}

// GET /system/inputs Get all inputs
func (ms *MockServer) handleGetInputs(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	ms.handleInit(w, r, false)
	arr, err := ms.InputList()
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err,
		}).Error("ms.InputList() is failure")
		write500Error(w)
	}
	inputs := &inputsBody{Inputs: arr, Total: len(arr)}
	writeOr500Error(w, inputs)
}
