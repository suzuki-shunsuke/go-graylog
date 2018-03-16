package mockserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasStream
func (ms *MockServer) HasStream(id string) (bool, error) {
	return ms.store.HasStream(id)
}

// GetStream returns a stream.
func (ms *MockServer) GetStream(id string) (*graylog.Stream, error) {
	return ms.store.GetStream(id)
}

// AddStream adds a stream to the MockServer.
func (ms *MockServer) AddStream(stream *graylog.Stream) (int, error) {
	if err := validator.CreateValidator.Struct(stream); err != nil {
		return 400, err
	}
	stream.ID = randStringBytesMaskImprSrc(24)
	if err := ms.store.AddStream(stream); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateStream updates a stream at the MockServer.
func (ms *MockServer) UpdateStream(stream *graylog.Stream) (int, error) {
	ok, err := ms.HasStream(stream.ID)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": stream.ID,
		}).Error("ms.HasStream() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("No stream found with id %s", stream.ID)
	}
	if err := validator.UpdateValidator.Struct(stream); err != nil {
		return 400, err
	}
	if err := ms.store.UpdateStream(stream); err != nil {
		return 500, err
	}
	return 200, nil
}

// DeleteStream deletes a stream from the MockServer.
func (ms *MockServer) DeleteStream(id string) (int, error) {
	ok, err := ms.HasStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.HasStream() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("No stream found with id %s", id)
	}
	if err := ms.store.DeleteStream(id); err != nil {
		return 500, err
	}
	return 200, nil
}

// StreamList returns a list of all streams.
func (ms *MockServer) StreamList() ([]graylog.Stream, error) {
	return ms.store.GetStreams()
}

// EnabledStreamList returns all enabled streams.
func (ms *MockServer) EnabledStreamList() ([]graylog.Stream, error) {
	return ms.store.GetEnabledStreams()
}

// GET /streams Get a list of all streams
func (ms *MockServer) handleGetStreams(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	arr, err := ms.StreamList()
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
func (ms *MockServer) handleCreateStream(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) (int, interface{}, error) {
	requiredFields := []string{"title", "index_set_id"}
	allowedFields := []string{
		"rules", "description", "content_pack",
		"matching_type", "remove_matches_from_default_stream"}
	sc, msg, body := validateRequestBody(r.Body, requiredFields, allowedFields, nil)
	if sc != 200 {
		return sc, nil, fmt.Errorf(msg)
	}

	stream := &graylog.Stream{}
	if err := msDecode(body, stream); err != nil {
		ms.logger.WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as stream")
		return 400, nil, err
	}

	sc, err := ms.AddStream(stream)
	if err != nil {
		return 400, nil, err
	}
	ret := map[string]string{"stream_id": stream.ID}
	return 201, ret, nil
}

// GET /streams/enabled Get a list of all enabled streams
func (ms *MockServer) handleGetEnabledStreams(
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
func (ms *MockServer) handleGetStream(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) (int, interface{}, error) {
	id := ps.ByName("streamID")
	if id == "enabled" {
		return ms.handleGetEnabledStreams(w, r, ps)
	}
	stream, err := ms.GetStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.GetStream() is failure")
		return 500, nil, err
	}
	if stream == nil {
		return 404, nil, fmt.Errorf("No stream found with id %s", id)
	}
	return 200, stream, nil
}

// PUT /streams/{streamID} Update a stream
func (ms *MockServer) handleUpdateStream(
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
		return 404, nil, fmt.Errorf("No stream found with id %s", id)
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
	ms.safeSave()
	return 200, stream, nil
}

// DELETE /streams/{streamID} Delete a stream
func (ms *MockServer) handleDeleteStream(
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
		return 404, nil, fmt.Errorf("No stream found with id %s", id)
	}
	ms.DeleteStream(id)
	ms.safeSave()
	return 204, nil, nil
}

// POST /streams/{streamID}/pause Pause a stream
func (ms *MockServer) handlePauseStream(
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
		return 404, nil, fmt.Errorf("No stream found with id %s", id)
	}
	// TODO pause
	return 200, nil, nil
}

func (ms *MockServer) handleResumeStream(
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
		return 404, nil, fmt.Errorf("No stream found with id %s", id)
	}
	// TODO resume
	return 200, nil, nil
}
