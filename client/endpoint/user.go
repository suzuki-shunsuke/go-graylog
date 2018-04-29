package endpoint

import (
	"net/url"
)

// User returns a User API's endpoint url.
func (ep *Endpoints) User(name string) (*url.URL, error) {
	return urlJoin(ep.users, name)
}

// Users returns a User API's endpoint url.
func (ep *Endpoints) Users() string {
	return ep.users.String()
}
