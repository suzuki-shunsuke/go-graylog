package graylog

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// HasInput
func (ms *MockServer) HasInput(id string) bool {
	_, ok := ms.inputs[id]
	return ok
}

// GetInput
func (ms *MockServer) GetInput(id string) (Input, bool) {
	s, ok := ms.inputs[id]
	return s, ok
}

// AddInput adds an input to the mock server.
func (ms *MockServer) AddInput(input *Input) (*Input, int, error) {
	if err := CreateValidator.Struct(input); err != nil {
		return nil, 400, err
	}
	s := *input
	s.Id = randStringBytesMaskImprSrc(24)
	ms.inputs[s.Id] = s
	return &s, 200, nil
}

// UpdateInput updates an input at the MockServer.
// Required: Title, Type, Configuration
// Allowed: Global, Node
func (ms *MockServer) UpdateInput(input *Input) (int, error) {
	u, ok := ms.GetInput(input.Id)
	if !ok {
		return 404, fmt.Errorf("The input is not found")
	}
	if err := UpdateValidator.Struct(input); err != nil {
		return 400, err
	}
	u.Title = input.Title
	u.Type = input.Type
	u.Configuration = &(*(input.Configuration))

	u.Global = input.Global
	u.Node = input.Node

	ms.inputs[u.Id] = u
	return 200, nil
}

// DeleteInput deletes a input from the mock server.
func (ms *MockServer) DeleteInput(id string) (int, error) {
	if !ms.HasInput(id) {
		return 404, fmt.Errorf("The input is not found")
	}
	delete(ms.inputs, id)
	return 200, nil
}

func (ms *MockServer) InputList() []Input {
	if ms.inputs == nil {
		return []Input{}
	}
	size := len(ms.inputs)
	arr := make([]Input, size)
	i := 0
	for _, input := range ms.inputs {
		arr[i] = input
		i++
	}
	return arr
}

// GET /system/inputs/{inputId} Get information of a single input on this node
func (ms *MockServer) handleGetInput(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	id := ps.ByName("inputId")
	input, ok := ms.GetInput(id)
	if !ok {
		writeApiError(w, 404, "No input found with id %s", id)
		return
	}
	writeOr500Error(w, &input)
}

// PUT /system/inputs/{inputId} Update input on this node
func (ms *MockServer) handleUpdateInput(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		write500Error(w)
		return
	}
	id := ps.ByName("inputId")
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

	input.Id = id
	if sc, err := ms.UpdateInput(input); err != nil {
		writeApiError(w, sc, err.Error())
		return
	}
	ms.safeSave()
	writeOr500Error(w, input)
}

// DELETE /system/inputs/{inputId} Terminate input on this node
func (ms *MockServer) handleDeleteInput(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	id := ps.ByName("inputId")
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
	d := map[string]string{"id": input.Id}
	writeOr500Error(w, &d)
}

// GET /system/inputs Get all inputs
func (ms *MockServer) handleGetInputs(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	ms.handleInit(w, r, false)
	arr := ms.InputList()
	inputs := &inputsBody{Inputs: arr, Total: len(arr)}
	writeOr500Error(w, inputs)
}
