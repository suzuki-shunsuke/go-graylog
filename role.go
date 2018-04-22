package graylog

import (
	"github.com/suzuki-shunsuke/go-ptr"
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

type RoleUpdateParams struct {
	Name        string  `json:"name,omitempty" v-create:"required" v-update:"required"`
	Description *string `json:"description,omitempty"`
	// ex. ["clusterconfigentry:read", "users:edit"]
	Permissions *set.StrSet `json:"permissions,omitempty" v-create:"required" v-update:"required"`
}

func (role *Role) NewUpdateParams() *RoleUpdateParams {
	return &RoleUpdateParams{
		Name:        role.Name,
		Description: ptr.PStr(role.Description),
		Permissions: role.Permissions,
	}
}

type RolesBody struct {
	Roles []Role `json:"roles"`
	Total int    `json:"total"`
}
