package test

import (
	"reflect"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestSetDefaultIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := testutil.IndexSet("hoge")
	is.Default = false
	is.Writable = true
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	is, _, err = client.SetDefaultIndexSet(is.ID)
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	if !is.Default {
		t.Fatal("updatedIndexSet.Default == false")
	}
	if _, _, err := client.SetDefaultIndexSet(""); err == nil {
		t.Fatal("index set id is required")
	}
	if _, _, err := client.SetDefaultIndexSet("h"); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}

	is.Default = false
	is.Writable = false

	if _, err := server.UpdateIndexSet(is); err != nil {
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
	defer server.Close()
	indexSet := testutil.IndexSet("hoge")
	if _, err = server.AddIndexSet(indexSet); err != nil {
		t.Fatal(err)
	}
	indexSetStats := testutil.DummyIndexSetStats()

	if _, err := server.SetIndexSetStats(indexSet.ID, indexSetStats); err != nil {
		t.Fatal(err)
	}
	isStats, _, err := client.GetIndexSetStats(indexSet.ID)
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	if !reflect.DeepEqual(*indexSetStats, *isStats) {
		t.Fatalf(
			"client.GetIndexSetStats() == %v, wanted %v", isStats, indexSetStats)
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
	defer server.Close()
	indexSet := testutil.IndexSet("hoge")
	if _, err = server.AddIndexSet(indexSet); err != nil {
		t.Fatal(err)
	}
	indexSetStats := testutil.DummyIndexSetStats()
	if _, err := server.SetIndexSetStats(indexSet.ID, indexSetStats); err != nil {
		t.Fatal(err)
	}
	isStats, _, err := client.GetAllIndexSetsStats()
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	if !reflect.DeepEqual(*indexSetStats, *isStats) {
		t.Fatalf(
			"client.GetAllIndexSetsStats() == %v, wanted %v", isStats, indexSetStats)
	}
}
