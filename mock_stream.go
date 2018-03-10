package graylog

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// HasStream
func (ms *MockServer) HasStream(id string) (bool, error) {
	return ms.store.HasStream(id)
}

// GetStream returns a stream.
func (ms *MockServer) GetStream(id string) (Stream, bool, error) {
	return ms.store.GetStream(id)
}

// AddStream adds a stream to the MockServer.
func (ms *MockServer) AddStream(stream *Stream) (*Stream, int, error) {
	if err := CreateValidator.Struct(stream); err != nil {
		return nil, 400, err
	}
	s := *stream
	s.Id = randStringBytesMaskImprSrc(24)
	return ms.store.AddStream(&s)
}

// UpdateStream updates a stream at the MockServer.
func (ms *MockServer) UpdateStream(stream *Stream) (int, error) {
	ok, err := ms.HasStream(stream.Id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": stream.Id,
		}).Error("ms.HasStream() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("No stream found with id %s", stream.Id)
	}
	if err := UpdateValidator.Struct(stream); err != nil {
		return 400, err
	}
	return ms.store.UpdateStream(stream)
}

// DeleteStream removes a stream from the MockServer.
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
	return ms.store.DeleteStream(id)
}

// StreamList returns a list of all streams.
func (ms *MockServer) StreamList() ([]Stream, error) {
	return ms.store.GetStreams()
}

// EnabledStreamList returns all enabled streams.
func (ms *MockServer) EnabledStreamList() ([]Stream, error) {
	return ms.store.GetEnabledStreams()
}

// GET /streams Get a list of all streams
func (ms *MockServer) handleGetStreams(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	ms.handleInit(w, r, false)
	arr, err := ms.StreamList()
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err,
		}).Error("ms.StreamList() is failure")
		writeApiError(w, 500, err.Error())
		return
	}

	streams := &streamsBody{Streams: arr, Total: len(arr)}
	writeOr500Error(w, streams)
}

// POST /streams Create index set
func (ms *MockServer) handleCreateStream(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		write500Error(w)
		return
	}

	requiredFields := []string{"title", "index_set_id"}
	allowedFields := []string{
		"rules", "description", "content_pack",
		"matching_type", "remove_matches_from_default_stream"}
	sc, msg, body := validateRequestBody(b, requiredFields, allowedFields, nil)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write([]byte(msg))
		return
	}

	stream := &Stream{}
	if err := msDecode(body, stream); err != nil {
		ms.logger.WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as stream")
		writeApiError(w, 400, "400 Bad Request")
		return
	}

	s, sc, err := ms.AddStream(stream)
	if err != nil {
		writeApiError(w, 400, err.Error())
		return
	}
	ret := map[string]string{"stream_id": s.Id}
	writeOr500Error(w, ret)
}

// GET /streams/enabled Get a list of all enabled streams
func (ms *MockServer) handleGetEnabledStreams(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	ms.handleInit(w, r, false)
	arr, err := ms.EnabledStreamList()
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err,
		}).Error("ms.EnabledStreamList() is failure")
		writeApiError(w, 500, err.Error())
		return
	}
	streams := &streamsBody{Streams: arr, Total: len(arr)}
	writeOr500Error(w, streams)
}

// GET /streams/{streamId} Get a single stream
func (ms *MockServer) handleGetStream(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	id := ps.ByName("streamId")
	if id == "enabled" {
		ms.handleGetEnabledStreams(w, r, ps)
		return
	}
	ms.handleInit(w, r, false)
	stream, ok, err := ms.GetStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.GetStream() is failure")
		writeApiError(w, 500, err.Error())
		return
	}
	if !ok {
		writeApiError(w, 404, "No stream found with id %s", id)
		return
	}
	writeOr500Error(w, &stream)
}

// PUT /streams/{streamId} Update a stream
func (ms *MockServer) handleUpdateStream(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	b, err := ms.handleInit(w, r, true)
	if err != nil {
		write500Error(w)
		return
	}
	id := ps.ByName("streamId")
	stream, ok, err := ms.GetStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.GetStream() is failure")
		writeApiError(w, 500, err.Error())
		return
	}
	if !ok {
		writeApiError(w, 404, "No stream found with id %s", id)
		return
	}
	data := map[string]interface{}{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		writeApiError(w, 400, "400 Bad Request")
		return
	}
	if title, ok := data["title"]; ok {
		t, ok := title.(string)
		if !ok {
			writeApiError(w, 400, "title must be string")
			return
		}
		stream.Title = t
	}
	if description, ok := data["description"]; ok {
		d, ok := description.(string)
		if !ok {
			writeApiError(w, 400, "description must be string")
			return
		}
		stream.Description = d
	}
	// TODO outputs
	if matchingType, ok := data["matching_type"]; ok {
		m, ok := matchingType.(string)
		if !ok {
			writeApiError(w, 400, "matching_type must be string")
			return
		}
		stream.MatchingType = m
	}
	if removeMathcesFromDefaultStream, ok := data["remove_matches_from_default_stream"]; ok {
		m, ok := removeMathcesFromDefaultStream.(bool)
		if !ok {
			writeApiError(w, 400, "remove_matches_from_default_stream must be bool")
			return
		}
		stream.RemoveMatchesFromDefaultStream = m
	}
	if indexSetId, ok := data["index_set_id"]; ok {
		m, ok := indexSetId.(string)
		if !ok {
			writeApiError(w, 400, "index_set_id must be string")
			return
		}
		stream.IndexSetId = m
	}
	stream.Id = id
	if sc, err := ms.UpdateStream(&stream); err != nil {
		writeApiError(w, sc, err.Error())
		return
	}
	ms.safeSave()
	writeOr500Error(w, &stream)
}

// DELETE /streams/{streamId} Delete a stream
func (ms *MockServer) handleDeleteStream(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	id := ps.ByName("streamId")
	ok, err := ms.HasStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.HasStream() is failure")
		writeApiError(w, 500, err.Error())
		return
	}

	if !ok {
		writeApiError(w, 404, "No stream found with id %s", id)
		return
	}
	ms.DeleteStream(id)
	ms.safeSave()
}

// POST /streams/{streamId}/pause Pause a stream
func (ms *MockServer) handlePauseStream(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	id := ps.ByName("streamId")
	ok, err := ms.HasStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.HasStream() is failure")
		writeApiError(w, 500, err.Error())
		return
	}
	if !ok {
		writeApiError(w, 404, "No stream found with id %s", id)
		return
	}
	// TODO pause
}

func (ms *MockServer) handleResumeStream(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.handleInit(w, r, false)
	id := ps.ByName("streamId")
	ok, err := ms.HasStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.HasStream() is failure")
		writeApiError(w, 500, err.Error())
		return
	}
	if !ok {
		writeApiError(w, 404, "No stream found with id %s", id)
		return
	}
	// TODO resume
}
