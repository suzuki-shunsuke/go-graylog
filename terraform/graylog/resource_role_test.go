package graylog

import (
	"testing"
)

// func testDeleteRole(
// 	ctx context.Context, cl *client.Client, name string,
// ) resource.TestCheckFunc {
// 	return func(tfState *terraform.State) error {
// 		if _, _, err := cl.GetRole(ctx, name); err == nil {
// 			return fmt.Errorf(`role "%s" must be deleted`, name)
// 		}
// 		return nil
// 	}
// }
//
// func testCreateRole(
// 	ctx context.Context, cl *client.Client, name string,
// ) resource.TestCheckFunc {
// 	return func(tfState *terraform.State) error {
// 		_, _, err := cl.GetRole(ctx, name)
// 		return err
// 	}
// }
//
// func testUpdateRole(
// 	ctx context.Context, cl *client.Client, name, description string,
// ) resource.TestCheckFunc {
// 	return func(tfState *terraform.State) error {
// 		role, _, err := cl.GetRole(ctx, name)
// 		if err != nil {
// 			return err
// 		}
// 		if role.Description != description {
// 			return errors.New("role.Description is not updated")
// 		}
// 		return nil
// 	}
// }

func TestAccRole(t *testing.T) {
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
	//	roleTf := `
	//resource "graylog_role" "test-terraform" {
	//  name = "test terraform name"
	//  description = "test terraform description"
	//  permissions = ["*"]
	//}`
	//	description := "test terraform description updated"
	//	updateTf := fmt.Sprintf(`
	//resource "graylog_role" "test-terraform" {
	//  name = "test terraform name"
	//  description = "%s"
	//  permissions = ["*"]
	//}`, description)
	//	name := "test terraform name"
	//	resource.Test(t, resource.TestCase{
	//		Providers:    testAccProviders,
	//		CheckDestroy: testDeleteRole(ctx, cl, name),
	//		Steps: []resource.TestStep{
	//			{
	//				Config: roleTf,
	//				Check: resource.ComposeTestCheckFunc(
	//					testCreateRole(ctx, cl, name),
	//				),
	//			},
	//			{
	//				Config: updateTf,
	//				Check: resource.ComposeTestCheckFunc(
	//					testUpdateRole(ctx, cl, name, description),
	//				),
	//			},
	//		},
	//	})
}
