package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-set"
)

// GET /system/inputs/{inputID} Get information of a single input on this node
func HandleGetInput(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("inputID")
	if sc, err := ms.Authorize(user, "inputs:read", id); err != nil {
		return sc, nil, err
	}
	input, sc, err := ms.GetInput(id)
	return sc, input, err
}

// PUT /system/inputs/{inputID} Update input on this node
func HandleUpdateInput(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("inputID")
	if sc, err := ms.Authorize(user, "inputs:edit", id); err != nil {
		return sc, nil, err
	}
	requiredFields := set.NewStrSet("title", "type", "configuration")
	allowedFields := set.NewStrSet("global", "node")
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if err != nil {
		return sc, nil, err
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
	if err := ms.Save(); err != nil {
		return 500, nil, err
	}
	return 200, input, nil
}

// DELETE /system/inputs/{inputID} Terminate input on this node
func HandleDeleteInput(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("inputID")
	if sc, err := ms.Authorize(user, "inputs:terminate", id); err != nil {
		return sc, nil, err
	}
	if sc, err := ms.DeleteInput(id); err != nil {
		return sc, nil, err
	}
	if err := ms.Save(); err != nil {
		return 500, nil, err
	}
	return 204, nil, nil
}

// POST /system/inputs Launch input on this node
func HandleCreateInput(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	if sc, err := ms.Authorize(user, "inputs:create"); err != nil {
		return sc, nil, err
	}
	requiredFields := set.NewStrSet("title", "type", "configuration")
	allowedFields := set.NewStrSet("global", "node")
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if err != nil {
		return sc, nil, err
	}

	input := &graylog.Input{}
	if err := msDecode(body, input); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as Input")
		return 400, nil, err
	}

	sc, err = ms.AddInput(input)
	if err != nil {
		return sc, nil, err
	}
	if err := ms.Save(); err != nil {
		return 500, nil, err
	}
	d := map[string]string{"id": input.ID}
	return 201, &d, nil
}

// GET /system/inputs Get all inputs
func HandleGetInputs(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	arr, sc, err := ms.GetInputs()
	if err != nil {
		return sc, arr, err
	}
	inputs := &graylog.InputsBody{Inputs: arr, Total: len(arr)}
	return sc, inputs, nil
}
