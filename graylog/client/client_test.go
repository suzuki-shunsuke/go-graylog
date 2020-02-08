package client_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/client"
)

const (
	endpoint = "http://localhost:9000/api"
)

func TestNewClient(t *testing.T) {
	client, err := client.NewClient(endpoint, "admin", "password")
	if err != nil {
		t.Fatal("Failed to NewClient", err)
	}
	if client == nil {
		t.Fatal("client == nil")
	}
}

func TestClient_Name(t *testing.T) {
	name := "admin"
	client, err := client.NewClient(endpoint, name, "password")
	if err != nil {
		t.Fatal("Failed to NewClient", err)
	}
	if client == nil {
		t.Fatal("client == nil")
	}
	act := client.Name()
	if act != name {
		t.Fatalf("client.Name() == %s, wanted %s", act, name)
	}

	exp := "http://localhost:9000/api/roles"
	act = client.Endpoints().Roles()
	if act != exp {
		t.Fatalf("client.Endpoints().Roles == %s, wanted %s", act, exp)
	}
}

func TestClient_Password(t *testing.T) {
	password := "password"
	client, err := client.NewClient(endpoint, "admin", password)
	if err != nil {
		t.Fatal("Failed to NewClient", err)
	}
	if client == nil {
		t.Fatal("client == nil")
	}
	real := client.Password()
	if real != password {
		t.Fatalf("client.Password() == %s, wanted %s", real, password)
	}
}
