package testdata

import (
	"github.com/suzuki-shunsuke/go-set/v6"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

var (
	Roles = &graylog.RolesBody{
		Roles: []graylog.Role{
			{
				Name:        "Views Manager",
				Description: "Allows reading and writing all views and extended searches (built-in)",
				Permissions: set.NewStrSet(
					"extendedsearch:create",
					"extendedsearch:use",
					"view:create",
					"view:edit",
					"view:read",
					"view:use",
				),
				ReadOnly: true,
			},
			{
				Name:        "Reader",
				Description: "Grants basic permissions for every Graylog user (built-in)",
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
				),
				ReadOnly: true,
			},
			{
				Name:        "terraform",
				Description: "terraform",
				Permissions: set.NewStrSet(
					"dashboards:*",
					"indexsets:*",
					"inputs:*",
					"pipeline_rule:*",
					"roles:*",
					"streams:*",
					"users:*",
				),
				ReadOnly: false,
			},
			{
				Name:        "Sidecar System (Internal)",
				Description: "Internal technical role. Grants access to register and pull configurations for a Sidecar node (built-in)",
				Permissions: set.NewStrSet(
					"sidecar_collector_configurations:read",
					"sidecar_collectors:read",
					"sidecars:update",
				),
				ReadOnly: true,
			},
			{
				Name:        "Admin",
				Description: "Grants all permissions for Graylog administrators (built-in)",
				Permissions: set.NewStrSet("*"),
				ReadOnly:    true,
			},
			{
				Name:        "Views User",
				Description: "Allows using views and extended searches (built-in)",
				Permissions: set.NewStrSet(
					"extendedsearch:create",
					"extendedsearch:use",
					"view:create",
					"view:use",
				),
				ReadOnly: true,
			},
			{
				Name:        "Alerts Manager",
				Description: "Allows reading and writing all event definitions and event notifications (built-in)",
				Permissions: set.NewStrSet(
					"eventdefinitions:create",
					"eventdefinitions:delete",
					"eventdefinitions:edit",
					"eventdefinitions:execute",
					"eventdefinitions:read",
					"eventnotifications:create",
					"eventnotifications:delete",
					"eventnotifications:edit",
					"eventnotifications:read",
				),
				ReadOnly: true,
			},
			{
				Name:        "terraform-read",
				Description: "terraform-read",
				Permissions: set.NewStrSet(
					"dashboards:read",
					"indexsets:read",
					"inputs:read",
					"pipeline_rule:read",
					"roles:read",
					"streams:read",
					"users:edit",
					"users:list",
					"users:tokencreate",
					"users:tokenlist",
				),
				ReadOnly: false,
			},
		},
		Total: 8,
	}
)
