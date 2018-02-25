package graylog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

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
func (ms *MockServer) handleGetInput(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("inputId")
	input, ok := ms.Inputs[id]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(`{"type": "ApiError", "message": "No input found with name %s"}`, id)))
		return
	}
	b, err := json.Marshal(&input)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

// PUT /system/inputs/{inputId} Update input on this node
func (ms *MockServer) handleUpdateInput(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	id := ps.ByName("inputId")
	if _, ok := ms.Inputs[id]; !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(`{"type": "ApiError", "message": "No input found with id %s"}`, id)))
		return
	}
	input := Input{}
	err = json.Unmarshal(b, &input)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	sc, msg := validateInput(&input)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	delete(ms.Inputs, id)
	ms.Inputs[input.Id] = input
	b, err = json.Marshal(&input)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

// DELETE /system/inputs/{inputId} Terminate input on this node
func (ms *MockServer) handleDeleteInput(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("inputId")
	_, ok := ms.Inputs[id]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(`{"type": "ApiError", "message": "No input found with id %s"}`, id)))
		return
	}
	delete(ms.Inputs, id)
}

func validateInput(input *Input) (int, []byte) {
	if input.Type == "" {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.models.system.inputs.requests.InputCreateRequest, problem: Null type\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@107566a4; line: 1, column: 17]"}`)
	}
	if input.Configuration == nil {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.models.system.inputs.requests.InputCreateRequest, problem: Null configuration\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@3d687f1; line: 1, column: 30]"}`)
	}
	if input.Node == "" {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.models.system.inputs.requests.InputCreateRequest, problem: Null node\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@3d687f1; line: 1, column: 30]"}`)
	}
	// skip configuration validation
	// 400 {"type": "ApiError", "message": "Missing or invalid input configuration."}
	// skip type validation
	// 404 {"type": "ApiError", "message": "There is no such input type registered."}
	return 200, []byte("")
}

// POST /system/inputs Launch input on this node
func (ms *MockServer) handleCreateInput(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	input := Input{}
	err = json.Unmarshal(b, &input)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	// generate id 24 ex: 5a90cee5c006c60001efbbf5
	input.Id = randStringBytesMaskImprSrc(24)
	sc, msg := validateInput(&input)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	ms.Inputs[input.Id] = input
	d := map[string]string{"id": input.Id}
	b, err = json.Marshal(&d)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

// GET /system/inputs Get all inputs
func (ms *MockServer) handleGetInputs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	arr := ms.InputList()
	inputs := inputsBody{Inputs: arr, Total: len(arr)}
	b, err := json.Marshal(&inputs)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}
