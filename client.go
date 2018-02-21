package graylog

import (
	"net/url"
	"path"
)

type Endpoints struct {
	Roles string
	Users string
}

type Client struct {
	name      string
	password  string
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
