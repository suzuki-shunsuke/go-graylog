package endpoint

// User returns a User API's endpoint url.
func (ep *Endpoints) User(name string) string {
	return ep.users + "/" + name
}

// Users returns a User API's endpoint url.
func (ep *Endpoints) Users() string {
	return ep.users
}

// UserTokens returns a User token API's endpoint url.
func (ep *Endpoints) UserTokens(name string) string {
	return ep.users + "/" + name + "/tokens"
}

// UserToken returns a User token API's endpoint url.
func (ep *Endpoints) UserToken(userName, tokenName string) string {
	return ep.users + "/" + userName + "/tokens/" + tokenName
}
