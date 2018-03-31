package logic_test

import (
	"reflect"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
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
	indexSets, err := server.GetIndexSets()
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
	is, _, err = server.GetIndexSet("")
	if err == nil {
		t.Fatal("index set id is empty")
	}
	is, _, err = server.GetIndexSet("h")
	if err == nil {
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
	is.Description = "changed!"

	if _, err := server.UpdateIndexSet(is); err != nil {
		t.Fatal("UpdateIndexSet is failure", err)
	}
	is.ID = ""
	if _, err := server.UpdateIndexSet(is); err == nil {
		t.Fatal("index set id is required")
	}
	is.ID = "h"
	if _, err := server.UpdateIndexSet(is); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}
	is.Title = ""
	if _, err := server.UpdateIndexSet(is); err == nil {
		t.Fatal("title is required")
	}
	if _, err := server.UpdateIndexSet(nil); err == nil {
		t.Fatal("index set is nil")
	}
}

func TestDeleteIndexSet(t *testing.T) {
	server, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	indexSets, err := server.GetIndexSets()
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
