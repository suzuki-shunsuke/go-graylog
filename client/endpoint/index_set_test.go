package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

func TestIndexSets(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s", apiURL, "system/indices/index_sets")
	if ep.IndexSets() != exp {
		t.Fatalf(`ep.IndexSets() = "%s", wanted "%s"`, ep.IndexSets(), exp)
	}
}

func TestIndexSet(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s/%s", apiURL, "system/indices/index_sets", ID)
	act, err := ep.IndexSet(ID)
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.IndexSet("%s") = "%s", wanted "%s"`, ID, act.String(), exp)
	}
}

func TestSetDefaultIndexSet(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s/%s/default", apiURL, "system/indices/index_sets", ID)
	act, err := ep.SetDefaultIndexSet(ID)
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.SetDefaultIndexSet("%s") = "%s", wanted "%s"`, ID, act.String(), exp)
	}
}

func TestIndexSetsStats(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s", apiURL, "system/indices/index_sets/stats")
	if ep.IndexSetsStats() != exp {
		t.Fatalf(`ep.IndexSetsStats() = "%s", wanted "%s"`, ep.IndexSetsStats(), exp)
	}
}

func TestIndexSetStats(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	if err != nil {
		t.Fatal(err)
	}
	exp := fmt.Sprintf("%s/%s/%s/stats", apiURL, "system/indices/index_sets", ID)
	act, err := ep.IndexSetStats(ID)
	if err != nil {
		t.Fatal(err)
	}
	if act.String() != exp {
		t.Fatalf(`ep.IndexSetStats("%s") = "%s", wanted "%s"`, ID, act.String(), exp)
	}
}
