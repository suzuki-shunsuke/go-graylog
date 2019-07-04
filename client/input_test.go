package client_test

import (
	"context"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestGetInputs(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	inputs, _, _, err := client.GetInputs(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if inputs == nil {
		t.Fatal("inputs is nil")
	}
	if len(inputs) == 0 {
		t.Fatal("inputs is empty")
	}
}

func TestGetInput(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	input := testutil.Input()
	if _, err := client.CreateInput(ctx, input); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteInput(ctx, input.ID)
	r, _, err := client.GetInput(ctx, input.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("input is nil")
	}
	if r.ID != input.ID {
		t.Fatalf(`input.ID = "%s", wanted "%s"`, r.ID, input.ID)
	}
	if _, _, err := client.GetInput(ctx, ""); err == nil {
		t.Fatal("input id is required")
	}
	if _, _, err := client.GetInput(ctx, "h"); err == nil {
		t.Fatal("input should not be found")
	}
}

func TestCreateInput(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	// nil check
	if _, err := client.CreateInput(ctx, nil); err == nil {
		t.Fatal("input is nil")
	}
	input := testutil.Input()
	if _, err := client.CreateInput(ctx, input); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteInput(ctx, input.ID)
	attrs := input.Attrs.(*graylog.InputBeatsAttrs)
	if attrs.BindAddress == "" {
		t.Fatal(`attrs.BindAddress == ""`)
	}
	// error check
	if _, err := client.CreateInput(ctx, input); err == nil {
		t.Fatal("input id should be empty")
	}
}

func TestUpdateInput(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	input := testutil.Input()
	if _, err := client.CreateInput(ctx, input); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteInput(ctx, input.ID)
	attrs := input.Attrs.(*graylog.InputBeatsAttrs)
	if attrs.BindAddress == "" {
		t.Fatal(`attrs.BindAddress == ""`)
	}
	if _, _, err := client.UpdateInput(ctx, input.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	input.ID = ""
	if _, _, err := client.UpdateInput(ctx, input.NewUpdateParams()); err == nil {
		t.Fatal("input id is required")
	}
	if _, _, err := client.UpdateInput(ctx, nil); err == nil {
		t.Fatal("input is required")
	}
	input.ID = "h"
	if _, _, err := client.UpdateInput(ctx, input.NewUpdateParams()); err == nil {
		t.Fatal("input should no be found")
	}
}

func TestDeleteInput(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, err := client.DeleteInput(ctx, ""); err == nil {
		t.Fatal("input id is required")
	}
	if _, err := client.DeleteInput(ctx, "h"); err == nil {
		t.Fatal(`no input with id "h" is found`)
	}
}
