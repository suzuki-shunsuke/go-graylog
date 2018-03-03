package graylog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("inputId")
	input, ok := ms.Inputs[id]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "No input found with name %s"}`, id)))
		return
	}
	writeOr500Error(w, &input)
}

// PUT /system/inputs/{inputId} Update input on this node
func (ms *MockServer) handleUpdateInput(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		write500Error(w)
		return
	}
	id := ps.ByName("inputId")
	if _, ok := ms.Inputs[id]; !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "No input found with id %s"}`, id)))
		return
	}

	input := &Input{}
	err = json.Unmarshal(b, input)
	if err != nil {
		ms.Logger.WithFields(log.Fields{
			"body": string(b), "id": id, "error": err,
		}).Debug("Bad Request")
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}

	ms.Logger.WithFields(log.Fields{
		"body": string(b), "input": input, "id": id,
	}).Debug("request body")

	input.Id = id

	// {"type": "ApiError", "message": "Unable to map property id.\nKnown properties include: title, type, global, configuration, node"}
	sc, msg := validateInput(input)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	ms.AddInput(input)
	writeOr500Error(w, input)
}

// DELETE /system/inputs/{inputId} Terminate input on this node
func (ms *MockServer) handleDeleteInput(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("inputId")
	_, ok := ms.Inputs[id]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "No input found with id %s"}`, id)))
		return
	}
	ms.DeleteInput(id)
}

func validateInput(input *Input) (int, []byte) {
	// Required
	// type, title configuration.bind_address, configuration.port
	// configuration.recv_buffer_size
	if input.Type == "" {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.models.system.inputs.requests.InputCreateRequest, problem: Null type\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@107566a4; line: 1, column: 17]"}`)
	}
	if input.Title == "" {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.models.system.inputs.requests.InputCreateRequest, problem: Null title\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@320397d1; line: 8, column: 1]"}`)
	}
	if input.Configuration == nil {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.models.system.inputs.requests.InputCreateRequest, problem: Null configuration\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@3d687f1; line: 1, column: 30]"}`)
	}
	if input.Configuration.BindAddress == "" {
		return 400, []byte(
			`{"type": "ApiError", "message": "Missing or invalid input configuration."}`)
	}
	if input.Configuration.Port == 0 {
		return 400, []byte(
			`{"type": "ApiError", "message": "Missing or invalid input configuration."}`)
	}
	if input.Configuration.RecvBufferSize == 0 {
		return 400, []byte(
			`{"type": "ApiError", "message": "Missing or invalid input configuration."}`)
	}
	// node optional
	// skip type validation
	// 404 {"type": "ApiError", "message": "There is no such input type registered."}
	return 200, []byte("")
}

// POST /system/inputs Launch input on this node
func (ms *MockServer) handleCreateInput(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		write500Error(w)
		return
	}
	input := &Input{}
	err = json.Unmarshal(b, input)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	sc, msg := validateInput(input)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
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
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	arr := ms.InputList()
	inputs := &inputsBody{Inputs: arr, Total: len(arr)}
	writeOr500Error(w, inputs)
}
