package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v9"
)

func CreateIndexSet() graylog.IndexSet {
	return graylog.IndexSet{
		Title:                 "test",
		IndexPrefix:           "1234-test",
		RotationStrategyClass: "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy",
		RotationStrategy: &graylog.RotationStrategy{
			Type:            "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
			MaxDocsPerIndex: 20000000,
		},
		RetentionStrategyClass: "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
		RetentionStrategy: &graylog.RetentionStrategy{
			Type:               "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
			MaxNumberOfIndices: 20,
		},
		CreationDate:                    "2019-09-20T11:59:32.219Z",
		IndexAnalyzer:                   "standard",
		Shards:                          4,
		IndexOptimizationMaxNumSegments: 1,
		FieldTypeRefreshInterval:        5000,
		Description:                     "The Graylog default index set",
		Writable:                        true,
	}
}
