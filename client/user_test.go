package client_test

import (
	"context"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestDeleteUser(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, err := client.DeleteUser(ctx, ""); err == nil {
		t.Fatal("username is required")
	}
	if _, err := client.DeleteUser(ctx, "h"); err == nil {
		t.Fatal(`no user with name "h" is found`)
	}
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	// nil check
	if _, err := client.CreateUser(ctx, nil); err == nil {
		t.Fatal("user is nil")
	}
	user := testutil.User()
	client.DeleteUser(ctx, user.Username)
	if _, err := client.CreateUser(ctx, user); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteUser(ctx, user.Username)
	// error check
	user.Username = ""
	if _, err := client.CreateUser(ctx, user); err == nil {
		t.Fatal("user name is empty")
	}
}

func TestGetUsers(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	users, _, err := client.GetUsers(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if users == nil {
		t.Fatal("users is nil")
	}
	if len(users) == 0 {
		t.Fatal("users is empty")
	}
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	user := testutil.User()
	client.DeleteUser(ctx, user.Username)
	if _, _, err := client.GetUser(ctx, user.Username); err == nil {
		t.Fatal("user should be deleted")
	}
	if _, err := client.CreateUser(ctx, user); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteUser(ctx, user.Username)
	u, _, err := client.GetUser(ctx, user.Username)
	if err != nil {
		t.Fatal(err)
	}
	if u == nil {
		t.Fatal("user is nil")
	}
	if u.Username != user.Username {
		t.Fatalf(`user.Username = "%s", wanted "%s"`, u.Username, user.Username)
	}
	if _, _, err := client.GetUser(ctx, ""); err == nil {
		t.Fatal("user name is required")
	}
}

func TestUpdateUser(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	user := testutil.User()
	client.DeleteUser(ctx, user.Username)
	if _, err := client.UpdateUser(ctx, user.NewUpdateParams()); err == nil {
		t.Fatal("user should be deleted")
	}
	if _, err := client.CreateUser(ctx, user); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteUser(ctx, user.Username)
	if _, err := client.UpdateUser(ctx, user.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	user.Username = ""
	if _, err := client.UpdateUser(ctx, user.NewUpdateParams()); err == nil {
		t.Fatal("user name is required")
	}
	if _, err := client.UpdateUser(ctx, nil); err == nil {
		t.Fatal("user is required")
	}
}
