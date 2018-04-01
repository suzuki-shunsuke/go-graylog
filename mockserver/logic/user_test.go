package logic_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
	"github.com/suzuki-shunsuke/go-set"
)

func TestAddUser(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("basic", func(t *testing.T) {
		user := testutil.User()
		if user.Roles == nil {
			user.Roles = set.NewStrSet("Admin")
		}
		if _, err := lgc.AddUser(user); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("username is required", func(t *testing.T) {
		user := testutil.User()
		user.Username = ""
		if _, err := lgc.AddUser(user); err == nil {
			t.Fatal("user.Username is required")
		}
	})
	t.Run("username duplicate", func(t *testing.T) {
		user := testutil.User()
		if _, err := lgc.AddUser(user); err == nil {
			t.Fatal("user name is duplicate")
		}
	})
	t.Run("unexisting role name", func(t *testing.T) {
		user := testutil.User()
		if user.Roles == nil {
			user.Roles = set.NewStrSet("aa")
		}
		user.Username += " changed"
		if _, err := lgc.AddUser(user); err == nil {
			t.Fatal("unexisting role name")
		}
	})
}

func TestGetUser(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("basic", func(t *testing.T) {
		if _, _, err := lgc.GetUser("admin"); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("username is required", func(t *testing.T) {
		if _, _, err := lgc.GetUser(""); err == nil {
			t.Fatal("username is required")
		}
	})
	t.Run("not found", func(t *testing.T) {
		if _, _, err := lgc.GetUser("h"); err == nil {
			t.Fatal("user not found")
		}
	})
}

func TestGetUsers(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	if _, _, err := lgc.GetUsers(); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateUser(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("nil", func(t *testing.T) {
		if _, err := lgc.UpdateUser(nil); err == nil {
			t.Fatal("user is nil")
		}
	})
	user := testutil.User()
	name := user.Username
	if user.Roles == nil {
		user.Roles = set.NewStrSet("Admin")
	}
	if _, err := lgc.AddUser(user); err != nil {
		t.Fatal(err)
	}
	t.Run("basic", func(t *testing.T) {
		if _, err := lgc.UpdateUser(user); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("username is required", func(t *testing.T) {
		user.Username = ""
		if _, err := lgc.UpdateUser(user); err == nil {
			t.Fatal("user.Username is required")
		}
	})
	t.Run("validation error", func(t *testing.T) {
		// TODO
	})
	t.Run("check role", func(t *testing.T) {
		user.Username = name
		user.Roles = set.NewStrSet("aa")
		if _, err := lgc.UpdateUser(user); err == nil {
			t.Fatal("unexisting role")
		}
	})
}

func TestDeleteUser(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("username is required", func(t *testing.T) {
		if _, err := lgc.DeleteUser(""); err == nil {
			t.Fatal("username is required")
		}
	})
	user := testutil.User()
	if _, err := lgc.AddUser(user); err != nil {
		t.Fatal(err)
	}
	t.Run("basic", func(t *testing.T) {
		if _, err := lgc.DeleteUser(user.Username); err != nil {
			t.Fatal(err)
		}
	})
}

func TestUserList(t *testing.T) {
	lgc, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("basic", func(t *testing.T) {
		if _, err := lgc.UserList(); err != nil {
			t.Fatal(err)
		}
	})
}
