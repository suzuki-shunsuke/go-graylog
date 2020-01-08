package graylog

import (
	"testing"
)

// func testDeleteStream(
// 	ctx context.Context, cl *client.Client, key string,
// ) resource.TestCheckFunc {
// 	return func(tfState *terraform.State) error {
// 		id, err := getIDFromTfState(tfState, key)
// 		if err != nil {
// 			return err
// 		}
// 		if _, _, err := cl.GetStream(ctx, id); err == nil {
// 			return fmt.Errorf(`stream "%s" must be deleted`, id)
// 		}
// 		return nil
// 	}
// }

// func testCreateStream(
// 	ctx context.Context, cl *client.Client, key string,
// ) resource.TestCheckFunc {
// 	return func(tfState *terraform.State) error {
// 		id, err := getIDFromTfState(tfState, key)
// 		if err != nil {
// 			return err
// 		}
//
// 		_, _, err = cl.GetStream(ctx, id)
// 		return err
// 	}
// }

// func testUpdateStream(
// 	ctx context.Context, cl *client.Client, key, title string,
// ) resource.TestCheckFunc {
// 	return func(tfState *terraform.State) error {
// 		id, err := getIDFromTfState(tfState, key)
// 		if err != nil {
// 			return err
// 		}
// 		stream, _, err := cl.GetStream(ctx, id)
// 		if err != nil {
// 			return err
// 		}
// 		if stream.Title != title {
// 			return fmt.Errorf("stream.Title == %s, wanted %s", stream.Title, title)
// 		}
// 		return nil
// 	}
// }

func TestAccStream(t *testing.T) {
	//	ctx := context.Background()
	//	cl, err := setEnv()
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//
	//	testAccProvider := Provider()
	//	testAccProviders := map[string]terraform.ResourceProvider{
	//		"graylog": testAccProvider,
	//	}
	//
	//	u, err := uuid.NewV4()
	//	if err != nil {
	//		t.Fatal(err)
	//	}
	//	prefix := u.String()
	//	roleTf := `
	//resource "graylog_index_set" "test" {
	//  title = "terraform test index set"
	//	description = "terraform test index set description"
	//	index_prefix = "%s"
	//	shards = 4
	//	replicas = 0
	//  rotation_strategy_class = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategy"
	//  rotation_strategy {
	//    type = "org.graylog2.indexer.rotation.strategies.MessageCountRotationStrategyConfig"
	//  }
	//  retention_strategy_class = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategy"
	//  retention_strategy {
	//    type = "org.graylog2.indexer.retention.strategies.DeletionRetentionStrategyConfig"
	//  }
	//  index_analyzer = "standard"
	//	writable = true
	//  index_optimization_max_num_segments = 1
	//}
	//
	//resource "graylog_stream" "test" {
	//  title = "%s"
	//	index_set_id = "${graylog_index_set.test.id}"
	//	matching_type = "AND"
	//}`
	//	createTitle := "terraform stream test"
	//	updateTitle := "terraform stream test updated"
	//
	//	key := "graylog_stream.test"
	//	resource.Test(t, resource.TestCase{
	//		Providers:    testAccProviders,
	//		CheckDestroy: testDeleteStream(ctx, cl, key),
	//		Steps: []resource.TestStep{
	//			{
	//				Config: fmt.Sprintf(roleTf, prefix, createTitle),
	//				Check: resource.ComposeTestCheckFunc(
	//					testCreateStream(ctx, cl, key),
	//				),
	//			},
	//			{
	//				Config: fmt.Sprintf(roleTf, prefix, updateTitle),
	//				Check: resource.ComposeTestCheckFunc(
	//					testUpdateStream(ctx, cl, key, updateTitle),
	//				),
	//			},
	//		},
	//	})
}
