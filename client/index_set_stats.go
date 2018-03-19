package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-graylog"
)

// GetIndexSetStats returns a given Index Set statistics.
func (client *Client) GetIndexSetStats(id string) (
	*graylog.IndexSetStats, *ErrorInfo, error,
) {
	return client.GetIndexSetStatsContext(context.Background(), id)
}

// GetIndexSetStatsContext returns a given Index Set statistics with a context.
func (client *Client) GetIndexSetStatsContext(
	ctx context.Context, id string,
) (*graylog.IndexSetStats, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}

	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.IndexSetStats(id), nil, true)
	if err != nil {
		return nil, ei, err
	}

	indexSetStats := &graylog.IndexSetStats{}
	if err := json.Unmarshal(ei.ResponseBody, indexSetStats); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSetStats: %s",
				string(ei.ResponseBody)))
	}
	return indexSetStats, ei, nil
}

// GetAllIndexSetsStats returns stats of all Index Sets.
func (client *Client) GetAllIndexSetsStats() (
	*graylog.IndexSetStats, *ErrorInfo, error,
) {
	return client.GetAllIndexSetsStatsContext(context.Background())
}

// GetAllIndexSetsStats returns stats of all Index Sets with a context.
func (client *Client) GetAllIndexSetsStatsContext(
	ctx context.Context,
) (*graylog.IndexSetStats, *ErrorInfo, error) {
	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.IndexSetsStats(), nil, true)
	if err != nil {
		return nil, ei, err
	}
	indexSetStats := &graylog.IndexSetStats{}
	if err := json.Unmarshal(ei.ResponseBody, indexSetStats); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf("Failed to parse response body as IndexSetStats: %s",
				string(ei.ResponseBody)))
	}
	return indexSetStats, ei, nil
}
