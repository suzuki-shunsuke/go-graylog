package plain

import (
	"github.com/suzuki-shunsuke/go-graylog"
)

func defaultLDAPSetting() *graylog.LDAPSetting {
	return &graylog.LDAPSetting{
		LDAPURI: "ldap://localhost:389",
	}
}

// GetLDAPSetting returns a LDAP Setting.
func (store *Store) GetLDAPSetting() (*graylog.LDAPSetting, error) {
	store.imutex.RLock()
	defer store.imutex.RUnlock()
	return store.ldapSetting, nil
}

// UpdateLDAPSetting updates a LDAP Setting at the store.
func (store *Store) UpdateLDAPSetting(prms *graylog.LDAPSettingUpdateParams) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	ls := store.ldapSetting
	if ls == nil {
		ls = defaultLDAPSetting()
	}

	if prms.Enabled != nil {
		ls.Enabled = *prms.Enabled
	}
	if prms.UseStartTLS != nil {
		ls.UseStartTLS = *prms.UseStartTLS
	}
	if prms.TrustAllCertificates != nil {
		ls.TrustAllCertificates = *prms.TrustAllCertificates
	}
	if prms.ActiveDirectory != nil {
		ls.ActiveDirectory = *prms.ActiveDirectory
	}
	if prms.SystemUsername != nil {
		ls.SystemUsername = *prms.SystemUsername
	}
	if prms.SystemPassword != nil {
		ls.SystemPassword = *prms.SystemPassword
	}
	if prms.LDAPURI != nil {
		ls.LDAPURI = *prms.LDAPURI
	}
	if prms.SearchBase != nil {
		ls.SearchBase = *prms.SearchBase
	}
	if prms.SearchPattern != nil {
		ls.SearchPattern = *prms.SearchPattern
	}
	if prms.DisplayNameAttribute != nil {
		ls.DisplayNameAttribute = *prms.DisplayNameAttribute
	}
	if prms.DefaultGroup != nil {
		ls.DefaultGroup = *prms.DefaultGroup
	}
	if prms.GroupSearchBase != nil {
		ls.GroupSearchBase = *prms.GroupSearchBase
	}
	if prms.GroupIDAttribute != nil {
		ls.GroupIDAttribute = *prms.GroupIDAttribute
	}
	if prms.GroupSearchPattern != nil {
		ls.GroupSearchPattern = *prms.GroupSearchPattern
	}

	store.ldapSetting = ls
	return nil
}

// DeleteLDAPSetting deletes a LDAP Setting from store.
func (store *Store) DeleteLDAPSetting() error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	store.ldapSetting = defaultLDAPSetting()
	return nil
}
