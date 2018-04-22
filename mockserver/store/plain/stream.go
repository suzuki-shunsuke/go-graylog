package plain

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
	st "github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

// HasStream
func (store *Store) HasStream(id string) (bool, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	_, ok := store.streams[id]
	return ok, nil
}

// GetStream returns a stream.
func (store *Store) GetStream(id string) (*graylog.Stream, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	s, ok := store.streams[id]
	if ok {
		return &s, nil
	}
	return nil, nil
}

// AddStream adds a stream to the store.
func (store *Store) AddStream(stream *graylog.Stream) error {
	if stream == nil {
		return fmt.Errorf("stream is nil")
	}
	if stream.ID == "" {
		stream.ID = st.NewObjectID()
	}

	store.imutex.Lock()
	defer store.imutex.Unlock()
	store.streams[stream.ID] = *stream
	return nil
}

// UpdateStream updates a stream at the store.
func (store *Store) UpdateStream(stream *graylog.Stream) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	orig, ok := store.streams[stream.ID]
	if !ok {
		return fmt.Errorf("the stream <%s> is not found", stream.ID)
	}
	stream.CreatorUserID = orig.CreatorUserID
	stream.CreatedAt = orig.CreatedAt
	stream.Disabled = orig.Disabled
	stream.IsDefault = orig.IsDefault
	if stream.Title == "" {
		stream.Title = orig.Title
	}
	if stream.IndexSetID == "" {
		stream.IndexSetID = orig.IndexSetID
	}
	if stream.Description == "" {
		stream.Description = orig.Description
	}
	if stream.Outputs == nil {
		stream.Description = orig.Description
	}
	if stream.MatchingType == "" {
		stream.MatchingType = orig.MatchingType
	}
	if stream.Rules == nil {
		stream.Rules = orig.Rules
	}
	if stream.AlertConditions == nil {
		stream.AlertConditions = orig.AlertConditions
	}
	if stream.AlertReceivers == nil {
		stream.AlertReceivers = orig.AlertReceivers
	}
	store.streams[stream.ID] = *stream
	return nil
}

// DeleteStream removes a stream from the store.
func (store *Store) DeleteStream(id string) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	delete(store.streams, id)
	return nil
}

// GetStreams returns a list of all streams.
func (store *Store) GetStreams() ([]graylog.Stream, int, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	total := len(store.streams)
	arr := make([]graylog.Stream, total)
	i := 0
	for _, index := range store.streams {
		arr[i] = index
		i++
	}
	return arr, total, nil
}

// GetEnabledStreams returns all enabled streams.
func (store *Store) GetEnabledStreams() ([]graylog.Stream, int, error) {
	arr := []graylog.Stream{}
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	for _, index := range store.streams {
		if index.Disabled {
			continue
		}
		arr = append(arr, index)
	}
	return arr, len(arr), nil
}
