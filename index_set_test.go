package graylog

import (
	"reflect"
	"testing"
)

func dummyIndexSet() *IndexSet {
	return &IndexSet{
		Id:                    "5a8c086fc006c600013ca6f5",
		Title:                 "Default index set",
		Description:           "The Graylog default index set",
		IndexPrefix:           "graylog",
		Shards:                4,
		Replicas:              0,
		RotationStrategyClass: "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy",
		RotationStrategy: &RotationStrategy{
			Type:            "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
			MaxDocsPerIndex: 20000000},
		RetentionStrategyClass: "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
		RetentionStrategy: &RetentionStrategy{
			Type:               "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
			MaxNumberOfIndices: 20},
		CreationDate:                    "2018-02-20T11:37:19.305Z",
		IndexAnalyzer:                   "standard",
		IndexOptimizationMaxNumSegments: 1,
		IndexOptimizationDisabled:       false,
		Writable:                        true,
		Default:                         true}
}

func dummyIndexSetStats() *IndexSetStats {
	return &IndexSetStats{
		Indices:   2,
		Documents: 0,
		Size:      1412,
	}
}

func TestGetIndexSets(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := dummyIndexSet()
	exp := []IndexSet{*indexSet}
	server.IndexSets[indexSet.Id] = *indexSet
	indexSets, _, _, err := client.GetIndexSets(0, 0)
	if err != nil {
		t.Fatal("Failed to GetIndexSets", err)
	}
	if !reflect.DeepEqual(indexSets, exp) {
		t.Fatalf("client.GetIndexSets() == %v, wanted %v", indexSets, exp)
	}
}

func TestGetIndexSet(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	exp := dummyIndexSet()
	server.IndexSets[exp.Id] = *exp
	act, _, err := client.GetIndexSet(exp.Id)
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
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	exp := dummyIndexSet()
	act, _, err := client.CreateIndexSet(exp)
	if err != nil {
		t.Fatal("Failed to CreateIndexSet", err)
	}
	if act == nil {
		t.Fatal("client.CreateIndexSet() == nil")
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
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := dummyIndexSet()
	server.IndexSets[indexSet.Id] = *indexSet
	indexSet.Description = "changed!"
	updatedIndexSet, _, err := client.UpdateIndexSet(indexSet.Id, indexSet)
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	if !reflect.DeepEqual(*updatedIndexSet, *indexSet) {
		t.Fatalf(
			"client.UpdateIndexSet() == %v, wanted %v", updatedIndexSet, indexSet)
	}
	if _, _, err := client.UpdateIndexSet("", indexSet); err == nil {
		t.Fatal("index set id is required")
	}
	if _, _, err := client.UpdateIndexSet("h", indexSet); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}
}

func TestDeleteIndexSet(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := dummyIndexSet()
	server.IndexSets[indexSet.Id] = *indexSet
	if _, err = client.DeleteIndexSet(indexSet.Id); err != nil {
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
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := dummyIndexSet()
	indexSet.Default = false
	indexSet.Writable = true
	server.IndexSets[indexSet.Id] = *indexSet
	updatedIndexSet, _, err := client.SetDefaultIndexSet(indexSet.Id)
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	if !updatedIndexSet.Default {
		t.Fatal("updatedIndexSet.Default == false")
	}
	indexSet.Default = true
	if !reflect.DeepEqual(*updatedIndexSet, *indexSet) {
		t.Fatalf(
			"client.SetDefaultIndexSet() == %v, wanted %v",
			updatedIndexSet, indexSet)
	}
	if _, _, err := client.SetDefaultIndexSet(""); err == nil {
		t.Fatal("index set id is required")
	}
	if _, _, err := client.SetDefaultIndexSet("h"); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}
}

func TestGetIndexSetStats(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := dummyIndexSet()
	indexSetStats := dummyIndexSetStats()
	server.IndexSets[indexSet.Id] = *indexSet
	server.IndexSetStats[indexSet.Id] = *indexSetStats
	isStats, _, err := client.GetIndexSetStats(indexSet.Id)
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
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := dummyIndexSet()
	indexSetStats := dummyIndexSetStats()
	server.IndexSets[indexSet.Id] = *indexSet
	server.IndexSetStats[indexSet.Id] = *indexSetStats
	isStats, _, err := client.GetAllIndexSetsStats()
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	if !reflect.DeepEqual(*indexSetStats, *isStats) {
		t.Fatalf(
			"client.GetAllIndexSetsStats() == %v, wanted %v", isStats, indexSetStats)
	}
}
