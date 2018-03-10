package graylog

// HasStream
func (store *InMemoryStore) HasStream(id string) (bool, error) {
	_, ok := store.streams[id]
	return ok, nil
}

// GetStream returns a stream.
func (store *InMemoryStore) GetStream(id string) (Stream, bool, error) {
	s, ok := store.streams[id]
	return s, ok, nil
}

// AddStream adds a stream to the store.
func (store *InMemoryStore) AddStream(stream *Stream) (*Stream, int, error) {
	store.streams[stream.ID] = *stream
	return stream, 200, nil
}

// UpdateStream updates a stream at the store.
func (store *InMemoryStore) UpdateStream(stream *Stream) (int, error) {
	store.streams[stream.ID] = *stream
	return 200, nil
}

// DeleteStream removes a stream from the store.
func (store *InMemoryStore) DeleteStream(id string) (int, error) {
	delete(store.streams, id)
	return 200, nil
}

// GetStreams returns a list of all streams.
func (store *InMemoryStore) GetStreams() ([]Stream, error) {
	arr := make([]Stream, len(store.streams))
	i := 0
	for _, index := range store.streams {
		arr[i] = index
		i++
	}
	return arr, nil
}

// GetEnabledStreams returns all enabled streams.
func (store *InMemoryStore) GetEnabledStreams() ([]Stream, error) {
	arr := []Stream{}
	for _, index := range store.streams {
		if index.Disabled {
			continue
		}
		arr = append(arr, index)
	}
	return arr, nil
}
