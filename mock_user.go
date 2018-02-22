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

func validateUser(user *User) (int, []byte) {
	if user.Username == "" {
		return 400, []byte(`{"type": "ApiError", "message": "Can not construct instance of org.graylog2.rest.models.users.responses.UserResponse, problem: Null name\n at [Source: org.glassfish.jersey.message.internal.ReaderInterceptorExecutor$UnCloseableInputStream@472db3c8; line: 1, column: 31]"}`)
	}
	return 200, []byte("")
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
	sc, msg := validateUser(&user)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	if _, ok := ms.Users[user.Username]; ok {
		w.WriteHeader(400)
		w.Write([]byte(fmt.Sprintf(`{"type": "ApiError", "message": "User %s already exists."}`, user.Username)))
		return
	}
	ms.Users[user.Username] = user
	b, err = json.Marshal(&user)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func (ms *MockServer) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	arr := ms.UserList()
	users := usersBody{Users: arr}
	b, err := json.Marshal(&users)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func (ms *MockServer) handleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := path.Base(r.URL.Path)
	user, ok := ms.Users[name]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(`{"type": "ApiError", "message": "No user found with name %s"}`, name)))
		return
	}
	b, err := json.Marshal(&user)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func (ms *MockServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	name := path.Base(r.URL.Path)
	if _, ok := ms.Users[name]; !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(`{"type": "ApiError", "message": "No user found with name %s"}`, name)))
		return
	}
	user := User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`{"message":"400 Bad Request"}`))
		return
	}
	sc, msg := validateUser(&user)
	if sc != 200 {
		w.WriteHeader(sc)
		w.Write(msg)
		return
	}
	delete(ms.Users, name)
	ms.Users[user.Username] = user
	b, err = json.Marshal(&user)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}

func (ms *MockServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := path.Base(r.URL.Path)
	_, ok := ms.Users[name]
	if !ok {
		w.WriteHeader(404)
		w.Write([]byte(fmt.Sprintf(`{"type": "ApiError", "message": "No user found with name %s"}`, name)))
		return
	}
	delete(ms.Users, name)
}
