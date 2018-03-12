package graylog

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("http://localhost:9000/api", "admin", "password")
	if err != nil {
		t.Fatal("Failed to NewClient", err)
	}
	if client == nil {
		t.Fatal("client == nil")
	}
}

func TestGetName(t *testing.T) {
	name := "admin"
	client, err := NewClient("http://localhost:9000/api", name, "password")
	if err != nil {
		t.Fatal("Failed to NewClient", err)
	}
	if client == nil {
		t.Fatal("client == nil")
	}
	act := client.GetName()
	if act != name {
		t.Fatalf("client.GetName() == %s, wanted %s", act, name)
	}

	exp := "http://localhost:9000/api/roles"
	act = client.Endpoints.Roles
	if act != exp {
		t.Fatalf("client.Endpoints.Roles == %s, wanted %s", act, exp)
	}
}

func TestGetPassword(t *testing.T) {
	password := "password"
	client, err := NewClient("http://localhost:9000/api", "admin", password)
	if err != nil {
		t.Fatal("Failed to NewClient", err)
	}
	if client == nil {
		t.Fatal("client == nil")
	}
	real := client.GetPassword()
	if real != password {
		t.Fatalf("client.GetPassword() == %s, wanted %s", real, password)
	}
}
