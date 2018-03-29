package graylog

import (
	"github.com/suzuki-shunsuke/go-set"
)

// User represents a user.
type User struct {
	// ex. "local:admin"
	Username    string      `json:"username,omitempty" v-create:"required" v-update:"required"`
	Email       string      `json:"email,omitempty" v-create:"required"`
	Permissions *set.StrSet `json:"permissions,omitempty" v-create:"required"`
	FullName    string      `json:"full_name,omitempty" v-create:"required"`
	Password    string      `json:"password,omitempty" v-create:"required"`

	ID          string       `json:"id,omitempty"`
	Preferences *Preferences `json:"preferences,omitempty"`
	// ex. "UTC"
	Timezone string `json:"timezone,omitempty"`
	// ex. 28800000
	SessionTimeoutMs int        `json:"session_timeout_ms,omitempty"`
	External         bool       `json:"external,omitempty"`
	Startpage        *Startpage `json:"startpage,omitempty"`
	// ex. ["Admin"]
	Roles         *set.StrSet `json:"roles,omitempty"`
	ReadOnly      bool        `json:"read_only,omitempty"`
	SessionActive bool        `json:"session_active,omitempty"`
	// ex. "2018-03-02T06:32:01.841+0000"
	LastActivity string `json:"last_activity,omitempty"`
	// ex. "192.168.192.1"
	ClientAddress string `json:"client_address,omitempty"`
}

// Preferences represents user's preferences.
type Preferences struct {
	UpdateUnfocussed  bool `json:"updateUnfocussed,omitempty"`
	EnableSmartSearch bool `json:"enableSmartSearch,omitempty"`
}

// Startpage represents a user's startpage.
type Startpage struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}

type UsersBody struct {
	Users []User `json:"users"`
}
