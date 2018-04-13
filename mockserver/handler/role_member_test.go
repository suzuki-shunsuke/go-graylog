package handler_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
)

func TestHandleRoleMembers(t *testing.T) {
	test.TestGetRoleMembers(t)
}

func TestHandleAddUserToRole(t *testing.T) {
	test.TestAddUserToRole(t)
}

func TestHandleRemoveUserFromRole(t *testing.T) {
	test.TestRemoveUserFromRole(t)
}
