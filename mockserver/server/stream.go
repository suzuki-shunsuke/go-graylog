package server

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasStream
func (ms *Server) HasStream(id string) (bool, error) {
	return ms.store.HasStream(id)
}

// GetStream returns a stream.
func (ms *Server) GetStream(id string) (*graylog.Stream, error) {
	return ms.store.GetStream(id)
}

// AddStream adds a stream to the Server.
func (ms *Server) AddStream(stream *graylog.Stream) (int, error) {
	if err := validator.CreateValidator.Struct(stream); err != nil {
		return 400, err
	}
	stream.ID = randStringBytesMaskImprSrc(24)
	if err := ms.store.AddStream(stream); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateStream updates a stream at the Server.
func (ms *Server) UpdateStream(stream *graylog.Stream) (int, error) {
	if stream == nil {
		return 400, fmt.Errorf("stream is nil")
	}
	if err := validator.UpdateValidator.Struct(stream); err != nil {
		return 400, err
	}
	ok, err := ms.HasStream(stream.ID)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": stream.ID,
		}).Error("ms.HasStream() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream found with id <%s>", stream.ID)
	}
	if err := ms.store.UpdateStream(stream); err != nil {
		return 500, err
	}
	return 200, nil
}

// DeleteStream deletes a stream from the Server.
func (ms *Server) DeleteStream(id string) (int, error) {
	ok, err := ms.HasStream(id)
	if err != nil {
		ms.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("ms.HasStream() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream found with id <%s>", id)
	}
	if err := ms.store.DeleteStream(id); err != nil {
		return 500, err
	}
	return 200, nil
}

// GetStreams returns a list of all streams.
func (ms *Server) GetStreams() ([]graylog.Stream, error) {
	return ms.store.GetStreams()
}

// EnabledStreamList returns all enabled streams.
func (ms *Server) EnabledStreamList() ([]graylog.Stream, error) {
	return ms.store.GetEnabledStreams()
}
