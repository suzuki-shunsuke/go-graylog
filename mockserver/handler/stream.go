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

// HandleGetStreams is the handler of Get Streams API.
func HandleGetStreams(
	user *graylog.User, lgc *logic.Logic,
	r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// GET /streams Get a list of all streams
	arr, total, sc, err := lgc.GetStreams()
	if err != nil {
		return nil, sc, err
	}

	return &graylog.StreamsBody{Streams: arr, Total: total}, sc, nil
}

// HandleGetStream is the handler of Get a Stream API.
func HandleGetStream(
	user *graylog.User, lgc *logic.Logic,
	r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// GET /streams/{streamID} Get a single stream
	id := ps.ByName("streamID")
	if id == "enabled" {
		return HandleGetEnabledStreams(user, lgc, r, ps)
	}
	if sc, err := lgc.Authorize(user, "streams:read", id); err != nil {
		return nil, sc, err
	}
	return lgc.GetStream(id)
}

// HandleCreateStream is the handler of Create a Stream API.
func HandleCreateStream(
	user *graylog.User, lgc *logic.Logic,
	r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// POST /streams Create index set
	if sc, err := lgc.Authorize(user, "streams:create"); err != nil {
		return nil, sc, err
	}
	// empty description is ignored
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required:     set.NewStrSet("title", "index_set_id"),
			Optional:     set.NewStrSet("rules", "description", "content_pack", "matching_type", "remove_matches_from_default_stream"),
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}

	stream := &graylog.Stream{}
	if err := util.MSDecode(body, stream); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as stream")
		return nil, 400, err
	}

	sc, err = lgc.AddStream(stream)
	if err != nil {
		return nil, sc, err
	}
	return map[string]string{"stream_id": stream.ID}, sc, nil
}

// HandleUpdateStream is the handler of Update a Stream API.
func HandleUpdateStream(
	user *graylog.User, lgc *logic.Logic,
	r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// PUT /streams/{streamID} Update a stream
	prms := &graylog.StreamUpdateParams{ID: ps.ByName("streamID")}
	if sc, err := lgc.Authorize(user, "streams:edit", prms.ID); err != nil {
		return nil, sc, err
	}

	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required: nil,
			Optional: set.NewStrSet(
				"title", "index_set_id", "description", "outputs", "matching_type",
				"rules", "alert_conditions", "alert_receivers",
				"remove_matches_from_default_stream"),
			Ignored:      set.NewStrSet("creator_user_id", "created_at", "disabled", "is_default"),
			ExtForbidden: false,
		})
	if err != nil {
		return nil, sc, err
	}

	if err := util.MSDecode(body, prms); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Warn("Failed to parse request body as stream")
		return nil, 400, err
	}

	stream, sc, err := lgc.UpdateStream(prms)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return stream, sc, nil
}

// HandleDeleteStream is the handler of Delete a Stream API.
func HandleDeleteStream(
	user *graylog.User, lgc *logic.Logic,
	r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// DELETE /streams/{streamID} Delete a stream
	id := ps.ByName("streamID")
	// TODO authorization
	sc, err := lgc.DeleteStream(id)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}

// HandleGetEnabledStreams is the handler of Get all enabled streams API.
func HandleGetEnabledStreams(
	user *graylog.User, lgc *logic.Logic,
	r *http.Request, _ httprouter.Params,
) (interface{}, int, error) {
	// GET /streams/enabled Get a list of all enabled streams
	arr, total, sc, err := lgc.GetEnabledStreams()
	if err != nil {
		return nil, sc, err
	}
	return &graylog.StreamsBody{Streams: arr, Total: total}, sc, nil
}

// HandlePauseStream is the handler of Pause a Stream API.
func HandlePauseStream(
	user *graylog.User, lgc *logic.Logic,
	r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	// POST /streams/{streamID}/pause Pause a stream
	id := ps.ByName("streamID")
	if sc, err := lgc.Authorize(user, "streams:changestate", id); err != nil {
		return nil, sc, err
	}
	sc, err := lgc.PauseStream(id)
	return nil, sc, err
}

// HandleResumeStream is the handler of Resume a Stream API.
func HandleResumeStream(
	user *graylog.User, lgc *logic.Logic,
	r *http.Request, ps httprouter.Params,
) (interface{}, int, error) {
	id := ps.ByName("streamID")
	if sc, err := lgc.Authorize(user, "streams:changestate", id); err != nil {
		return nil, sc, err
	}
	sc, err := lgc.ResumeStream(id)
	return nil, sc, err
}
