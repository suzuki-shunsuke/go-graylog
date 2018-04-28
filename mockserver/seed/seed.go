package seed

import (
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-set"
)

// Role returns a Role.
func Role() *graylog.Role {
	return &graylog.Role{
		Name:        "Admin",
		Description: "Grants all permissions for Graylog administrators (built-in)",
		Permissions: set.NewStrSet("*"),
		ReadOnly:    true}
}

// User returns a user.
func User() *graylog.User {
	return &graylog.User{
		Username:    "admin",
		Email:       "hoge@example.com",
		FullName:    "Administrator",
		Password:    "admin",
		Permissions: set.NewStrSet("*"),
	}
}

// Nobody returns a user who has no permission.
// This user is used to test the authorization.
func Nobody() *graylog.User {
	return &graylog.User{
		Username:    "nobody",
		Email:       "nobody@example.com",
		FullName:    "No Body",
		Password:    "password",
		Permissions: set.NewStrSet(),
	}
}

// Input returns an input.
func Input() *graylog.Input {
	return &graylog.Input{
		Title: "test",
		Node:  "2ad6b340-3e5f-4a96-ae81-040cfb8b6024",
		Attrs: &graylog.InputBeatsAttrs{
			BindAddress:    "0.0.0.0",
			Port:           514,
			RecvBufferSize: 262144,
		}}
}

// IndexSet returns an index set.
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

// IndexSetStats returns an index set statistics.
func IndexSetStats() *graylog.IndexSetStats {
	return &graylog.IndexSetStats{
		Indices:   2,
		Documents: 0,
		Size:      1412,
	}
}

// Stream returns a stream.
func Stream() *graylog.Stream {
	return &graylog.Stream{
		MatchingType: "AND",
		Description:  "Stream containing all messages",
		Rules:        []graylog.StreamRule{},
		Title:        "All messages",
	}
}

// StreamRule returns a stream rule.
func StreamRule() *graylog.StreamRule {
	return &graylog.StreamRule{
		Type:  1,
		Value: "test",
		Field: "tag",
	}
}
