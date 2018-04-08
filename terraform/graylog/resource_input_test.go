package graylog

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/client"
)

func testDeleteInput(
	cl *client.Client, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIdFromTfState(tfState, key)
		if err != nil {
			return err
		}
		if _, _, err := cl.GetInput(id); err == nil {
			return fmt.Errorf(`input "%s" must be deleted`, id)
		}
		return nil
	}
}

func testCreateInput(
	cl *client.Client, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIdFromTfState(tfState, key)
		if err != nil {
			return err
		}

		if _, _, err := cl.GetInput(id); err != nil {
			return err
		}
		return nil
	}
}

func testUpdateInput(
	cl *client.Client, key, title string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIdFromTfState(tfState, key)
		if err != nil {
			return err
		}
		input, _, err := cl.GetInput(id)
		if err != nil {
			return err
		}
		if input.Title != title {
			return fmt.Errorf("input.Title == %s, wanted %s", input.Title, title)
		}
		return nil
	}
}

func TestAccInput(t *testing.T) {
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

	input := &graylog.Input{
		Title: "terraform test input title",
		Type:  "org.graylog2.inputs.syslog.udp.SyslogUDPInput",
		Configuration: &graylog.InputConfiguration{
			BindAddress:    "0.0.0.0",
			Port:           514,
			RecvBufferSize: 262144,
		},
	}
	tc := &tfConf{
		Resource: map[string]map[string]interface{}{
			"graylog_input": {"test": input}},
	}

	b, err := json.Marshal(tc)
	if err != nil {
		t.Fatal(err)
	}

	input.Title = "terraform test input title updated"
	tc.Resource["graylog_input"]["test"] = input
	ub, err := json.Marshal(tc)
	if err != nil {
		t.Fatal(err)
	}

	key := "graylog_input.test"
	if server != nil {
		server.Start()
		defer server.Close()
	}
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDeleteInput(cl, key),
		Steps: []resource.TestStep{
			{
				Config: string(b),
				Check: resource.ComposeTestCheckFunc(
					testCreateInput(cl, key),
				),
			},
			{
				Config: string(ub),
				Check: resource.ComposeTestCheckFunc(
					testUpdateInput(cl, key, input.Title),
				),
			},
		},
	})
}
