package inmemory_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/inmemory"
)

func TestGetIndexSetStats(t *testing.T) {
	store := inmemory.NewStore("")
	_, err := store.GetIndexSetStats("foo")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTotalIndexSetStats(t *testing.T) {
	store := inmemory.NewStore("")
	_, err := store.GetTotalIndexSetStats()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetIndexSetStats(t *testing.T) {
	store := inmemory.NewStore("")
	if err := store.SetIndexSetStats("hoge", &graylog.IndexSetStats{}); err != nil {
		t.Fatal(err)
	}
}
