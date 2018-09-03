package graylog

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/suzuki-shunsuke/go-graylog/client"
)

func testDeleteLDAPSetting(
	cl *client.Client,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		ls, _, err := cl.GetLDAPSetting()
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
	cl *client.Client, exp string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		ls, _, err := cl.GetLDAPSetting()
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
	cl *client.Client, exp string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		ls, _, err := cl.GetLDAPSetting()
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
  enabled = true
  system_username = ""
  system_password = ""
  ldap_uri = "ldap://localhost:389"
  use_start_tls = false
  trust_all_certificates = false
	display_name_attribute = "displayname"
}`
	updateTf := `
resource "graylog_ldap_setting" "test-terraform" {
  enabled = false
  system_username = ""
  system_password = ""
  ldap_uri = "ldap://localhost:389"
  use_start_tls = false
  trust_all_certificates = false
	display_name_attribute = "displayname_updated"
}`
	if server != nil {
		server.Start()
		defer server.Close()
	}
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDeleteLDAPSetting(cl),
		Steps: []resource.TestStep{
			{
				Config: createTf,
				Check: resource.ComposeTestCheckFunc(
					testCreateLDAPSetting(cl, "displayname"),
				),
			},
			{
				Config: updateTf,
				Check: resource.ComposeTestCheckFunc(
					testUpdateLDAPSetting(cl, "displayname_updated"),
				),
			},
		},
	})
}
