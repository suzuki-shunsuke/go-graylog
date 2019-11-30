package graylog

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/suzuki-shunsuke/go-graylog/v8/client"
	"github.com/suzuki-shunsuke/go-graylog/v8/testutil"
	"github.com/suzuki-shunsuke/graylog-mock-server/mockserver"
)

func testDeleteIndexSet(
	ctx context.Context, cl *client.Client, server *mockserver.Server, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIDFromTfState(tfState, key)
		if err != nil {
			return err
		}
		testutil.WaitAfterDeleteIndexSet(server)
		if _, _, err := cl.GetIndexSet(ctx, id); err == nil {
			return fmt.Errorf(`indexSet "%s" must be deleted`, id)
		}
		return nil
	}
}

func testCreateIndexSet(
	ctx context.Context,
	cl *client.Client, server *mockserver.Server, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIDFromTfState(tfState, key)
		if err != nil {
			return err
		}
		testutil.WaitAfterCreateIndexSet(server)

		_, _, err = cl.GetIndexSet(ctx, id)
		return err
	}
}

func testUpdateIndexSet(
	ctx context.Context,
	cl *client.Client, key, title string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIDFromTfState(tfState, key)
		if err != nil {
			return err
		}
		indexSet, _, err := cl.GetIndexSet(ctx, id)
		if err != nil {
			return err
		}
		if indexSet.Title != title {
			return fmt.Errorf("indexSet.Title == %s, wanted %s", indexSet.Title, title)
		}
		return nil
	}
}

func TestAccIndexSet(t *testing.T) {
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

	u, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	prefix := u.String()
	roleTf := `
resource "graylog_index_set" "test" {
  title = "%s"
	description = "terraform test index set description"
	index_prefix = "%s"
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
  index_optimization_max_num_segments = 1
}`

	updateTitle := "terraform test index set title updated"
	key := "graylog_index_set.test"
	if server != nil {
		server.Start()
		defer server.Close()
	}
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDeleteIndexSet(ctx, cl, server, key),
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(roleTf, "terraform test index set title", prefix),
				Check: resource.ComposeTestCheckFunc(
					testCreateIndexSet(ctx, cl, server, key),
				),
			},
			{
				Config: fmt.Sprintf(roleTf, updateTitle, prefix),
				Check: resource.ComposeTestCheckFunc(
					testUpdateIndexSet(ctx, cl, key, updateTitle),
				),
			},
		},
	})
}
