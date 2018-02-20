package graylog

import (
	"testing"
)

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
	real := client.GetName()
	if real != name {
		t.Errorf("client.GetName() == %s, wanted %s", real, name)
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

func TestGetUrl(t *testing.T) {
	endpoint := "http://localhost:9000/api"
	path := "/roles"
	client, err := NewClient(endpoint, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	if client == nil {
		t.Error("client == nil")
		return
	}
	wanted := "http://localhost:9000/api/roles"
	real, err := client.getUrl(path)
	if err != nil {
		t.Error("Failed to getUrl", path, err)
		return
	}
	if real != wanted {
		t.Errorf("client.GetUrl() == %s, wanted %s", real, wanted)
	}

	real, err = client.getUrl(path)
	if err != nil {
		t.Error("Failed to getUrl", path, err)
		return
	}
	if real != wanted {
		t.Errorf("client.GetUrl() == %s, wanted %s", real, wanted)
	}
}
