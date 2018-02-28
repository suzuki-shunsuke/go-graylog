package graylog

import (
	"reflect"
	"testing"
)

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
	params := dummyInput()
	params.Id = ""
	id, err := client.CreateInput(params)
	if err != nil {
		t.Fatal("Failed to CreateInput", err)
	}
	if id == "" {
		t.Fatal(`client.CreateInput() == ""`)
	}
}

func TestGetInputs(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := dummyInput()
	exp := []Input{*input}
	server.Inputs[input.Id] = *input
	act, err := client.GetInputs()
	if err != nil {
		t.Fatal("Failed to GetInputs", err)
	}
	if !reflect.DeepEqual(act, exp) {
		t.Fatalf("client.GetInputs() == %v, wanted %v", act, exp)
	}
}

func TestGetInput(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	exp := dummyInput()
	server.Inputs[exp.Id] = *exp
	act, err := client.GetInput(exp.Id)
	if err != nil {
		t.Fatal("Failed to GetInput", err)
	}
	if !reflect.DeepEqual(*exp, *act) {
		t.Fatalf("client.GetInput() == %v, wanted %v", act, exp)
	}
}

func TestUpdateInput(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	exp := dummyInput()
	server.Inputs[exp.Id] = *exp
	exp.Global = true
	act, err := client.UpdateInput(exp.Id, exp)
	if err != nil {
		t.Fatal("Failed to UpdateInput", err)
	}
	if !reflect.DeepEqual(*act, *exp) {
		t.Fatalf("client.UpdateInput() == %v, wanted %v", act, exp)
	}
}

func TestDeleteInput(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := dummyInput()
	server.Inputs[input.Id] = *input
	err = client.DeleteInput(input.Id)
	if err != nil {
		t.Fatal("Failed to DeleteInput", err)
	}
}
