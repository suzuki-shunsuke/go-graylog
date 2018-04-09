package handler_test

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestServerHandleUpdateIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := testutil.IndexSet("hoge")
	if _, err = server.AddIndexSet(indexSet); err != nil {
		t.Fatal(err)
	}
	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPut, client.Endpoints.IndexSet(indexSet.ID), body)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth(client.Name(), client.Password())
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode == %d, wanted 400", resp.StatusCode)
	}
}

func TestServerHandleCreateIndexSet(t *testing.T) {
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
	req.SetBasicAuth(client.Name(), client.Password())
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode == %d, wanted 400", resp.StatusCode)
	}
}

func TestServerAddIndexSet(t *testing.T) {
	server, _, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := testutil.IndexSet("hoge")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	is.ID = ""
	if _, err := server.AddIndexSet(is); err == nil {
		t.Fatal("index prefix should conflict")
	}
}

func TestServerUpdateIndexSet(t *testing.T) {
	server, _, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	if _, err := server.UpdateIndexSet(nil); err == nil {
		t.Fatal("index set is nil")
	}
	is := testutil.IndexSet("hoge")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	id := is.ID
	is.ID = ""
	if _, err := server.UpdateIndexSet(is); err == nil {
		t.Fatal("index set id is required")
	}
	is.ID = id
	is.IndexPrefix = "graylog"
	if _, err := server.UpdateIndexSet(is); err == nil {
		t.Fatal("index prefix should be conflict")
	}
}

func TestGetIndexSets(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSets, _, _, _, err := client.GetIndexSets(0, 0, false)
	if err != nil {
		t.Fatal("Failed to GetIndexSets", err)
	}
	if indexSets == nil {
		t.Fatal("indexSets == nil")
	}
	if len(indexSets) != 1 {
		t.Fatalf("len(indexSets) == %d, wanted %d", len(indexSets), 1)
	}
}

func TestGetIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	is := testutil.IndexSet("hoge")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	act, _, err := client.GetIndexSet(is.ID)
	if err != nil {
		t.Fatal("Failed to GetIndexSet", err)
	}
	if !reflect.DeepEqual(*act, *is) {
		t.Fatalf("client.GetIndexSet() == %v, wanted %v", act, is)
	}
	if _, _, err := client.GetIndexSet(""); err == nil {
		t.Fatal("index set id is required")
	}
	if _, _, err := client.GetIndexSet("h"); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}
}

func TestCreateIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	exp := testutil.IndexSet("hoge")
	if _, err := client.CreateIndexSet(exp); err != nil {
		t.Fatal("Failed to CreateIndexSet", err)
	}
	if exp.ID == "" {
		t.Fatal("IndexSet's id is empty")
	}
	exp.IndexPrefix = "fuga"
	act := *exp
	exp.Title = ""
	if _, err := client.CreateIndexSet(exp); err == nil {
		t.Fatal("title is required")
	}
	exp.Title = act.Title
	exp.IndexPrefix = ""
	if _, err := client.CreateIndexSet(exp); err == nil {
		t.Fatal("indexPrefix is required")
	}
	exp.IndexPrefix = "fuga"
	exp.RotationStrategyClass = ""
	if _, err := client.CreateIndexSet(exp); err == nil {
		t.Fatal("rotationStrategyClass is required")
	}
	exp.RotationStrategyClass = act.RotationStrategyClass
	exp.RotationStrategy = nil
	if _, err := client.CreateIndexSet(exp); err == nil {
		t.Fatal("rotationStrategy is required")
	}
	if _, err := client.CreateIndexSet(nil); err == nil {
		t.Fatal("index set is nil")
	}
}

func TestUpdateIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := testutil.IndexSet("fuga")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	is.Description = "changed!"

	if _, err := client.UpdateIndexSet(is); err != nil {
		t.Fatal("UpdateIndexSet is failure", err)
	}
	is.ID = ""
	if _, err := client.UpdateIndexSet(is); err == nil {
		t.Fatal("index set id is required")
	}
	is.ID = "h"
	if _, err := client.UpdateIndexSet(is); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}
	is.Title = ""
	if _, err := client.UpdateIndexSet(is); err == nil {
		t.Fatal("title is required")
	}
	if _, err := client.UpdateIndexSet(nil); err == nil {
		t.Fatal("index set is nil")
	}
}

func TestDeleteIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSets, _, err := server.GetIndexSets()
	if err != nil {
		t.Fatal(err)
	}
	indexSet := indexSets[0]
	if _, err = client.DeleteIndexSet(indexSet.ID); err == nil {
		t.Fatal("default index set should not be deleted")
	}
	is := testutil.IndexSet("hoge")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	if _, err = client.DeleteIndexSet(is.ID); err != nil {
		t.Fatal("Failed to DeleteIndexSet", err)
	}
	if _, err = client.DeleteIndexSet(""); err == nil {
		t.Fatal("index set id is required")
	}
	if _, err = client.DeleteIndexSet("h"); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}
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
