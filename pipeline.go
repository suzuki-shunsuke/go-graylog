package graylog

// Pipeline represents a Graylog's Pipeline.
// http://docs.graylog.org/en/3.0/pages/pipelines/pipelines.html
type Pipeline struct {
	// required
	Source string          `json:"source,omitempty" v-create:"required" v-update:"required"`
	ID     string          `json:"id,omitempty" v-create:"isdefault" v-update:"required"`
	Title  string          `json:"title,omitempty"`
	Stages []PipelineStage `json:"stages"`

	Description string `json:"description,omitempty"`
	// CreatedAt   string `json:"created_at,omitempty"`
	// ModifiedAt   string `json:"modified_at,omitempty"`
	// Errors string `json:"errors"`
}

// PipelineStage is a stage of pipelines.
type PipelineStage struct {
	Stage    int      `json:"stage"`
	MatchAll bool     `json:"match_all"`
	Rules    []string `json:"rules"`
}
