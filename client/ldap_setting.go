package client

import (
	"context"
	"errors"

	"github.com/suzuki-shunsuke/go-graylog/v10"
)

// GetLDAPSetting returns the LDAP setting.
func (client *Client) GetLDAPSetting(ctx context.Context) (
	*graylog.LDAPSetting, *ErrorInfo, error,
) {
	// GET /system/ldap/settings Get the LDAP configuration if it is configured
	ls := &graylog.LDAPSetting{}
	ei, err := client.callGet(
		ctx, client.Endpoints().LDAPSetting(), nil, ls)
	return ls, ei, err
}

// UpdateLDAPSetting updates the LDAP setting.
func (client *Client) UpdateLDAPSetting(
	ctx context.Context, prms *graylog.LDAPSetting,
) (*ErrorInfo, error) {
	if prms == nil {
		return nil, errors.New("ldap setting is nil")
	}
	return client.callPut(ctx, client.Endpoints().LDAPSetting(), prms, nil)
}

// DeleteLDAPSetting deletes the LDAP setting.
func (client *Client) DeleteLDAPSetting(ctx context.Context) (*ErrorInfo, error) {
	return client.callDelete(ctx, client.Endpoints().LDAPSetting(), nil, nil)
}
