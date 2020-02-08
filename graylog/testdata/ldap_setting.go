package testdata

func CreateLDAPSettingMap() map[string]interface{} {
	return map[string]interface{}{
		"enabled":                true,
		"system_username":        "",
		"system_password":        "",
		"ldap_uri":               "ldap://localhost:389",
		"use_start_tls":          false,
		"trust_all_certificates": false,
		"active_directory":       false,
		"search_base":            "",
		"search_pattern":         "",
		"display_name_attribute": "",
		"default_group":          "",
	}
}
