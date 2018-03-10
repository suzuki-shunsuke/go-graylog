package graylog

import (
	"reflect"
	"testing"
)

func dummyNewIndexSet() *IndexSet {
	return &IndexSet{
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
	indexSet := dummyNewIndexSet()
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
	if indexSets[0].Id != is.Id {
		t.Fatalf("indexSets[0].Id == %s, wanted %s", indexSets[0].Id, is.Id)
	}
}

func TestGetIndexSet(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := dummyNewIndexSet()
	exp, _, err := server.AddIndexSet(is)
	if err != nil {
		t.Fatal(err)
	}
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
	exp := dummyNewIndexSet()
	act, _, err := client.CreateIndexSet(exp)
	if err != nil {
		t.Fatal("Failed to CreateIndexSet", err)
	}
	if act == nil {
		t.Fatal("client.CreateIndexSet() == nil")
	}
	if act.Id == "" {
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
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := dummyNewIndexSet()
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
	indexSet.Id = ""
	if _, _, err := client.UpdateIndexSet(indexSet); err == nil {
		t.Fatal("index set id is required")
	}
	indexSet.Id = "h"
	if _, _, err := client.UpdateIndexSet(indexSet); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}
	indexSet.Title = ""
	if _, _, err := client.UpdateIndexSet(indexSet); err == nil {
		t.Fatal("title is required")
	}
}

func TestDeleteIndexSet(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := dummyNewIndexSet()
	indexSet, _, err := server.AddIndexSet(is)
	if err != nil {
		t.Fatal(err)
	}
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
	indexSet := dummyNewIndexSet()
	indexSet.Default = false
	indexSet.Writable = true
	is, _, err := server.AddIndexSet(indexSet)
	if err != nil {
		t.Fatal(err)
	}
	updatedIndexSet, _, err := client.SetDefaultIndexSet(is.Id)
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	if !updatedIndexSet.Default {
		t.Fatal("updatedIndexSet.Default == false")
	}
	if updatedIndexSet.Id != is.Id {
		t.Fatalf(
			"updatedIndexSet.Id == %v, wanted %v", updatedIndexSet.Id, is.Id)
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
	if _, _, err := client.SetDefaultIndexSet(is.Id); err == nil {
		t.Fatal("Default index set must be writable.")
	}
}

func TestGetIndexSetStats(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSet := dummyNewIndexSet()
	indexSet, _, err = server.AddIndexSet(indexSet)
	if err != nil {
		t.Fatal(err)
	}
	indexSetStats := dummyIndexSetStats()
	server.indexSetStats[indexSet.Id] = *indexSetStats
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
	indexSet := dummyNewIndexSet()
	indexSet, _, err = server.AddIndexSet(indexSet)
	if err != nil {
		t.Fatal(err)
	}
	indexSetStats := dummyIndexSetStats()
	server.indexSetStats[indexSet.Id] = *indexSetStats
	isStats, _, err := client.GetAllIndexSetsStats()
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	if !reflect.DeepEqual(*indexSetStats, *isStats) {
		t.Fatalf(
			"client.GetAllIndexSetsStats() == %v, wanted %v", isStats, indexSetStats)
	}
}
