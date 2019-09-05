package endpoint

// StreamAlertCondition returns Stream Alert Condition API's endpoint url.
func (ep *Endpoints) StreamAlertCondition(streamID, id string) string {
	// /streams/{streamId}/alerts/conditions/{conditionId}
	return ep.streams + "/" + streamID + "/alerts/conditions/" + id
}

// StreamAlertConditions returns Stream Alert Condition API's endpoint url.
func (ep *Endpoints) StreamAlertConditions(streamID string) string {
	// /streams/{streamId}/alerts/conditions
	return ep.streams + "/" + streamID + "/alerts/conditions"
}
