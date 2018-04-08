package client_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestCreateInput(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	// nil check
	if _, err := client.CreateInput(nil); err == nil {
		t.Fatal("input is nil")
	}
	input := testutil.Input()
	if _, err := client.CreateInput(input); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteInput(input.ID)
	// error check
	if _, err := client.CreateInput(input); err == nil {
		t.Fatal("input id should be empty")
	}
}

func TestGetInputs(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	inputs, _, _, err := client.GetInputs()
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
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	input := testutil.Input()
	if _, err := client.CreateInput(input); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteInput(input.ID)
	r, _, err := client.GetInput(input.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("input is nil")
	}
	if r.ID != input.ID {
		t.Fatalf(`input.ID = "%s", wanted "%s"`, r.ID, input.ID)
	}
	if _, _, err := client.GetInput(""); err == nil {
		t.Fatal("input id is required")
	}
	if _, _, err := client.GetInput("h"); err == nil {
		t.Fatal("input should not be found")
	}
}

func TestUpdateInput(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	input := testutil.Input()
	if _, err := client.CreateInput(input); err != nil {
		t.Fatal(err)
	}
	defer client.DeleteInput(input.ID)
	if _, err := client.UpdateInput(input); err != nil {
		t.Fatal(err)
	}
	input.ID = ""
	if _, err := client.UpdateInput(input); err == nil {
		t.Fatal("input id is required")
	}
	if _, err := client.UpdateInput(nil); err == nil {
		t.Fatal("input is required")
	}
	input.ID = "h"
	if _, err := client.UpdateInput(nil); err == nil {
		t.Fatal("input should no be found")
	}
}

func TestDeleteInput(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	if _, err := client.DeleteInput(""); err == nil {
		t.Fatal("input id is required")
	}
	if _, err := client.DeleteInput("h"); err == nil {
		t.Fatal(`no input with id "h" is found`)
	}
}
