package testutil

import (
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-set"
)

func Role() *graylog.Role {
	return &graylog.Role{
		Name:        "Writer",
		Description: "writer",
		Permissions: set.NewStrSet("*"),
		ReadOnly:    true}
}

func User() *graylog.User {
	return &graylog.User{
		Username:    "foo",
		Email:       "foo@example.com",
		FullName:    "foo bar",
		Password:    "password",
		Permissions: set.NewStrSet("*"),
	}
}

func DummyAdmin() *graylog.User {
	return &graylog.User{
		ID:          "local:admin",
		Username:    "admin",
		Email:       "hoge@example.com",
		FullName:    "Administrator",
		Password:    "password",
		Permissions: set.NewStrSet("*"),
		Preferences: &graylog.Preferences{
			UpdateUnfocussed:  false,
			EnableSmartSearch: true,
		},
		Timezone:         "UTC",
		SessionTimeoutMs: 28800000,
		External:         false,
		Startpage:        nil,
		Roles:            set.NewStrSet("Admin"),
		ReadOnly:         true,
		SessionActive:    true,
		LastActivity:     "2018-02-21T07:35:45.926+0000",
		ClientAddress:    "172.18.0.1",
	}
}

func Input() *graylog.Input {
	return &graylog.Input{
		Title: "test",
		Type:  "org.graylog2.inputs.gelf.tcp.GELFTCPInput",
		Node:  "2ad6b340-3e5f-4a96-ae81-040cfb8b6024",
		Configuration: &graylog.InputConfiguration{
			BindAddress:    "0.0.0.0",
			Port:           514,
			RecvBufferSize: 262144,
		}}
}

func IndexSet(prefix string) *graylog.IndexSet {
	return &graylog.IndexSet{
		Title:                 "Default index set",
		Description:           "The Graylog default index set",
		IndexPrefix:           prefix,
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

func DummyIndexSetStats() *graylog.IndexSetStats {
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

func DummyStream() *graylog.Stream {
	return &graylog.Stream{
		ID:              "000000000000000000000001",
		CreatorUserID:   "local:admin",
		Outputs:         []graylog.Output{},
		MatchingType:    "AND",
		Description:     "Stream containing all messages",
		CreatedAt:       "2018-02-20T11:37:19.371Z",
		Rules:           []graylog.StreamRule{},
		AlertConditions: []graylog.AlertCondition{},
		AlertReceivers: &graylog.AlertReceivers{
			Emails: []string{},
			Users:  []string{},
		},
		Title:      "All messages",
		IndexSetID: "5a8c086fc006c600013ca6f5",
		// "content_pack": null,
	}
}

func StreamRule() *graylog.StreamRule {
	return &graylog.StreamRule{
		Type:  1,
		Value: "test",
		Field: "tag",
	}
}

func DummyNewStreamRule() *graylog.StreamRule {
	return &graylog.StreamRule{
		Type:  1,
		Value: "test",
		Field: "tag",
	}
}

func DummyStreamRule() *graylog.StreamRule {
	return &graylog.StreamRule{
		ID:       "5a9b53c7c006c6000127f965",
		Type:     1,
		Value:    "test",
		StreamID: "5a94abdac006c60001f04fc1",
		Field:    "tag",
	}
}
