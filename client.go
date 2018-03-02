// Package graylog provides Graylog API client and mock server.
package graylog

import (
	"fmt"
	"net/url"
	"path"
)

// Endpoints represents each API's endpoint URLs.
type Endpoints struct {
	Endpoint       *url.URL
	Roles          string
	Users          string
	Inputs         string
	IndexSets      string
	Streams        string
	EnabledStreams string
}

func (endpoint *Endpoints) SetDefaultIndexSet(id string) string {
	return fmt.Sprintf("%s/%s/default", endpoint.IndexSets, id)
}

func (endpoint *Endpoints) IndexSetsStats() string {
	return fmt.Sprintf("%s/stats", endpoint.IndexSets)
}

func (endpoint *Endpoints) IndexSetStats(id string) string {
	return fmt.Sprintf("%s/%s/stats", endpoint.IndexSets, id)
}

func (endpoint *Endpoints) IndexSet(id string) string {
	return fmt.Sprintf("%s/%s", endpoint.IndexSets, id)
}

func (endpoint *Endpoints) Input(id string) string {
	return fmt.Sprintf("%s/%s", endpoint.Inputs, id)
}

func (endpoint *Endpoints) User(name string) string {
	return fmt.Sprintf("%s/%s", endpoint.Users, name)
}

func (endpoint *Endpoints) Role(name string) string {
	return fmt.Sprintf("%s/%s", endpoint.Roles, name)
}

func (endpoint *Endpoints) Stream(id string) string {
	return fmt.Sprintf("%s/%s", endpoint.Streams, id)
}

func (endpoint *Endpoints) PauseStream(id string) string {
	return fmt.Sprintf("%s/%s/pause", endpoint.Streams, id)
}

func (endpoint *Endpoints) ResumeStream(id string) string {
	return fmt.Sprintf("%s/%s/resume", endpoint.Streams, id)
}

// Client represents a Graylog API client.
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

// RoleMembersEndpoint returns given role's member endpoint url.
func (client *Client) RoleMembersEndpoint(name string) string {
	u := *(client.endpoint)
	u.Path = path.Join(u.Path, fmt.Sprintf("/roles/%s/members", name))
	return u.String()
}

// RoleMemberEndpoint returns given role member endpoint url.
func (client *Client) RoleMemberEndpoint(userName, roleName string) string {
	u := *(client.endpoint)
	u.Path = path.Join(u.Path, fmt.Sprintf("/roles/%s/members/%s", roleName, userName))
	return u.String()
}
