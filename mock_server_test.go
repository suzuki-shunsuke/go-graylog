package graylog

import (
	"fmt"
	"net/http"
	"testing"
)

func TestMockServerLoad(t *testing.T) {
	server, _, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	if err := server.Load(); err != nil {
		t.Fatal(err)
	}
}

func TestMockServerHandleNotFound(t *testing.T) {
	server, _, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	endpoint := fmt.Sprintf("%s/dummy", server.Endpoint)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}
	hc := &http.Client{}
	if _, err := hc.Do(req); err != nil {
		t.Fatal(err)
	}
}
