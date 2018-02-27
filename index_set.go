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

type IndexSet struct {
	Id                              string             `json:"id,omitempty"`
	Title                           string             `json:"title,omitempty"`
	Description                     string             `json:"description,omitempty"`
	IndexPrefix                     string             `json:"index_prefix,omitempty"`
	Shards                          int                `json:"shards,omitempty"`
	Replicas                        int                `json:"replicas,omitempty"`
	RotationStrategyClass           string             `json:"rotation_strategy_class,omitempty"`
	RotationStrategy                *RotationStrategy  `json:"rotation_strategy,omitempty"`
	RetentionStrategyClass          string             `json:"retention_strategy_class,omitempty"`
	RetentionStrategy               *RetentionStrategy `json:"retention_strategy,omitempty"`
	CreationDate                    string             `json:"creation_date,omitempty"`
	IndexAnalyzer                   string             `json:"index_analyzer,omitempty"`
	IndexOptimizationMaxNumSegments int                `json:"index_optimization_max_num_segments,omitempty"`
	IndexOptimizationDisabled       bool               `json:"index_optimization_disabled,omitempty"`
	Writable                        bool               `json:"writable,omitempty"`
	Default                         bool               `json:"default,omitempty"`
}

type IndexSetStats struct {
	Indices   int `json:"indices"`
	Documents int `json:"documents"`
	Size      int `json:"size"`
}

type RotationStrategy struct {
	Type            string `json:"type,omitempty"`
	MaxDocsPerIndex int    `json:"max_docs_per_index,omitempty"`
}

type RetentionStrategy struct {
	Type               string `json:"type,omitempty"`
	MaxNumberOfIndices int    `json:"max_number_of_indices,omitempty"`
}

type indexSetsBody struct {
	IndexSets []IndexSet     `json:"index_sets"`
	Stats     *IndexSetStats `json:"stats"`
	Total     int            `json:"total"`
}

// GET /system/indices/index_sets Get a list of all index sets
func (client *Client) GetIndexSets(
	skip, limit int,
) ([]IndexSet, *IndexSetStats, error) {
	return client.GetIndexSetsContext(context.Background(), skip, limit)
}

// GET /system/indices/index_sets Get a list of all index sets
func (client *Client) GetIndexSetsContext(
	ctx context.Context, skip, limit int,
) ([]IndexSet, *IndexSetStats, error) {
	req, err := http.NewRequest(http.MethodGet, client.endpoints.IndexSets, nil)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, nil, errors.Wrap(
			err, "Failed to call GET /system/indices/index_sets API")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to read response body")
	}
	if resp.StatusCode >= 400 {
		e := Error{}
		err = json.Unmarshal(b, &e)
		if err != nil {
			return nil, nil, errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as Error: %s", string(b)))
		}
		return nil, nil, errors.New(e.Message)
	}
	indexSets := &indexSetsBody{}
	err = json.Unmarshal(b, indexSets)
	if err != nil {
		return nil, nil, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as indexSetsBody: %s", string(b)))
	}
	return indexSets.IndexSets, indexSets.Stats, nil
}

// GET /system/indices/index_sets/{id} Get index set
func (client *Client) GetIndexSet(id string) (*IndexSet, error) {
	return client.GetIndexSetContext(context.Background(), id)
}

// GET /system/indices/index_sets/{id} Get index set
func (client *Client) GetIndexSetContext(
	ctx context.Context, id string,
) (*IndexSet, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	req, err := http.NewRequest(
		http.MethodGet, fmt.Sprintf("%s/%s", client.endpoints.IndexSets, id), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(
			err, "Failed to call GET /system/indices/index_sets/{id} API")
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
	indexSet := &IndexSet{}
	err = json.Unmarshal(b, indexSet)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s", string(b)))
	}
	return indexSet, nil
}

// POST /system/indices/index_sets Create index set
func (client *Client) CreateIndexSet(indexSet *IndexSet) (*IndexSet, error) {
	return client.CreateIndexSetContext(context.Background(), indexSet)
}

// POST /system/indices/index_sets Create index set
func (client *Client) CreateIndexSetContext(
	ctx context.Context, indexSet *IndexSet,
) (*IndexSet, error) {
	b, err := json.Marshal(indexSet)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(indexSet)")
	}
	req, err := http.NewRequest(
		http.MethodPost, client.endpoints.IndexSets, bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(
			err, "Failed to call POST /system/indices/index_sets API")
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
	is := &IndexSet{}
	err = json.Unmarshal(b, is)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s", string(b)))
	}
	return is, nil
}

