package graylog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"reflect"
	"sync"
	"testing"
)

var (
	once sync.Once
)

func handlerFuncs() {
	http.HandleFunc("/api/roles", handleRoles)
	http.HandleFunc("/api/roles/", handleGetRoleByName)
}

// /roles
func handleRoles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRoles(w, r)
	case http.MethodPost:
		handleCreateRole(w, r)
	}
}

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
	once.Do(handlerFuncs)
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

func handleGetRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	roles := rolesBody{
		Roles: []Role{
			{
				Name:        "Admin",
				Description: "Grants all permissions for Graylog administrators (built-in)",
				Permissions: []string{"*"},
				ReadOnly:    true},
		},
		Total: 1}
	b, err := json.Marshal(&roles)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func TestGetRoles(t *testing.T) {
	once.Do(handlerFuncs)
	server := httptest.NewServer(nil)
	defer server.Close()
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	client, err := NewClient(u, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	roles, err := client.GetRoles()
	if err != nil {
		t.Error("Failed to GetRoles", err)
		return
	}
	exp := []Role{
		{
			Name:        "Admin",
			Description: "Grants all permissions for Graylog administrators (built-in)",
			Permissions: []string{"*"},
			ReadOnly:    true},
	}
	if !reflect.DeepEqual(roles, exp) {
		t.Errorf("client.GetRoles() == %v, wanted %v", roles, exp)
	}
}

func handleGetRoleByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	admin := Role{
		Name:        "Admin",
		Description: "Grants all permissions for Graylog administrators (built-in)",
		Permissions: []string{"*"},
		ReadOnly:    true,
	}
	name := path.Base(r.URL.Path)
	if name == "Admin" {
		b, err := json.Marshal(&admin)
		if err != nil {
			w.Write([]byte(`{"message":"500 Internal Server Error"}`))
			return
		}
		w.Write(b)
		return
	}
	t := Error{
		Message: fmt.Sprintf("No role found with name %s", name),
		Type:    "ApiError"}
	b, err := json.Marshal(&t)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func TestGetRoleByName(t *testing.T) {
	once.Do(handlerFuncs)
	server := httptest.NewServer(nil)
	defer server.Close()
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	client, err := NewClient(u, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	role, err := client.GetRoleByName("Admin")
	if err != nil {
		t.Error("Failed to GetRoleByName", err)
		return
	}
	exp := &Role{
		Name:        "Admin",
		Description: "Grants all permissions for Graylog administrators (built-in)",
		Permissions: []string{"*"},
		ReadOnly:    true,
	}
	if !reflect.DeepEqual(role, exp) {
		t.Errorf("client.GetRoleByName() == %v, wanted %v", role, exp)
	}
}
