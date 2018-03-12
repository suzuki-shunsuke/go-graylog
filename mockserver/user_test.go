package mockserver_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestMockServerHandleCreateUser(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPost, client.Endpoints.Users, body)
	if err != nil {
		t.Fatal(err)
	}
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode == %d, wanted 400", resp.StatusCode)
	}
}

func TestMockServerHandleUpdateUser(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	user, _, err := server.AddUser(testutil.DummyNewUser())
	if err != nil {
		t.Fatal(err)
	}
	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPut, client.Endpoints.User(user.Username), body)
	if err != nil {
		t.Fatal(err)
	}
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode == %d, wanted 400", resp.StatusCode)
	}
}

func TestCreateUser(t *testing.T) {
	test.TestCreateUser(t)
}

func TestGetUsers(t *testing.T) {
	test.TestGetUsers(t)
}

func TestGetUser(t *testing.T) {
	test.TestGetUser(t)
}

func TestUpdateUser(t *testing.T) {
	test.TestUpdateUser(t)
}

func TestDeleteUser(t *testing.T) {
	test.TestDeleteUser(t)
}
