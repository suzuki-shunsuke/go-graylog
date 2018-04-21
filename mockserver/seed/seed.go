package seed

import (
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-set"
)

func Role() *graylog.Role {
	return &graylog.Role{
		Name:        "Admin",
		Description: "Grants all permissions for Graylog administrators (built-in)",
		Permissions: set.NewStrSet("*"),
		ReadOnly:    true}
}

func User() *graylog.User {
	return &graylog.User{
		Username:    "admin",
		Email:       "hoge@example.com",
		FullName:    "Administrator",
		Password:    "admin",
		Permissions: set.NewStrSet("*"),
	}
}

func Nobody() *graylog.User {
	return &graylog.User{
		Username:    "nobody",
		Email:       "nobody@example.com",
		FullName:    "No Body",
		Password:    "password",
		Permissions: set.NewStrSet(),
	}
}

func Input() *graylog.Input {
	ba := "0.0.0.0"
	port := 514
	rbs := 262144
	return &graylog.Input{
		Title: "test",
		Type:  "org.graylog2.inputs.gelf.tcp.GELFTCPInput",
		Node:  "2ad6b340-3e5f-4a96-ae81-040cfb8b6024",
		Attributes: &graylog.InputAttributes{
			BindAddress:    &ba,
			Port:           &port,
			RecvBufferSize: &rbs,
		}}
}

func IndexSet() *graylog.IndexSet {
	return &graylog.IndexSet{
		Title:                 "Default index set",
		Description:           "The Graylog default index set",
		IndexPrefix:           "graylog",
		Shards:                4,
		Replicas:              0,
		RotationStrategyClass: "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy",
		RotationStrategy: &graylog.RotationStrategy{
			Type:            "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
			MaxDocsPerIndex: 20000000},
		RetentionStrategyClass: "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
		RetentionStrategy: &graylog.RetentionStrategy{
			Type:               "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
			MaxNumberOfIndices: 20},
		CreationDate:                    "2018-02-20T11:37:19.305Z",
		IndexAnalyzer:                   "standard",
		IndexOptimizationMaxNumSegments: 1,
		IndexOptimizationDisabled:       false,
		Writable:                        true,
		Default:                         true}
}

func IndexSetStats() *graylog.IndexSetStats {
	return &graylog.IndexSetStats{
		Indices:   2,
		Documents: 0,
		Size:      1412,
	}
}

func Stream() *graylog.Stream {
	return &graylog.Stream{
		MatchingType: "AND",
		Description:  "Stream containing all messages",
		Rules:        []graylog.StreamRule{},
		Title:        "All messages",
	}
}

func StreamRule() *graylog.StreamRule {
	return &graylog.StreamRule{
		Type:  1,
		Value: "test",
		Field: "tag",
	}
}
