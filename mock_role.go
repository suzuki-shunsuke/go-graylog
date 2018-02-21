package graylog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// /roles
func (ms *MockServer) handleRoles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ms.handleGetRoles(w, r)
	case http.MethodPost:
		ms.handleCreateRole(w, r)
	}
}

func (ms *MockServer) handleCreateRole(w http.ResponseWriter, r *http.Request) {
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

func (ms *MockServer) handleGetRoles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	size := len(ms.Roles)
	arr := make([]Role, size)
	i := 0
	for _, role := range ms.Roles {
		arr[i] = role
		i++
	}
	roles := rolesBody{
		Roles: arr,
		Total: size}
	b, err := json.Marshal(&roles)
	if err != nil {
		w.Write([]byte(`{"message":"500 Internal Server Error"}`))
		return
	}
	w.Write(b)
}
