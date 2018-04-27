package endpoint

import (
	"net/url"
	"path"
)

// Streams returns a Stream API's endpoint url.
func (ep *Endpoints) Streams() string {
	return ep.streams.String()
}

// Stream returns a Stream API's endpoint url.
func (ep *Endpoints) Stream(id string) (*url.URL, error) {
	return urlJoin(ep.streams, id)
}

// PauseStream returns PauseStream API's endpoint url.
func (ep *Endpoints) PauseStream(id string) (*url.URL, error) {
	return urlJoin(ep.streams, path.Join(id, "pause"))
}

// ResumeStream returns ResumeStream API's endpoint url.
func (ep *Endpoints) ResumeStream(id string) (*url.URL, error) {
	return urlJoin(ep.streams, path.Join(id, "resume"))
}

// EnabledStreams returns GetEnabledStreams API's endpoint url.
func (ep *Endpoints) EnabledStreams() string {
	return ep.enabledStreams.String()
}
