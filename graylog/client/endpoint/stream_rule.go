package endpoint

// StreamRules returns Stream Rules API's endpoint url.
func (ep *Endpoints) StreamRules(streamID string) string {
	// /streams/{streamid}/rules
	return ep.streams + "/" + streamID + "/rules"
}

// StreamRuleTypes returns Stream Rule Types API's endpoint url.
func (ep *Endpoints) StreamRuleTypes(streamID string) string {
	// /streams/{streamid}/rules/types
	return ep.streams + "/" + streamID + "/rules/types"
}

// StreamRule returns a Stream Rule API's endpoint url.
func (ep *Endpoints) StreamRule(streamID, streamRuleID string) string {
	// /streams/{streamid}/rules/{streamRuleID}
	return ep.streams + "/" + streamID + "/rules/" + streamRuleID
}
