package graylog

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/suzuki-shunsuke/go-graylog/client"
)

func testDeleteUser(
	ctx context.Context, cl *client.Client, name string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		if _, _, err := cl.GetUser(ctx, name); err == nil {
			return fmt.Errorf(`user "%s" must be deleted`, name)
		}
		return nil
	}
}

func testCreateUser(
	ctx context.Context, cl *client.Client, name string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		_, _, err := cl.GetUser(ctx, name)
		return err
	}
}

func testUpdateUser(
	ctx context.Context, cl *client.Client, name, fullName string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		user, _, err := cl.GetUser(ctx, name)
		if err != nil {
			return err
		}
		if user.FullName != fullName {
			return fmt.Errorf("user.FullName is not updated")
		}
		return nil
	}
}

func TestAccUser(t *testing.T) {
	ctx := context.Background()
	cl, server, err := setEnv()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer os.Unsetenv("GRAYLOG_WEB_ENDPOINT_URI")
	}

	testAccProvider := Provider()
	testAccProviders := map[string]terraform.ResourceProvider{
		"graylog": testAccProvider,
	}

	name := "test terraform name"

	// TODO: "users:edit:{name}" and "users:passwordchange:{name}" is automatically added
	userTf := fmt.Sprintf(`
resource "graylog_user" "zoo" {
  username = "%s"
  password = "password"
  email = "zoo@example.com"
  full_name = "zooull"
  permissions = [
	  "users:read:zoo",
		"users:edit:%s",
		"users:passwordchange:%s"
  ]
}`, name, name, name)
	fullName := "new full name"
	updateTf := fmt.Sprintf(`
resource "graylog_user" "zoo" {
  username = "%s"
  password = "password"
  email = "zoo@example.com"
  full_name = "%s"
  permissions = [
	  "users:read:zoo",
		"users:edit:%s",
		"users:passwordchange:%s"
  ]
}`, name, fullName, name, name)
	if server != nil {
		server.Start()
		defer server.Close()
	}
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDeleteUser(ctx, cl, name),
		Steps: []resource.TestStep{
			{
				Config: userTf,
				Check: resource.ComposeTestCheckFunc(
					testCreateUser(ctx, cl, name),
				),
			},
			{
				Config: updateTf,
				Check: resource.ComposeTestCheckFunc(
					testUpdateUser(ctx, cl, name, fullName),
				),
			},
		},
	})
}
