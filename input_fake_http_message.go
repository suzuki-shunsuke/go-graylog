package graylog

const (
	// InputTypeFakeHTTPMessage is one of input types.
	InputTypeFakeHTTPMessage string = "org.graylog2.inputs.random.FakeHttpMessageInput"
)

// NewInputFakeHTTPMessageAttrs is the constructor of InputFakeHTTPMessageAttrs.
func NewInputFakeHTTPMessageAttrs() InputAttrs {
	return &InputFakeHTTPMessageAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputFakeHTTPMessageAttrs) InputType() string {
	return InputTypeFakeHTTPMessage
}

// InputFakeHTTPMessageAttrs represents fake HTTP message Input's attributes.
type InputFakeHTTPMessageAttrs struct {
	Sleep             int    `json:"sleep,omitempty"`
	SleepDeviation    int    `json:"sleep_deviation,omitempty"`
	Source            string `json:"source,omitempty"`
	OverrideSource    string `json:"override_source,omitempty"`
	ThrottlingAllowed bool   `json:"throttling_allowed,omitempty"`
}
