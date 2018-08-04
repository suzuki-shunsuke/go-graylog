package logic_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

func TestGetAlerts(t *testing.T) {
	server, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	_, _, _, err = server.GetAlerts(0, 1)
	if err != nil {
		t.Fatal("Failed to GetAlerts", err)
	}
}

func TestGetAlert(t *testing.T) {
	server, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := server.GetAlert(""); err == nil {
		t.Fatal("alert id is required")
	}

	if _, _, err := server.GetAlert("h"); err == nil {
		t.Fatal(`no alert whose id is "h"`)
	}
}
