package graylog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// IndexSet represents a Graylog's Index Set.
type IndexSet struct {
	// required
	Title       string `json:"title,omitempty" v-create:"required" v-update:"required"`
	IndexPrefix string `json:"index_prefix,omitempty" v-create:"required" v-update:"required"`
	// ex. "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy"
	RotationStrategyClass string            `json:"rotation_strategy_class,omitempty" v-create:"required" v-update:"required"`
	RotationStrategy      *RotationStrategy `json:"rotation_strategy,omitempty" v-create:"required" v-update:"required"`
	// ex. "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy"
	RetentionStrategyClass string             `json:"retention_strategy_class,omitempty" v-create:"required" v-update:"required"`
	RetentionStrategy      *RetentionStrategy `json:"retention_strategy,omitempty" v-create:"required" v-update:"required"`
	// ex. "2018-02-20T11:37:19.305Z"
	CreationDate                    string `json:"creation_date,omitempty" v-create:"required" v-update:"required"`
	IndexAnalyzer                   string `json:"index_analyzer,omitempty" v-create:"required" v-update:"required"`
	Shards                          int    `json:"shards,omitempty" v-create:"required" v-update:"required"`
	IndexOptimizationMaxNumSegments int    `json:"index_optimization_max_num_segments,omitempty" v-create:"required" v-update:"required"`

	ID string `json:"id,omitempty" v-create:"isdefault" v-update:"required"`

	Description               string `json:"description,omitempty"`
	Replicas                  int    `json:"replicas,omitempty"`
	IndexOptimizationDisabled bool   `json:"index_optimization_disabled,omitempty"`
	Writable                  bool   `json:"writable,omitempty"`
	Default                   bool   `json:"default,omitempty"`
}

// RotationStrategy represents a Graylog's Index Set Rotation Strategy.
type RotationStrategy struct {
	// ex. "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"
	Type string `json:"type,omitempty"`
	// ex. 20000000
	MaxDocsPerIndex int `json:"max_docs_per_index,omitempty"`
}

// RetentionStrategy represents a Graylog's Index Set Retention Strategy.
type RetentionStrategy struct {
	// ex. "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"
	Type string `json:"type,omitempty"`
	// ex. 20
	MaxNumberOfIndices int `json:"max_number_of_indices,omitempty"`
}

type IndexSetsBody struct {
	IndexSets []IndexSet     `json:"index_sets"`
	Stats     *IndexSetStats `json:"stats"`
	Total     int            `json:"total"`
}

// GetIndexSets returns a list of all index sets.
func (client *Client) GetIndexSets(
	skip, limit int,
) ([]IndexSet, *IndexSetStats, *ErrorInfo, error) {
	return client.GetIndexSetsContext(context.Background(), skip, limit)
}

// GetIndexSetsContext returns a list of all index sets with a context.
func (client *Client) GetIndexSetsContext(
	ctx context.Context, skip, limit int,
) ([]IndexSet, *IndexSetStats, *ErrorInfo, error) {
	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.IndexSets, nil, true)
	if err != nil {
		return nil, nil, ei, err
	}
	indexSets := &IndexSetsBody{}
	err = json.Unmarshal(ei.ResponseBody, indexSets)
	if err != nil {
		return nil, nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSetsBody: %s",
				string(ei.ResponseBody)))
	}
	return indexSets.IndexSets, indexSets.Stats, ei, nil
}

// GetIndexSet returns a given index set.
func (client *Client) GetIndexSet(id string) (*IndexSet, *ErrorInfo, error) {
	return client.GetIndexSetContext(context.Background(), id)
}

// GetIndexSetContext returns a given index set with a context.
func (client *Client) GetIndexSetContext(
	ctx context.Context, id string,
) (*IndexSet, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}
	ei, err := client.callReq(
		ctx, http.MethodGet, client.Endpoints.IndexSet(id), nil, true)
	if err != nil {
		return nil, ei, err
	}
	indexSet := &IndexSet{}
	if err := json.Unmarshal(ei.ResponseBody, indexSet); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s",
				string(ei.ResponseBody)))
	}
	return indexSet, ei, nil
}

// CreateIndexSet creates a Index Set.
func (client *Client) CreateIndexSet(indexSet *IndexSet) (
	*ErrorInfo, error,
) {
	return client.CreateIndexSetContext(context.Background(), indexSet)
}

// CreateIndexSetContext creates a Index Set with a context.
func (client *Client) CreateIndexSetContext(
	ctx context.Context, indexSet *IndexSet,
) (*ErrorInfo, error) {
	b, err := json.Marshal(indexSet)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(indexSet)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPost, client.Endpoints.IndexSets, b, true)
	if err != nil {
		return ei, err
	}

	if err := json.Unmarshal(ei.ResponseBody, indexSet); err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s",
				string(ei.ResponseBody)))
	}
	return ei, nil
}

// UpdateIndexSet updates a given Index Set.
func (client *Client) UpdateIndexSet(indexSet *IndexSet) (*ErrorInfo, error) {
	return client.UpdateIndexSetContext(context.Background(), indexSet)
}

// UpdateIndexSetContext updates a given Index Set with a context.
func (client *Client) UpdateIndexSetContext(
	ctx context.Context, is *IndexSet,
) (*ErrorInfo, error) {
	id := is.ID
	if id == "" {
		return nil, errors.New("id is empty")
	}
	copiedIndexSet := *is
	copiedIndexSet.ID = ""
	b, err := json.Marshal(&copiedIndexSet)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(indexSet)")
	}

	ei, err := client.callReq(
		ctx, http.MethodPut, client.Endpoints.IndexSet(id), b, true)
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
	*IndexSet, *ErrorInfo, error,
) {
	return client.SetDefaultIndexSetContext(context.Background(), id)
}

// SetDefaultIndexSet sets default Index Set with a context.
func (client *Client) SetDefaultIndexSetContext(
	ctx context.Context, id string,
) (*IndexSet, *ErrorInfo, error) {
	if id == "" {
		return nil, nil, errors.New("id is empty")
	}

	ei, err := client.callReq(
		ctx, http.MethodPut, client.Endpoints.SetDefaultIndexSet(id), nil, true)
	if err != nil {
		return nil, ei, err
	}

	is := &IndexSet{}
	if err := json.Unmarshal(ei.ResponseBody, is); err != nil {
		return nil, ei, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s",
				string(ei.ResponseBody)))
	}
	return is, ei, nil
}
