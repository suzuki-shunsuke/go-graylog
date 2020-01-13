package graylog

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/v9/testdata"
)

func TestAccIndexSet(t *testing.T) {
	setEnv()

	indexSet := testdata.IndexSet()

	creationDate := "2019-09-20T11:59:32.219Z"

	tc := &testCase{
		t:          t,
		Name:       "index set",
		CreatePath: "/api/system/indices/index_sets",
		GetPath:    "/api/system/indices/index_sets/" + indexSet.ID,

		ConvertReqBody: func(body io.Reader) (map[string]interface{}, error) {
			m := map[string]interface{}{}
			if err := json.NewDecoder(body).Decode(&m); err != nil {
				return nil, err
			}
			if _, ok := m["creation_date"]; ok {
				m["creation_date"] = creationDate
			}
			return m, nil
		},

		CreateReqBodyMap: map[string]interface{}{
			"title":       "test",
			"description": "The Graylog default index set",

			"index_prefix":            "1234-test",
			"rotation_strategy_class": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy",
			"rotation_strategy": map[string]interface{}{
				"type":               "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
				"max_docs_per_index": float64(20000000),
			},
			"retention_strategy_class": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
			"retention_strategy": map[string]interface{}{
				"type":                  "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
				"max_number_of_indices": float64(20),
			},
			"creation_date":                       creationDate,
			"index_analyzer":                      "standard",
			"shards":                              float64(4),
			"index_optimization_max_num_segments": float64(1),
			"field_type_refresh_interval":         float64(5000),
			"writable":                            true,
			"index_optimization_disabled":         false,
			"default":                             false,
		},
		UpdateReqBodyMap: map[string]interface{}{
			"title":       "updated title",
			"description": "updated description",

			"index_prefix":            "1234-test",
			"rotation_strategy_class": "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy",
			"rotation_strategy": map[string]interface{}{
				"type":               "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
				"max_docs_per_index": float64(20000000),
			},
			"retention_strategy_class": "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
			"retention_strategy": map[string]interface{}{
				"type":                  "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
				"max_number_of_indices": float64(20),
			},
			"index_analyzer":                      "standard",
			"shards":                              float64(4),
			"index_optimization_max_num_segments": float64(1),
			"replicas":                            float64(0),
			"field_type_refresh_interval":         float64(5000),
			"writable":                            true,
			"index_optimization_disabled":         false,
		},
		CreatedDataPath:    "index_set/create_index_set_response.json",
		UpdatedDataPath:    "index_set/update_response.json",
		CreateRespBodyPath: "index_set/create_index_set_response.json",
		CreateTFPath:       "index_set/create.tf",
		UpdateTFPath:       "index_set/update.tf",

		UpdateRespBodyPath: "index_set/update_response.json",
	}
	tc.Test()
}
