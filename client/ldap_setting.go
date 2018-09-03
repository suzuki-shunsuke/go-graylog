package client

import (
	"context"
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog"
)

// GetLDAPSetting returns the LDAP setting.
func (client *Client) GetLDAPSetting() (*graylog.LDAPSetting, *ErrorInfo, error) {
	return client.GetLDAPSettingContext(context.Background())
}

// GetLDAPSettingContext returns the LDAP setting with a context.
func (client *Client) GetLDAPSettingContext(ctx context.Context) (
	*graylog.LDAPSetting, *ErrorInfo, error,
) {
	// GET /system/ldap/settings Get the LDAP configuration if it is configured
	ls := &graylog.LDAPSetting{}
	ei, err := client.callGet(
		ctx, client.Endpoints().LDAPSetting(), nil, ls)
	return ls, ei, err
}

// UpdateLDAPSetting updates the LDAP setting.
func (client *Client) UpdateLDAPSetting(ldapSetting *graylog.LDAPSettingUpdateParams) (
	*ErrorInfo, error,
) {
	return client.UpdateLDAPSettingContext(context.Background(), ldapSetting)
}

// UpdateLDAPSettingContext updates the LDAP setting with a context.
func (client *Client) UpdateLDAPSettingContext(
	ctx context.Context, prms *graylog.LDAPSettingUpdateParams,
) (*ErrorInfo, error) {
	if prms == nil {
		return nil, fmt.Errorf("ldap setting is nil")
	}
	return client.callPut(ctx, client.Endpoints().LDAPSetting(), prms, nil)
}

// DeleteLDAPSetting deletes the LDAP setting.
func (client *Client) DeleteLDAPSetting() (*ErrorInfo, error) {
	return client.DeleteLDAPSettingContext(context.Background())
}

// DeleteLDAPSettingContext deletes the LDAP setting with a context.
func (client *Client) DeleteLDAPSettingContext(
	ctx context.Context,
) (*ErrorInfo, error) {
	return client.callDelete(ctx, client.Endpoints().LDAPSetting(), nil, nil)
}
