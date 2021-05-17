package client

import (
	"context"
)

// Generic calls a generic API Endpoint by a relative endpoint URL
func (client *Client) Generic(ctx context.Context, method, endpoint string, input, output interface{}) (
	*ErrorInfo, error,
) {
	ei, err := client.callAPI(ctx, method, client.Endpoints().Generic(endpoint), input, output)
	return ei, err
}
