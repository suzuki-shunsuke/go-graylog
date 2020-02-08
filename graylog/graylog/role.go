package graylog

import (
	"github.com/suzuki-shunsuke/go-ptr"
	"github.com/suzuki-shunsuke/go-set/v6"
)

type (
	// Role represents a role.
	Role struct {
		Name        string `json:"name,omitempty" v-create:"required" v-update:"required"`
		Description string `json:"description,omitempty"`
		// ex. ["clusterconfigentry:read", "users:edit"]
		Permissions set.StrSet `json:"permissions,omitempty" v-create:"required" v-update:"required"`
		ReadOnly    bool       `json:"read_only"`
	}

	// RoleUpdateParams represents Update Role API's parameters.
	RoleUpdateParams struct {
		Name        string  `json:"name,omitempty" v-create:"required" v-update:"required"`
		Description *string `json:"description,omitempty"`
		// ex. ["clusterconfigentry:read", "users:edit"]
		Permissions set.StrSet `json:"permissions,omitempty" v-create:"required" v-update:"required"`
		ReadOnly    *bool      `json:"read_only"`
	}

	// RolesBody represents Get Roles API's response body.
	// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
	RolesBody struct {
		Roles []Role `json:"roles"`
		Total int    `json:"total"`
	}
)

// NewUpdateParams returns Update Role API's parameters.
func (role *Role) NewUpdateParams() *RoleUpdateParams {
	return &RoleUpdateParams{
		Name:        role.Name,
		Description: ptr.PStr(role.Description),
		Permissions: role.Permissions,
		ReadOnly:    ptr.PBool(role.ReadOnly),
	}
}
