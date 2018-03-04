package graylog

import (
	"net/url"
)

// Client represents a Graylog API client.
type Client struct {
	name      string
	password  string
	endpoint  *url.URL
	endpoints *Endpoints
}

// NewClient returns a new Graylog API Client.
// endpoint is API endpoint url (ex. http://localhost:9000/api).
// name and password are authentication name and password.
// If you use an access token instead of password, name is access token and password is literal password "token".
// If you use a session token instead of password, name is session token and password is literal password "session".
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
	endpoints.Streams = getEndpoint(*base, "/streams")
	endpoints.EnabledStreams = getEndpoint(*base, "/streams/enabled")
	endpoints.Endpoint = base

	return &Client{
		name: name, password: password, endpoints: endpoints,
		endpoint: base,
	}, nil
}

// GetName returns authentication name.
func (client *Client) GetName() string {
	return client.name
}

// GetName returns authentication password.
func (client *Client) GetPassword() string {
	return client.password
}
