package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
)

func TestCreateInput(t *testing.T) {
	test.TestCreateInput(t)
}

func TestGetInputs(t *testing.T) {
	test.TestGetInputs(t)
}

func TestGetInput(t *testing.T) {
	test.TestGetInput(t)
}

func TestUpdateInput(t *testing.T) {
	test.TestUpdateInput(t)
}

func TestDeleteInput(t *testing.T) {
	test.TestDeleteInput(t)
}
