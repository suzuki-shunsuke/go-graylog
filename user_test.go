package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
)

func TestCreateUser(t *testing.T) {
	test.TestCreateUser(t)
}

func TestGetUsers(t *testing.T) {
	test.TestGetUsers(t)
}

func TestGetUser(t *testing.T) {
	test.TestGetUser(t)
}

func TestUpdateUser(t *testing.T) {
	test.TestUpdateUser(t)
}

func TestDeleteUser(t *testing.T) {
	test.TestDeleteUser(t)
}
