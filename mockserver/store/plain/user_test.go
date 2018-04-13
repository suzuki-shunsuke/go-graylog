package plain_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/plain"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHasUser(t *testing.T) {
	store := plain.NewStore("")
	ok, err := store.HasUser("foo")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("user foo should not exist")
	}
}

func TestGetUser(t *testing.T) {
	store := plain.NewStore("")
	user, err := store.GetUser("foo")
	if err != nil {
		t.Fatal(err)
	}
	if user != nil {
		t.Fatal("user foo should not exist")
	}
}

func TestGetUsers(t *testing.T) {
	store := plain.NewStore("")
	users, err := store.GetUsers()
	if err != nil {
		t.Fatal(err)
	}
	if users != nil && len(users) != 0 {
		t.Fatal("users should be nil or empty array")
	}
}

func TestAddUser(t *testing.T) {
	store := plain.NewStore("")
	user := testutil.User()
	if err := store.AddUser(user); err != nil {
		t.Fatal(err)
	}
	r, err := store.GetUser(user.Username)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("user is nil")
	}
}

func TestUpdateUser(t *testing.T) {
	store := plain.NewStore("")
	user := testutil.User()
	if err := store.AddUser(user); err != nil {
		t.Fatal(err)
	}
	r, err := store.GetUser(user.Username)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("user is nil")
	}
	user.FullName += " changed"
	if err := store.UpdateUser(user); err != nil {
		t.Fatal(err)
	}
	r, err = store.GetUser(user.Username)
	if err != nil {
		t.Fatal(err)
	}
	if user.FullName != r.FullName {
		t.Fatalf(`user.FullName = "%s", wanted "%s"`, r.FullName, user.FullName)
	}
}

func TestDeleteUser(t *testing.T) {
	store := plain.NewStore("")
	if err := store.DeleteUser("foo"); err != nil {
		t.Fatal(err)
	}
	user := testutil.User()
	if err := store.AddUser(user); err != nil {
		t.Fatal(err)
	}
	if err := store.DeleteUser(user.Username); err != nil {
		t.Fatal(err)
	}
	ok, err := store.HasUser(user.Username)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("user should be deleted")
	}
}
