package client

import (
	"net/url"
)

// Client represents a Graylog API client.
type Client struct {
	name      string
	password  string
	endpoint  *url.URL
	Endpoints *Endpoints
}

// NewClient returns a new Graylog API Client.
// endpoint is API endpoint url (ex. http://localhost:9000/api).
// name and password are authentication name and password.
// If you use an access token instead of password, name is access token and password is literal password "token".
// If you use a session token instead of password, name is session token and password is literal password "session".
func NewClient(endpoint, name, password string) (*Client, error) {
	endpoints, err := NewEndpoints(endpoint)
	if err != nil {
		return nil, err
	}
	return &Client{
		name: name, password: password, Endpoints: endpoints,
	}, nil
}

// GetName returns authentication name.
func (client *Client) Name() string {
	return client.name
}

// GetName returns authentication password.
func (client *Client) Password() string {
	return client.password
}
