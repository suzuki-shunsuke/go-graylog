package handler_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHandleGetInput(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := testutil.Input()
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

func TestHandleGetInputs(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	act, _, _, err := client.GetInputs()
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

func TestHandleCreateInput(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := testutil.Input()
	if _, err := client.CreateInput(input); err != nil {
		t.Fatal("Failed to CreateInput", err)
	}
	if input.ID == "" {
		t.Fatal(`client.CreateInput() == ""`)
	}

	input.ID = ""
	input.Type = ""
	if _, err := client.CreateInput(input); err == nil {
		t.Fatal("input type is required")
	}
	if _, err := client.CreateInput(nil); err == nil {
		t.Fatal("input is nil")
	}

	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPost, client.Endpoints.Inputs, body)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth(client.Name(), client.Password())
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode == %d, wanted 400", resp.StatusCode)
	}
}

func TestHandleUpdateInput(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := testutil.Input()
	if _, err := server.AddInput(input); err != nil {
		t.Fatal(err)
	}
	id := input.ID
	input.Title += " updated"
	if _, err := client.UpdateInput(input); err != nil {
		t.Fatal("Failed to UpdateInput", err)
	}
	act, _, err := server.GetInput(id)
	if err != nil {
		t.Fatal(err)
	}
	if act == nil {
		t.Fatal("input is not found")
	}
	if act.Title != input.Title {
		t.Fatalf(`UpdateInput title "%s" != "%s"`, act.Title, input.Title)
	}

	input.ID = ""
	if _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input id is required")
	}

	input.ID = "h"
	if _, err := client.UpdateInput(input); err == nil {
		t.Fatal(`no input whose id is "h"`)
	}

	input.ID = id
	input.Type = ""
	if _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input type is required")
	}
	input.Type = act.Type
	input.Attributes = nil
	if _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input attributes is required")
	}
	input.Attributes = act.Attributes
	input.Title = ""
	if _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input title is required")
	}

	input.Title = act.Title
	input.Attributes.BindAddress = nil
	if _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input bind_address is required")
	}
	input.Attributes.BindAddress = act.Attributes.BindAddress
	input.Attributes.Port = nil
	if _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input port is required")
	}
	input.Attributes.Port = act.Attributes.Port
	input.Attributes.RecvBufferSize = nil
	if _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input recv_buffer_size is required")
	}

	if _, err := client.UpdateInput(nil); err == nil {
		t.Fatal("input is required")
	}

	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPut, client.Endpoints.Input(id), body)
	if err != nil {
		t.Fatal(err)
	}
	req.SetBasicAuth(client.Name(), client.Password())
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode == %d, wanted 400", resp.StatusCode)
	}
}

func TestHandleDeleteInput(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	input := testutil.Input()
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
