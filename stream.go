package graylog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// CloneStream
// POST /streams/{streamId}/clone Clone a stream
// TestMatchStream
// POST /streams/{streamId}/testMatch Test matching of a stream against a supplied message

// Stream represents a steram.
type Stream struct {
	// \d{24}
	Id string `json:"id,omitempty"`
	// "2018-02-20T11:37:19.371Z"
	CreatedAt string `json:"created_at,omitempty"`
	Title     string `json:"title,omitempty"`
	// local:admin
	CreatorUserId                  string           `json:"creator_user_id,omitempty"`
	Description                    string           `json:"description,omitempty"`
	Outputs                        []Output         `json:"outputs,omitempty"`
	MatchingType                   string           `json:"matching_type,omitempty"`
	Disabled                       bool             `json:"disabled,omitempty"`
	Rules                          []StreamRule     `json:"rules,omitempty"`
	AlertConditions                []AlertCondition `json:"alert_conditions,omitempty"`
	AlertReceivers                 *AlertReceivers  `json:"alert_receivers,omitempty"`
	RemoveMatchesFromDefaultStream bool             `json:"remove_matches_from_default_stream,omitempty"`
	IndexSetId                     string           `json:"index_set_id,omitempty"`
	IsDefault                      bool             `json:"is_default,omitempty"`
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
func (client *Client) GetStreams() (streams []Stream, total int, err error) {
	return client.GetStreamsContext(context.Background())
}

// GetStreamsContext returns all streams with a context.
func (client *Client) GetStreamsContext(
	ctx context.Context,
) (streams []Stream, total int, err error) {
	req, err := http.NewRequest(http.MethodGet, client.endpoints.Streams, nil)
	if err != nil {
		return nil, 0, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "Failed to call GET /streams API")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, errors.Wrap(err, "Failed to read response body")
	}
	if resp.StatusCode >= 400 {
		e := Error{}
		err = json.Unmarshal(b, &e)
		if err != nil {
			return nil, 0, errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return nil, 0, errors.New(e.Message)
	}
	streamsBody := &streamsBody{}
	err = json.Unmarshal(b, streamsBody)
	if err != nil {
		return nil, 0, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as streamsBody: %s", string(b)))
	}
	return streamsBody.Streams, streamsBody.Total, nil
}

// CreateStream creates a stream.
func (client *Client) CreateStream(stream *Stream) (string, error) {
	return client.CreateStreamContext(context.Background(), stream)
}

// CreateStreamContext creates a stream with a context.
func (client *Client) CreateStreamContext(
	ctx context.Context, stream *Stream,
) (string, error) {
	b, err := json.Marshal(stream)
	if err != nil {
		return "", errors.Wrap(err, "Failed to json.Marshal(stream)")
	}
	req, err := http.NewRequest(
		http.MethodPost, client.endpoints.Streams, bytes.NewBuffer(b))
	if err != nil {
		return "", errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return "", errors.Wrap(err, "Failed to call POST /streams API")
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "Failed to read response body")
	}
	if resp.StatusCode >= 400 {
		e := Error{}
		err = json.Unmarshal(b, &e)
		if err != nil {
			return "", errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return "", errors.New(e.Message)
	}
	ret := map[string]string{}
	err = json.Unmarshal(b, &ret)
	if err != nil {
		return "", errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as map[string]string: %s", string(b)))
	}
	if id, ok := ret["stream_id"]; ok {
		return id, nil
	}
	return "", errors.New(`response doesn't have the field "stream_id"`)
}

// GetEnabledStreams returns all enabled streams.
func (client *Client) GetEnabledStreams() (
	streams []Stream, total int, err error,
) {
	return client.GetEnabledStreamsContext(context.Background())
}

// GetEnabledStreamsContext returns all enabled streams with a context.
func (client *Client) GetEnabledStreamsContext(
	ctx context.Context,
) (streams []Stream, total int, err error) {
	req, err := http.NewRequest(
		http.MethodGet, fmt.Sprintf("%s/enabled", client.endpoints.Streams), nil)
	if err != nil {
		return nil, 0, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, 0, errors.Wrap(err, "Failed to call GET /streams/enabled API")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, errors.Wrap(err, "Failed to read response body")
	}
	if resp.StatusCode >= 400 {
		e := Error{}
		err = json.Unmarshal(b, &e)
		if err != nil {
			return nil, 0, errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return nil, 0, errors.New(e.Message)
	}
	streamsBody := &streamsBody{}
	err = json.Unmarshal(b, streamsBody)
	if err != nil {
		return nil, 0, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as streamsBody: %s", string(b)))
	}
	return streamsBody.Streams, streamsBody.Total, nil
}

// GetStream returns a given stream.
func (client *Client) GetStream(id string) (*Stream, error) {
	return client.GetStreamContext(context.Background(), id)
}

// GetStream returns a given stream with a context.
func (client *Client) GetStreamContext(
	ctx context.Context, id string,
) (*Stream, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	req, err := http.NewRequest(
		http.MethodGet, fmt.Sprintf("%s/%s", client.endpoints.Streams, id), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to call GET /streams/{streamId} API")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read response body")
	}
	if resp.StatusCode >= 400 {
		e := Error{}
		err = json.Unmarshal(b, &e)
		if err != nil {
			return nil, errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return nil, errors.New(e.Message)
	}
	stream := &Stream{}
	err = json.Unmarshal(b, stream)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as Stream: %s", string(b)))
	}
	return stream, nil
}

// UpdateStream updates a stream.
func (client *Client) UpdateStream(id string, stream *Stream) (*Stream, error) {
	return client.UpdateStreamContext(context.Background(), id, stream)
}

// UpdateStreamContext updates a stream with a context.
func (client *Client) UpdateStreamContext(
	ctx context.Context, id string, stream *Stream,
) (*Stream, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	b, err := json.Marshal(stream)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(stream)")
	}
	req, err := http.NewRequest(
		http.MethodPut, fmt.Sprintf("%s/%s", client.endpoints.Streams, id),
		bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to call PUT /streams/{streamId} API")
	}
	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read response body")
	}
	if resp.StatusCode >= 400 {
		e := Error{}
		err = json.Unmarshal(b, &e)
		if err != nil {
			return nil, errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return nil, errors.New(e.Message)
	}
	ret := &Stream{}
	err = json.Unmarshal(b, ret)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Stream: %s", string(b)))
	}
	return ret, nil
}

