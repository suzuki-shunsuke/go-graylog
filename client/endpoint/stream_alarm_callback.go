package endpoint

import (
	"net/url"
	"path"
)

// StreamAlarmCallback returns Stream Alarm Callback API's endpoint url.
func (ep *Endpoints) StreamAlarmCallback(streamID, id string) (*url.URL, error) {
	// /streams/{streamid}/alarmcallbacks/{alarmCallbackId}
	return urlJoin(ep.streams, path.Join(streamID, "alarmcallbacks", id))
}

// StreamAlarmCallbacks returns Stream Alarm Callback API's endpoint url.
func (ep *Endpoints) StreamAlarmCallbacks(streamID string) (*url.URL, error) {
	// /streams/{streamid}/alarmcallbacks
	return urlJoin(ep.streams, path.Join(streamID, "alarmcallbacks"))
}
