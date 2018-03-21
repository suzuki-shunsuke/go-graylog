package client_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/test"
)

func TestGetStreamRules(t *testing.T) {
	test.TestGetStreamRules(t)
}

func TestCreateStreamRule(t *testing.T) {
	test.TestCreateStreamRule(t)
}

func TestUpdateStreamRule(t *testing.T) {
	test.TestUpdateStreamRule(t)
}

func TestDeleteStreamRule(t *testing.T) {
	test.TestDeleteStreamRule(t)
}