package mockserver_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestServerHandleCreateInput(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPost, client.Endpoints.Inputs, body)
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

func TestServerHandleUpdateInput(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := testutil.Input()

	if _, err := server.AddInput(input); err != nil {
		t.Fatal(err)
	}
	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPut, client.Endpoints.Input(input.ID), body)
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

func TestCreateInput(t *testing.T) {
	test.TestCreateInput(t)
}

func TestGetInputs(t *testing.T) {
	test.TestGetInputs(t)
}

func TestGetInput(t *testing.T) {
	test.TestGetInput(t)
}

func TestUpdateInput(t *testing.T) {
	test.TestUpdateInput(t)
}

func TestDeleteInput(t *testing.T) {
	test.TestDeleteInput(t)
}
