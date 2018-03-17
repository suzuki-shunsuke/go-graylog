package graylog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// InputAttributes represents Input's attributes.
type InputAttributes struct {
	// OverrideSource string `json:"override_source,omitempty"`
	RecvBufferSize      int    `json:"recv_buffer_size,omitempty"`
	BindAddress         string `json:"bind_address,omitempty"`
	Port                int    `json:"port,omitempty"`
	DecompressSizeLimit int    `json:"decompress_size_limit,omitempty"`
}

// InputConfiguration represents Input's configuration.
type InputConfiguration struct {
	// ex. 0.0.0.0
	BindAddress string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	Port        int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	// ex. 262144
	RecvBufferSize int `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
}

// Input represents Graylog Input.
type Input struct {
	// required
	Title         string              `json:"title,omitempty" v-create:"required" v-update:"required"`
	Type          string              `json:"type,omitempty" v-create:"required" v-update:"required"`
	Configuration *InputConfiguration `json:"configuration,omitempty" v-create:"required" v-update:"required"`

	// ex. "5a90d5c2c006c60001efc368"
	ID string `json:"id,omitempty" v-create:"isdefault" v-update:"required"`

	Global bool `json:"global,omitempty"`
	// ex. "2ad6b340-3e5f-4a96-ae81-040cfb8b6024"
	Node string `json:"node,omitempty"`
	// ex. 2018-02-24T03:02:26.001Z
	CreatedAt string `json:"created_at,omitempty" v-create:"isdefault" v-update:"isdefault"`
	// ex. "admin"
	CreatorUserID string           `json:"creator_user_id,omitempty" v-create:"isdefault" v-update:"isdefault"`
	Attributes    *InputAttributes `json:"attributes,omitempty" v-create:"isdefault"`
	// ContextPack `json:"context_pack,omitempty"`
	// StaticFields `json:"static_fields,omitempty"`
}

// CreateInput creates an input.
func (client *Client) CreateInput(input *Input) (
	ei *ErrorInfo, err error,
) {
	return client.CreateInputContext(context.Background(), input)
}

// CreateInputContext creates an input with a context.
func (client *Client) CreateInputContext(
	ctx context.Context, input *Input,
) (ei *ErrorInfo, err error) {
	if input == nil {
		return nil, fmt.Errorf("input is nil")
	}
	b, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(input)")
	}

	ei, err = client.callReq(
		ctx, http.MethodPost, client.Endpoints.Inputs, b, true)
	if err != nil {
		return ei, err
	}

	if err := json.Unmarshal(ei.ResponseBody, input); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Input: %s",
				string(ei.ResponseBody)))
	}
	return ei, nil
}

type InputsBody struct {
	Inputs []Input `json:"inputs"`
	Total  int     `json:"total"`
}

// GetInputs returns all inputs.
func (client *Client) GetInputs() ([]Input, *ErrorInfo, error) {
	return client.GetInputsContext(context.Background())
}

// GetInputsContext returns all inputs with a context.
func (client *Client) GetInputsContext(ctx context.Context) (
	[]Input, *ErrorInfo, error,
) {
	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.Inputs, nil, true)
	if err != nil {
		return nil, ei, err
	}

	inputs := &InputsBody{}
	if err := json.Unmarshal(ei.ResponseBody, inputs); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as Inputs: %s",
				string(ei.ResponseBody)))
	}
	return inputs.Inputs, ei, nil
}

// GetInput returns a given input.
func (client *Client) GetInput(id string) (*Input, *ErrorInfo, error) {
	return client.GetInputContext(context.Background(), id)
}

// GetInputContext returns a given input with a context.
func (client *Client) GetInputContext(
	ctx context.Context, id string,
) (*Input, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}

	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.Input(id), nil, true)
	if err != nil {
		return nil, ei, err
	}

	input := &Input{}
	if err := json.Unmarshal(ei.ResponseBody, input); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as Input: %s", string(ei.ResponseBody)))
	}
	return input, ei, nil
}

// UpdateInput updates an given input.
func (client *Client) UpdateInput(input *Input) (
	*ErrorInfo, error,
) {
	return client.UpdateInputContext(context.Background(), input)
}

// UpdateInputContext updates an given input with a context.
func (client *Client) UpdateInputContext(
	ctx context.Context, input *Input,
) (*ErrorInfo, error) {
	if input == nil {
		return nil, fmt.Errorf("input is nil")
	}
	if input.ID == "" {
		return nil, errors.New("id is empty")
	}
	copiedInput := *input
	copiedInput.ID = ""
	b, err := json.Marshal(copiedInput)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(input)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPut, client.Endpoints.Input(input.ID), b, true)
	if err != nil {
		return ei, err
	}

	if err := json.Unmarshal(ei.ResponseBody, input); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as Input: %s", string(ei.ResponseBody)))
	}
	return ei, nil
}

// DeleteInput deletes an given input.
func (client *Client) DeleteInput(id string) (*ErrorInfo, error) {
	return client.DeleteInputContext(context.Background(), id)
}

// DeleteInputContext deletes an given input with a context.
func (client *Client) DeleteInputContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}

	return client.callReq(
		ctx, http.MethodDelete, client.Endpoints.Input(id), nil, false)
}
