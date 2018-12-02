package client

import (
	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

// Client represents a Graylog API client.
type Client struct {
	name         string
	password     string
	xRequestedBy string
	endpoints    *endpoint.Endpoints
}

// NewClient returns a new Graylog API Client.
// ep is API endpoint url (ex. http://localhost:9000/api).
// name and password are authentication name and password.
// If you use an access token instead of password, name is access token and password is literal password "token".
// If you use a session token instead of password, name is session token and password is literal password "session".
func NewClient(ep, name, password string) (*Client, error) {
	endpoints, err := endpoint.NewEndpoints(ep)
	if err != nil {
		return nil, err
	}
	return &Client{
		name: name, password: password,
		xRequestedBy: "go-graylog", endpoints: endpoints}, nil
}

// Endpoints returns endpoints.
func (client *Client) Endpoints() *endpoint.Endpoints {
	return client.endpoints
}

// Name returns authentication name.
func (client *Client) Name() string {
	return client.name
}

// Password returns authentication password.
func (client *Client) Password() string {
	return client.password
}

// SetXRequestedBy sets a custom header "X-Requested-By".
// The default value is "go-graylog".
func (client *Client) SetXRequestedBy(x string) {
	client.xRequestedBy = x
}
