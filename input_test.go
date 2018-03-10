package graylog

import (
	"testing"
)

func dummyNewInput() *Input {
	return &Input{
		Title: "test",
		Type:  "org.graylog2.inputs.gelf.tcp.GELFTCPInput",
		Node:  "2ad6b340-3e5f-4a96-ae81-040cfb8b6024",
		Configuration: &InputConfiguration{
			BindAddress:    "0.0.0.0",
			Port:           514,
			RecvBufferSize: 262144,
		}}
}

func dummyInput() *Input {
	return &Input{
		Id:    "5a90cee5c006c60001efbbf5",
		Title: "test",
		Type:  "org.graylog2.inputs.gelf.tcp.GELFTCPInput",
		Node:  "2ad6b340-3e5f-4a96-ae81-040cfb8b6024",
		Configuration: &InputConfiguration{
			BindAddress:    "0.0.0.0",
			Port:           514,
			RecvBufferSize: 262144,
		}}
}

func TestCreateInput(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	params := dummyNewInput()
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
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := dummyNewInput()
	if _, _, err := server.AddInput(input); err != nil {
		t.Fatal(err)
	}
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
	if act[0].Node != input.Node {
		t.Fatalf("Node == %s, wanted %s", act[0].Node, input.Node)
	}
}

func TestGetInput(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := dummyNewInput()
	exp, _, err := server.AddInput(input)
	if err != nil {
		t.Fatal(err)
	}
	act, _, err := client.GetInput(exp.Id)
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
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := dummyNewInput()
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
	act2, ok, err := server.GetInput(exp.Id)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("input is not found")
	}
	if act2.Title != exp.Title {
		t.Fatalf(`UpdateInput title "%s" != "%s"`, act2.Title, exp.Title)
	}

	exp.Id = ""
	if _, _, err := client.UpdateInput(exp); err == nil {
		t.Fatal("input id is required")
	}

	exp.Id = "h"
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
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := dummyNewInput()
	input, _, err = server.AddInput(input)
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.DeleteInput(input.Id)
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
