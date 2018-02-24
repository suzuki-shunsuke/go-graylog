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

func TestGetIndexSets(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	indexSet := dummyIndexSet()
	exp := []IndexSet{*indexSet}
	server.IndexSets[indexSet.Id] = *indexSet
	indexSets, _, err := client.GetIndexSets(0, 0)
	if err != nil {
		t.Error("Failed to GetIndexSets", err)
		return
	}
	if !reflect.DeepEqual(indexSets, exp) {
		t.Errorf("client.GetIndexSets() == %v, wanted %v", indexSets, exp)
	}
}

func TestGetIndexSet(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	exp := dummyIndexSet()
	server.IndexSets[exp.Id] = *exp
	act, err := client.GetIndexSet(exp.Id)
	if err != nil {
		t.Error("Failed to GetIndexSet", err)
		return
	}
	if !reflect.DeepEqual(*act, *exp) {
		t.Errorf("client.GetIndexSet() == %v, wanted %v", act, exp)
	}
}

func TestCreateIndexSet(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	exp := dummyIndexSet()
	act, err := client.CreateIndexSet(exp)
	if err != nil {
		t.Error("Failed to CreateIndexSet", err)
		return
	}
	if act == nil {
		t.Error("client.CreateIndexSet() == nil")
		return
	}
	if act.Title != exp.Title {
		t.Errorf("indexSet.Title == %s, wanted %s", act.Title, exp.Title)
	}
}

func TestUpdateIndexSet(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	indexSet := dummyIndexSet()
	server.IndexSets[indexSet.Id] = *indexSet
	indexSet.Description = "changed!"
	updatedIndexSet, err := client.UpdateIndexSet(indexSet.Id, indexSet)
	if err != nil {
		t.Error("Failed to UpdateIndexSet", err)
		return
	}
	if !reflect.DeepEqual(*updatedIndexSet, *indexSet) {
		t.Errorf("client.UpdateIndexSet() == %v, wanted %v", updatedIndexSet, indexSet)
	}
}

func TestDeleteIndexSet(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	indexSet := dummyIndexSet()
	server.IndexSets[indexSet.Id] = *indexSet
	err = client.DeleteIndexSet(indexSet.Id)
	if err != nil {
		t.Error("Failed to DeleteIndexSet", err)
		return
	}
}
