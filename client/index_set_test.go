package client_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetIndexSets(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, _, _, _, err := client.GetIndexSets(0, 0, false); err != nil {
		t.Fatal(err)
	}
}

func TestGetIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	is, f, err := testutil.GetIndexSet(client, server, "hoge")
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(is.ID)
	}

	// success
	r, _, err := client.GetIndexSet(is.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("indexSet is nil")
	}
	if r.ID != is.ID {
		t.Fatalf(`indexSet.ID = "%s", wanted "%s"`, r.ID, is.ID)
	}
	// id is required
	if _, _, err := client.GetIndexSet(""); err == nil {
		t.Fatal("id is required")
	}
	// invalid id
	if _, _, err := client.GetIndexSet("h"); err == nil {
		t.Fatal("index set should not be found")
	}
}

func TestCreateIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	// nil check
	if _, err := client.CreateIndexSet(nil); err == nil {
		t.Fatal("index set is nil")
	}
	is := testutil.IndexSet("hoge")
	// success
	if _, err := client.CreateIndexSet(is); err != nil {
		t.Fatal(err)
	}
	testutil.WaitAfterCreateIndexSet(server)
	// clean
	defer func() {
		if _, err := client.DeleteIndexSet(is.ID); err != nil {
			t.Fatal(err)
		}
		testutil.WaitAfterDeleteIndexSet(server)
	}()
}

func TestUpdateIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	is, f, err := testutil.GetIndexSet(client, server, "hoge")
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(is.ID)
	}
	// success
	if _, err := client.UpdateIndexSet(is); err != nil {
		t.Fatal(err)
	}
	// id required
	is.ID = ""
	if _, err := client.UpdateIndexSet(is); err == nil {
		t.Fatal("index set id is required")
	}
	// nil check
	if _, err := client.UpdateIndexSet(nil); err == nil {
		t.Fatal("index set is required")
	}
	// invalid id
	is.ID = "h"
	if _, err := client.UpdateIndexSet(is); err == nil {
		t.Fatal("index set should no be found")
	}
}

func TestDeleteIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	// id required
	if _, err := client.DeleteIndexSet(""); err == nil {
		t.Fatal("id is required")
	}
	// invalid id
	if _, err := client.DeleteIndexSet("h"); err == nil {
		t.Fatal(`no index set with id "h" is found`)
	}
}
