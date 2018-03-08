package graylog

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func (ms *MockServer) AddInput(input *Input) {
	if input.Id == "" {
		input.Id = randStringBytesMaskImprSrc(24)
	}
	ms.Inputs[input.Id] = *input
	ms.safeSave()
}

func (ms *MockServer) DeleteInput(id string) {
	delete(ms.Inputs, id)
	ms.safeSave()
}

func (ms *MockServer) InputList() []Input {
	if ms.Inputs == nil {
		return []Input{}
	}
	size := len(ms.Inputs)
	arr := make([]Input, size)
	i := 0
	for _, input := range ms.Inputs {
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
	input, ok := ms.Inputs[id]
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
	if _, ok := ms.Inputs[id]; !ok {
		writeApiError(w, 404, "No input found with id %s", id)
		return
	}

	requiredFields := []string{"title", "type", "configuration"}
	allowedFields := []string{
		"title", "type", "global", "configuration", "node"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	input := &Input{}
	if err := msDecode(body, input); err != nil {
		ms.Logger.WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as Input")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	input.Id = id
	if err := UpdateValidator.Struct(input); err != nil {
		writeApiError(w, 400, err.Error())
		return
	}

	ms.Logger.WithFields(log.Fields{
		"body": string(b), "input": input, "id": id,
	}).Debug("request body")
	ms.AddInput(input)
	writeOr500Error(w, input)
}

// DELETE /system/inputs/{inputId} Terminate input on this node
func (ms *MockServer) handleDeleteInput(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	id := ps.ByName("inputId")
	_, ok := ms.Inputs[id]
	if !ok {
		writeApiError(w, 404, "No input found with id %s", id)
		return
	}
	ms.DeleteInput(id)
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
	allowedFields := []string{
		"title", "type", "global", "configuration", "node"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	input := &Input{}
	if err := msDecode(body, input); err != nil {
		ms.Logger.WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as Input")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	if err := CreateValidator.Struct(input); err != nil {
		writeApiError(w, 400, err.Error())
		return
	}

	ms.AddInput(input)
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
