package logic_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
)

func TestAuthenticate(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := lgc.Authenticate("", ""); err == nil {
		t.Fatal("name and password are required")
	}
	if _, _, err := lgc.Authenticate("", "session"); err == nil {
		t.Fatal("session token is not supported")
	}
	if _, _, err := lgc.Authenticate("hoge", "token"); err == nil {
		t.Fatal(`token "hoge" should not been found`)
	}
	if _, _, err := lgc.Authenticate("admin", "admin"); err != nil {
		t.Fatal(err)
	}
}
