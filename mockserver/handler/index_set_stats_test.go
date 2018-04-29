package handler_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHandleGetIndexSetStats(t *testing.T) {
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
	is := testutil.IndexSet("hoge")
	if len(iss) == 0 {
		if _, err := client.CreateIndexSet(is); err != nil {
			t.Fatal(err)
		}
		// clean
		defer func(id string) {
			if _, err := client.DeleteIndexSet(id); err != nil {
				t.Fatal(err)
			}
		}(is.ID)
	} else {
		is = &(iss[0])
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

func TestGetTotalIndexSetsStats(t *testing.T) {
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
	if _, _, err := client.GetTotalIndexSetsStats(); err != nil {
		t.Fatal(err)
	}
}
