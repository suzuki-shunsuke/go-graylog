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

// HandleGetStreams
func HandleGetStreams(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// GET /streams Get a list of all streams
	arr, total, sc, err := ms.GetStreams()
	if err != nil {
		return nil, sc, err
	}

	return &graylog.StreamsBody{Streams: arr, Total: total}, sc, nil
}

// HandleGetStream
func HandleGetStream(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /streams/{streamID} Get a single stream
	id := ps.ByName("streamID")
	if id == "enabled" {
		return HandleGetEnabledStreams(user, ms, w, r, ps)
	}
	if sc, err := ms.Authorize(user, "streams:read", id); err != nil {
		return nil, sc, err
	}
	return ms.GetStream(id)
}

// HandleCreateStream
func HandleCreateStream(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// POST /streams Create index set
	if sc, err := ms.Authorize(user, "streams:create"); err != nil {
		return nil, sc, err
	}
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required: set.NewStrSet("title", "index_set_id"),
			Optional: set.NewStrSet("rules", "description", "content_pack", "matching_type", "remove_matches_from_default_stream"),
		})
	if err != nil {
		return nil, sc, err
	}

	stream := &graylog.Stream{}
	if err := util.MSDecode(body, stream); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as stream")
		return nil, 400, err
	}

	sc, err = ms.AddStream(stream)
	if err != nil {
		return nil, 400, err
	}
	return map[string]string{"stream_id": stream.ID}, 201, nil
}

// HandleUpdateStream
func HandleUpdateStream(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /streams/{streamID} Update a stream
	stream := &graylog.Stream{ID: ps.ByName("streamID")}
	if sc, err := ms.Authorize(user, "streams:edit", stream.ID); err != nil {
		return nil, sc, err
	}

	// requiredFields is nothing
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Optional: set.NewStrSet(
				"title", "index_set_id", "description", "outputs", "matching_type",
				"rules", "alert_conditions", "alert_receivers",
				"remove_matches_from_default_stream"),
		})
	if err != nil {
		return nil, sc, err
	}

	if err := util.MSDecode(body, stream); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Warn("Failed to parse request body as stream")
		return nil, 400, err
	}

	if sc, err := ms.UpdateStream(stream); err != nil {
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return stream, 200, nil
}

// HandleDeleteStream
func HandleDeleteStream(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// DELETE /streams/{streamID} Delete a stream
	id := ps.ByName("streamID")
	// TODO authorization
	sc, err := ms.DeleteStream(id)
	if err != nil {
		return nil, sc, err
	}
	if err := ms.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}

// HandleGetEnabledStreams
func HandleGetEnabledStreams(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// GET /streams/enabled Get a list of all enabled streams
	arr, total, sc, err := ms.GetEnabledStreams()
	if err != nil {
		return nil, sc, err
	}
	return &graylog.StreamsBody{Streams: arr, Total: total}, sc, nil
}

// HandlePauseStream
func HandlePauseStream(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// POST /streams/{streamID}/pause Pause a stream
	id := ps.ByName("streamID")
	if sc, err := ms.Authorize(user, "streams:changestate", id); err != nil {
		return nil, sc, err
	}
	sc, err := ms.PauseStream(id)
	return nil, sc, err
}

// HandleResumeStream
func HandleResumeStream(
	user *graylog.User, ms *logic.Logic,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	id := ps.ByName("streamID")
	if sc, err := ms.Authorize(user, "streams:changestate", id); err != nil {
		return nil, sc, err
	}
	sc, err := ms.ResumeStream(id)
	return nil, sc, err
}
