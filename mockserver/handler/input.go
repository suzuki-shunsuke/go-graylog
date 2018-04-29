package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/util"
	"github.com/suzuki-shunsuke/go-set"
)

// HandleGetInput is the handler of Get an Input API.
func HandleGetInput(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /system/inputs/{inputID} Get information of a single input on this node
	id := ps.ByName("inputID")
	if sc, err := lgc.Authorize(user, "inputs:read", id); err != nil {
		return nil, sc, err
	}
	return lgc.GetInput(id)
}

// HandleGetInputs is the handler of Get Inputs API.
func HandleGetInputs(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// GET /system/inputs Get all inputs
	arr, total, sc, err := lgc.GetInputs()
	if err != nil {
		return arr, sc, err
	}
	inputs := &graylog.InputsBody{Inputs: arr, Total: total}
	return inputs, sc, nil
}

// HandleCreateInput is the handler of Create an Input API.
func HandleCreateInput(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// POST /system/inputs Launch input on this node
	if sc, err := lgc.Authorize(user, "inputs:create"); err != nil {
		return nil, sc, err
	}
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("title", "type", "configuration"),
			Optional:     set.NewStrSet("global", "node"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}
	// change configuration to attributes
	// https://github.com/Graylog2/graylog2-server/issues/3480
	body["attributes"] = body["configuration"]
	delete(body, "configuration")
	d := &graylog.InputData{}
	if err := util.MSDecode(body, d); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as InputData")
		return nil, 400, err
	}
	input := &graylog.Input{}
	if err := d.ToInput(input); err != nil {
		return nil, 400, err
	}
	sc, err = lgc.AddInput(input)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return &map[string]string{"id": input.ID}, 201, nil
}

// HandleUpdateInput is the handler of Update an Input API.
func HandleUpdateInput(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /system/inputs/{inputID} Update input on this node
	id := ps.ByName("inputID")
	if sc, err := lgc.Authorize(user, "inputs:edit", id); err != nil {
		return nil, sc, err
	}
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("title", "type", "configuration"),
			Optional:     set.NewStrSet("global", "node"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}
	// change configuration to attributes
	// https://github.com/Graylog2/graylog2-server/issues/3480
	body["attributes"] = body["configuration"]
	delete(body, "configuration")
	d := &graylog.InputUpdateParamsData{}
	if err := util.MSDecode(body, d); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as InputUpdateParamsData")
		return nil, 400, err
	}
	prms := &graylog.InputUpdateParams{}
	if err := d.ToInputUpdateParams(prms); err != nil {
		return nil, 400, err
	}

	lgc.Logger().WithFields(log.Fields{
		"body": body, "input": prms, "id": id,
	}).Debug("request body")

	prms.ID = id
	input, sc, err := lgc.UpdateInput(prms)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return input, 200, nil
}

// HandleDeleteInput is the handler of Delete an Input API.
func HandleDeleteInput(
	user *graylog.User, lgc *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// DELETE /system/inputs/{inputID} Terminate input on this node
	id := ps.ByName("inputID")
	if sc, err := lgc.Authorize(user, "inputs:terminate", id); err != nil {
		return nil, sc, err
	}
	if sc, err := lgc.DeleteInput(id); err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, 204, nil
}
