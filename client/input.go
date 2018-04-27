package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-graylog"
)

// GetInputs returns all inputs.
func (client *Client) GetInputs() ([]graylog.Input, int, *ErrorInfo, error) {
	return client.GetInputsContext(context.Background())
}

// GetInputsContext returns all inputs with a context.
func (client *Client) GetInputsContext(ctx context.Context) (
	[]graylog.Input, int, *ErrorInfo, error,
) {
	inputs := &graylog.InputsBody{}
	ei, err := client.callGet(
		ctx, client.Endpoints().Inputs(), nil, inputs)
	return inputs.Inputs, inputs.Total, ei, err
}

// GetInput returns a given input.
func (client *Client) GetInput(id string) (*graylog.Input, *ErrorInfo, error) {
	return client.GetInputContext(context.Background(), id)
}

// GetInputContext returns a given input with a context.
func (client *Client) GetInputContext(
	ctx context.Context, id string,
) (*graylog.Input, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().Input(id)
	if err != nil {
		return nil, nil, err
	}
	input := &graylog.Input{}
	ei, err := client.callGet(
		ctx, u.String(), nil, input)
	return input, ei, err
}

// CreateInput creates an input.
func (client *Client) CreateInput(input *graylog.Input) (
	ei *ErrorInfo, err error,
) {
	return client.CreateInputContext(context.Background(), input)
}

// CreateInputContext creates an input with a context.
func (client *Client) CreateInputContext(
	ctx context.Context, input *graylog.Input,
) (ei *ErrorInfo, err error) {
	if input == nil {
		return nil, fmt.Errorf("input is nil")
	}
	if input.ID != "" {
		return nil, fmt.Errorf("input id should be empty")
	}
	// change attributes to configuration
	// https://github.com/Graylog2/graylog2-server/issues/3480
	d := map[string]interface{}{
		"title":         input.Title,
		"type":          input.Type,
		"configuration": input.Attributes,
		"global":        input.Global,
	}
	if input.Node != "" {
		d["node"] = input.Node
	}

	return client.callPost(ctx, client.Endpoints().Inputs(), &d, input)
}

// UpdateInput updates an given input.
func (client *Client) UpdateInput(input *graylog.InputUpdateParams) (*graylog.Input, *ErrorInfo, error) {
	return client.UpdateInputContext(context.Background(), input)
}

// UpdateInputContext updates an given input with a context.
func (client *Client) UpdateInputContext(
	ctx context.Context, prms *graylog.InputUpdateParams,
) (*graylog.Input, *ErrorInfo, error) {
	if prms == nil {
		return nil, nil, fmt.Errorf("input is nil")
	}
	if prms.ID == "" {
		return nil, nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().Input(prms.ID)
	if err != nil {
		return nil, nil, err
	}
	// change attributes to configuration
	// https://github.com/Graylog2/graylog2-server/issues/3480
	d := map[string]interface{}{
		"title":         prms.Title,
		"type":          prms.Type,
		"configuration": prms.Attributes,
		"global":        prms.Global,
	}
	if prms.Node != "" {
		d["node"] = prms.Node
	}
	input := &graylog.Input{}
	ei, err := client.callPut(ctx, u.String(), &d, input)
	return input, ei, err
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
	u, err := client.Endpoints().Input(id)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}
