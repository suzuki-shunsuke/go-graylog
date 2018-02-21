package graylog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"reflect"
	"testing"
)

func dummyAdmin() *User {
	return &User{
		Id:          "local:admin",
		Username:    "admin",
		Email:       "",
		FullName:    "Administrator",
		Permissions: []string{"*"},
		Preferences: &Preferences{
			UpdateUnfocussed:  false,
			EnableSmartSearch: true,
		},
		Timezone:         "UTC",
		SessionTimeoutMs: 28800000,
		External:         false,
		Startpage:        nil,
		Roles:            []string{"Admin"},
		ReadOnly:         true,
		SessionActive:    true,
		LastActivity:     "2018-02-21T07:35:45.926+0000",
		ClientAddress:    "172.18.0.1",
	}
}

// /users
func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetUsers(w, r)
	case http.MethodPost:
		handleCreateUser(w, r)
	}
}

// /users/{username}
func handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetUser(w, r)
	case http.MethodPut:
		handleUpdateUser(w, r)
	case http.MethodDelete:
		handleDeleteUser(w, r)
	}
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	user := User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	b, err = json.Marshal(&user)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func TestCreateUser(t *testing.T) {
	once.Do(handlerFuncs)
	server := httptest.NewServer(nil)
	defer server.Close()
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	client, err := NewClient(u, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	params := &User{Username: "foo"}
	user, err := client.CreateUser(params)
	if err != nil {
		t.Error("Failed to CreateUser", err)
		return
	}
	if user == nil {
		t.Error("client.CreateUser() == nil")
		return
	}
	if user.Username != "foo" {
		t.Errorf("user.Username == %s, wanted foo", user.Username)
	}
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	admin := dummyAdmin()
	users := usersBody{
		Users: []User{*admin}}
	b, err := json.Marshal(&users)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func TestGetUsers(t *testing.T) {
	once.Do(handlerFuncs)
	server := httptest.NewServer(nil)
	defer server.Close()
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	client, err := NewClient(u, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	users, err := client.GetUsers()
	if err != nil {
		t.Error("Failed to GetUsers", err)
		return
	}
	admin := dummyAdmin()
	exp := []User{*admin}
	if !reflect.DeepEqual(users, exp) {
		t.Errorf("client.GetUsers() == %v, wanted %v", users, exp)
	}
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	admin := dummyAdmin()
	name := path.Base(r.URL.Path)
	if name == "Admin" {
		b, err := json.Marshal(admin)
		if err != nil {
			w.Write([]byte(`{"message":"500 Internal Server Error"}`))
			return
		}
		w.Write(b)
		return
	}
	t := Error{
		Message: fmt.Sprintf("No user found with name %s", name),
		Type:    "ApiError"}
	b, err := json.Marshal(&t)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func TestGetUser(t *testing.T) {
	once.Do(handlerFuncs)
	server := httptest.NewServer(nil)
	defer server.Close()
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	client, err := NewClient(u, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	user, err := client.GetUser("Admin")
	if err != nil {
		t.Error("Failed to GetUser", err)
		return
	}
	exp := dummyAdmin()
	if !reflect.DeepEqual(*user, *exp) {
		t.Errorf("client.GetUser() == %v, wanted %v", user, exp)
	}
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := path.Base(r.URL.Path)
	if name != "Admin" {
		t := Error{
			Message: fmt.Sprintf("No user found with name %s", name),
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
	user := User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	b, err = json.Marshal(&user)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func TestUpdateUser(t *testing.T) {
	once.Do(handlerFuncs)
	server := httptest.NewServer(nil)
	defer server.Close()
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	client, err := NewClient(u, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	user := dummyAdmin()
	updatedUser, err := client.UpdateUser("Admin", user)
	if err != nil {
		t.Error("Failed to UpdateUser", err)
		return
	}
	if !reflect.DeepEqual(*updatedUser, *user) {
		t.Errorf("client.UpdateUser() == %v, wanted %v", user, updatedUser)
	}
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {}

func TestDeleteUser(t *testing.T) {
	once.Do(handlerFuncs)
	server := httptest.NewServer(nil)
	defer server.Close()
	u := fmt.Sprintf("http://%s/api", server.Listener.Addr().String())
	client, err := NewClient(u, "admin", "password")
	if err != nil {
		t.Error("Failed to NewClient", err)
		return
	}
	err = client.DeleteUser("Admin")
	if err != nil {
		t.Error("Failed to DeleteUser", err)
		return
	}
}
