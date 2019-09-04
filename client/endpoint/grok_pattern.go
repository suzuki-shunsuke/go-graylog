package endpoint

// GrokPatterns returns a GrokPattern API's endpoint url.
func (ep *Endpoints) GrokPatterns() string {
	return ep.grokPatterns
}

// GrokPattern returns a GrokPattern API's endpoint url.
func (ep *Endpoints) GrokPattern(id string) string {
	return ep.grokPatterns + "/" + id
}

// GrokPatternMembers returns /system/grok/test endpoint url.
func (ep *Endpoints) GrokPatternTest() string {
	return ep.grokPatterns + "/test"
}
