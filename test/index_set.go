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
	indexSet := testutil.DummyNewIndexSet()
	is, _, err := server.AddIndexSet(indexSet)
	if err != nil {
		t.Fatal(err)
	}
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
	if indexSets[0].ID != is.ID {
		t.Fatalf("indexSets[0].ID == %s, wanted %s", indexSets[0].ID, is.ID)
	}
}

func TestGetIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := testutil.DummyNewIndexSet()
	exp, _, err := server.AddIndexSet(is)
	if err != nil {
		t.Fatal(err)
	}
	act, _, err := client.GetIndexSet(exp.ID)
	if err != nil {
		t.Fatal("Failed to GetIndexSet", err)
	}
	if !reflect.DeepEqual(*act, *exp) {
		t.Fatalf("client.GetIndexSet() == %v, wanted %v", act, exp)
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
	exp := testutil.DummyNewIndexSet()
	act, _, err := client.CreateIndexSet(exp)
	if err != nil {
		t.Fatal("Failed to CreateIndexSet", err)
	}
	if act == nil {
		t.Fatal("client.CreateIndexSet() == nil")
	}
	if act.ID == "" {
		t.Fatal("returned IndexSet's id is empty")
	}
	if act.Title != exp.Title {
		t.Fatalf("indexSet.Title == %s, wanted %s", act.Title, exp.Title)
	}
	exp.Title = ""
	if _, _, err := client.CreateIndexSet(exp); err == nil {
		t.Fatal("title is required")
	}
	exp.Title = act.Title
	exp.IndexPrefix = ""
	if _, _, err := client.CreateIndexSet(exp); err == nil {
		t.Fatal("indexPrefix is required")
	}
	exp.IndexPrefix = act.IndexPrefix
	exp.RotationStrategyClass = ""
	if _, _, err := client.CreateIndexSet(exp); err == nil {
		t.Fatal("rotationStrategyClass is required")
	}
	exp.RotationStrategyClass = act.RotationStrategyClass
	exp.RotationStrategy = nil
	if _, _, err := client.CreateIndexSet(exp); err == nil {
		t.Fatal("rotationStrategy is required")
	}
}

func TestUpdateIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := testutil.DummyNewIndexSet()
	indexSet, _, err := server.AddIndexSet(is)
	if err != nil {
		t.Fatal(err)
	}
	indexSet.Description = "changed!"
	updatedIndexSet, _, err := client.UpdateIndexSet(indexSet)
	if err != nil {
		t.Fatal("UpdateIndexSet is failure", err)
	}
	if !reflect.DeepEqual(*updatedIndexSet, *indexSet) {
		t.Fatalf(
			"client.UpdateIndexSet() == %v, wanted %v", updatedIndexSet, indexSet)
	}
	indexSet.ID = ""
	if _, _, err := client.UpdateIndexSet(indexSet); err == nil {
		t.Fatal("index set id is required")
	}
	indexSet.ID = "h"
	if _, _, err := client.UpdateIndexSet(indexSet); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}
	indexSet.Title = ""
	if _, _, err := client.UpdateIndexSet(indexSet); err == nil {
		t.Fatal("title is required")
	}
}

func TestDeleteIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := testutil.DummyNewIndexSet()
	indexSet, _, err := server.AddIndexSet(is)
	if err != nil {
		t.Fatal(err)
	}
	if _, err = client.DeleteIndexSet(indexSet.ID); err != nil {
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
	indexSet := testutil.DummyNewIndexSet()
	indexSet.Default = false
	indexSet.Writable = true
	is, _, err := server.AddIndexSet(indexSet)
	if err != nil {
		t.Fatal(err)
	}
	updatedIndexSet, _, err := client.SetDefaultIndexSet(is.ID)
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	if !updatedIndexSet.Default {
		t.Fatal("updatedIndexSet.Default == false")
	}
	if updatedIndexSet.ID != is.ID {
		t.Fatalf(
			"updatedIndexSet.ID == %v, wanted %v", updatedIndexSet.ID, is.ID)
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
	indexSet := testutil.DummyNewIndexSet()
	indexSet, _, err = server.AddIndexSet(indexSet)
	if err != nil {
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
	indexSet := testutil.DummyNewIndexSet()
	indexSet, _, err = server.AddIndexSet(indexSet)
	if err != nil {
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
