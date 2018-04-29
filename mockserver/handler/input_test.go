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
		t.Fatal(err)
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
		t.Fatal(err)
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
		t.Fatal(err)
	}
	if input.ID == "" {
		t.Fatal(`client.CreateInput() == ""`)
	}
	if _, err := client.CreateInput(nil); err == nil {
		t.Fatal("input is nil")
	}

	body := bytes.NewBuffer([]byte("hoge"))
	req, err := http.NewRequest(
		http.MethodPost, client.Endpoints().Inputs(), body)
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
	if _, _, err := client.UpdateInput(input.NewUpdateParams()); err != nil {
		t.Fatal(err)
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
	if _, _, err := client.UpdateInput(input.NewUpdateParams()); err == nil {
		t.Fatal("input id is required")
	}

	input.ID = "h"
	if _, _, err := client.UpdateInput(input.NewUpdateParams()); err == nil {
		t.Fatal(`no input whose id is "h"`)
	}

	input.ID = id
	input.Attrs = nil
	if _, _, err := client.UpdateInput(input.NewUpdateParams()); err == nil {
		t.Fatal("input attributes is required")
	}
	input.Attrs = act.Attrs
	input.Title = ""
	if _, _, err := client.UpdateInput(input.NewUpdateParams()); err == nil {
		t.Fatal("input title is required")
	}

	input.Title = act.Title

	if _, _, err := client.UpdateInput(nil); err == nil {
		t.Fatal("input is required")
	}

	body := bytes.NewBuffer([]byte("hoge"))
	u, err := client.Endpoints().Input(id)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(
		http.MethodPut, u.String(), body)
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
