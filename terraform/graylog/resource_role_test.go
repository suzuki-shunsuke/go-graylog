package graylog

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/suzuki-shunsuke/go-graylog/client"
)

func testDeleteRole(
	cl *client.Client, name string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		if _, _, err := cl.GetRole(name); err == nil {
			return fmt.Errorf(`role "%s" must be deleted`, name)
		}
		return nil
	}
}

func testCreateRole(
	cl *client.Client, name string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		if _, _, err := cl.GetRole(name); err != nil {
			return err
		}
		return nil
	}
}

func testUpdateRole(
	cl *client.Client, name, description string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		role, _, err := cl.GetRole(name)
		if err != nil {
			return err
		}
		if role.Description != description {
			return fmt.Errorf("role.Description is not updated")
		}
		return nil
	}
}

func TestAccRole(t *testing.T) {
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

	roleTf := `
resource "graylog_role" "test-terraform" {
  name = "test terraform name"
  description = "test terraform description"
  permissions = ["*"]
}`
	description := "test terraform description updated"
	updateTf := fmt.Sprintf(`
resource "graylog_role" "test-terraform" {
  name = "test terraform name"
  description = "%s"
  permissions = ["*"]
}`, description)
	name := "test terraform name"
	if server != nil {
		server.Start()
		defer server.Close()
	}
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDeleteRole(cl, name),
		Steps: []resource.TestStep{
			{
				Config: roleTf,
				Check: resource.ComposeTestCheckFunc(
					testCreateRole(cl, name),
				),
			},
			{
				Config: updateTf,
				Check: resource.ComposeTestCheckFunc(
					testUpdateRole(cl, name, description),
				),
			},
		},
	})
}