// DeleteStream deletes a stream.
func (client *Client) DeleteStream(id string) error {
	return client.DeleteStreamContext(context.Background(), id)
}

// DeleteStreamContext deletes a stream with a context.
func (client *Client) DeleteStreamContext(
	ctx context.Context, id string,
) error {
	if id == "" {
		return errors.New("id is empty")
	}
	req, err := http.NewRequest(
		http.MethodDelete, fmt.Sprintf("%s/%s", client.endpoints.Streams, id), nil)
	if err != nil {
		return errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to call DELETE /streams/{streamId} API")
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "Failed to read response body")
		}
		e := &Error{}
		err = json.Unmarshal(b, e)
		if err != nil {
			return errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return errors.New(e.Message)
	}
	return nil
}

// PauseStream pauses a stream.
func (client *Client) PauseStream(id string) error {
	return client.PauseStreamContext(context.Background(), id)
}

// PauseStreamContext pauses a stream with a context.
func (client *Client) PauseStreamContext(
	ctx context.Context, id string,
) error {
	if id == "" {
		return errors.New("id is empty")
	}
	req, err := http.NewRequest(
		http.MethodPost, fmt.Sprintf("%s/%s/pause", client.endpoints.Streams, id), nil)
	if err != nil {
		return errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to call POST /streams/{streamId}/pause API")
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "Failed to read response body")
		}
		e := &Error{}
		err = json.Unmarshal(b, e)
		if err != nil {
			return errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return errors.New(e.Message)
	}
	return nil
}

// ResumeStream resumes a stream.
func (client *Client) ResumeStream(id string) error {
	return client.ResumeStreamContext(context.Background(), id)
}

// ResumeStreamContext resumes a stream with a context.
func (client *Client) ResumeStreamContext(
	ctx context.Context, id string,
) error {
	if id == "" {
		return errors.New("id is empty")
	}
	req, err := http.NewRequest(
		http.MethodPost, fmt.Sprintf(
			"%s/%s/resume", client.endpoints.Streams, id), nil)
	if err != nil {
		return errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return errors.Wrap(
			err, "Failed to call POST /streams/{streamId}/resume API")
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "Failed to read response body")
		}
		e := &Error{}
		err = json.Unmarshal(b, e)
		if err != nil {
			return errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return errors.New(e.Message)
	}
	return nil
}
