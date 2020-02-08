package testdata

import (
	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

func IndexSet() *graylog.IndexSet {
	return &graylog.IndexSet{
		Title:                 "Default index set",
		IndexPrefix:           "graylog",
		RotationStrategyClass: "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy",
		RotationStrategy: &graylog.RotationStrategy{
			Type:            "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
			MaxDocsPerIndex: 20000000,
			RotationPeriod:  "",
			MaxSize:         0,
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
		ID:                              "5d84bf242ab79c000d691b7f",
		Description:                     "The Graylog default index set",
		Replicas:                        0,
		IndexOptimizationDisabled:       false,
		Writable:                        true,
		Default:                         true,
		Stats:                           nil,
	}
}

func IndexSetUpdateParams() *graylog.IndexSetUpdateParams {
	return &graylog.IndexSetUpdateParams{
		Title:       "updated title",
		Description: ptr.PStr("updated description"),

		IndexPrefix:           "1234-test",
		RotationStrategyClass: "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy",
		RotationStrategy: &graylog.RotationStrategy{
			Type:            "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
			MaxDocsPerIndex: 20000000,
			RotationPeriod:  "",
			MaxSize:         0,
		},
		RetentionStrategyClass: "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
		RetentionStrategy: &graylog.RetentionStrategy{
			Type:               "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
			MaxNumberOfIndices: 20,
		},
		IndexAnalyzer:                   "standard",
		Shards:                          4,
		IndexOptimizationMaxNumSegments: 1,
		FieldTypeRefreshInterval:        5000,
		ID:                              "5d84bf242ab79c000d691b7f",
		Replicas:                        ptr.PInt(0),
		IndexOptimizationDisabled:       ptr.PBool(false),
		Writable:                        ptr.PBool(true),
	}
}
