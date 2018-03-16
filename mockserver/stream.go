package mockserver

import (
	"fmt"

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

// GetStreams returns a list of all streams.
func (ms *MockServer) GetStreams() ([]graylog.Stream, error) {
	return ms.store.GetStreams()
}

// EnabledStreamList returns all enabled streams.
func (ms *MockServer) EnabledStreamList() ([]graylog.Stream, error) {
	return ms.store.GetEnabledStreams()
}
