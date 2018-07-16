package plain_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/plain"
)

func TestHasAlert(t *testing.T) {
	store := plain.NewStore("")
	ok, err := store.HasAlert("foo")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("alert foo should not exist")
	}
}

func TestGetAlert(t *testing.T) {
	store := plain.NewStore("")
	alert, err := store.GetAlert("foo")
	if err != nil {
		t.Fatal(err)
	}
	if alert != nil {
		t.Fatal("alert foo should not exist")
	}
}

func TestGetAlerts(t *testing.T) {
	store := plain.NewStore("")
	alerts, _, err := store.GetAlerts(0, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(alerts) != 0 {
		t.Fatal("alerts should be nil or empty array")
	}
}
