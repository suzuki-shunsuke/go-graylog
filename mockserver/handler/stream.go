package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-set"
)

// GET /streams Get a list of all streams
func HandleGetStreams(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	arr, err := ms.GetStreams()
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err,
		}).Error("ms.StreamList() is failure")
		return 500, nil, err
	}

	streams := &graylog.StreamsBody{Streams: arr, Total: len(arr)}
	return 200, streams, nil
}

// POST /streams Create index set
func HandleCreateStream(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	requiredFields := set.NewStrSet("title", "index_set_id")
	allowedFields := set.NewStrSet(
		"rules", "description", "content_pack",
		"matching_type", "remove_matches_from_default_stream")
	body, sc, err := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if err != nil {
		return sc, nil, err
	}

	stream := &graylog.Stream{}
	if err := msDecode(body, stream); err != nil {
		ms.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as stream")
		return 400, nil, err
	}

	sc, err = ms.AddStream(stream)
	if err != nil {
		return 400, nil, err
	}
	ret := map[string]string{"stream_id": stream.ID}
	return 201, ret, nil
}

// GET /streams/enabled Get a list of all enabled streams
func HandleGetEnabledStreams(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	arr, err := ms.EnabledStreamList()
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err,
		}).Error("ms.EnabledStreamList() is failure")
		return 500, nil, err
	}
	streams := &graylog.StreamsBody{Streams: arr, Total: len(arr)}
	return 200, streams, nil
}

// GET /streams/{streamID} Get a single stream
func HandleGetStream(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("streamID")
	if id == "enabled" {
		return HandleGetEnabledStreams(user, ms, w, r, ps)
	}
	stream, err := ms.GetStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.GetStream() is failure")
		return 500, nil, err
	}
	if stream == nil {
		return 404, nil, fmt.Errorf("no stream found with id <%s>", id)
	}
	return 200, stream, nil
}

// PUT /streams/{streamID} Update a stream
func HandleUpdateStream(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("streamID")
	stream, err := ms.GetStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.GetStream() is failure")
		return 500, nil, err
	}
	if stream == nil {
		return 404, nil, fmt.Errorf("no stream found with id <%s>", id)
	}
	data := map[string]interface{}{}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&data); err != nil {
		return 400, nil, err
	}
	if title, ok := data["title"]; ok {
		t, ok := title.(string)
		if !ok {
			return 400, nil, fmt.Errorf("title must be string")
		}
		stream.Title = t
	}
	if description, ok := data["description"]; ok {
		d, ok := description.(string)
		if !ok {
			return 400, nil, fmt.Errorf("description must be string")
		}
		stream.Description = d
	}
	// TODO outputs
	if matchingType, ok := data["matching_type"]; ok {
		m, ok := matchingType.(string)
		if !ok {
			return 400, nil, fmt.Errorf("matching_type must be string")
		}
		stream.MatchingType = m
	}
	if removeMathcesFromDefaultStream, ok := data["remove_matches_from_default_stream"]; ok {
		m, ok := removeMathcesFromDefaultStream.(bool)
		if !ok {
			return 400, nil, fmt.Errorf("remove_matches_from_default_stream must be bool")
		}
		stream.RemoveMatchesFromDefaultStream = m
	}
	if indexSetID, ok := data["index_set_id"]; ok {
		m, ok := indexSetID.(string)
		if !ok {
			return 400, nil, fmt.Errorf("index_set_id must be string")
		}
		stream.IndexSetID = m
	}
	stream.ID = id
	if sc, err := ms.UpdateStream(stream); err != nil {
		return sc, nil, err
	}
	ms.SafeSave()
	return 200, stream, nil
}

// DELETE /streams/{streamID} Delete a stream
func HandleDeleteStream(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("streamID")
	ok, err := ms.HasStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.HasStream() is failure")
		return 500, nil, err
	}

	if !ok {
		return 404, nil, fmt.Errorf("no stream found with id <%s>", id)
	}
	ms.DeleteStream(id)
	ms.SafeSave()
	return 204, nil, nil
}

// POST /streams/{streamID}/pause Pause a stream
func HandlePauseStream(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("streamID")
	ok, err := ms.HasStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.HasStream() is failure")
		return 500, nil, err
	}
	if !ok {
		return 404, nil, fmt.Errorf("no stream found with id <%s>", id)
	}
	// TODO pause
	return 200, nil, nil
}

func HandleResumeStream(
	user *graylog.User, ms *logic.Server,
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("streamID")
	ok, err := ms.HasStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.HasStream() is failure")
		return 500, nil, err
	}
	if !ok {
		return 404, nil, fmt.Errorf("no stream found with id <%s>", id)
	}
	// TODO resume
	return 200, nil, nil
}
