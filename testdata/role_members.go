package testdata

import (
	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

var (
	RoleMembers = &graylog.UsersBody{
		Users: []graylog.User{
			{
				Username:         "admin",
				Email:            "",
				FullName:         "Administrator",
				Password:         "",
				ID:               "local:admin",
				Timezone:         "UTC",
				LastActivity:     "2019-09-21T01:28:20.552+0000",
				ClientAddress:    "172.18.0.1",
				SessionTimeoutMs: 28800000,
				External:         false,
				ReadOnly:         true,
				SessionActive:    true,
				Preferences: &graylog.Preferences{
					UpdateUnfocussed:  false,
					EnableSmartSearch: true,
				},
				Startpage:   nil,
				Roles:       set.NewStrSet("Admin"),
				Permissions: set.NewStrSet("*"),
			},
			{
				Username:         "graylog-sidecar",
				Email:            "sidecar@graylog.local",
				FullName:         "Sidecar System User (built-in)",
				Password:         "",
				ID:               "5d84bfbb2ab79c000d35d402",
				Timezone:         "UTC",
				LastActivity:     "",
				ClientAddress:    "",
				SessionTimeoutMs: 28800000,
				External:         false,
				ReadOnly:         false,
				SessionActive:    false,
				Preferences: &graylog.Preferences{
					UpdateUnfocussed:  false,
					EnableSmartSearch: true,
				},
				Startpage: nil,
				Roles:     set.NewStrSet("Reader", "Sidecar System (Internal)"),
				Permissions: set.NewStrSet(
					"buffers:read",
					"clusterconfigentry:read",
					"decorators:read",
					"fieldnames:read",
					"indexercluster:read",
					"inputs:read",
					"journal:read",
					"jvmstats:read",
					"messagecount:read",
					"messages:analyze",
					"messages:read",
					"metrics:read",
					"savedsearches:create",
					"savedsearches:edit",
					"savedsearches:read",
					"sidecar_collector_configurations:read",
					"sidecar_collectors:read",
					"sidecars:update",
					"system:read",
					"throughput:read",
					"users:edit:graylog-sidecar",
					"users:passwordchange:graylog-sidecar",
					"users:tokencreate:graylog-sidecar",
					"users:tokenlist:graylog-sidecar",
					"users:tokenremove:graylog-sidecar",
				),
			},
			{
				Username:         "test",
				Email:            "test@example.com",
				FullName:         "test test",
				Password:         "",
				ID:               "5d84c1a92ab79c000d35d6cb",
				Timezone:         "",
				LastActivity:     "",
				ClientAddress:    "",
				SessionTimeoutMs: 28800000,
				External:         false,
				ReadOnly:         false,
				SessionActive:    false,
				Preferences: &graylog.Preferences{
					UpdateUnfocussed:  false,
					EnableSmartSearch: true,
				},
				Startpage: nil,
				Roles:     set.NewStrSet("Reader"),
				Permissions: set.NewStrSet(
					"buffers:read",
					"clusterconfigentry:read",
					"decorators:read",
					"fieldnames:read",
					"indexercluster:read",
					"inputs:read",
					"journal:read",
					"jvmstats:read",
					"messagecount:read",
					"messages:analyze",
					"messages:read",
					"metrics:read",
					"savedsearches:create",
					"savedsearches:edit",
					"savedsearches:read",
					"system:read",
					"throughput:read",
					"users:edit:test",
					"users:passwordchange:test",
					"users:tokencreate:test",
					"users:tokenlist:test",
					"users:tokenremove:test",
				),
			},
		},
	}
)
