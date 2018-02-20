package graylog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func handleCreateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	role := Role{}
	err = json.Unmarshal(b, &role)
	if err != nil {
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	b, err = json.Marshal(&role)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func TestCreateRole(t *testing.T) {
	http.HandleFunc("/api/roles", handleCreateRole)
	server := httptest.NewServer(nil)
	defer server.Close()
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	client, err := NewClient(u, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	params := &Role{Name: "foo"}
	role, err := client.CreateRole(params)
	if err != nil {
		t.Error("Failed to CreateRole", err)
		return
	}
	if role == nil {
		t.Error("client.CreateRole() == nil")
		return
	}
	if role.Name != "foo" {
		t.Errorf("role.Name == %s, wanted foo", role.Name)
	}
}
