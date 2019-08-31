package endpoint

import (
	"net/url"
)

// GrokPatterns returns a GrokPattern API's endpoint url.
func (ep *Endpoints) GrokPatterns() string {
	return ep.grokPatterns.String()
}

// GrokPattern returns a GrokPattern API's endpoint url.
func (ep *Endpoints) GrokPattern(id string) (*url.URL, error) {
	return urlJoin(ep.grokPatterns, id)
}

// GrokPatternMembers returns /system/grok/test endpoint url.
func (ep *Endpoints) GrokPatternTest() string {
	return ep.roles.String() + "/test"
}
