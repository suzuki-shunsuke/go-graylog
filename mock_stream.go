package graylog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// AddStream adds a stream to the MockServer.
func (ms *MockServer) AddStream(stream *Stream) {
	if stream.Id == "" {
		stream.Id = randStringBytesMaskImprSrc(24)
	}
	ms.Streams[stream.Id] = *stream
	ms.safeSave()
}

// DeleteStream removes a stream from the MockServer.
func (ms *MockServer) DeleteStream(id string) {
	delete(ms.Streams, id)
	ms.safeSave()
}

// StreamList returns a list of all streams.
func (ms *MockServer) StreamList() []Stream {
	if ms.Streams == nil {
		return []Stream{}
	}
	arr := make([]Stream, len(ms.Streams))
	i := 0
	for _, index := range ms.Streams {
		arr[i] = index
		i++
	}
	return arr
}

// EnabledStreamList returns all enabled streams.
func (ms *MockServer) EnabledStreamList() []Stream {
	if ms.Streams == nil {
		return []Stream{}
	}
	arr := []Stream{}
	for _, index := range ms.Streams {
		if index.Disabled {
			continue
		}
		arr = append(arr, index)
	}
	return arr
}

// CreateStream
// {"type": "ApiError", "message": "Unable to map property id.\nKnown properties include: index_set_id, rules, title, description, content_pack, matching_type, remove_matches_from_default_stream"}
// not allowed id, creator_user_id, outputs, created_at, disabled, alert_conditions, alert_receivers, is_default
// Assigned index set must be writable!
func validateCreateStream(stream *Stream) (int, []byte) {
	key := ""
	switch {
	case stream.Id != "":
		key = "id"
	case stream.CreatorUserId != "":
		key = "creator_user_id"
	case stream.Outputs != nil && len(stream.Outputs) != 0:
		key = "outputs"
	case stream.CreatedAt != "":
		key = "created_at"
	case stream.Disabled:
		key = "disabled"
	case stream.AlertConditions != nil && len(stream.AlertConditions) != 0:
		key = "alert_conditions"
	case stream.AlertReceivers != nil:
		key = "alert_receivers"
	case stream.IsDefault:
		key = "is_default"
	}
	if key != "" {
		return 400, []byte(fmt.Sprintf(`{"type": "ApiError", "message": "Unable to map property %s.\nKnown properties include: index_set_id, rules, title, description, content_pack, matching_type, remove_matches_from_default_stream"}`, key))
	}
	if stream.Title == "" {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.resources.streams.requests.CreateStreamRequest, problem: Null title\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@53a6a093; line: 1, column: 2]" }`)
	}
	if stream.IndexSetId == "" {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.resources.streams.requests.CreateStreamRequest, problem: Null indexSetId\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@3b7194f4; line: 1, column: 17]"}`)
	}
	// 500, {"type": "ApiError", "message": "invalid hexadecimal representation of an ObjectId: [%s]"}

	return 200, []byte("")
}

// GET /streams Get a list of all streams
func (ms *MockServer) handleGetStreams(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	arr := ms.StreamList()
	streams := &streamsBody{Streams: arr, Total: len(arr)}
	writeOr500Error(w, streams)
}

// POST /streams Create index set
func (ms *MockServer) handleCreateStream(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		write500Error(w)
		return
	}
	stream := &Stream{}
	err = json.Unmarshal(b, stream)
	if err != nil {
		ms.Logger.WithFields(log.Fields{
			"body": string(b), "error": err,
		}).Info("Failed to parse request body as Stream")
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	ms.Logger.WithFields(log.Fields{
		"body": string(b), "stream": stream,
	}).Debug("request body")
	sc, msg := validateCreateStream(stream)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	ms.AddStream(stream)
	ret := map[string]string{"stream_id": stream.Id}
	writeOr500Error(w, ret)
}

// GET /streams/enabled Get a list of all enabled streams
func (ms *MockServer) handleGetEnabledStreams(
	w http.ResponseWriter, r *http.Request, _ httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	arr := ms.EnabledStreamList()
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
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	stream, ok := ms.Streams[id]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "No stream found with id %s"}`, id)))
		return
	}
	writeOr500Error(w, &stream)
}

// PUT /streams/{streamId} Update a stream
func (ms *MockServer) handleUpdateStream(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		write500Error(w)
		return
	}
	id := ps.ByName("streamId")
	stream, ok := ms.Streams[id]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "No stream found with id %s"}`, id)))
		return
	}
	data := map[string]interface{}{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	if title, ok := data["title"]; ok {
		t, ok := title.(string)
		if !ok {
			w.WriteHeader(400)
			w.Write([]byte(`{"message":"title must be string"}`))
			return
		}
		stream.Title = t
	}
	if description, ok := data["description"]; ok {
		d, ok := description.(string)
		if !ok {
			w.WriteHeader(400)
			w.Write([]byte(`{"message":"description must be string"}`))
			return
		}
		stream.Description = d
	}
	// outputs
	if matchingType, ok := data["matching_type"]; ok {
		m, ok := matchingType.(string)
		if !ok {
			w.WriteHeader(400)
			w.Write([]byte(`{"message":"matching_type must be string"}`))
			return
		}
		stream.MatchingType = m
	}
	if matchingType, ok := data["matching_type"]; ok {
		m, ok := matchingType.(string)
		if !ok {
			w.WriteHeader(400)
			w.Write([]byte(`{"message":"matching_type must be string"}`))
			return
		}
		stream.MatchingType = m
	}
	if removeMathcesFromDefaultStream, ok := data["remove_matches_from_default_stream"]; ok {
		m, ok := removeMathcesFromDefaultStream.(bool)
		if !ok {
			w.WriteHeader(400)
			w.Write([]byte(
				`{"message":"remove_matches_from_default_stream must be bool"}`))
			return
		}
		stream.RemoveMatchesFromDefaultStream = m
	}
	if indexSetId, ok := data["index_set_id"]; ok {
		m, ok := indexSetId.(string)
		if !ok {
			w.WriteHeader(400)
			w.Write([]byte(`{"message":"index_set_id must be string"}`))
			return
		}
		stream.IndexSetId = m
	}
	stream.Id = id
	ms.AddStream(&stream)
	writeOr500Error(w, &stream)
}

// DELETE /streams/{streamId} Delete a stream
func (ms *MockServer) handleDeleteStream(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("streamId")
	_, ok := ms.Streams[id]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "No stream found with id %s"}`, id)))
		return
	}
	ms.DeleteStream(id)
}

// POST /streams/{streamId}/pause Pause a stream
func (ms *MockServer) handlePauseStream(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("streamId")
	_, ok := ms.Streams[id]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "No stream found with id %s"}`, id)))
		return
	}
	// TODO pause
}

func (ms *MockServer) handleResumeStream(
	w http.ResponseWriter, r *http.Request, ps httprouter.Params,
) {
	ms.Logger.WithFields(log.Fields{
		"path": r.URL.Path, "method": r.Method,
	}).Info("request start")
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("streamId")
	_, ok := ms.Streams[id]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(
			`{"type": "ApiError", "message": "No stream found with id %s"}`, id)))
		return
	}
	// TODO resume
}
