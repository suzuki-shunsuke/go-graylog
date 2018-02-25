package graylog

import (
	"reflect"
	"testing"
)

func dummyInput() *Input {
	return &Input{
		Id:            "5a90cee5c006c60001efbbf5",
		Title:         "test",
		Type:          "org.graylog2.inputs.gelf.tcp.GELFTCPInput",
		Node:          "2ad6b340-3e5f-4a96-ae81-040cfb8b6024",
		Configuration: &InputConfiguration{}}
}

func TestCreateInput(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	params := dummyInput()
	params.Id = ""
	id, err := client.CreateInput(params)
	if err != nil {
		t.Error("Failed to CreateInput", err)
		return
	}
	if id == "" {
		t.Error(`client.CreateInput() == ""`)
		return
	}
}

func TestGetInputs(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	input := dummyInput()
	exp := []Input{*input}
	server.Inputs[input.Id] = *input
	act, err := client.GetInputs()
	if err != nil {
		t.Error("Failed to GetInputs", err)
		return
	}
	if !reflect.DeepEqual(act, exp) {
		t.Errorf("client.GetInputs() == %v, wanted %v", act, exp)
	}
}

func TestGetInput(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	exp := dummyInput()
	server.Inputs[exp.Id] = *exp
	act, err := client.GetInput(exp.Id)
	if err != nil {
		t.Error("Failed to GetInput", err)
		return
	}
	if !reflect.DeepEqual(*exp, *act) {
		t.Errorf("client.GetInput() == %v, wanted %v", act, exp)
	}
}

func TestUpdateInput(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	exp := dummyInput()
	server.Inputs[exp.Id] = *exp
	exp.Global = true
	act, err := client.UpdateInput(exp.Id, exp)
	if err != nil {
		t.Error("Failed to UpdateInput", err)
		return
	}
	if !reflect.DeepEqual(*act, *exp) {
		t.Errorf("client.UpdateInput() == %v, wanted %v", act, exp)
	}
}

func TestDeleteInput(t *testing.T) {
	server, client, err := getServerAndClient()
	if err != nil {
		t.Error(err)
		return
	}
	defer server.Server.Close()
	input := dummyInput()
	server.Inputs[input.Id] = *input
	err = client.DeleteInput(input.Id)
	if err != nil {
		t.Error("Failed to DeleteInput", err)
		return
	}
}
