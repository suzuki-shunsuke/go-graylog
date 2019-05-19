package client

import (
	"context"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetPipelineConnections returns all pipeline connections.
func (client *Client) GetPipelineConnections() ([]graylog.PipelineConnection, *ErrorInfo, error) {
	return client.GetPipelineConnectionsContext(context.Background())
}

// GetPipelineConnectionsContext returns all pipeline connections with a context.
func (client *Client) GetPipelineConnectionsContext(ctx context.Context) (
	[]graylog.PipelineConnection, *ErrorInfo, error,
) {
	// GET /plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/connections
	conns := []graylog.PipelineConnection{}
	ei, err := client.callGet(ctx, client.Endpoints().PipelineConnections(), nil, &conns)
	return conns, ei, err
}

// GetPipeline returns a pipeline connection for a given stream.
func (client *Client) GetPipelineConnectionsOfStream(id string) (
	*graylog.PipelineConnection, *ErrorInfo, error,
) {
	return client.GetPipelineConnectionsOfStreamContext(context.Background(), id)
}

// GetPipelineContext returns a pipeline connection for a given stream with a context.
func (client *Client) GetPipelineConnectionsOfStreamContext(ctx context.Context, id string) (
	*graylog.PipelineConnection, *ErrorInfo, error,
) {
	u, err := client.Endpoints().PipelineConnectionsOfStream(id)
	if err != nil {
		return nil, nil, err
	}
	conn := &graylog.PipelineConnection{}
	ei, err := client.callGet(ctx, u.String(), nil, conn)
	return conn, ei, err
}

// ConnectStreamsToPipeline connects streams to a pipeline.
func (client *Client) ConnectStreamsToPipeline(
	pipelineID string, streamIDs []string,
) ([]graylog.PipelineConnection, *ErrorInfo, error) {
	return client.ConnectStreamsToPipelineContext(context.Background(), pipelineID, streamIDs)
}

// ConnectStreamsToPipeline connects streams to a pipeline with a context.
func (client *Client) ConnectStreamsToPipelineContext(
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
	conn *graylog.PipelineConnection,
) (*ErrorInfo, error) {
	return client.ConnectPipelinesToStreamContext(context.Background(), conn)
}

// ConnectStreamsToPipelineContext connects processing pipelines to a stream with a context.
func (client *Client) ConnectPipelinesToStreamContext(
	ctx context.Context, conn *graylog.PipelineConnection,
) (*ErrorInfo, error) {
	// POST /plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/connections/to_stream Connect processing pipelines to a stream
	return client.callPost(
		ctx, client.Endpoints().ConnectPipelinesToStream(), conn, conn)
}
