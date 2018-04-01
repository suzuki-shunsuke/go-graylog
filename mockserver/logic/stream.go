package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasStream
func (ms *Logic) HasStream(id string) (bool, error) {
	return ms.store.HasStream(id)
}

// GetStream returns a stream.
func (ms *Logic) GetStream(id string) (*graylog.Stream, int, error) {
	stream, err := ms.store.GetStream(id)
	if err != nil {
		return nil, 500, err
	}
	if stream == nil {
		return nil, 404, fmt.Errorf("no stream found with id <%s>", id)
	}
	return stream, 200, nil
}

// AddStream adds a stream to the Server.
func (ms *Logic) AddStream(stream *graylog.Stream) (int, error) {
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
func (ms *Logic) UpdateStream(stream *graylog.Stream) (int, error) {
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
func (ms *Logic) DeleteStream(id string) (int, error) {
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
func (ms *Logic) GetStreams() ([]graylog.Stream, int, error) {
	streams, err := ms.store.GetStreams()
	if err != nil {
		return nil, 500, err
	}
	return streams, 200, nil
}

// GetEnabledStreams returns all enabled streams.
func (ms *Logic) GetEnabledStreams() ([]graylog.Stream, int, error) {
	streams, err := ms.store.GetEnabledStreams()
	if err != nil {
		return nil, 500, err
	}
	return streams, 200, nil
}

// PauseStream pauses a stream.
func (ms *Logic) PauseStream(id string) (int, error) {
	ok, err := ms.HasStream(id)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream found with id <%s>", id)
	}
	// TODO pause
	return 200, nil
}

// ResumeStream resumes a stream.
func (ms *Logic) ResumeStream(id string) (int, error) {
	ok, err := ms.HasStream(id)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream found with id <%s>", id)
	}
	// TODO resume
	return 200, nil
}
