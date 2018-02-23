package graylog

import (
	"fmt"
	"net/url"
	"path"
)

type Endpoints struct {
	Roles  string
	Users  string
	Inputs string
}

type Client struct {
	name      string
	password  string
	endpoint  *url.URL
	endpoints *Endpoints
}

func NewClient(endpoint, name, password string) (*Client, error) {
	base, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	endpoints := &Endpoints{}

	u := *base
	u.Path = path.Join(u.Path, "/roles")
	endpoints.Roles = u.String()

	u = *base
	u.Path = path.Join(u.Path, "/users")
	endpoints.Users = u.String()

	u = *base
	u.Path = path.Join(u.Path, "/system/inputs")
	endpoints.Inputs = u.String()

	return &Client{
		name: name, password: password, endpoints: endpoints,
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
