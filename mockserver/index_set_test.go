package mockserver_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestMockServerHandleUpdateIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := testutil.DummyNewIndexSet("hoge")
	if _, err = server.AddIndexSet(indexSet); err != nil {
		t.Fatal(err)
	}
	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPut, client.Endpoints.IndexSet(indexSet.ID), body)
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

func TestMockServerHandleCreateIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPost, client.Endpoints.IndexSets, body)
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

func TestGetIndexSets(t *testing.T) {
	test.TestGetIndexSets(t)
}

func TestGetIndexSet(t *testing.T) {
	test.TestGetIndexSet(t)
}

func TestCreateIndexSet(t *testing.T) {
	test.TestCreateIndexSet(t)
}

func TestUpdateIndexSet(t *testing.T) {
	test.TestUpdateIndexSet(t)
}

func TestDeleteIndexSet(t *testing.T) {
	test.TestDeleteIndexSet(t)
}

func TestSetDefaultIndexSet(t *testing.T) {
	test.TestSetDefaultIndexSet(t)
}

func TestGetIndexSetStats(t *testing.T) {
	test.TestGetIndexSetStats(t)
}

func TestGetAllIndexSetsStats(t *testing.T) {
	test.TestGetAllIndexSetsStats(t)
}
