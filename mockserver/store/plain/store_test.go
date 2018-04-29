package plain_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/plain"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
	"github.com/suzuki-shunsuke/go-set"
)

func TestNewStore(t *testing.T) {
	store := plain.NewStore("")
	if store == nil {
		t.Fatal("store is nil")
	}
}

func TestSave(t *testing.T) {
	store := plain.NewStore("")
	if err := store.Save(); err != nil {
		t.Fatal(err)
	}
	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	store = plain.NewStore(tmpfile.Name())
	if err := store.Save(); err != nil {
		t.Fatal(err)
	}
}

func TestLoad(t *testing.T) {
	store := plain.NewStore("")
	if err := store.Load(); err != nil {
		t.Fatal(err)
	}
	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	store = plain.NewStore(tmpfile.Name())
	if err := store.Save(); err != nil {
		t.Fatal(err)
	}
	if err := store.Load(); err != nil {
		t.Fatal(err)
	}
}

func TestAuthorize(t *testing.T) {
	store := plain.NewStore("")
	ok, err := store.Authorize(nil, "users:read")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("not allowed")
	}

	admin := testutil.User()
	admin.Permissions.Add("*")
	ok, err = store.Authorize(admin, "users:read")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("not allowed")
	}

	admin.Permissions = nil
	ok, err = store.Authorize(admin, "users:read")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("not allowed")
	}

	adminRole := testutil.Role()
	if admin.Permissions == nil {
		admin.Permissions = set.NewStrSet()
	}
	adminRole.Permissions.Add("*")
	if admin.Roles == nil {
		admin.Roles = set.NewStrSet()
	}
	admin.Roles.Add(adminRole.Name)
	if err := store.AddRole(adminRole); err != nil {
		t.Fatal(err)
	}
	ok, err = store.Authorize(admin, "users:read")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("not allowed")
	}
	ok, err = store.Authorize(admin, "users:read", "foo")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("not allowed")
	}
	adminRole.Permissions = nil
	if _, err := store.UpdateRole(adminRole.Name, adminRole.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	ok, err = store.Authorize(admin, "users:read", "foo")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("not allowed")
	}
}
