package testdata

import (
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog"
)

var (
	Roles = &graylog.RolesBody{
		Roles: []graylog.Role{
			{
				Name:        "Views Manager",
				Description: "Allows reading and writing all views and extended searches (built-in)",
				Permissions: set.StrSet{
					"extendedsearch:create": struct{}{},
					"extendedsearch:use":    struct{}{},
					"view:create":           struct{}{},
					"view:edit":             struct{}{},
					"view:read":             struct{}{},
					"view:use":              struct{}{},
				},
				ReadOnly: true,
			},
			{
				Name:        "Reader",
				Description: "Grants basic permissions for every Graylog user (built-in)",
				Permissions: set.StrSet{
					"buffers:read":            struct{}{},
					"clusterconfigentry:read": struct{}{},
					"decorators:read":         struct{}{},
					"fieldnames:read":         struct{}{},
					"indexercluster:read":     struct{}{},
					"inputs:read":             struct{}{},
					"journal:read":            struct{}{},
					"jvmstats:read":           struct{}{},
					"messagecount:read":       struct{}{},
					"messages:analyze":        struct{}{},
					"messages:read":           struct{}{},
					"metrics:read":            struct{}{},
					"savedsearches:create":    struct{}{},
					"savedsearches:edit":      struct{}{},
					"savedsearches:read":      struct{}{},
					"system:read":             struct{}{},
					"throughput:read":         struct{}{},
				},
				ReadOnly: true,
			},
			{
				Name:        "terraform",
				Description: "terraform",
				Permissions: set.StrSet{
					"dashboards:*":    struct{}{},
					"indexsets:*":     struct{}{},
					"inputs:*":        struct{}{},
					"pipeline_rule:*": struct{}{},
					"roles:*":         struct{}{},
					"streams:*":       struct{}{},
					"users:*":         struct{}{},
				},
				ReadOnly: false,
			},
			{
				Name:        "Sidecar System (Internal)",
				Description: "Internal technical role. Grants access to register and pull configurations for a Sidecar node (built-in)",
				Permissions: set.StrSet{
					"sidecar_collector_configurations:read": struct{}{},
					"sidecar_collectors:read":               struct{}{},
					"sidecars:update":                       struct{}{},
				},
				ReadOnly: true,
			},
			{
				Name:        "Admin",
				Description: "Grants all permissions for Graylog administrators (built-in)",
				Permissions: set.StrSet{
					"*": struct{}{},
				},
				ReadOnly: true,
			},
			{
				Name:        "Views User",
				Description: "Allows using views and extended searches (built-in)",
				Permissions: set.StrSet{
					"extendedsearch:create": struct{}{},
					"extendedsearch:use":    struct{}{},
					"view:create":           struct{}{},
					"view:use":              struct{}{},
				},
				ReadOnly: true,
			},
			{
				Name:        "Alerts Manager",
				Description: "Allows reading and writing all event definitions and event notifications (built-in)",
				Permissions: set.StrSet{
					"eventdefinitions:create":   struct{}{},
					"eventdefinitions:delete":   struct{}{},
					"eventdefinitions:edit":     struct{}{},
					"eventdefinitions:execute":  struct{}{},
					"eventdefinitions:read":     struct{}{},
					"eventnotifications:create": struct{}{},
					"eventnotifications:delete": struct{}{},
					"eventnotifications:edit":   struct{}{},
					"eventnotifications:read":   struct{}{},
				},
				ReadOnly: true,
			},
			{
				Name:        "terraform-read",
				Description: "terraform-read",
				Permissions: set.StrSet{
					"dashboards:read":    struct{}{},
					"indexsets:read":     struct{}{},
					"inputs:read":        struct{}{},
					"pipeline_rule:read": struct{}{},
					"roles:read":         struct{}{},
					"streams:read":       struct{}{},
					"users:edit":         struct{}{},
					"users:list":         struct{}{},
					"users:tokencreate":  struct{}{},
					"users:tokenlist":    struct{}{},
				},
				ReadOnly: false,
			},
		},
		Total: 8,
	}
)
