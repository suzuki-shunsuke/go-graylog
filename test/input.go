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
	exp, _, err := server.AddInput(input)
	if err != nil {
		t.Fatal(err)
	}
	act, _, err := client.GetInput(exp.ID)
	if err != nil {
		t.Fatal("Failed to GetInput", err)
	}
	if exp.Node != act.Node {
		t.Fatalf("Node == %s, wanted %s", act.Node, exp.Node)
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
	exp, _, err := server.AddInput(input)
	if err != nil {
		t.Fatal(err)
	}
	exp.Title += " updated"
	act, _, err := client.UpdateInput(exp)
	if err != nil {
		t.Fatal("Failed to UpdateInput", err)
	}
	if act.Title != exp.Title {
		t.Fatalf(`UpdateInput title "%s" != "%s"`, act.Title, exp.Title)
	}
	act2, err := server.GetInput(exp.ID)
	if err != nil {
		t.Fatal(err)
	}
	if act2 == nil {
		t.Fatal("input is not found")
	}
	if act2.Title != exp.Title {
		t.Fatalf(`UpdateInput title "%s" != "%s"`, act2.Title, exp.Title)
	}

	exp.ID = ""
	if _, _, err := client.UpdateInput(exp); err == nil {
		t.Fatal("input id is required")
	}

	exp.ID = "h"
	if _, _, err := client.UpdateInput(exp); err == nil {
		t.Fatal(`no input whose id is "h"`)
	}

	exp.Type = ""
	if _, _, err := client.UpdateInput(exp); err == nil {
		t.Fatal("input type is required")
	}
	exp.Type = act.Type
	exp.Configuration = nil
	if _, _, err := client.UpdateInput(exp); err == nil {
		t.Fatal("input configuration is required")
	}
	exp.Configuration = act.Configuration
	exp.Title = ""
	if _, _, err := client.UpdateInput(exp); err == nil {
		t.Fatal("input title is required")
	}

	exp.Title = act.Title
	exp.Configuration.BindAddress = ""
	if _, _, err := client.UpdateInput(exp); err == nil {
		t.Fatal("input bind_address is required")
	}
	exp.Configuration.BindAddress = "0.0.0.0"
	exp.Configuration.Port = 0
	if _, _, err := client.UpdateInput(exp); err == nil {
		t.Fatal("input port is required")
	}
	exp.Configuration.Port = 514
	exp.Configuration.RecvBufferSize = 0
	if _, _, err := client.UpdateInput(exp); err == nil {
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
	input, _, err = server.AddInput(input)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.DeleteInput(input.ID)
	if err != nil {
		t.Fatal("Failed to DeleteInput", err)
	}

	if _, err := client.DeleteInput(""); err == nil {
		t.Fatal("input id is required")
	}

	if _, err := client.DeleteInput("h"); err == nil {
		t.Fatal(`no input whose id is "h"`)
	}
}
