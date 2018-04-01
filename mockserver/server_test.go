package mockserver_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver"
)

func TestNewServer(t *testing.T) {
	server, err := mockserver.NewServer("", nil)
	if err != nil {
		t.Fatal(err)
	}
	if server == nil {
		t.Fatal("server is nil")
	}
	server.Start()
	defer server.Close()
	if ep := server.Endpoint(); ep == "" {
		t.Fatal("endpoint is empty")
	}
}
