package graylog

// import (
// 	"fmt"
// 	"os"
// 	"testing"
//
// 	"github.com/hashicorp/terraform/helper/resource"
// 	"github.com/hashicorp/terraform/terraform"
//
// 	"github.com/suzuki-shunsuke/go-graylog/v8/client"
// )
//
// var (
// 	terraformTestInputID string
// )
//
// func testDeleteInput(cl *client.Client, key string) resource.TestCheckFunc {
// 	return func(tfState *terraform.State) error {
// 		if _, _, err := cl.GetInput(terraformTestInputID); err == nil {
// 			return fmt.Errorf(`input "%s" must be deleted`, terraformTestInputID)
// 		}
// 		return nil
// 	}
// }
//
// func testCreateInput(cl *client.Client, key string) resource.TestCheckFunc {
// 	return func(tfState *terraform.State) error {
// 		id, err := getIDFromTfState(tfState, key)
// 		if err != nil {
// 			return err
// 		}
// 		terraformTestInputID = id
//
// 		_, _, err = cl.GetInput(id)
// 		return err
// 	}
// }
//
// func testUpdateInput(cl *client.Client, key, title string) resource.TestCheckFunc {
// 	return func(tfState *terraform.State) error {
// 		id, err := getIDFromTfState(tfState, key)
// 		if err != nil {
// 			return err
// 		}
// 		input, _, err := cl.GetInput(id)
// 		if err != nil {
// 			return err
// 		}
// 		if input.Title != title {
// 			return fmt.Errorf("input.Title == %s, wanted %s", input.Title, title)
// 		}
// 		return nil
// 	}
// }
//
// func TestAccInput(t *testing.T) {
// 	cl, server, err := setEnv()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if server != nil {
// 		defer os.Unsetenv("GRAYLOG_WEB_ENDPOINT_URI")
// 	}
//
// 	testAccProvider := Provider()
// 	testAccProviders := map[string]terraform.ResourceProvider{
// 		"graylog": testAccProvider,
// 	}
//
// 	roleTf := `
// resource "graylog_input" "test" {
//   title = "%s"
//   type = "org.graylog2.inputs.syslog.udp.SyslogUDPInput"
//   attributes {
//     bind_address = "0.0.0.0"
//     port = 514
//     recv_buffer_size = 262144
//   }
// }`
// 	createTitle := "terraform test input title"
// 	updateTitle := "terraform test input title updated"
//
// 	key := "graylog_input.test"
// 	if server != nil {
// 		server.Start()
// 		defer server.Close()
// 	}
// 	resource.Test(t, resource.TestCase{
// 		Providers:    testAccProviders,
// 		CheckDestroy: testDeleteInput(cl, key),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: fmt.Sprintf(roleTf, createTitle),
// 				Check:  testCreateInput(cl, key),
// 			},
// 			{
// 				Config: fmt.Sprintf(roleTf, updateTitle),
// 				Check:  testUpdateInput(cl, key, updateTitle),
// 			},
// 		},
// 	})
// }
