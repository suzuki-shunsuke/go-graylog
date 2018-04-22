package graylog

func (attrs *InputJSONPathAttrs) InputType() string {
	return InputTypeJSONPath
}

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
