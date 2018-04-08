package test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestSetDefaultIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	iss, _, _, err := client.GetIndexSets(0, 0)
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

func TestGetIndexSetStats(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	iss, _, _, err := client.GetIndexSets(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet("hoge")
	if len(iss) == 0 {
		if _, err := client.CreateIndexSet(is); err != nil {
			t.Fatal(err)
		}
		testutil.WaitAfterCreateIndexSet(server)
		// clean
		defer func(id string) {
			if _, err := client.DeleteIndexSet(id); err != nil {
				t.Fatal(err)
			}
			testutil.WaitAfterDeleteIndexSet(server)
		}(is.ID)
	} else {
		is = &(iss[0])
	}

	if server != nil {
		indexSetStats := testutil.DummyIndexSetStats()
		if _, err := server.SetIndexSetStats(is.ID, indexSetStats); err != nil {
			t.Fatal(err)
		}
	}

	if _, _, err := client.GetIndexSetStats(is.ID); err != nil {
		t.Fatal(err)
	}
	if _, _, err := client.GetIndexSetStats(""); err == nil {
		t.Fatal("index set id is required")
	}
	// if _, _, err := client.GetIndexSetStats("h"); err == nil {
	// 	t.Fatal(`no index set whose id is "h"`)
	// }
}

func TestGetAllIndexSetsStats(t *testing.T) {
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
	if server != nil {
		indexSetStats := testutil.DummyIndexSetStats()
		if _, err := server.SetIndexSetStats(is.ID, indexSetStats); err != nil {
			t.Fatal(err)
		}
	}
	if _, _, err := client.GetAllIndexSetsStats(); err != nil {
		t.Fatal(err)
	}
}
