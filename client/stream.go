package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetStreams returns all streams.
func (client *Client) GetStreams(
	ctx context.Context,
) (streams []graylog.Stream, total int, ei *ErrorInfo, err error) {
	streamsBody := &graylog.StreamsBody{}
	ei, err = client.callGet(
		ctx, client.Endpoints().Streams(), nil, streamsBody)
	return streamsBody.Streams, streamsBody.Total, ei, err
}

// GetStream returns a given stream.
func (client *Client) GetStream(
	ctx context.Context, id string,
) (*graylog.Stream, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	stream := &graylog.Stream{}
	ei, err := client.callGet(ctx, client.Endpoints().Stream(id), nil, stream)
	return stream, ei, err
}

// CreateStream creates a stream.
func (client *Client) CreateStream(
	ctx context.Context, stream *graylog.Stream,
) (*ErrorInfo, error) {
	if stream == nil {
		return nil, errors.New("stream is nil")
	}
	ret := map[string]string{}
	ei, err := client.callPost(ctx, client.Endpoints().Streams(), stream, &ret)
	if err != nil {
		return ei, err
	}
	if id, ok := ret["stream_id"]; ok {
		stream.ID = id
		return ei, nil
	}
	return ei, errors.New(`response doesn't have the field "stream_id"`)
}

// GetEnabledStreams returns all enabled streams.
func (client *Client) GetEnabledStreams(
	ctx context.Context,
) (streams []graylog.Stream, total int, ei *ErrorInfo, err error) {
	streamsBody := &graylog.StreamsBody{}
	ei, err = client.callGet(
		ctx, client.Endpoints().EnabledStreams(), nil, streamsBody)
	return streamsBody.Streams, streamsBody.Total, ei, err
}

// UpdateStream updates a stream.
func (client *Client) UpdateStream(
	ctx context.Context, stream *graylog.Stream,
) (*ErrorInfo, error) {
	if stream == nil {
		return nil, errors.New("stream is nil")
	}
	if stream.ID == "" {
		return nil, errors.New("id is empty")
	}
	body := *stream
	body.ID = ""
	return client.callPut(ctx, client.Endpoints().Stream(stream.ID), &body, stream)
}

// DeleteStream deletes a stream.
func (client *Client) DeleteStream(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return client.callDelete(ctx, client.Endpoints().Stream(id), nil, nil)
}

// PauseStream pauses a stream.
func (client *Client) PauseStream(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return client.callPost(ctx, client.Endpoints().PauseStream(id), nil, nil)
}

// ResumeStream resumes a stream.
func (client *Client) ResumeStream(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return client.callPost(ctx, client.Endpoints().ResumeStream(id), nil, nil)
}
