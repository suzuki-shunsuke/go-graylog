package client

import (
	"context"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetPipelines returns all pipelines.
func (client *Client) GetPipelines() ([]graylog.Pipeline, *ErrorInfo, error) {
	return client.GetPipelinesContext(context.Background())
}

// GetPipelinesContext returns all pipeline with a context.
func (client *Client) GetPipelinesContext(ctx context.Context) (
	[]graylog.Pipeline, *ErrorInfo, error,
) {
	pipe := []graylog.Pipeline{}
	ei, err := client.callGet(ctx, client.Endpoints().Pipelines(), nil, &pipe)
	return pipe, ei, err
}

// GetPipeline returns a pipeline.
func (client *Client) GetPipeline(id string) (*graylog.Pipeline, *ErrorInfo, error) {
	return client.GetPipelineContext(context.Background(), id)
}

// GetPipelineContext returns a pipeline with a context.
func (client *Client) GetPipelineContext(ctx context.Context, id string) (
	*graylog.Pipeline, *ErrorInfo, error,
) {
	u, err := client.Endpoints().Pipeline(id)
	if err != nil {
		return nil, nil, err
	}
	pipe := &graylog.Pipeline{}
	ei, err := client.callGet(ctx, u.String(), nil, pipe)
	return pipe, ei, err
}

// CreatePipeline creates a pipeline.
func (client *Client) CreatePipeline(
	pipeline *graylog.Pipeline,
) (*ErrorInfo, error) {
	return client.CreatePipelineContext(context.Background(), pipeline)
}

// CreatePipelineContext creates a pipeline with a context.
func (client *Client) CreatePipelineContext(
	ctx context.Context, pipeline *graylog.Pipeline,
) (*ErrorInfo, error) {
	return client.callPost(
		ctx, client.Endpoints().Pipelines(), pipeline, &pipeline)
}

// UpdatePipeline updates a pipeline.
func (client *Client) UpdatePipeline(
	pipeline *graylog.Pipeline,
) (*ErrorInfo, error) {
	return client.UpdatePipelineContext(context.Background(), pipeline)
}

// UpdatePipelineContext updates a pipeline with a context.
func (client *Client) UpdatePipelineContext(
	ctx context.Context, pipeline *graylog.Pipeline,
) (*ErrorInfo, error) {
	u, err := client.Endpoints().Pipeline(pipeline.ID)
	if err != nil {
		return nil, err
	}
	return client.callPut(ctx, u.String(), map[string]interface{}{
		"source":      pipeline.Source,
		"description": pipeline.Description,
	}, pipeline)
}

// DeletePipeline deletes a pipeline.
func (client *Client) DeletePipeline(id string) (*ErrorInfo, error) {
	return client.DeletePipelineContext(context.Background(), id)
}

// DeletePipelineContext deletes a pipeline with a context.
func (client *Client) DeletePipelineContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	u, err := client.Endpoints().Pipeline(id)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
