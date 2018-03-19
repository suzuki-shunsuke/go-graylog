package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-graylog"
)

// GetIndexSets returns a list of all index sets.
func (client *Client) GetIndexSets(
	skip, limit int,
) ([]graylog.IndexSet, *graylog.IndexSetStats, *ErrorInfo, error) {
	return client.GetIndexSetsContext(context.Background(), skip, limit)
}

// GetIndexSetsContext returns a list of all index sets with a context.
func (client *Client) GetIndexSetsContext(
	ctx context.Context, skip, limit int,
) ([]graylog.IndexSet, *graylog.IndexSetStats, *ErrorInfo, error) {
	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.IndexSets, nil, true)
	if err != nil {
		return nil, nil, ei, err
	}
	indexSets := &graylog.IndexSetsBody{}
	if err := json.Unmarshal(ei.ResponseBody, indexSets); err != nil {
		return nil, nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSetsBody: %s",
				string(ei.ResponseBody)))
	}
	return indexSets.IndexSets, indexSets.Stats, ei, nil
}

// GetIndexSet returns a given index set.
func (client *Client) GetIndexSet(id string) (*graylog.IndexSet, *ErrorInfo, error) {
	return client.GetIndexSetContext(context.Background(), id)
}

// GetIndexSetContext returns a given index set with a context.
func (client *Client) GetIndexSetContext(
	ctx context.Context, id string,
) (*graylog.IndexSet, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.IndexSet(id), nil, true)
	if err != nil {
		return nil, ei, err
	}
	indexSet := &graylog.IndexSet{}
	if err := json.Unmarshal(ei.ResponseBody, indexSet); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s",
				string(ei.ResponseBody)))
	}
	return indexSet, ei, nil
}

// CreateIndexSet creates a Index Set.
func (client *Client) CreateIndexSet(indexSet *graylog.IndexSet) (*ErrorInfo, error) {
	return client.CreateIndexSetContext(context.Background(), indexSet)
}

// CreateIndexSetContext creates a Index Set with a context.
func (client *Client) CreateIndexSetContext(
	ctx context.Context, is *graylog.IndexSet,
) (*ErrorInfo, error) {
	if is == nil {
		return nil, fmt.Errorf("IndexSet is nil")
	}
	b, err := json.Marshal(is)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(indexSet)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPost, client.Endpoints.IndexSets, b, true)
	if err != nil {
		return ei, err
	}

	if err := json.Unmarshal(ei.ResponseBody, is); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s",
				string(ei.ResponseBody)))
	}
	return ei, nil
}

// UpdateIndexSet updates a given Index Set.
func (client *Client) UpdateIndexSet(is *graylog.IndexSet) (*ErrorInfo, error) {
	return client.UpdateIndexSetContext(context.Background(), is)
}

// UpdateIndexSetContext updates a given Index Set with a context.
func (client *Client) UpdateIndexSetContext(
	ctx context.Context, is *graylog.IndexSet,
) (*ErrorInfo, error) {
	if is == nil {
		return nil, fmt.Errorf("IndexSet is nil")
	}
	if is.ID == "" {
		return nil, errors.New("id is empty")
	}
	copiedIndexSet := *is
	copiedIndexSet.ID = ""
	b, err := json.Marshal(&copiedIndexSet)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(indexSet)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPut, client.Endpoints.IndexSet(is.ID), b, true)
	if err != nil {
		return ei, err
	}

	if err := json.Unmarshal(ei.ResponseBody, is); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s",
				string(ei.ResponseBody)))
	}
	return ei, nil
}

// DeleteIndexSet deletes a given Index Set.
func (client *Client) DeleteIndexSet(id string) (*ErrorInfo, error) {
	return client.DeleteIndexSetContext(context.Background(), id)
}

// DeleteIndexSet deletes a given Index Set with a context.
func (client *Client) DeleteIndexSetContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}

	return client.callReq(
		ctx, http.MethodDelete, client.Endpoints.IndexSet(id), nil, false)
}

// SetDefaultIndexSet sets default Index Set.
func (client *Client) SetDefaultIndexSet(id string) (
	*graylog.IndexSet, *ErrorInfo, error,
) {
	return client.SetDefaultIndexSetContext(context.Background(), id)
}

// SetDefaultIndexSet sets default Index Set with a context.
func (client *Client) SetDefaultIndexSetContext(
	ctx context.Context, id string,
) (*graylog.IndexSet, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}

	ei, err := client.callReq(
		ctx, http.MethodPut, client.Endpoints.SetDefaultIndexSet(id), nil, true)
	if err != nil {
		return nil, ei, err
	}

	is := &graylog.IndexSet{}
	if err := json.Unmarshal(ei.ResponseBody, is); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s",
				string(ei.ResponseBody)))
	}
	return is, ei, nil
}
