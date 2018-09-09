package graylog

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/suzuki-shunsuke/go-graylog/client"
)

func testDeleteCollectorConfiguration(
	cl *client.Client, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIDFromTfState(tfState, key)
		if err != nil {
			return err
		}
		if _, _, err := cl.GetCollectorConfiguration(id); err == nil {
			return fmt.Errorf(`collector configuration "%s" must be deleted`, id)
		}
		return nil
	}
}

func testCreateCollectorConfiguration(
	cl *client.Client, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIDFromTfState(tfState, key)
		if err != nil {
			return err
		}
		_, _, err = cl.GetCollectorConfiguration(id)
		return err
	}
}

func testUpdateCollectorConfiguration(
	cl *client.Client, key, name string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIDFromTfState(tfState, key)
		if err != nil {
			return err
		}
		cfg, _, err := cl.GetCollectorConfiguration(id)
		if err != nil {
			return err
		}
		if cfg.Name != name {
			return fmt.Errorf("collector configuration name is not updated")
		}
		return nil
	}
}

func TestAccCollectorConfiguration(t *testing.T) {
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

	name := "test-cfg"
	updatedName := "test-cfg changed"

	dbTf := fmt.Sprintf(`
resource "graylog_collector_configuration" "zoo" {
  name = "%s"
	tags = ["test"]
}`, name)
	updateTf := fmt.Sprintf(`
resource "graylog_collector_configuration" "zoo" {
  name = "%s"
	tags = ["test"]
}`, updatedName)
	if server != nil {
		server.Start()
		defer server.Close()
	}
	key := "graylog_collector_configuration.zoo"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDeleteCollectorConfiguration(cl, key),
		Steps: []resource.TestStep{
			{
				Config: dbTf,
				Check: resource.ComposeTestCheckFunc(
					testCreateCollectorConfiguration(cl, key),
				),
			},
			{
				Config: updateTf,
				Check: resource.ComposeTestCheckFunc(
					testUpdateCollectorConfiguration(cl, key, updatedName),
				),
			},
		},
	})
}
