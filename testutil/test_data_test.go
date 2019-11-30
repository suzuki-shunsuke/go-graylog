package testutil_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil/v8"
)

func TestRole(t *testing.T) {
	if testutil.Role() == nil {
		t.Fatal("role is nil")
	}
}

func TestUser(t *testing.T) {
	if testutil.User() == nil {
		t.Fatal("user is nil")
	}
}

func TestDummyAdmin(t *testing.T) {
	if testutil.DummyAdmin() == nil {
		t.Fatal("user is nil")
	}
}

func TestInput(t *testing.T) {
	if testutil.Input() == nil {
		t.Fatal("input is nil")
	}
}

func TestIndexSet(t *testing.T) {
	if testutil.IndexSet("hoge") == nil {
		t.Fatal("index set is nil")
	}
}

func TestStream(t *testing.T) {
	if testutil.Stream() == nil {
		t.Fatal("stream is nil")
	}
}

func TestDummyStream(t *testing.T) {
	if testutil.DummyStream() == nil {
		t.Fatal("stream is nil")
	}
}

func TestStreamRule(t *testing.T) {
	if testutil.StreamRule() == nil {
		t.Fatal("stream rule is nil")
	}
}
