package graylog

// StreamRule represents a stream rule.
type StreamRule struct {
	// ex. "5a9b53c7c006c6000127f965"
	ID    string `json:"id,omitempty" v-create:"isdefault" v-update:"required,objectid"`
	Field string `json:"field,omitempty" v-create:"required" v-update:"required"`
	// ex. "5a94abdac006c60001f04fc1"
	StreamID    string `json:"stream_id,omitempty" v-create:"required" v-update:"required,objectid"`
	Description string `json:"description,omitempty"`
	Type        int    `json:"type,omitempty"`
	Inverted    *bool  `json:"inverted,omitempty"`
	Value       string `json:"value,omitempty" v-create:"required" v-update:"required"`
}

type StreamRulesBody struct {
	Total       int          `json:"total"`
	StreamRules []StreamRule `json:"stream_rules"`
}
