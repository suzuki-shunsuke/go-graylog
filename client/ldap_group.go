package client

import (
	"context"
)

// GetLDAPGroups returns the available LDAP groups.
func (client *Client) GetLDAPGroups() ([]string, *ErrorInfo, error) {
	return client.GetLDAPGroupsContext(context.Background())
}

// GetLDAPGroupsContext returns the available LDAP groups with a context.
func (client *Client) GetLDAPGroupsContext(ctx context.Context) ([]string, *ErrorInfo, error) {
	// GET /system/ldap/groups Get the available LDAP groups
	groups := []string{}
	ei, err := client.callGet(
		ctx, client.Endpoints().LDAPGroups(), nil, &groups)
	return groups, ei, err
}

// GetLDAPGroupRoleMapping returns the LDAP group and role mapping.
func (client *Client) GetLDAPGroupRoleMapping() (map[string]string, *ErrorInfo, error) {
	return client.GetLDAPGroupRoleMappingContext(context.Background())
}

// GetLDAPGroupRoleMappingContext returns the LDAP group and role mapping with a context.
func (client *Client) GetLDAPGroupRoleMappingContext(ctx context.Context) (map[string]string, *ErrorInfo, error) {
	// GET /system/ldap/settings/groups Get the LDAP group to Graylog role mapping
	m := map[string]string{}
	ei, err := client.callGet(
		ctx, client.Endpoints().LDAPGroupRoleMapping(), nil, &m)
	return m, ei, err
}

// UpdateLDAPGroupRoleMapping updates the LDAP group and role mapping.
func (client *Client) UpdateLDAPGroupRoleMapping(mapping map[string]string) (*ErrorInfo, error) {
	return client.UpdateLDAPGroupRoleMappingContext(context.Background(), mapping)
}

// UpdateLDAPGroupRoleMappingContext returns the LDAP group and role mapping with a context.
func (client *Client) UpdateLDAPGroupRoleMappingContext(ctx context.Context, mapping map[string]string) (*ErrorInfo, error) {
	// PUT /system/ldap/settings/groups Update the LDAP group to Graylog role mapping
	return client.callPut(
		ctx, client.Endpoints().LDAPGroupRoleMapping(), mapping, nil)
}
