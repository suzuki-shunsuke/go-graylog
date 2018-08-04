package client

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/pkg/errors"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetIndexSets returns a list of all index sets.
func (client *Client) GetIndexSets(
	skip, limit int, stats bool,
) ([]graylog.IndexSet, map[string]graylog.IndexSetStats, int, *ErrorInfo, error) {
	return client.GetIndexSetsContext(context.Background(), skip, limit, stats)
}

// GetIndexSetsContext returns a list of all index sets with a context.
func (client *Client) GetIndexSetsContext(
	ctx context.Context, skip, limit int, stats bool,
) ([]graylog.IndexSet, map[string]graylog.IndexSetStats, int, *ErrorInfo, error) {
	indexSets := &graylog.IndexSetsBody{}
	v := url.Values{
		"skip":  []string{strconv.Itoa(skip)},
		"limit": []string{strconv.Itoa(limit)},
		"stats": []string{strconv.FormatBool(stats)},
	}
	u := fmt.Sprintf("%s?%s", client.Endpoints().IndexSets(), v.Encode())
	ei, err := client.callGet(ctx, u, nil, indexSets)
	return indexSets.IndexSets, indexSets.Stats, indexSets.Total, ei, err
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
	is := &graylog.IndexSet{}
	u, err := client.Endpoints().IndexSet(id)
	if err != nil {
		return nil, nil, err
	}
	ei, err := client.callGet(
		ctx, u.String(), nil, is)
	return is, ei, err
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
		return nil, fmt.Errorf("index set is nil")
	}
	is.SetCreateDefaultValues()

	return client.callPost(ctx, client.Endpoints().IndexSets(), is, is)
}

// UpdateIndexSet updates a given Index Set.
func (client *Client) UpdateIndexSet(is *graylog.IndexSetUpdateParams) (*graylog.IndexSet, *ErrorInfo, error) {
	return client.UpdateIndexSetContext(context.Background(), is)
}

// UpdateIndexSetContext updates a given Index Set with a context.
func (client *Client) UpdateIndexSetContext(
	ctx context.Context, prms *graylog.IndexSetUpdateParams,
) (*graylog.IndexSet, *ErrorInfo, error) {
	if prms == nil {
		return nil, nil, fmt.Errorf("index set is nil")
	}
	if prms.ID == "" {
		return nil, nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().IndexSet(prms.ID)
	if err != nil {
		return nil, nil, err
	}
	a := *prms
	a.ID = ""
	is := &graylog.IndexSet{}
	ei, err := client.callPut(ctx, u.String(), &a, is)
	return is, ei, err
}

// DeleteIndexSet deletes a given Index Set.
func (client *Client) DeleteIndexSet(id string) (*ErrorInfo, error) {
	return client.DeleteIndexSetContext(context.Background(), id)
}

// DeleteIndexSetContext deletes a given Index Set with a context.
func (client *Client) DeleteIndexSetContext(
	ctx context.Context, id string,
) (*ErrorInfo, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().IndexSet(id)
	if err != nil {
		return nil, err
	}
	return client.callDelete(ctx, u.String(), nil, nil)
}

// SetDefaultIndexSet sets default Index Set.
func (client *Client) SetDefaultIndexSet(id string) (
	*graylog.IndexSet, *ErrorInfo, error,
) {
	return client.SetDefaultIndexSetContext(context.Background(), id)
}

// SetDefaultIndexSetContext sets default Index Set with a context.
func (client *Client) SetDefaultIndexSetContext(
	ctx context.Context, id string,
) (*graylog.IndexSet, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	u, err := client.Endpoints().SetDefaultIndexSet(id)
	if err != nil {
		return nil, nil, err
	}
	is := &graylog.IndexSet{}
	ei, err := client.callPut(ctx, u.String(), nil, is)
	return is, ei, err
}
