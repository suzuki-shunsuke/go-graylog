package logic_test

import (
	"reflect"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
	"github.com/suzuki-shunsuke/go-ptr"
)

func TestAddIndexSet(t *testing.T) {
	server, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet("hoge")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	is.ID = ""
	if _, err := server.AddIndexSet(is); err == nil {
		t.Fatal("index prefix should conflict")
	}
}

func TestGetIndexSets(t *testing.T) {
	server, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	indexSets, _, _, err := server.GetIndexSets(0, 0)
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
	server, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}

	is := testutil.IndexSet("hoge")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	act, _, err := server.GetIndexSet(is.ID)
	if err != nil {
		t.Fatal("Failed to GetIndexSet", err)
	}
	if !reflect.DeepEqual(*act, *is) {
		t.Fatalf("server.GetIndexSet() == %v, wanted %v", act, is)
	}
	if _, _, err = server.GetIndexSet(""); err == nil {
		t.Fatal("index set id is empty")
	}
	if _, _, err = server.GetIndexSet("h"); err == nil {
		t.Fatal("no index set <h> is found")
	}
}

func TestUpdateIndexSet(t *testing.T) {
	server, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet("fuga")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	prms := is.NewUpdateParams()
	prms.Description = ptr.PStr("changed!")
	if _, _, err := server.UpdateIndexSet(prms); err != nil {
		t.Fatal("UpdateIndexSet is failure", err)
	}
	prms.ID = ""
	if _, _, err := server.UpdateIndexSet(prms); err == nil {
		t.Fatal("index set id is required")
	}
	prms.ID = "h"
	if _, _, err := server.UpdateIndexSet(prms); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}
	prms.ID = is.ID
	prms.Title = ""
	if _, _, err := server.UpdateIndexSet(prms); err == nil {
		t.Fatal("title is required")
	}
	if _, _, err := server.UpdateIndexSet(nil); err == nil {
		t.Fatal("index set is nil")
	}
}

func TestDeleteIndexSet(t *testing.T) {
	server, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	indexSets, _, _, err := server.GetIndexSets(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	indexSet := indexSets[0]
	if _, err = server.DeleteIndexSet(indexSet.ID); err == nil {
		t.Fatal("default index set should not be deleted")
	}
	is := testutil.IndexSet("hoge")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	if _, err := server.DeleteIndexSet(is.ID); err != nil {
		t.Fatal("Failed to DeleteIndexSet", err)
	}
	if _, err := server.DeleteIndexSet(""); err == nil {
		t.Fatal("index set id is required")
	}
	if _, err := server.DeleteIndexSet("h"); err == nil {
		t.Fatal(`no index set whose id is "h"`)
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

	prms := is.NewUpdateParams()
	prms.Writable = ptr.PBool(false)

	if _, _, err := client.UpdateIndexSet(prms); err == nil {
		t.Fatal("Default index set must be writable.")
	}
}
