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

// RoleUpdateParams represents Update Role API's parameteres.
type RoleUpdateParams struct {
	Name        string  `json:"name,omitempty" v-create:"required" v-update:"required"`
	Description *string `json:"description,omitempty"`
	// ex. ["clusterconfigentry:read", "users:edit"]
	Permissions *set.StrSet `json:"permissions,omitempty" v-create:"required" v-update:"required"`
}

// NewUpdateParams returns Update Role API's parameteres.
func (role *Role) NewUpdateParams() *RoleUpdateParams {
	return &RoleUpdateParams{
		Name:        role.Name,
		Description: ptr.PStr(role.Description),
		Permissions: role.Permissions,
	}
}

// RolesBody represents Get Roles API's response body.
// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
type RolesBody struct {
	Roles []Role `json:"roles"`
	Total int    `json:"total"`
}
