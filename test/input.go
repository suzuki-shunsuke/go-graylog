package test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestCreateInput(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	params := testutil.DummyNewInput()
	id, _, err := client.CreateInput(params)
	if err != nil {
		t.Fatal("Failed to CreateInput", err)
	}
	if id == "" {
		t.Fatal(`client.CreateInput() == ""`)
	}

	params.Type = ""
	if _, _, err := client.CreateInput(params); err == nil {
		t.Fatal("input type is required")
	}
}

func TestGetInputs(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	act, _, err := client.GetInputs()
	if err != nil {
		t.Fatal("Failed to GetInputs", err)
	}
	if act == nil {
		t.Fatal("client.GetInputs() returns nil")
	}
	if len(act) != 1 {
		t.Fatalf("len(act) == %d, wanted 1", len(act))
	}
}

func TestGetInput(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := testutil.DummyNewInput()
	if _, err := server.AddInput(input); err != nil {
		t.Fatal(err)
	}
	act, _, err := client.GetInput(input.ID)
	if err != nil {
		t.Fatal("Failed to GetInput", err)
	}
	if input.Node != act.Node {
		t.Fatalf("Node == %s, wanted %s", act.Node, input.Node)
	}

	if _, _, err := client.GetInput(""); err == nil {
		t.Fatal("input id is required")
	}

	if _, _, err := client.GetInput("h"); err == nil {
		t.Fatal(`no input whose id is "h"`)
	}
}

func TestUpdateInput(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := testutil.DummyNewInput()
	if _, err := server.AddInput(input); err != nil {
		t.Fatal(err)
	}
	input.Title += " updated"
	act, _, err := client.UpdateInput(input)
	if err != nil {
		t.Fatal("Failed to UpdateInput", err)
	}
	if act.Title != input.Title {
		t.Fatalf(`UpdateInput title "%s" != "%s"`, act.Title, input.Title)
	}
	act2, err := server.GetInput(input.ID)
	if err != nil {
		t.Fatal(err)
	}
	if act2 == nil {
		t.Fatal("input is not found")
	}
	if act2.Title != input.Title {
		t.Fatalf(`UpdateInput title "%s" != "%s"`, act2.Title, input.Title)
	}

	input.ID = ""
	if _, _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input id is required")
	}

	input.ID = "h"
	if _, _, err := client.UpdateInput(input); err == nil {
		t.Fatal(`no input whose id is "h"`)
	}

	input.Type = ""
	if _, _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input type is required")
	}
	input.Type = act.Type
	input.Configuration = nil
	if _, _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input configuration is required")
	}
	input.Configuration = act.Configuration
	input.Title = ""
	if _, _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input title is required")
	}

	input.Title = act.Title
	input.Configuration.BindAddress = ""
	if _, _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input bind_address is required")
	}
	input.Configuration.BindAddress = "0.0.0.0"
	input.Configuration.Port = 0
	if _, _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input port is required")
	}
	input.Configuration.Port = 514
	input.Configuration.RecvBufferSize = 0
	if _, _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input recv_buffer_size is required")
	}
}

func TestDeleteInput(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := testutil.DummyNewInput()
	if _, err = server.AddInput(input); err != nil {
		t.Fatal(err)
	}
	if _, err = client.DeleteInput(input.ID); err != nil {
		t.Fatal("Failed to DeleteInput", err)
	}

	if _, err := client.DeleteInput(""); err == nil {
		t.Fatal("input id is required")
	}

	if _, err := client.DeleteInput("h"); err == nil {
		t.Fatal(`no input whose id is "h"`)
	}
}
