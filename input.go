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

type Attributes struct {
	// OverrideSource string `json:"override_source,omitempty"`
	RecvBufferSize      int    `json:"recv_buffer_size,omitempty"`
	BindAddress         string `json:"bind_address,omitempty"`
	Port                int    `json:"port,omitempty"`
	DecompressSizeLimit int    `json:"decompress_size_limit,omitempty"`
}

type Input struct {
	Id            string `json:"id,omitempty"`
	Title         string `json:"title,omitempty"`
	Type          string `json:"type,omitempty"`
	Global        bool   `json:"global,omitempty"`
	Node          string `json:"node,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	CreatorUserId string `json:"creator_user_id,omitempty"`
	Attributes    string `json:"attributes,omitempty"`
	// ContextPack `json:"context_pack,omitempty"`
	// Configuration `json:"configuration,omitempty"`
	// StaticFields `json:"static_fields,omitempty"`
}

// CreateInput
// POST /system/inputs Launch input on this node
func (client *Client) CreateInput(params *Input) (*Input, error) {
	return client.CreateInputContext(context.Background(), params)
}

// CreateInputContext
// POST /system/inputs Launch input on this node
func (client *Client) CreateInputContext(
	ctx context.Context, params *Input,
) (*Input, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(params)")
	}
	req, err := http.NewRequest(
		http.MethodPost, client.endpoints.Inputs, bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to call POST /inputs API")
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
	input := Input{}
	err = json.Unmarshal(b, &input)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Input: %s", string(b)))
	}
	return &input, nil
}

type inputsBody struct {
	Inputs []Input `json:"inputs"`
	Total  int     `json:"total"`
}

// GetInputs
// GET /system/inputs Get all inputs
func (client *Client) GetInputs() ([]Input, error) {
	return client.GetInputsContext(context.Background())
}

// GetInputsContext
// GET /system/inputs Get all inputs
func (client *Client) GetInputsContext(
	ctx context.Context,
) ([]Input, error) {
	req, err := http.NewRequest(http.MethodGet, client.endpoints.Inputs, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to call GET /inputs API")
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
	inputs := inputsBody{}
	err = json.Unmarshal(b, &inputs)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Inputs: %s", string(b)))
	}
	return inputs.Inputs, nil
}

// GetInput
// GET /system/inputs/{inputId} Get information of a single input on this node
func (client *Client) GetInput(id string) (*Input, error) {
	return client.GetInputContext(context.Background(), id)
}

// GetInputContext
// GET /system/inputs/{inputId} Get information of a single input on this node
func (client *Client) GetInputContext(
	ctx context.Context, id string,
) (*Input, error) {
	req, err := http.NewRequest(
		http.MethodGet, fmt.Sprintf("%s/%s", client.endpoints.Inputs, id), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to call GET /system/inputs/{inputId} API")
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
	input := Input{}
	err = json.Unmarshal(b, &input)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Input: %s", string(b)))
	}
	return &input, nil
}

// UpdateInput
// PUT /system/inputs/{inputId} Update input on this node
func (client *Client) UpdateInput(id string, params *Input) (*Input, error) {
	return client.UpdateInputContext(context.Background(), id, params)
}

// UpdateInputContext
// PUT /system/inputs/{inputId} Update input on this node
func (client *Client) UpdateInputContext(
	ctx context.Context, id string, params *Input,
) (*Input, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(params)")
	}
	req, err := http.NewRequest(
		http.MethodPut, fmt.Sprintf("%s/%s", client.endpoints.Inputs, id),
		bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to call PUT /system/inputs/{inputId} API")
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
	input := Input{}
	err = json.Unmarshal(b, &input)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as Input: %s", string(b)))
	}
	return &input, nil
}

// DeleteInput
// DELETE /system/inputs/{inputId} Terminate input on this node
func (client *Client) DeleteInput(id string) error {
	return client.DeleteInputContext(context.Background(), id)
}

// DeleteInputContext
// DELETE /system/inputs/{inputId} Terminate input on this node
func (client *Client) DeleteInputContext(
	ctx context.Context, id string,
) error {
	req, err := http.NewRequest(
		http.MethodDelete, fmt.Sprintf("%s/%s", client.endpoints.Inputs, id), nil)
	if err != nil {
		return errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to call DELETE /system/inputs/{inputId} API")
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "Failed to read response body")
		}
		e := Error{}
		err = json.Unmarshal(b, &e)
		if err != nil {
			return errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return errors.New(e.Message)
	}
	return nil
}
