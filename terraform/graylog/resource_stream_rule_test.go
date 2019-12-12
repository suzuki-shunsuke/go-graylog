package graylog

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/suzuki-shunsuke/go-graylog/v8/client"
	"github.com/suzuki-shunsuke/go-graylog/v8/testutil"
	"github.com/suzuki-shunsuke/graylog-mock-server/mockserver"
)

func testDeleteStreamRule(
	ctx context.Context, cl *client.Client, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		rs, ok := tfState.RootModule().Resources[key]
		if !ok {
			return errors.New("Not found: " + key)
		}
		id := rs.Primary.ID
		streamID, ok := rs.Primary.Attributes["stream_id"]
		if !ok {
			return errors.New("stream_id is not found: " + key)
		}
		if _, _, err := cl.GetStreamRule(ctx, streamID, id); err == nil {
			return fmt.Errorf(`stream rule "%s" must be deleted`, id)
		}
		return nil
	}
}

func testCreateStreamRule(
	ctx context.Context, cl *client.Client, server *mockserver.Server, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		rs, ok := tfState.RootModule().Resources[key]
		if !ok {
			return errors.New("Not found: " + key)
		}
		id := rs.Primary.ID
		streamID, ok := rs.Primary.Attributes["stream_id"]
		if !ok {
			return errors.New("stream_id is not found: " + key)
		}
		testutil.WaitAfterCreateIndexSet(server)
		_, _, err := cl.GetStreamRule(ctx, streamID, id)
		return err
	}
}

func testUpdateStreamRule(
	ctx context.Context, cl *client.Client, key, desc string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		rs, ok := tfState.RootModule().Resources[key]
		if !ok {
			return errors.New("Not found: " + key)
		}
		id := rs.Primary.ID
		streamID, ok := rs.Primary.Attributes["stream_id"]
		if !ok {
			return errors.New("stream_id is not found: " + key)
		}
		rule, _, err := cl.GetStreamRule(ctx, streamID, id)
		if err != nil {
			return err
		}
		if rule.Description != desc {
			return fmt.Errorf("rule.Description == %s, wanted %s", rule.Description, desc)
		}
		return nil
	}
}

func TestAccStreamRule(t *testing.T) {
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

	roleTf := `
resource "graylog_index_set" "test" {
  title = "terraform test index set"
	description = "terraform test index set description"
	index_prefix = "test"
	shards = 4
	replicas = 0
  rotation_strategy_class = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy"
  rotation_strategy {
    type = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"
  }
  retention_strategy_class = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy"
  retention_strategy {
    type = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"
  }
  index_analyzer = "standard"
	writable = true
  index_optimization_max_num_segments = 1
}

resource "graylog_stream" "test" {
  title = "stream test"
	index_set_id = "${graylog_index_set.test.id}"
	matching_type = "AND"
}

resource "graylog_stream_rule" "test" {
  field       = "tag"
  stream_id   = "${graylog_stream.test.id}"
  description = "%s"
  type        = 1
  value       = "stream_rule.test"
}`
	createDesc := "terraform stream rule test"
	updateDesc := "terraform stream rule test updated"

	key := "graylog_stream_rule.test"
	if server != nil {
		server.Start()
		defer server.Close()
	}
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDeleteStreamRule(ctx, cl, key),
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(roleTf, createDesc),
				Check: resource.ComposeTestCheckFunc(
					testCreateStreamRule(ctx, cl, server, key),
				),
			},
			{
				Config: fmt.Sprintf(roleTf, updateDesc),
				Check: resource.ComposeTestCheckFunc(
					testUpdateStreamRule(ctx, cl, key, updateDesc),
				),
			},
		},
	})
}
