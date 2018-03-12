package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
)

func TestGetIndexSets(t *testing.T) {
	test.TestGetIndexSets(t)
}

func TestGetIndexSet(t *testing.T) {
	test.TestGetIndexSet(t)
}

func TestCreateIndexSet(t *testing.T) {
	test.TestCreateIndexSet(t)
}

func TestUpdateIndexSet(t *testing.T) {
	test.TestUpdateIndexSet(t)
}

func TestDeleteIndexSet(t *testing.T) {
	test.TestDeleteIndexSet(t)
}

func TestSetDefaultIndexSet(t *testing.T) {
	test.TestSetDefaultIndexSet(t)
}

func TestGetIndexSetStats(t *testing.T) {
	test.TestGetIndexSetStats(t)
}

func TestGetAllIndexSetsStats(t *testing.T) {
	test.TestGetAllIndexSetsStats(t)
}
