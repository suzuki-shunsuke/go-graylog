package mockserver

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
)

// GET /system/inputs/{inputID} Get information of a single input on this node
func (ms *MockServer) handleGetInput(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
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
	id := ps.ByName("inputID")
	requiredFields := []string{"title", "type", "configuration"}
	allowedFields := []string{"global", "node"}
	sc, msg, body := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if sc != 200 {
		return sc, nil, fmt.Errorf(msg)
	}

	input := &graylog.Input{}
	if err := msDecode(body, input); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as Input")
		return 400, nil, err
	}

	ms.Logger().WithFields(log.Fields{
		"body": body, "input": input, "id": id,
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
	id := ps.ByName("inputID")
	if sc, err := ms.DeleteInput(id); err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	return 204, nil, nil
}

// POST /system/inputs Launch input on this node
func (ms *MockServer) handleCreateInput(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	requiredFields := []string{"title", "type", "configuration"}
	allowedFields := []string{"global", "node"}
	sc, msg, body := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if sc != 200 {
		return sc, nil, fmt.Errorf(msg)
	}

	input := &graylog.Input{}
	if err := msDecode(body, input); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as Input")
		return 400, nil, err
	}

	sc, err := ms.AddInput(input)
	if err != nil {
		return sc, nil, err
	}
	ms.safeSave()
	d := map[string]string{"id": input.ID}
	return 201, &d, nil
}

// GET /system/inputs Get all inputs
func (ms *MockServer) handleGetInputs(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
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