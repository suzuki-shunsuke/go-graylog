package graylog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"reflect"
	"testing"
)

// /roles
func handleRoles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRoles(w, r)
	case http.MethodPost:
		handleCreateRole(w, r)
	}
}

// /roles/{rolename}
func handleRole(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRole(w, r)
	case http.MethodPut:
		handleUpdateRole(w, r)
	case http.MethodDelete:
		handleDeleteRole(w, r)
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
	server, err := GetMockServer()
	if err != nil {
		t.Error("Failed to Get Mock Server", err)
		return
	}
	defer server.Server.Close()
	client, err := NewClient(server.Endpoint, "admin", "password")
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
	server, err := GetMockServer()
	if err != nil {
		t.Error("Failed to Get Mock Server", err)
		return
	}
	defer server.Server.Close()
	client, err := NewClient(server.Endpoint, "admin", "password")
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

func handleGetRole(w http.ResponseWriter, r *http.Request) {
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

func TestGetRole(t *testing.T) {
	server, err := GetMockServer()
	if err != nil {
		t.Error("Failed to Get Mock Server", err)
		return
	}
	defer server.Server.Close()
	client, err := NewClient(server.Endpoint, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	role, err := client.GetRole("Admin")
	if err != nil {
		t.Error("Failed to GetRole", err)
		return
	}
	exp := &Role{
		Name:        "Admin",
		Description: "Grants all permissions for Graylog administrators (built-in)",
		Permissions: []string{"*"},
		ReadOnly:    true,
	}
	if !reflect.DeepEqual(role, exp) {
		t.Errorf("client.GetRole() == %v, wanted %v", role, exp)
	}
}

func handleUpdateRole(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := path.Base(r.URL.Path)
	if name != "Admin" {
		t := Error{
			Message: fmt.Sprintf("No role found with name %s", name),
			Type:    "ApiError"}
		b, err := json.Marshal(&t)
		if err != nil {
			w.Write([]byte(`{"message":"500 Internal Server Error"}`))
			return
		}
		w.Write(b)
		return
	}
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

func TestUpdateRole(t *testing.T) {
	server, err := GetMockServer()
	if err != nil {
		t.Error("Failed to Get Mock Server", err)
		return
	}
	defer server.Server.Close()
	client, err := NewClient(server.Endpoint, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	role := Role{
		Name:        "foo",
		Description: "",
		Permissions: []string{"users:edit"},
		ReadOnly:    false,
	}
	updatedRole, err := client.UpdateRole("Admin", &role)
	if err != nil {
		t.Error("Failed to UpdateRole", err)
		return
	}
	if !reflect.DeepEqual(*updatedRole, role) {
		t.Errorf("client.UpdateRole() == %v, wanted %v", role, updatedRole)
	}
}

func handleDeleteRole(w http.ResponseWriter, r *http.Request) {}

func TestDeleteRole(t *testing.T) {
	server, err := GetMockServer()
	if err != nil {
		t.Error("Failed to Get Mock Server", err)
		return
	}
	defer server.Server.Close()
	client, err := NewClient(server.Endpoint, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	err = client.DeleteRole("Admin")
	if err != nil {
		t.Error("Failed to DeleteRole", err)
		return
	}
}
