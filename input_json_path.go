package graylog

const (
	InputTypeJSONPath string = "org.graylog2.inputs.misc.jsonpath.JsonPathInput"
)

// NewInputJSONPathAttrs is the constructor of InputJSONPathAttrs.
func NewInputJSONPathAttrs() InputAttrs {
	return &InputJSONPathAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputJSONPathAttrs) InputType() string {
	return InputTypeJSONPath
}

// InputJSONPathAttrs represents JSON path Input's attributes.
type InputJSONPathAttrs struct {
	ThrottlingAllowed bool   `json:"throttling_allowed,omitempty"`
	OverrideSource    string `json:"override_source,omitempty"`
	Headers           string `json:"headers,omitempty"`
	Path              string `json:"path,omitempty"`
	TargetURL         string `json:"target_url,omitempty"`
	Interval          int    `json:"interval,omitempty"`
	Source            string `json:"source,omitempty"`
	Timeunit          string `json:"timeunit,omitempty"`
}
