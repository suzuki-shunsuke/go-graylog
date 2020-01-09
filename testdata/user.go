package testdata

import (
	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

var (
	User = &graylog.User{
		Username:         "admin",
		Email:            "",
		FullName:         "Administrator",
		Password:         "",
		ID:               "local:admin",
		Timezone:         "UTC",
		LastActivity:     "",
		ClientAddress:    "",
		SessionTimeoutMs: 28800000,
		External:         false,
		ReadOnly:         true,
		SessionActive:    false,
		Preferences: &graylog.Preferences{
			UpdateUnfocussed:  false,
			EnableSmartSearch: true,
		},
		Startpage: nil,
		Roles: set.StrSet{
			"Admin": struct{}{},
		},
		Permissions: set.StrSet{
			"*": struct{}{},
		},
	}
)
