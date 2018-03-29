package graylog

import (
	"github.com/suzuki-shunsuke/go-set"
)

// Role represents a role.
type Role struct {
	Name        string `json:"name,omitempty" v-create:"required" v-update:"required"`
	Description string `json:"description,omitempty"`
	// ex. ["clusterconfigentry:read", "users:edit"]
	Permissions *set.StrSet `json:"permissions,omitempty" v-create:"required" v-update:"required"`
	ReadOnly    bool        `json:"read_only,omitempty"`
}

type RolesBody struct {
	Roles []Role `json:"roles"`
	Total int    `json:"total"`
}
