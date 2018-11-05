package graylog

import (
	"github.com/suzuki-shunsuke/go-ptr"
)

// LDAPSetting represents a ldap settings.
type LDAPSetting struct {
	Enabled                 bool              `json:"enabled"`
	UseStartTLS             bool              `json:"use_start_tls"`
	TrustAllCertificates    bool              `json:"trust_all_certificates"`
	ActiveDirectory         bool              `json:"active_directory"`
	SystemUsername          string            `json:"system_username,omitempty"`
	SystemPassword          string            `json:"system_password"`
	LDAPURI                 string            `json:"ldap_uri"`
	SearchBase              string            `json:"search_base"`
	SearchPattern           string            `json:"search_pattern"`
	DisplayNameAttribute    string            `json:"display_name_attribute"`
	DefaultGroup            string            `json:"default_group"`
	GroupSearchBase         string            `json:"group_search_base"`
	GroupIDAttribute        string            `json:"group_id_attribute"`
	GroupSearchPattern      string            `json:"group_search_pattern"`
	GroupMapping            map[string]string `json:"group_mapping"`
	AdditionalDefaultGroups []string          `json:"additional_default_groups"`
}

// LDAPSettingUpdateParams represents Update LDAP Setting API's parameters.
type LDAPSettingUpdateParams struct {
	Enabled                 *bool             `json:"enabled"`
	UseStartTLS             *bool             `json:"use_start_tls"`
	TrustAllCertificates    *bool             `json:"trust_all_certificates"`
	ActiveDirectory         *bool             `json:"active_directory"`
	SystemUsername          *string           `json:"system_username,omitempty"`
	SystemPassword          *string           `json:"system_password"`
	LDAPURI                 *string           `json:"ldap_uri"`
	SearchBase              *string           `json:"search_base"`
	SearchPattern           *string           `json:"search_pattern"`
	DisplayNameAttribute    *string           `json:"display_name_attribute"`
	DefaultGroup            *string           `json:"default_group"`
	GroupSearchBase         *string           `json:"group_search_base"`
	GroupIDAttribute        *string           `json:"group_id_attribute"`
	GroupSearchPattern      *string           `json:"group_search_pattern"`
	GroupMapping            map[string]string `json:"group_mapping"`
	AdditionalDefaultGroups []string          `json:"additional_default_groups"`
}

// NewUpdateParams returns Update LDAP Setting API's parameters.
func (ls *LDAPSetting) NewUpdateParams() *LDAPSettingUpdateParams {
	return &LDAPSettingUpdateParams{
		Enabled:              ptr.PBool(ls.Enabled),
		UseStartTLS:          ptr.PBool(ls.UseStartTLS),
		TrustAllCertificates: ptr.PBool(ls.TrustAllCertificates),
		ActiveDirectory:      ptr.PBool(ls.ActiveDirectory),
		SystemUsername:       ptr.PStr(ls.SystemUsername),
		SystemPassword:       ptr.PStr(ls.SystemPassword),
		LDAPURI:              ptr.PStr(ls.LDAPURI),
		SearchBase:           ptr.PStr(ls.SearchBase),
		SearchPattern:        ptr.PStr(ls.SearchPattern),
		DisplayNameAttribute: ptr.PStr(ls.DisplayNameAttribute),
		DefaultGroup:         ptr.PStr(ls.DefaultGroup),
		GroupSearchBase:      ptr.PStr(ls.GroupSearchBase),
		GroupIDAttribute:     ptr.PStr(ls.GroupIDAttribute),
		GroupSearchPattern:   ptr.PStr(ls.GroupSearchPattern),
	}
}
