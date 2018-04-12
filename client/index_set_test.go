package client_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
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

func TestSetDefaultIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	iss, _, _, _, err := client.GetIndexSets(0, 0, false)
	if err != nil {
		t.Fatal(err)
	}
	var defIs, is *graylog.IndexSet
	for _, i := range iss {
		if i.Default {
			defIs = &i
		} else {
			is = &i
		}
	}
	if is == nil {
		is = testutil.IndexSet("hoge")
		is.Default = false
		is.Writable = true
		if _, err := client.CreateIndexSet(is); err != nil {
			t.Fatal(err)
		}
		testutil.WaitAfterCreateIndexSet(server)
		defer func(id string) {
			if _, err := client.DeleteIndexSet(id); err != nil {
				t.Fatal(err)
			}
			testutil.WaitAfterDeleteIndexSet(server)
		}(is.ID)
	}
	is, _, err = client.SetDefaultIndexSet(is.ID)
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	defer func(id string) {
		if _, _, err = client.SetDefaultIndexSet(id); err != nil {
			t.Fatal(err)
		}
	}(defIs.ID)
	if !is.Default {
		t.Fatal("updatedIndexSet.Default == false")
	}
	if _, _, err := client.SetDefaultIndexSet(""); err == nil {
		t.Fatal("index set id is required")
	}
	if _, _, err := client.SetDefaultIndexSet("h"); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}

	is.Writable = false

	if _, err := client.UpdateIndexSet(is); err != nil {
		t.Fatal(err)
	}
	if _, _, err := client.SetDefaultIndexSet(is.ID); err == nil {
		t.Fatal("Default index set must be writable.")
	}
}
