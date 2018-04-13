package plain_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/plain"
)

func TestGetIndexSetStats(t *testing.T) {
	store := plain.NewStore("")
	_, err := store.GetIndexSetStats("foo")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTotalIndexSetStats(t *testing.T) {
	store := plain.NewStore("")
	_, err := store.GetTotalIndexSetStats()
	if err != nil {
		t.Fatal(err)
	}
}
