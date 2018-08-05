package graylog

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/suzuki-shunsuke/go-graylog/client"
)

func testDeleteDashboard(
	cl *client.Client, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIDFromTfState(tfState, key)
		if err != nil {
			return err
		}
		if _, _, err := cl.GetDashboard(id); err == nil {
			return fmt.Errorf(`dashboard "%s" must be deleted`, id)
		}
		return nil
	}
}

func testCreateDashboard(
	cl *client.Client, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIDFromTfState(tfState, key)
		if err != nil {
			return err
		}
		_, _, err = cl.GetDashboard(id)
		return err
	}
}

func testUpdateDashboard(
	cl *client.Client, key, title string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIDFromTfState(tfState, key)
		if err != nil {
			return err
		}
		db, _, err := cl.GetDashboard(id)
		if err != nil {
			return err
		}
		if db.Title != title {
			return fmt.Errorf("db.Title is not updated")
		}
		return nil
	}
}

func TestAccDashboard(t *testing.T) {
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

	title := "test-dashboard"
	updatedTitle := "test-dashboard changed"

	dbTf := fmt.Sprintf(`
resource "graylog_dashboard" "zoo" {
  title = "%s"
  description = "test dashboard"
}`, title)
	updateTf := fmt.Sprintf(`
resource "graylog_dashboard" "zoo" {
  title = "%s"
  description = "test dashboard"
}`, updatedTitle)
	if server != nil {
		server.Start()
		defer server.Close()
	}
	key := "graylog_dashboard.zoo"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDeleteDashboard(cl, key),
		Steps: []resource.TestStep{
			{
				Config: dbTf,
				Check: resource.ComposeTestCheckFunc(
					testCreateDashboard(cl, key),
				),
			},
			{
				Config: updateTf,
				Check: resource.ComposeTestCheckFunc(
					testUpdateDashboard(cl, key, updatedTitle),
				),
			},
		},
	})
}
