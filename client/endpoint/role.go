package endpoint

// Roles returns a Role API's endpoint url.
func (ep *Endpoints) Roles() string {
	return ep.roles
}

// Role returns a Role API's endpoint url.
func (ep *Endpoints) Role(name string) string {
	return ep.roles + "/" + name
}

// RoleMembers returns given role's member endpoint url.
func (ep *Endpoints) RoleMembers(name string) string {
	return ep.roles + "/" + name + "/members"
}

// RoleMember returns given role member endpoint url.
func (ep *Endpoints) RoleMember(userName, roleName string) string {
	return ep.roles + "/" + roleName + "/members/" + userName
}
