package logic

import (
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// GetLDAPSetting returns a LDAP Setting.
func (lgc *Logic) GetLDAPSetting() (*graylog.LDAPSetting, int, error) {
	ls, err := lgc.store.GetLDAPSetting()
	if err != nil {
		return ls, 500, err
	}
	return ls, 200, nil
}

// UpdateLDAPSetting updates a LDAP Setting.
func (lgc *Logic) UpdateLDAPSetting(prms *graylog.LDAPSetting) (int, error) {
	if err := validator.UpdateValidator.Struct(prms); err != nil {
		return 400, err
	}
	if err := lgc.store.UpdateLDAPSetting(prms); err != nil {
		return 500, err
	}
	return 201, nil
}

// DeleteLDAPSetting deletes a LDAP Setting.
func (lgc *Logic) DeleteLDAPSetting() (int, error) {
	if err := lgc.store.DeleteLDAPSetting(); err != nil {
		return 500, err
	}
	return 200, nil
}
