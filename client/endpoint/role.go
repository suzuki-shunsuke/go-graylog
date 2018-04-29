package endpoint

import (
	"net/url"
	"path"
)

// Roles returns a Role API's endpoint url.
func (ep *Endpoints) Roles() string {
	return ep.roles.String()
}

// Role returns a Role API's endpoint url.
func (ep *Endpoints) Role(name string) (*url.URL, error) {
	return urlJoin(ep.roles, name)
}

// RoleMembers returns given role's member endpoint url.
func (ep *Endpoints) RoleMembers(name string) (*url.URL, error) {
	return urlJoin(ep.roles, path.Join(name, "members"))
}

// RoleMember returns given role member endpoint url.
func (ep *Endpoints) RoleMember(userName, roleName string) (*url.URL, error) {
	return urlJoin(ep.roles, path.Join(roleName, "members", userName))
}
