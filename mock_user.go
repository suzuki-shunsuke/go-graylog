package graylog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

// /users
func (ms *MockServer) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ms.handleGetUsers(w, r)
	case http.MethodPost:
		ms.handleCreateUser(w, r)
	}
}

// /users/{username}
func (ms *MockServer) handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ms.handleGetUser(w, r)
	case http.MethodPut:
		ms.handleUpdateUser(w, r)
	case http.MethodDelete:
		ms.handleDeleteUser(w, r)
	}
}

func (ms *MockServer) handleCreateUser(w http.ResponseWriter, r *http.Request) {
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

func (ms *MockServer) handleGetUsers(w http.ResponseWriter, r *http.Request) {
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

func (ms *MockServer) handleGetUser(w http.ResponseWriter, r *http.Request) {
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

func (ms *MockServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
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

func (ms *MockServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) {}
