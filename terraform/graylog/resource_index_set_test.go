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

func testDeleteIndexSet(
	cl *client.Client, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIdFromTfState(tfState, key)
		if err != nil {
			return err
		}
		if _, _, err := cl.GetIndexSet(id); err == nil {
			return fmt.Errorf(`indexSet "%s" must be deleted`, id)
		}
		return nil
	}
}

func testCreateIndexSet(
	cl *client.Client, key string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIdFromTfState(tfState, key)
		if err != nil {
			return err
		}

		if _, _, err := cl.GetIndexSet(id); err != nil {
			return err
		}
		return nil
	}
}

func testUpdateIndexSet(
	cl *client.Client, key, title string,
) resource.TestCheckFunc {
	return func(tfState *terraform.State) error {
		id, err := getIdFromTfState(tfState, key)
		if err != nil {
			return err
		}
		indexSet, _, err := cl.GetIndexSet(id)
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

	is := &graylog.IndexSet{
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

	tfConf := &TFConf{
		Resource: map[string]map[string]interface{}{
			"graylog_index_set": {"test": is}},
	}

	b, err := json.Marshal(tfConf)
	if err != nil {
		t.Fatal(err)
	}

	is.Title = "terraform test index set title updated"
	tfConf.Resource["graylog_index_set"]["test"] = is

	ub, err := json.Marshal(tfConf)
	if err != nil {
		t.Fatal(err)
	}

	key := "graylog_index_set.test"
	if server != nil {
		server.Start()
		defer server.Close()
	}
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testDeleteIndexSet(cl, key),
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: string(b),
				Check: resource.ComposeTestCheckFunc(
					testCreateIndexSet(cl, key),
				),
			},
			resource.TestStep{
				Config: string(ub),
				Check: resource.ComposeTestCheckFunc(
					testUpdateIndexSet(cl, key, is.Title),
				),
			},
		},
	})
}
