package graylog

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/suzuki-shunsuke/go-graylog/client"
)

func testDeleteUser(
	cl *client.Client, name string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		if _, _, err := cl.GetUser(name); err == nil {
			return fmt.Errorf(`user "%s" must be deleted`, name)
		}
		return nil
	}
}

func testCreateUser(
	cl *client.Client, name string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		if _, _, err := cl.GetUser(name); err != nil {
			return err
		}
		return nil
	}
}

func testUpdateUser(
	cl *client.Client, name, fullName string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		user, _, err := cl.GetUser(name)
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

	userTf := fmt.Sprintf(`
resource "graylog_user" "zoo" {
  username = "%s"
  password = "password"
  email = "zoo@example.com"
  full_name = "zooull"
  permissions = ["users:read:zoo"]
	session_timeout_ms = 28800000
	timezone = "UTC"
}`, name)
	fullName := "new full name"
	updateTf := fmt.Sprintf(`
resource "graylog_user" "zoo" {
  username = "%s"
  password = "password"
  email = "zoo@example.com"
  full_name = "%s"
  permissions = ["users:read:zoo"]
	session_timeout_ms = 28800000
	timezone = "UTC"
}`, name, fullName)
	if server != nil {
		server.Start()
		defer server.Close()
	}
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDeleteUser(cl, name),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: userTf,
				Check: resource.ComposeTestCheckFunc(
					testCreateUser(cl, name),
				),
			},
			resource.TestStep{
				Config: updateTf,
				Check: resource.ComposeTestCheckFunc(
					testUpdateUser(cl, name, fullName),
				),
			},
		},
	})
}
