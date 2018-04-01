package testutil

import (
	"os"

	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/mockserver"
)

// GetServerAndClient returns server and client.
// If you want to use mock server, pass "mock" as endpoint.
// If you want to use real server, pass "real" as endpoint.
// If endpoint is "" and GRAYLOG_WEB_ENDPOINT_URI is set, returns real server.
func GetServerAndClient() (*mockserver.Server, *client.Client, error) {
	var (
		server *mockserver.Server
		err    error
	)
	authName := os.Getenv("GRAYLOG_AUTH_NAME")
	authPass := os.Getenv("GRAYLOG_AUTH_PASSWORD")
	if authName == "" {
		authName = "admin"
	}
	if authPass == "" {
		authPass = "admin"
	}
	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")
	if endpoint == "" {
		server, err = mockserver.NewServer("", nil)
		if err != nil {
			return nil, nil, errors.Wrap(err, "Failed to get Mock Server")
		}
		server.SetAuth(true)
		endpoint = server.Endpoint()
	}
	client, err := client.NewClient(endpoint, authName, authPass)
	if err != nil {
		server.Close()
		return nil, nil, errors.Wrap(err, "NewClient is failure")
	}
	if server != nil {
		server.Start()
	}
	return server, client, nil
}
