package client

import (
	"context"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetPipelines returns all pipeline.
func (client *Client) GetPipelines(ctx context.Context) (
	[]graylog.Pipeline, *ErrorInfo, error,
) {
	pipe := []graylog.Pipeline{}
	ei, err := client.callGet(ctx, client.Endpoints().Pipelines(), nil, &pipe)
	return pipe, ei, err
}

// GetPipeline returns a pipeline.
func (client *Client) GetPipeline(ctx context.Context, id string) (
	*graylog.Pipeline, *ErrorInfo, error,
) {
	pipe := &graylog.Pipeline{}
	ei, err := client.callGet(ctx, client.Endpoints().Pipeline(id), nil, pipe)
	return pipe, ei, err
}

// CreatePipeline creates a pipeline.
func (client *Client) CreatePipeline(
	ctx context.Context, pipeline *graylog.Pipeline,
) (*ErrorInfo, error) {
	return client.callPost(
		ctx, client.Endpoints().Pipelines(), pipeline, &pipeline)
}

// UpdatePipeline updates a pipeline.
func (client *Client) UpdatePipeline(
	ctx context.Context, pipeline *graylog.Pipeline,
) (*ErrorInfo, error) {
	return client.callPut(ctx, client.Endpoints().Pipeline(pipeline.ID), map[string]interface{}{
		"source":      pipeline.Source,
		"description": pipeline.Description,
	}, pipeline)
}

// DeletePipeline deletes a pipeline.
func (client *Client) DeletePipeline(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	return client.callDelete(ctx, client.Endpoints().Pipeline(id), nil, nil)
}
