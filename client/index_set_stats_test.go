package client_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
)

func TestGetIndexSetStats(t *testing.T) {
	test.TestGetIndexSetStats(t)
}

func TestGetAllIndexSetsStats(t *testing.T) {
	test.TestGetAllIndexSetsStats(t)
}
