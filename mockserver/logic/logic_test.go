package logic_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/plain"
)

func TestLogger(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	if logger := lgc.Logger(); logger == nil {
		t.Fatal("logger is nil")
	}
}

func TestSave(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	lgc.SetStore(plain.NewStore(tmpfile.Name()))
	if err := lgc.Save(); err != nil {
		t.Fatal(err)
	}
	if err := lgc.Load(); err != nil {
		t.Fatal(err)
	}
}

func TestLoad(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := lgc.Load(); err != nil {
		t.Fatal(err)
	}
	lgc.SetStore(plain.NewStore("hoge"))
	if err := lgc.Load(); err != nil {
		t.Fatal(err)
	}
}

func TestSetAuth(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	lgc.SetAuth(true)
	if !lgc.Auth() {
		t.Fatal("auth should be true")
	}
	lgc.SetAuth(false)
	if lgc.Auth() {
		t.Fatal("auth should be false")
	}
}

func TestAuthorize(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := lgc.Authorize(nil, "users:read", "admin"); err != nil {
		t.Fatal(err)
	}
	user, _, err := lgc.GetUser("admin")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := lgc.Authorize(user, "users:read", "admin"); err != nil {
		t.Fatal(err)
	}
	nobody, _, err := lgc.GetUser("nobody")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := lgc.Authorize(nobody, "users:read", "admin"); err == nil {
		t.Fatal("authorization should be failure")
	}
}
