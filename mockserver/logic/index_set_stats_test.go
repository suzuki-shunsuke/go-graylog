package logic_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

func TestGetIndexSetStats(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	iss, _, _, err := lgc.GetIndexSets(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(iss) == 0 {
		t.Fatal("len(iss) == 0")
	}
	is := iss[0]
	if _, _, err := lgc.GetIndexSetStats(is.ID); err != nil {
		t.Fatal(err)
	}
}

func TestGetTotalIndexSetStats(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := lgc.GetTotalIndexSetStats(); err != nil {
		t.Fatal(err)
	}
}
