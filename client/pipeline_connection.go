package client

import (
	"context"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

// GetPipelineConnections returns all pipeline connections.
func (client *Client) GetPipelineConnections(ctx context.Context) (
	[]graylog.PipelineConnection, *ErrorInfo, error,
) {
	// GET /plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/connections
	conns := []graylog.PipelineConnection{}
	ei, err := client.callGet(ctx, client.Endpoints().PipelineConnections(), nil, &conns)
	return conns, ei, err
}

// GetPipeline returns a pipeline connection for a given stream.
func (client *Client) GetPipelineConnectionsOfStream(ctx context.Context, id string) (
	*graylog.PipelineConnection, *ErrorInfo, error,
) {
	conn := &graylog.PipelineConnection{}
	ei, err := client.callGet(ctx, client.Endpoints().PipelineConnectionsOfStream(id), nil, conn)
	return conn, ei, err
}

// ConnectStreamsToPipeline connects streams to a pipeline.
func (client *Client) ConnectStreamsToPipeline(
	ctx context.Context, pipelineID string, streamIDs []string,
) ([]graylog.PipelineConnection, *ErrorInfo, error) {
	// POST /plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/connections/to_pipeline Connect streams to a processing pipeline
	conns := []graylog.PipelineConnection{}
	ei, err := client.callPost(
		ctx, client.Endpoints().ConnectStreamsToPipeline(), map[string]interface{}{
			"pipeline_id": pipelineID,
			"stream_ids":  streamIDs,
		}, &conns)
	return conns, ei, err
}

// ConnectStreamsToPipeline connects processing pipelines to a stream.
func (client *Client) ConnectPipelinesToStream(
	ctx context.Context, conn *graylog.PipelineConnection,
) (*ErrorInfo, error) {
	// POST /plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/connections/to_stream Connect processing pipelines to a stream
	return client.callPost(
		ctx, client.Endpoints().ConnectPipelinesToStream(), conn, conn)
}
