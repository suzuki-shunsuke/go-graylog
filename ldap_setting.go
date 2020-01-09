package graylog

import (
	"github.com/suzuki-shunsuke/go-set/v6"
)

// LDAPSetting represents a ldap settings.
type LDAPSetting struct {
	Enabled                 bool              `json:"enabled"`
	UseStartTLS             bool              `json:"use_start_tls"`
	TrustAllCertificates    bool              `json:"trust_all_certificates"`
	ActiveDirectory         bool              `json:"active_directory"`
	SystemUsername          string            `json:"system_username" v-create:"required" v-update:"required"`
	SystemPassword          string            `json:"system_password" v-create:"required" v-update:"required"`
	LDAPURI                 string            `json:"ldap_uri" v-create:"required" v-update:"required"`
	SearchBase              string            `json:"search_base" v-create:"required" v-update:"required"`
	SearchPattern           string            `json:"search_pattern" v-create:"required" v-update:"required"`
	DisplayNameAttribute    string            `json:"display_name_attribute" v-create:"required" v-update:"required"`
	DefaultGroup            string            `json:"default_group" v-create:"required" v-update:"required"`
	GroupSearchBase         string            `json:"group_search_base,omitempty"`
	GroupIDAttribute        string            `json:"group_id_attribute,omitempty"`
	GroupSearchPattern      string            `json:"group_search_pattern,omitempty"`
	GroupMapping            map[string]string `json:"group_mapping,omitempty"`
	AdditionalDefaultGroups set.StrSet        `json:"additional_default_groups,omitempty"`
}
