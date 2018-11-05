package testutil

import (
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-ptr"
	"github.com/suzuki-shunsuke/go-set"
)

// Role returns a new Role.
func Role() *graylog.Role {
	return &graylog.Role{
		Name:        "Writer",
		Description: "writer",
		Permissions: set.NewStrSet("*"),
		ReadOnly:    true}
}

// User returns a new User.
func User() *graylog.User {
	return &graylog.User{
		Username:    "foo",
		Email:       "foo@example.com",
		FullName:    "foo bar",
		Password:    "password",
		Permissions: set.NewStrSet("*"),
	}
}

// DummyAdmin returns a new Admin user.
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

// Input returns a new Input.
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

// IndexSet returns a new IndexSet.
func IndexSet(prefix string) *graylog.IndexSet {
	return &graylog.IndexSet{
		Title:                           "Default index set",
		Description:                     "The Graylog default index set",
		IndexPrefix:                     prefix,
		Replicas:                        0,
		RotationStrategyClass:           graylog.MessageCountRotationStrategy,
		RotationStrategy:                graylog.NewMessageCountRotationStrategy(0),
		RetentionStrategyClass:          graylog.DeletionRetentionStrategy,
		RetentionStrategy:               graylog.NewDeletionRetentionStrategy(0),
		IndexOptimizationMaxNumSegments: 1,
		IndexOptimizationDisabled:       false,
		Writable:                        true,
		Default:                         true}
}

// Stream returns a new Stream.
func Stream() *graylog.Stream {
	return &graylog.Stream{
		MatchingType: "AND",
		Description:  "Stream containing all messages",
		Rules:        []graylog.StreamRule{},
		Title:        "All messages",
	}
}

// DummyStream returns a Stream.
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

// StreamRule returns a new StreamRule.
func StreamRule() *graylog.StreamRule {
	return &graylog.StreamRule{
		Type:  1,
		Value: "test",
		Field: "tag",
	}
}

// Dashboard returns a new Dashboard.
func Dashboard() *graylog.Dashboard {
	return &graylog.Dashboard{
		Title:       "test-dashboard",
		Description: "test dashboard",
	}
}

// FullDashboard returns a new Dashboard.
func FullDashboard() *graylog.Dashboard {
	return &graylog.Dashboard{
		Title:       "test-dashboard",
		Description: "test dashboard",
		CreatedAt:   "2018-02-20T11:37:19.305Z",
		Widgets: []graylog.Widget{
			{
				Type:          "STREAM_SEARCH_RESULT_COUNT",
				CreatorUserID: "admin",
				CacheTime:     ptr.PInt(0),
				Config: &graylog.WidgetConfig{
					Timerange: &graylog.Timerange{
						Type:  "",
						Range: 0,
					},
					LowerIsBetter: true,
					Trend:         true,
					Query:         "",
				},
			},
		},
	}
}
