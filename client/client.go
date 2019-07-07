package client

import (
	"net/http"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

// Client represents a Graylog API client.
type Client struct {
	name         string
	password     string
	xRequestedBy string
	apiVersion   string
	endpoints    *endpoint.Endpoints
	httpClient   *http.Client
}

// NewClient returns a new Graylog API Client.
// ep is API endpoint url (ex. http://localhost:9000/api).
// name and password are authentication name and password.
// If you use an access token instead of password, name is access token and password is literal password "token".
// If you use a session token instead of password, name is session token and password is literal password "session".
func NewClient(ep, name, password string) (*Client, error) {
	return newClient(ep, name, password, "")
}

// NewClientV3 returns a new Graylog v3 API Client.
// ep is API endpoint url (ex. http://localhost:9000/api).
// name and password are authentication name and password.
// If you use an access token instead of password, name is access token and password is literal password "token".
// If you use a session token instead of password, name is session token and password is literal password "session".
func NewClientV3(ep, name, password string) (*Client, error) {
	return newClient(ep, name, password, "v3")
}

func newClient(ep, name, password, version string) (*Client, error) {
	var (
		endpoints *endpoint.Endpoints
		err       error
	)
	if version == "v3" {
		endpoints, err = endpoint.NewEndpointsV3(ep)
	} else {
		endpoints, err = endpoint.NewEndpoints(ep)
	}
	if err != nil {
		return nil, err
	}
	return &Client{
		name: name, password: password,
		xRequestedBy: "go-graylog", endpoints: endpoints,
		apiVersion: version}, nil
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

// SetHTTPClient sets a custom *http.Client.
// If you don't set *http.Client by this method, the default value is http.DefaultClient.
func (client *Client) SetHTTPClient(c *http.Client) {
	client.httpClient = c
}
