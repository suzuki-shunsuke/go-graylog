package client

import (
	"context"
	"errors"
	"net/url"
	"strconv"

	"github.com/suzuki-shunsuke/go-graylog/v9"
)

// GetIndexSets returns a list of all index sets.
func (client *Client) GetIndexSets(
	ctx context.Context, skip, limit int, stats bool,
) ([]graylog.IndexSet, map[string]graylog.IndexSetStats, int, *ErrorInfo, error) {
	indexSets := &graylog.IndexSetsBody{}
	v := url.Values{
		"skip":  []string{strconv.Itoa(skip)},
		"limit": []string{strconv.Itoa(limit)},
		"stats": []string{strconv.FormatBool(stats)},
	}
	u := client.Endpoints().IndexSets() + "?" + v.Encode()
	ei, err := client.callGet(ctx, u, nil, indexSets)
	return indexSets.IndexSets, indexSets.Stats, indexSets.Total, ei, err
}

// GetIndexSet returns a given index set.
func (client *Client) GetIndexSet(
	ctx context.Context, id string,
) (*graylog.IndexSet, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	is := &graylog.IndexSet{}
	ei, err := client.callGet(ctx, client.Endpoints().IndexSet(id), nil, is)
	return is, ei, err
}

// CreateIndexSet creates a Index Set.
func (client *Client) CreateIndexSet(
	ctx context.Context, is *graylog.IndexSet,
) (*ErrorInfo, error) {
	if is == nil {
		return nil, errors.New("index set is nil")
	}
	is.SetCreateDefaultValues()

	return client.callPost(ctx, client.Endpoints().IndexSets(), is, is)
}

// UpdateIndexSet updates a given Index Set.
func (client *Client) UpdateIndexSet(
	ctx context.Context, prms *graylog.IndexSetUpdateParams,
) (*graylog.IndexSet, *ErrorInfo, error) {
	if prms == nil {
		return nil, nil, errors.New("index set is nil")
	}
	if prms.ID == "" {
		return nil, nil, errors.New("id is empty")
	}
	u := client.Endpoints().IndexSet(prms.ID)
	a := *prms
	a.ID = ""
	is := &graylog.IndexSet{}
	ei, err := client.callPut(ctx, u, &a, is)
	return is, ei, err
}

// DeleteIndexSet deletes a given Index Set.
func (client *Client) DeleteIndexSet(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	return client.callDelete(ctx, client.Endpoints().IndexSet(id), nil, nil)
}

// SetDefaultIndexSet sets default Index Set.
func (client *Client) SetDefaultIndexSet(
	ctx context.Context, id string,
) (*graylog.IndexSet, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	is := &graylog.IndexSet{}
	ei, err := client.callPut(ctx, client.Endpoints().SetDefaultIndexSet(id), nil, is)
	return is, ei, err
}
