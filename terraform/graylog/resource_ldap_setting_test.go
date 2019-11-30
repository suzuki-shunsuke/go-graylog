package graylog

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/suzuki-shunsuke/go-graylog/client/v8"
)

func testDeleteLDAPSetting(
	ctx context.Context, cl *client.Client,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		ls, _, err := cl.GetLDAPSetting(ctx)
		if err != nil {
			return err
		}
		if ls.DisplayNameAttribute != "" {
			return fmt.Errorf(
				`display_name_attribute = "%s", wanted ""`,
				ls.DisplayNameAttribute)
		}
		return nil
	}
}

func testCreateLDAPSetting(
	ctx context.Context, cl *client.Client, exp string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		ls, _, err := cl.GetLDAPSetting(ctx)
		if err != nil {
			return err
		}
		if ls.DisplayNameAttribute != exp {
			return fmt.Errorf(
				`display_name_attribute = "%s", wanted "%s"`,
				ls.DisplayNameAttribute, exp)
		}
		return nil
	}
}

func testUpdateLDAPSetting(
	ctx context.Context, cl *client.Client, exp string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		ls, _, err := cl.GetLDAPSetting(ctx)
		if err != nil {
			return err
		}
		if ls.DisplayNameAttribute != exp {
			return fmt.Errorf(
				`display_name_attribute = "%s", wanted "%s"`,
				ls.DisplayNameAttribute, exp)
		}
		return nil
	}
}

func TestAccLDAPSetting(t *testing.T) {
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

	createTf := `
resource "graylog_ldap_setting" "test-terraform" {
  system_username = "admin"
  system_password = "password"
  ldap_uri = "ldap://localhost:389"
	display_name_attribute = "displayname"
	search_base = "OU=user,OU=foo,DC=example,DC=com"
	search_pattern = "(cn={0})"
	default_group = "Reader"
}`
	updateTf := `
resource "graylog_ldap_setting" "test-terraform" {
  system_username = "admin"
  system_password = "password"
  ldap_uri = "ldap://localhost:389"
	display_name_attribute = "displayname_updated"
	search_base = "OU=user,OU=foo,DC=example,DC=com"
	search_pattern = "(cn={0})"
	default_group = "Reader"
}`
	if server != nil {
		server.Start()
		defer server.Close()
	}
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDeleteLDAPSetting(ctx, cl),
		Steps: []resource.TestStep{
			{
				Config: createTf,
				Check: resource.ComposeTestCheckFunc(
					testCreateLDAPSetting(ctx, cl, "displayname"),
				),
			},
			{
				Config: updateTf,
				Check: resource.ComposeTestCheckFunc(
					testUpdateLDAPSetting(ctx, cl, "displayname_updated"),
				),
			},
		},
	})
}
