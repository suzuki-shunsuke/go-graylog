package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

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
	ei, err = client.callReq(
		ctx, http.MethodGet, client.Endpoints.Streams, nil, true)
	if err != nil {
		return nil, 0, ei, err
	}

	streamsBody := &graylog.StreamsBody{}
	err = json.Unmarshal(ei.ResponseBody, streamsBody)
	if err != nil {
		return nil, 0, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as streamsBody: %s",
				string(ei.ResponseBody)))
	}
	return streamsBody.Streams, streamsBody.Total, ei, nil
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
	b, err := json.Marshal(stream)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(stream)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPost, client.Endpoints.Streams, b, true)
	if err != nil {
		return ei, err
	}

	ret := map[string]string{}
	if err := json.Unmarshal(ei.ResponseBody, &ret); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as map[string]string: %s",
				string(ei.ResponseBody)))
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
	ei, err = client.callReq(
		ctx, http.MethodGet, client.Endpoints.EnabledStreams, nil, true)
	if err != nil {
		return nil, 0, ei, err
	}
	streamsBody := &graylog.StreamsBody{}
	err = json.Unmarshal(ei.ResponseBody, streamsBody)
	if err != nil {
		return nil, 0, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as streamsBody: %s",
				string(ei.ResponseBody)))
	}
	return streamsBody.Streams, streamsBody.Total, ei, nil
}

// GetStream returns a given stream.
func (client *Client) GetStream(id string) (*graylog.Stream, *ErrorInfo, error) {
	return client.GetStreamContext(context.Background(), id)
}

// GetStream returns a given stream with a context.
func (client *Client) GetStreamContext(
	ctx context.Context, id string,
) (*graylog.Stream, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}

	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.Stream(id), nil, true)
	if err != nil {
		return nil, ei, err
	}

	stream := &graylog.Stream{}
	err = json.Unmarshal(ei.ResponseBody, stream)
	if err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as Stream: %s",
				string(ei.ResponseBody)))
	}
	return stream, ei, nil
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
	body := *stream
	body.ID = ""
	b, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(stream)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPut, client.Endpoints.Stream(stream.ID), b, true)
	if err != nil {
		return ei, err
	}
	if err := json.Unmarshal(ei.ResponseBody, stream); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Stream: %s",
				string(ei.ResponseBody)))
	}
	return ei, nil
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

	return client.callReq(
		ctx, http.MethodDelete, client.Endpoints.Stream(id), nil, false)
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

	return client.callReq(
		ctx, http.MethodPost, client.Endpoints.PauseStream(id), nil, false)
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

	return client.callReq(
		ctx, http.MethodPost, client.Endpoints.ResumeStream(id), nil, false)
}
