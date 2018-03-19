package client_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
)

func TestGetRoleMembers(t *testing.T) {
	test.TestGetRoleMembers(t)
}

func TestAddUserToRole(t *testing.T) {
	test.TestAddUserToRole(t)
}

func TestRemoveUserFromRole(t *testing.T) {
	test.TestRemoveUserFromRole(t)
}
