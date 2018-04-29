package store_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/store"
)

func TestNewObjectID(t *testing.T) {
	if id := store.NewObjectID(); len(id) != 24 {
		t.Fatalf(`len(id) = %d, wanted 24: %s`, len(id), id)
	}
}
