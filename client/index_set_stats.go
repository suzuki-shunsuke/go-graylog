package client

import (
	"context"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetIndexSetStats returns a given Index Set statistics.
func (client *Client) GetIndexSetStats(
	ctx context.Context, id string,
) (*graylog.IndexSetStats, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	indexSetStats := &graylog.IndexSetStats{}
	u, err := client.Endpoints().IndexSetStats(id)
	if err != nil {
		return nil, nil, err
	}
	ei, err := client.callGet(
		ctx, u.String(), nil, indexSetStats)
	return indexSetStats, ei, err
}

// GetTotalIndexSetsStats returns stats of all Index Sets.
func (client *Client) GetTotalIndexSetsStats(
	ctx context.Context,
) (*graylog.IndexSetStats, *ErrorInfo, error) {
	indexSetStats := &graylog.IndexSetStats{}
	ei, err := client.callGet(
		ctx, client.Endpoints().IndexSetsStats(), nil, indexSetStats)
	return indexSetStats, ei, err
}
