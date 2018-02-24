package graylog

import (
	"fmt"
	"net/url"
	"path"
)

type Endpoints struct {
	Roles     string
	Users     string
	Inputs    string
	IndexSets string
}

type Client struct {
	name      string
	password  string
	endpoint  *url.URL
	endpoints *Endpoints
}

func getEndpoint(u url.URL, p string) string {
	u.Path = path.Join(u.Path, p)
	return u.String()
}

func NewClient(endpoint, name, password string) (*Client, error) {
	base, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	endpoints := &Endpoints{}

	endpoints.Roles = getEndpoint(*base, "/roles")
	endpoints.Users = getEndpoint(*base, "/users")
	endpoints.Inputs = getEndpoint(*base, "/system/inputs")
	endpoints.IndexSets = getEndpoint(*base, "/system/indices/index_sets")

	return &Client{
		name: name, password: password, endpoints: endpoints,
		endpoint: base,
	}, nil
}

func (client *Client) GetName() string {
	return client.name
}

func (client *Client) GetPassword() string {
	return client.password
}

func (client *Client) RoleMembersEndpoint(name string) string {
	u := *(client.endpoint)
	u.Path = path.Join(u.Path, fmt.Sprintf("/roles/%s/members", name))
	return u.String()
}

func (client *Client) RoleMemberEndpoint(userName, roleName string) string {
	u := *(client.endpoint)
	u.Path = path.Join(u.Path, fmt.Sprintf("/roles/%s/members/%s", roleName, userName))
	return u.String()
}
