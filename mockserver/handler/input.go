package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-set"
)

// HandleGetInput
func HandleGetInput(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /system/inputs/{inputID} Get information of a single input on this node
	id := ps.ByName("inputID")
	if sc, err := ms.Authorize(user, "inputs:read", id); err != nil {
		return nil, sc, err
	}
	return ms.GetInput(id)
}

// HandleGetInputs
func HandleGetInputs(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// GET /system/inputs Get all inputs
	arr, total, sc, err := ms.GetInputs()
	if err != nil {
		return arr, sc, err
	}
	inputs := &graylog.InputsBody{Inputs: arr, Total: total}
	return inputs, sc, nil
}

// HandleCreateInput
func HandleCreateInput(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// POST /system/inputs Launch input on this node
	if sc, err := ms.Authorize(user, "inputs:create"); err != nil {
		return nil, sc, err
	}
	requiredFields := set.NewStrSet("title", "type", "configuration")
	allowedFields := set.NewStrSet("global", "node")
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if err != nil {
		return nil, sc, err
	}

	input := &graylog.Input{}
	if err := msDecode(body, input); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as Input")
		return nil, 400, err
	}

	sc, err = ms.AddInput(input)
	if err != nil {
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	d := map[string]string{"id": input.ID}
	return &d, 201, nil
}

// HandleUpdateInput
func HandleUpdateInput(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /system/inputs/{inputID} Update input on this node
	id := ps.ByName("inputID")
	if sc, err := ms.Authorize(user, "inputs:edit", id); err != nil {
		return nil, sc, err
	}
	requiredFields := set.NewStrSet("title", "type", "configuration")
	allowedFields := set.NewStrSet("global", "node")
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if err != nil {
		return nil, sc, err
	}

	input := &graylog.Input{}
	if err := msDecode(body, input); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as Input")
		return nil, 400, err
	}

	ms.Logger().WithFields(log.Fields{
		"body": body, "input": input, "id": id,
	}).Debug("request body")

	input.ID = id
	u, sc, err := ms.UpdateInput(input)
	if err != nil {
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return u, 200, nil
}

// HandleDeleteInput
func HandleDeleteInput(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// DELETE /system/inputs/{inputID} Terminate input on this node
	id := ps.ByName("inputID")
	if sc, err := ms.Authorize(user, "inputs:terminate", id); err != nil {
		return nil, sc, err
	}
	if sc, err := ms.DeleteInput(id); err != nil {
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return nil, 204, nil
}
