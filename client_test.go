package graylog

import (
	"net/http"
	"sync"
	"testing"
)

var (
	once sync.Once
)

func handlerFuncs() {
	http.HandleFunc("/api/roles", handleRoles)
	http.HandleFunc("/api/roles/", handleRole)

	http.HandleFunc("/api/users", handleUsers)
	http.HandleFunc("/api/users/", handleUser)
}

func TestNewClient(t *testing.T) {
	client, err := NewClient("http://localhost:9000/api", "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	if client == nil {
		t.Error("client == nil")
	}
}

func TestGetName(t *testing.T) {
	name := "admin"
	client, err := NewClient("http://localhost:9000/api", name, "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	if client == nil {
		t.Error("client == nil")
		return
	}
	act := client.GetName()
	if act != name {
		t.Errorf("client.GetName() == %s, wanted %s", act, name)
	}

	exp := "http://localhost:9000/api/roles"
	act = client.endpoints.Roles
	if act != exp {
		t.Errorf("client.endpoints.Roles == %s, wanted %s", act, exp)
	}
}

func TestGetPassword(t *testing.T) {
	password := "password"
	client, err := NewClient("http://localhost:9000/api", "admin", password)
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	if client == nil {
		t.Error("client == nil")
		return
	}
	real := client.GetPassword()
	if real != password {
		t.Errorf("client.GetPassword() == %s, wanted %s", real, password)
	}
}
