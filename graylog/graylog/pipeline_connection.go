package graylog

type (
	// PipelineConnection is a pipeline connection.
	PipelineConnection struct {
		ID          string   `json:"id"`
		StreamID    string   `json:"stream_id"`
		PipelineIDs []string `json:"pipeline_ids"`
	}
)
