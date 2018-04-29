package endpoint

import (
	"net/url"
	"path"
)

// StreamRules returns Stream Rules API's endpoint url.
func (ep *Endpoints) StreamRules(streamID string) (*url.URL, error) {
	// /streams/{streamid}/rules
	return urlJoin(ep.streams, path.Join(streamID, "rules"))
}

// StreamRuleTypes returns Stream Rule Types API's endpoint url.
func (ep *Endpoints) StreamRuleTypes(streamID string) (*url.URL, error) {
	// /streams/{streamid}/rules/types
	return urlJoin(ep.streams, path.Join(streamID, "rules/types"))
}

// StreamRule returns a Stream Rule API's endpoint url.
func (ep *Endpoints) StreamRule(streamID, streamRuleID string) (*url.URL, error) {
	// /streams/{streamid}/rules/{streamRuleID}
	return urlJoin(ep.streams, path.Join(streamID, "rules", streamRuleID))
}
