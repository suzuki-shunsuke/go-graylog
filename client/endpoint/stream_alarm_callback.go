package endpoint

// StreamAlarmCallback returns Stream Alarm Callback API's endpoint url.
func (ep *Endpoints) StreamAlarmCallback(streamID, id string) string {
	// /streams/{streamid}/alarmcallbacks/{alarmCallbackId}
	return ep.streams + "/" + streamID + "/alarmcallbacks/" + id
}

// StreamAlarmCallbacks returns Stream Alarm Callback API's endpoint url.
func (ep *Endpoints) StreamAlarmCallbacks(streamID string) string {
	// /streams/{streamid}/alarmcallbacks
	return ep.streams + "/" + streamID + "/alarmcallbacks"
}