// PUT /system/indices/index_sets/{id} Update index set
func (client *Client) UpdateIndexSet(
	id string, indexSet *IndexSet,
) (*IndexSet, error) {
	return client.UpdateIndexSetContext(context.Background(), id, indexSet)
}

// PUT /system/indices/index_sets/{id} Update index set
func (client *Client) UpdateIndexSetContext(
	ctx context.Context, id string, indexSet *IndexSet,
) (*IndexSet, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	b, err := json.Marshal(indexSet)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to json.Marshal(indexSet)")
	}
	req, err := http.NewRequest(
		http.MethodPut, fmt.Sprintf("%s/%s", client.endpoints.IndexSets, id),
		bytes.NewBuffer(b))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(
			err, "Failed to call PUT /indexSets/{indexSetid} API")
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
	is := &IndexSet{}
	err = json.Unmarshal(b, is)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s", string(b)))
	}
	return is, nil
}

// DELETE /system/indices/index_sets/{id} Delete index set
func (client *Client) DeleteIndexSet(id string) error {
	return client.DeleteIndexSetContext(context.Background(), id)
}

// DELETE /system/indices/index_sets/{id} Delete index set
func (client *Client) DeleteIndexSetContext(
	ctx context.Context, id string,
) error {
	if id == "" {
		return errors.New("id is empty")
	}
	req, err := http.NewRequest(
		http.MethodDelete, fmt.Sprintf(
			"%s/%s", client.endpoints.IndexSets, id), nil)
	if err != nil {
		return errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return errors.Wrap(err, "Failed to call DELETE /indexSets API")
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

// PUT /system/indices/index_sets/{id}/default Set default index set
func (client *Client) SetDefaultIndexSet(id string) (*IndexSet, error) {
	return client.SetDefaultIndexSetContext(context.Background(), id)
}

// PUT /system/indices/index_sets/{id}/default Set default index set
func (client *Client) SetDefaultIndexSetContext(
	ctx context.Context, id string,
) (*IndexSet, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	req, err := http.NewRequest(
		http.MethodPut, fmt.Sprintf(
			"%s/%s/default", client.endpoints.IndexSets, id), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(
			err, "Failed to call PUT /system/indices/index_sets/{indexSetid} API")
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
	is := &IndexSet{}
	err = json.Unmarshal(b, is)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSet: %s", string(b)))
	}
	return is, nil
}

// GET /system/indices/index_sets/{id}/stats Get index set statistics
func (client *Client) GetIndexSetStats(id string) (*IndexSetStats, error) {
	return client.GetIndexSetStatsContext(context.Background(), id)
}

// GET /system/indices/index_sets/{id}/stats Get index set statistics
func (client *Client) GetIndexSetStatsContext(
	ctx context.Context, id string,
) (*IndexSetStats, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	req, err := http.NewRequest(
		http.MethodGet, fmt.Sprintf(
			"%s/%s/stats", client.endpoints.IndexSets, id), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(
			err, "Failed to call GET /system/indices/index_sets/{id}/stats API")
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
	indexSetStats := &IndexSetStats{}
	err = json.Unmarshal(b, indexSetStats)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSetStats: %s", string(b)))
	}
	return indexSetStats, nil
}

// GET /system/indices/index_sets/stats Get stats of all index sets
func (client *Client) GetAllIndexSetsStats() (*IndexSetStats, error) {
	return client.GetAllIndexSetsStatsContext(context.Background())
}

// GET /system/indices/index_sets/stats Get stats of all index sets
func (client *Client) GetAllIndexSetsStatsContext(
	ctx context.Context,
) (*IndexSetStats, error) {
	req, err := http.NewRequest(
		http.MethodGet, fmt.Sprintf(
			"%s/stats", client.endpoints.IndexSets), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	resp, err := callRequest(req, client, ctx)
	if err != nil {
		return nil, errors.Wrap(
			err, "Failed to call GET /system/indices/index_sets/stats API")
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
	indexSetStats := &IndexSetStats{}
	err = json.Unmarshal(b, indexSetStats)
	if err != nil {
		return nil, errors.Wrap(
			err, fmt.Sprintf(
				"Failed to parse response body as IndexSetStats: %s", string(b)))
	}
	return indexSetStats, nil
}
