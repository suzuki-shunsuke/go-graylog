package endpoint

import (
	"net/url"
	"path"
)

// StreamAlertCondition returns Stream Alert Condition API's endpoint url.
func (ep *Endpoints) StreamAlertCondition(streamID, id string) (*url.URL, error) {
	// /streams/{streamId}/alerts/conditions/{conditionId}
	return urlJoin(ep.streams, path.Join(streamID, "alerts/conditions", id))
}

// StreamAlertConditions returns Stream Alert Condition API's endpoint url.
func (ep *Endpoints) StreamAlertConditions(streamID string) (*url.URL, error) {
	// /streams/{streamId}/alerts/conditions
	return urlJoin(ep.streams, path.Join(streamID, "alerts/conditions"))
}
