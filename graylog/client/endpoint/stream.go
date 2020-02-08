package endpoint

// Streams returns a Stream API's endpoint url.
func (ep *Endpoints) Streams() string {
	return ep.streams
}

// Stream returns a Stream API's endpoint url.
func (ep *Endpoints) Stream(id string) string {
	return ep.streams + "/" + id
}

// PauseStream returns PauseStream API's endpoint url.
func (ep *Endpoints) PauseStream(id string) string {
	return ep.streams + "/" + id + "/pause"
}

// ResumeStream returns ResumeStream API's endpoint url.
func (ep *Endpoints) ResumeStream(id string) string {
	return ep.streams + "/" + id + "/resume"
}

// EnabledStreams returns GetEnabledStreams API's endpoint url.
func (ep *Endpoints) EnabledStreams() string {
	return ep.enabledStreams
}
