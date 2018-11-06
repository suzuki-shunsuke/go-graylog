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
func (store *Store) UpdateLDAPSetting(prms *graylog.LDAPSetting) error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	p := *prms
	store.ldapSetting = &p
	return nil
}

// DeleteLDAPSetting deletes a LDAP Setting from store.
func (store *Store) DeleteLDAPSetting() error {
	store.imutex.Lock()
	defer store.imutex.Unlock()
	store.ldapSetting = defaultLDAPSetting()
	return nil
}
