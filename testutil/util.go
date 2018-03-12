package testutil

import (
	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver"
)

func GetServerAndClient() (*mockserver.MockServer, *graylog.Client, error) {
	server, err := mockserver.NewMockServer("", mockserver.NewInMemoryStore(""))
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to Get Mock Server")
	}
	client, err := graylog.NewClient(server.GetEndpoint(), "admin", "password")
	if err != nil {
		server.Close()
		return nil, nil, errors.Wrap(err, "Failed to NewClient")
	}
	server.Start()
	return server, client, nil
}
