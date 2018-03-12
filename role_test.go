package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
)

func TestCreateRole(t *testing.T) {
	test.TestCreateRole(t)
}

func TestGetRoles(t *testing.T) {
	test.TestGetRoles(t)
}

func TestGetRole(t *testing.T) {
	test.TestGetRole(t)
}

func TestUpdateRole(t *testing.T) {
	test.TestUpdateRole(t)
}

func TestDeleteRole(t *testing.T) {
	test.TestDeleteRole(t)
}
