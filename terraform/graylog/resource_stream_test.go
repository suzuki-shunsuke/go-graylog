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
	"github.com/suzuki-shunsuke/go-graylog/mockserver"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func testDeleteStream(
	cl *client.Client, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIdFromTfState(tfState, key)
		if err != nil {
			return err
		}
		if _, _, err := cl.GetStream(id); err == nil {
			return fmt.Errorf(`stream "%s" must be deleted`, id)
		}
		return nil
	}
}

func testCreateStream(
	cl *client.Client, server *mockserver.Server, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIdFromTfState(tfState, key)
		if err != nil {
			return err
		}
		testutil.WaitAfterCreateIndexSet(server)

		if _, _, err := cl.GetStream(id); err != nil {
			return err
		}
		return nil
	}
}

func testUpdateStream(
	cl *client.Client, key, title string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIdFromTfState(tfState, key)
		if err != nil {
			return err
		}
		stream, _, err := cl.GetStream(id)
		if err != nil {
			return err
		}
		if stream.Title != title {
			return fmt.Errorf("stream.Title == %s, wanted %s", stream.Title, title)
		}
		return nil
	}
}

func TestAccStream(t *testing.T) {
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

	indexSet := &graylog.IndexSet{
		Title:                 "terraform test index set title",
		Description:           "terraform test index set description",
		IndexPrefix:           "terraform-test",
		Shards:                4,
		Replicas:              0,
		RotationStrategyClass: "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy",
		RotationStrategy: &graylog.RotationStrategy{
			Type:            "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig",
			MaxDocsPerIndex: 20000000},
		RetentionStrategyClass: "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy",
		RetentionStrategy: &graylog.RetentionStrategy{
			Type:               "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig",
			MaxNumberOfIndices: 20},
		CreationDate:                    "2018-02-20T11:37:19.305Z",
		IndexAnalyzer:                   "standard",
		IndexOptimizationMaxNumSegments: 1,
		IndexOptimizationDisabled:       false,
		Writable:                        true,
		Default:                         false}
	stream := &graylog.Stream{
		Title:        "terraform test",
		IndexSetID:   "${graylog_index_set.test.id}",
		MatchingType: "AND",
	}

	tfConf := &TFConf{
		Resource: map[string]map[string]interface{}{
			"graylog_stream":    {"test": stream},
			"graylog_index_set": {"test": indexSet}},
	}

	b, err := json.Marshal(tfConf)
	if err != nil {
		t.Fatal(err)
	}

	updateConf := *tfConf
	stream.Title = "terraform test updated"
	updateConf.Resource["graylog_stream"]["test"] = stream

	updateByte, err := json.Marshal(updateConf)
	if err != nil {
		t.Fatal(err)
	}

	key := "graylog_stream.test"
	if server != nil {
		server.Start()
		defer server.Close()
	}
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDeleteStream(cl, key),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: string(b),
				Check: resource.ComposeTestCheckFunc(
					testCreateStream(cl, server, key),
				),
			},
			resource.TestStep{
				Config: string(updateByte),
				Check: resource.ComposeTestCheckFunc(
					testUpdateStream(cl, key, stream.Title),
				),
			},
		},
	})
}
