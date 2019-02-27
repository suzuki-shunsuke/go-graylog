package graylog

// PipelineRule represents a Graylog's Pipeline Rule.
// http://docs.graylog.org/en/3.0/pages/pipelines/rules.html
type PipelineRule struct {
	// required
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Source      string `json:"source,omitempty" v-create:"required"`
	ID          string `json:"id,omitempty" v-create:"isdefault"`
	// CreatedAt   string `json:"created_at,omitempty"`
	// ModifiedAt   string `json:"modified_at,omitempty"`
	// Errors string `json:"errors"`
}
