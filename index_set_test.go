package graylog_test

import (
	"testing"
	"time"

	"github.com/suzuki-shunsuke/go-graylog/mockserver"
	"github.com/suzuki-shunsuke/go-graylog/test"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func waitAfterCreateIndexSet(server *mockserver.MockServer) {
	// At real graylog API we need to sleep
	// 404 Index set not found.
	if server == nil {
		// time.Sleep(1 * time.Second)
		time.Sleep(875 * time.Millisecond)
	}
}

func waitAfterDeleteIndexSet(server *mockserver.MockServer) {
	// At real graylog API we need to sleep
	// 404 Index set not found.
	if server == nil {
		// time.Sleep(1 * time.Second)
		time.Sleep(935 * time.Millisecond)
	}
}

func TestGetIndexSets(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, _, _, err := client.GetIndexSets(0, 0); err != nil {
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

	// success
	is := testutil.IndexSet("hoge")
	if _, err := client.CreateIndexSet(is); err != nil {
		t.Fatal(err)
	}
	waitAfterCreateIndexSet(server)
	id := is.ID
	// clean
	defer func(id string) {
		if _, err := client.DeleteIndexSet(id); err != nil {
			t.Fatal(err)
		}
		waitAfterDeleteIndexSet(server)
	}(id)
	r, _, err := client.GetIndexSet(id)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("indexSet is nil")
	}
	if r.ID != id {
		t.Fatalf(`indexSet.ID = "%s", wanted "%s"`, r.ID, id)
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
	waitAfterCreateIndexSet(server)
	// clean
	defer func() {
		if _, err := client.DeleteIndexSet(is.ID); err != nil {
			t.Fatal(err)
		}
		waitAfterDeleteIndexSet(server)
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

	is := testutil.IndexSet("hoge")
	if _, err := client.CreateIndexSet(is); err != nil {
		t.Fatal(err)
	}
	waitAfterCreateIndexSet(server)
	id := is.ID
	// clean
	defer func(id string) {
		if _, err := client.DeleteIndexSet(id); err != nil {
			t.Fatal(err)
		}
		waitAfterDeleteIndexSet(server)
	}(id)
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

func TestSetDefaultIndexSet(t *testing.T) {
	test.TestSetDefaultIndexSet(t)
}
