package graylog

import (
	"testing"
)

func TestCreateRole(t *testing.T) {
	client, err := NewClient("http://localhost:9000/api", "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	role := &Role{Name: "name"}
	_, err := client.CreateRole(role)
	if err != nil {
		t.Error("Failed to CreateRole", err)
	}
}
