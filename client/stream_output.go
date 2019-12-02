package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

// CreateOutput creates a new output.
func (client *Client) CreateStreamOutputs(
	ctx context.Context, streamID string, outputIDs []string,
) (*ErrorInfo, error) {
	if streamID == "" {
		return nil, errors.New("stream id is empty")
	}
	return client.callPost(
		ctx, client.Endpoints().StreamOutputs(streamID),
		map[string]interface{}{
			"outputs": outputIDs,
		}, nil)
}

// GetOutputs returns all outputs.
func (client *Client) GetStreamOutputs(ctx context.Context, streamID string) (
	[]graylog.Output, int, *ErrorInfo, error,
) {
	outputs := &graylog.OutputsBody{}
	ei, err := client.callGet(ctx, client.Endpoints().StreamOutputs(streamID), nil, outputs)
	return outputs.Outputs, outputs.Total, ei, err
}

// DeleteOutput deletes a given output.
func (client *Client) DeleteStreamOutput(
	ctx context.Context, streamID, outputID string,
) (*ErrorInfo, error) {
	if streamID == "" {
		return nil, errors.New("stream id is empty")
	}
	if outputID == "" {
		return nil, errors.New("output id is empty")
	}
	return client.callDelete(ctx, client.Endpoints().StreamOutput(streamID, outputID), nil, nil)
}
