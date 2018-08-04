package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasStream returns whether the stream exists.
func (lgc *Logic) HasStream(id string) (bool, error) {
	return lgc.store.HasStream(id)
}

// GetStream returns a stream.
func (lgc *Logic) GetStream(id string) (*graylog.Stream, int, error) {
	if err := ValidateObjectID(id); err != nil {
		// unfortunately graylog returns not 400 but 404.
		return nil, 404, err
	}
	stream, err := lgc.store.GetStream(id)
	if err != nil {
		return nil, 500, err
	}
	if stream == nil {
		return nil, 404, fmt.Errorf("no stream found with id <%s>", id)
	}
	return stream, 200, nil
}

// AddStream adds a stream to the Server.
func (lgc *Logic) AddStream(stream *graylog.Stream) (int, error) {
	if err := validator.CreateValidator.Struct(stream); err != nil {
		return 400, err
	}
	// check index set existence
	is, sc, err := lgc.GetIndexSet(stream.IndexSetID)
	if err != nil {
		LogWE(sc, lgc.Logger().WithFields(log.Fields{
			"error": err, "index_set_id": stream.IndexSetID, "status_code": sc,
		}), "failed to get an index set")
		return sc, err
	}
	if !is.Writable {
		return 400, fmt.Errorf("assigned index set must be writable")
	}
	if err := lgc.store.AddStream(stream); err != nil {
		return 500, err
	}
	return 200, nil
}

// UpdateStream updates a stream at the Server.
func (lgc *Logic) UpdateStream(prms *graylog.StreamUpdateParams) (*graylog.Stream, int, error) {
	if prms == nil {
		return nil, 400, fmt.Errorf("stream is nil")
	}
	if err := validator.UpdateValidator.Struct(prms); err != nil {
		return nil, 400, err
	}
	stream, sc, err := lgc.GetStream(prms.ID)
	if err != nil {
		LogWE(sc, lgc.Logger().WithFields(log.Fields{
			"error": err, "id": prms.ID, "status_code": sc,
		}), "failed to get a stream")
		return nil, sc, err
	}
	if stream.IsDefault {
		return nil, 400, fmt.Errorf("the default stream cannot be edited")
	}
	// check index set existence
	if prms.IndexSetID != "" {
		is, sc, err := lgc.GetIndexSet(prms.IndexSetID)
		if err != nil {
			LogWE(sc, lgc.Logger().WithFields(log.Fields{
				"error": err, "index_set_id": prms.IndexSetID, "status_code": sc,
			}), "failed to get an index set")
			return nil, sc, err
		}
		if !is.Writable {
			return nil, 400, fmt.Errorf("assigned index set must be writable")
		}
	}
	s, err := lgc.store.UpdateStream(prms)
	if err != nil {
		return nil, 500, err
	}
	return s, 200, nil
}

// DeleteStream deletes a stream from the Server.
func (lgc *Logic) DeleteStream(id string) (int, error) {
	ok, err := lgc.HasStream(id)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("lgc.HasStream() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream found with id <%s>", id)
	}
	if err := lgc.store.DeleteStream(id); err != nil {
		return 500, err
	}
	return 200, nil
}

// GetStreams returns a list of all streams.
func (lgc *Logic) GetStreams() ([]graylog.Stream, int, int, error) {
	streams, total, err := lgc.store.GetStreams()
	if err != nil {
		return nil, 0, 500, err
	}
	return streams, total, 200, nil
}

// GetEnabledStreams returns all enabled streams.
func (lgc *Logic) GetEnabledStreams() ([]graylog.Stream, int, int, error) {
	streams, total, err := lgc.store.GetEnabledStreams()
	if err != nil {
		return nil, 0, 500, err
	}
	return streams, total, 200, nil
}

// PauseStream pauses a stream.
func (lgc *Logic) PauseStream(id string) (int, error) {
	ok, err := lgc.HasStream(id)
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
func (lgc *Logic) ResumeStream(id string) (int, error) {
	ok, err := lgc.HasStream(id)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("no stream found with id <%s>", id)
	}
	// TODO resume
	return 200, nil
}
