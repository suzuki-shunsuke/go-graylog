package test

import (
	"reflect"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetIndexSets(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSets, _, _, err := client.GetIndexSets(0, 0)
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

	is := testutil.DummyNewIndexSet("hoge")
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
	exp := testutil.DummyNewIndexSet("hoge")
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
	is := testutil.DummyNewIndexSet("fuga")
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
	is := testutil.DummyNewIndexSet("hoge")
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
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := testutil.DummyNewIndexSet("hoge")
	indexSet.Default = false
	indexSet.Writable = true
	if _, err := server.AddIndexSet(indexSet); err != nil {
		t.Fatal(err)
	}
	updatedIndexSet, _, err := client.SetDefaultIndexSet(indexSet.ID)
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	if !updatedIndexSet.Default {
		t.Fatal("updatedIndexSet.Default == false")
	}
	if updatedIndexSet.ID != indexSet.ID {
		t.Fatalf(
			"updatedIndexSet.ID == %v, wanted %v", updatedIndexSet.ID, indexSet.ID)
	}
	if _, _, err := client.SetDefaultIndexSet(""); err == nil {
		t.Fatal("index set id is required")
	}
	if _, _, err := client.SetDefaultIndexSet("h"); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}

	indexSet.Default = false
	indexSet.Writable = false

	if _, err := server.UpdateIndexSet(indexSet); err != nil {
		t.Fatal(err)
	}
	if _, _, err := client.SetDefaultIndexSet(indexSet.ID); err == nil {
		t.Fatal("Default index set must be writable.")
	}
}

func TestGetIndexSetStats(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := testutil.DummyNewIndexSet("hoge")
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
	if _, _, err := client.GetIndexSetStats("h"); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}
}

func TestGetAllIndexSetsStats(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := testutil.DummyNewIndexSet("hoge")
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
