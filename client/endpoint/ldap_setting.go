package endpoint

// LDAPSetting returns the LDAP Setting API's endpoint url.
func (ep *Endpoints) LDAPSetting() string {
	return ep.ldapSetting
}

// LDAPGroups returns the LDAP Setting API's endpoint url.
func (ep *Endpoints) LDAPGroups() string {
	// /system/ldap/groups
	return ep.ldapGroups
}

// LDAPGroupRoleMapping returns the LDAP Group and role mapping API's endpoint url.
func (ep *Endpoints) LDAPGroupRoleMapping() string {
	// /system/ldap/settings/groups
	return ep.ldapGroupRoleMapping
}
