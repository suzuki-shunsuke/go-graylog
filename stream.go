package graylog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// CloneStream
// POST /streams/{streamId}/clone Clone a stream
// TestMatchStream
// POST /streams/{streamId}/testMatch Test matching of a stream against a supplied message

// Stream represents a steram.
type Stream struct {
	// ex. "5a94abdac006c60001f04fc1"
	Id string `json:"id,omitempty"`
	// ex. "2018-02-20T11:37:19.371Z"
	CreatedAt string `json:"created_at,omitempty"`
	Title     string `json:"title,omitempty"`
	// ex. local:admin
	CreatorUserId string   `json:"creator_user_id,omitempty"`
	Description   string   `json:"description,omitempty"`
	Outputs       []Output `json:"outputs,omitempty"`
	// ex. "AND"
	MatchingType                   string           `json:"matching_type,omitempty"`
	Disabled                       bool             `json:"disabled,omitempty"`
	Rules                          []StreamRule     `json:"rules,omitempty"`
	AlertConditions                []AlertCondition `json:"alert_conditions,omitempty"`
	AlertReceivers                 *AlertReceivers  `json:"alert_receivers,omitempty"`
	RemoveMatchesFromDefaultStream bool             `json:"remove_matches_from_default_stream,omitempty"`
	// ex. "5a8c086fc006c600013ca6f5"
	IndexSetId string `json:"index_set_id,omitempty"`
	IsDefault  bool   `json:"is_default,omitempty"`
	// ContentPack `json:"content_pack,omitempty"`
}

// Output represents an output.
type Output struct{}

// StreamRule represents a stream rule.
type StreamRule struct{}

// AlertReceivers represents alert receivers.
type AlertReceivers struct {
	Emails []string `json:"emails,omitempty"`
	Users  []string `json:"users,omitempty"`
}

// AlertCondition represents an alert condition.
type AlertCondition struct{}

type streamsBody struct {
	Total   int      `json:"total,omitempty"`
	Streams []Stream `json:"streams,omitempty"`
}

// GetStreams returns all streams.
func (client *Client) GetStreams() (
	streams []Stream, total int, ei *ErrorInfo, err error,
) {
	return client.GetStreamsContext(context.Background())
}

// GetStreamsContext returns all streams with a context.
func (client *Client) GetStreamsContext(
	ctx context.Context,
) (streams []Stream, total int, ei *ErrorInfo, err error) {
	ei, err = client.callReq(
		ctx, http.MethodGet, client.endpoints.Streams, nil, true)
	if err != nil {
		return nil, 0, ei, err
	}

	streamsBody := &streamsBody{}
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
func (client *Client) CreateStream(stream *Stream) (
	string, *ErrorInfo, error,
) {
	return client.CreateStreamContext(context.Background(), stream)
}

// CreateStreamContext creates a stream with a context.
func (client *Client) CreateStreamContext(
	ctx context.Context, stream *Stream,
) (string, *ErrorInfo, error) {
	b, err := json.Marshal(stream)
	if err != nil {
		return "", nil, errors.Wrap(err, "Failed to json.Marshal(stream)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPost, client.endpoints.Streams, b, true)
	if err != nil {
		return "", ei, err
	}

	ret := map[string]string{}
	err = json.Unmarshal(ei.ResponseBody, &ret)
	if err != nil {
		return "", ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as map[string]string: %s",
				string(ei.ResponseBody)))
	}
	if id, ok := ret["stream_id"]; ok {
		return id, ei, nil
	}
	return "", ei, errors.New(`response doesn't have the field "stream_id"`)
}

// GetEnabledStreams returns all enabled streams.
func (client *Client) GetEnabledStreams() (
	streams []Stream, total int, ei *ErrorInfo, err error,
) {
	return client.GetEnabledStreamsContext(context.Background())
}

// GetEnabledStreamsContext returns all enabled streams with a context.
func (client *Client) GetEnabledStreamsContext(
	ctx context.Context,
) (streams []Stream, total int, ei *ErrorInfo, err error) {
	ei, err = client.callReq(
		ctx, http.MethodGet,
		fmt.Sprintf("%s/enabled", client.endpoints.Streams), nil, true)
	if err != nil {
		return nil, 0, ei, err
	}
	streamsBody := &streamsBody{}
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
func (client *Client) GetStream(id string) (*Stream, *ErrorInfo, error) {
	return client.GetStreamContext(context.Background(), id)
}

// GetStream returns a given stream with a context.
func (client *Client) GetStreamContext(
	ctx context.Context, id string,
) (*Stream, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}

	ei, err := client.callReq(
		ctx, http.MethodGet,
		fmt.Sprintf("%s/%s", client.endpoints.Streams, id), nil, true)
	if err != nil {
		return nil, ei, err
	}

	stream := &Stream{}
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
func (client *Client) UpdateStream(id string, stream *Stream) (
	*Stream, *ErrorInfo, error,
) {
	return client.UpdateStreamContext(context.Background(), id, stream)
}

// UpdateStreamContext updates a stream with a context.
func (client *Client) UpdateStreamContext(
	ctx context.Context, id string, stream *Stream,
) (*Stream, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	b, err := json.Marshal(stream)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to json.Marshal(stream)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPut,
		fmt.Sprintf("%s/%s", client.endpoints.Streams, id), b, true)
	if err != nil {
		return nil, ei, err
	}
	ret := &Stream{}

	if err := json.Unmarshal(ei.ResponseBody, ret); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Stream: %s",
				string(ei.ResponseBody)))
	}
	return ret, ei, nil
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
		ctx, http.MethodDelete,
		fmt.Sprintf("%s/%s", client.endpoints.Streams, id), nil, false)
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
		ctx, http.MethodPost,
		fmt.Sprintf("%s/%s/pause", client.endpoints.Streams, id), nil, false)
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
		ctx, http.MethodPost,
		fmt.Sprintf("%s/%s/resume", client.endpoints.Streams, id), nil, false)
}
