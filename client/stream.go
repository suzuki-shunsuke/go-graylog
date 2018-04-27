package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-graylog"
)

// GetStreams returns all streams.
func (client *Client) GetStreams() (
	streams []graylog.Stream, total int, ei *ErrorInfo, err error,
) {
	return client.GetStreamsContext(context.Background())
}

// GetStreamsContext returns all streams with a context.
func (client *Client) GetStreamsContext(
	ctx context.Context,
) (streams []graylog.Stream, total int, ei *ErrorInfo, err error) {
	streamsBody := &graylog.StreamsBody{}
	ei, err = client.callGet(
		ctx, client.Endpoints().Streams(), nil, streamsBody)
	return streamsBody.Streams, streamsBody.Total, ei, err
}

// GetStream returns a given stream.
func (client *Client) GetStream(id string) (*graylog.Stream, *ErrorInfo, error) {
	return client.GetStreamContext(context.Background(), id)
}

// GetStreamContext returns a given stream with a context.
func (client *Client) GetStreamContext(
	ctx context.Context, id string,
) (*graylog.Stream, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().Stream(id)
	if err != nil {
		return nil, nil, err
	}
	stream := &graylog.Stream{}
	ei, err := client.callGet(ctx, u.String(), nil, stream)
	return stream, ei, err
}

// CreateStream creates a stream.
func (client *Client) CreateStream(stream *graylog.Stream) (*ErrorInfo, error) {
	return client.CreateStreamContext(context.Background(), stream)
}

// CreateStreamContext creates a stream with a context.
func (client *Client) CreateStreamContext(
	ctx context.Context, stream *graylog.Stream,
) (*ErrorInfo, error) {
	if stream == nil {
		return nil, fmt.Errorf("stream is nil")
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
func (client *Client) GetEnabledStreams() (
	streams []graylog.Stream, total int, ei *ErrorInfo, err error,
) {
	return client.GetEnabledStreamsContext(context.Background())
}

// GetEnabledStreamsContext returns all enabled streams with a context.
func (client *Client) GetEnabledStreamsContext(
	ctx context.Context,
) (streams []graylog.Stream, total int, ei *ErrorInfo, err error) {
	streamsBody := &graylog.StreamsBody{}
	ei, err = client.callGet(
		ctx, client.Endpoints().EnabledStreams(), nil, streamsBody)
	return streamsBody.Streams, streamsBody.Total, ei, err
}

// UpdateStream updates a stream.
func (client *Client) UpdateStream(stream *graylog.Stream) (*ErrorInfo, error) {
	return client.UpdateStreamContext(context.Background(), stream)
}

// UpdateStreamContext updates a stream with a context.
func (client *Client) UpdateStreamContext(
	ctx context.Context, stream *graylog.Stream,
) (*ErrorInfo, error) {
	if stream == nil {
		return nil, fmt.Errorf("stream is nil")
	}
	if stream.ID == "" {
		return nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().Stream(stream.ID)
	if err != nil {
		return nil, err
	}
	body := *stream
	body.ID = ""
	return client.callPut(ctx, u.String(), &body, stream)
}

// DeleteStream deletes a stream.
func (client *Client) DeleteStream(id string) (*ErrorInfo, error) {
	return client.DeleteStreamContext(context.Background(), id)
}

// DeleteStreamContext deletes a stream with a context.
func (client *Client) DeleteStreamContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().Stream(id)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}

// PauseStream pauses a stream.
func (client *Client) PauseStream(id string) (*ErrorInfo, error) {
	return client.PauseStreamContext(context.Background(), id)
}

// PauseStreamContext pauses a stream with a context.
func (client *Client) PauseStreamContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().PauseStream(id)
	if err != nil {
		return nil, err
	}
	return client.callPost(ctx, u.String(), nil, nil)
}

// ResumeStream resumes a stream.
func (client *Client) ResumeStream(id string) (*ErrorInfo, error) {
	return client.ResumeStreamContext(context.Background(), id)
}

// ResumeStreamContext resumes a stream with a context.
func (client *Client) ResumeStreamContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().ResumeStream(id)
	if err != nil {
		return nil, err
	}
	return client.callPost(ctx, u.String(), nil, nil)
}
