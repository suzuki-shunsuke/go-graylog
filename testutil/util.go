package testutil

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/graylog-mock-server/mockserver"
)

const (
	adminName string = "admin"
)

// GetNonAdminUser returns a user whose name is not "admin".
func GetNonAdminUser(cl *client.Client) (*graylog.User, error) {
	users, _, err := cl.GetUsers()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.Username != adminName {
			return &user, nil
		}
	}
	return nil, nil
}

// GetRoleOrCreate gets a given name's role.
// If no role whose name is a given name exists, create a role with a given name and returns it.
func GetRoleOrCreate(cl *client.Client, name string) (*graylog.Role, error) {
	role, ei, err := cl.GetRole(name)
	if err == nil {
		return role, nil
	}
	if ei == nil || ei.Response == nil || ei.Response.StatusCode != 404 {
		return nil, err
	}
	role = Role()
	role.Name = name
	if _, err := cl.CreateRole(role); err != nil {
		return nil, err
	}
	return role, nil
}

// GetIndexSet returns an IndexSet.
func GetIndexSet(cl *client.Client, server *mockserver.Server, prefix string) (*graylog.IndexSet, func(string), error) {
	iss, _, _, _, err := cl.GetIndexSets(0, 0, false)
	if err != nil {
		return nil, nil, err
	}
	if len(iss) != 0 {
		return &(iss[0]), nil, nil
	}
	is := IndexSet(prefix)
	if _, err := cl.CreateIndexSet(is); err != nil {
		return nil, nil, err
	}
	WaitAfterCreateIndexSet(server)
	return is, func(id string) {
		if _, err := cl.DeleteIndexSet(id); err == nil {
			WaitAfterDeleteIndexSet(server)
		}
	}, nil
}

// GetStream returns a stream.
func GetStream(cl *client.Client, server *mockserver.Server, mode int) (*graylog.Stream, func(string), error) {
	streams, _, _, err := cl.GetStreams()
	if err != nil {
		return nil, nil, err
	}
	if len(streams) != 0 {
		if mode == 1 {
			for _, stream := range streams {
				if stream.IsDefault {
					return &stream, nil, nil
				}
			}
			return nil, nil, fmt.Errorf("default stream is not found")
		}
		if mode == 2 {
			for _, stream := range streams {
				if !stream.IsDefault {
					return &stream, nil, nil
				}
			}
			return nil, nil, fmt.Errorf("not default stream is not found")
		}
		return &(streams[0]), nil, nil
	}
	is, f, err := GetIndexSet(cl, server, "hoge")
	if err != nil {
		return nil, nil, err
	}
	stream := Stream()
	stream.IndexSetID = is.ID
	if _, err := cl.CreateStream(stream); err != nil {
		if f != nil {
			f(is.ID)
		}
		return nil, nil, err
	}
	return stream, func(id string) {
		cl.DeleteStream(id)
		if f != nil {
			f(is.ID)
		}
	}, nil
}

// WaitAfterCreateIndexSet waits for a while after creates an index set
// if server is not a mock server.
func WaitAfterCreateIndexSet(server *mockserver.Server) {
	// At real graylog API we need to sleep
	// 404 Index set not found.
	if server == nil {
		time.Sleep(1 * time.Second)
	}
}

// WaitAfterDeleteIndexSet waits for a while after deletes an index set
// if server is not a mock server.
func WaitAfterDeleteIndexSet(server *mockserver.Server) {
	// At real graylog API we need to sleep
	// 404 Index set not found.
	if server == nil {
		time.Sleep(1 * time.Second)
	}
}

// GetServerAndClient returns server and client.
func GetServerAndClient() (*mockserver.Server, *client.Client, error) {
	var (
		server *mockserver.Server
		err    error
	)
	authName := os.Getenv("GRAYLOG_AUTH_NAME")
	authPass := os.Getenv("GRAYLOG_AUTH_PASSWORD")
	if authName == "" {
		authName = adminName
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
		if server != nil {
			server.Close()
		}
		return nil, nil, errors.Wrap(err, "NewClient is failure")
	}
	if server != nil {
		server.Start()
	}
	return server, client, nil
}
