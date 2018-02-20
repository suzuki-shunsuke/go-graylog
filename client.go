package graylog

import (
	"errors"
	"net/url"
	"path"
)

type Client struct {
	name     string
	password string
	endpoint *url.URL
}

func NewClient(endpoint, name, password string) (*Client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &Client{name: name, password: password, endpoint: u}, nil
}

func (client *Client) GetName() string {
	return client.name
}

func (client *Client) GetPassword() string {
	return client.password
}

func (client *Client) getUrl(p string) (string, error) {
	if client.endpoint == nil {
		return "", errors.New("Client.endpoint == nil")
	}
	u := *client.endpoint
	u.Path = path.Join(u.Path, p)
	return u.String(), nil
}
