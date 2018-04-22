package graylog

// InputType is the implementation of the InputAttributes interface.
func (attrs *InputFakeHTTPMessageAttrs) InputType() string {
	return InputTypeFakeHTTPMessage
}

// InputFakeHTTPMessageAttrs
type InputFakeHTTPMessageAttrs struct {
	Sleep             int    `json:"sleep,omitempty"`
	SleepDeviation    int    `json:"sleep_deviation,omitempty"`
	Source            string `json:"source,omitempty"`
	OverrideSource    string `json:"override_source,omitempty"`
	ThrottlingAllowed bool   `json:"throttling_allowed,omitempty"`
}
