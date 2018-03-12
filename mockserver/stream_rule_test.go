package mockserver_test

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
