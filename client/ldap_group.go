package client

import (
	"context"
)

// GetLDAPGroups returns the available LDAP groups.
func (client *Client) GetLDAPGroups(ctx context.Context) ([]string, *ErrorInfo, error) {
	// GET /system/ldap/groups Get the available LDAP groups
	groups := []string{}
	ei, err := client.callGet(ctx, client.Endpoints().LDAPGroups(), nil, &groups)
	return groups, ei, err
}

// GetLDAPGroupRoleMapping returns the LDAP group and role mapping.
func (client *Client) GetLDAPGroupRoleMapping(ctx context.Context) (map[string]string, *ErrorInfo, error) {
	// GET /system/ldap/settings/groups Get the LDAP group to Graylog role mapping
	m := map[string]string{}
	ei, err := client.callGet(
		ctx, client.Endpoints().LDAPGroupRoleMapping(), nil, &m)
	return m, ei, err
}

// UpdateLDAPGroupRoleMapping returns the LDAP group and role mapping.
func (client *Client) UpdateLDAPGroupRoleMapping(ctx context.Context, mapping map[string]string) (*ErrorInfo, error) {
	// PUT /system/ldap/settings/groups Update the LDAP group to Graylog role mapping
	return client.callPut(
		ctx, client.Endpoints().LDAPGroupRoleMapping(), mapping, nil)
}
