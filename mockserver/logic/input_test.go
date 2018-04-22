package logic_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestAddInput(t *testing.T) {
	server, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	input := testutil.Input()
	if _, err := server.AddInput(input); err != nil {
		t.Fatal("Failed to AddInput", err)
	}
	if input.ID == "" {
		t.Fatal(`server.AddInput() == ""`)
	}

	input.ID = ""
	input.Type = ""
	if _, err := server.AddInput(input); err == nil {
		t.Fatal("input type is required")
	}
	if _, err := server.AddInput(nil); err == nil {
		t.Fatal("input is nil")
	}
}

func TestGetInputs(t *testing.T) {
	server, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	act, _, _, err := server.GetInputs()
	if err != nil {
		t.Fatal("Failed to GetInputs", err)
	}
	if act == nil {
		t.Fatal("server.GetInputs() returns nil")
	}
	if len(act) != 1 {
		t.Fatalf("len(act) == %d, wanted 1", len(act))
	}
}

func TestGetInput(t *testing.T) {
	server, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	input := testutil.Input()
	if _, err := server.AddInput(input); err != nil {
		t.Fatal(err)
	}
	act, _, err := server.GetInput(input.ID)
	if err != nil {
		t.Fatal("Failed to GetInput", err)
	}
	if input.Node != act.Node {
		t.Fatalf("Node == %s, wanted %s", act.Node, input.Node)
	}

	if _, _, err := server.GetInput(""); err == nil {
		t.Fatal("input id is required")
	}

	if _, _, err := server.GetInput("h"); err == nil {
		t.Fatal(`no input whose id is "h"`)
	}
}

func TestUpdateInput(t *testing.T) {
	server, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	input := testutil.Input()
	if _, err := server.AddInput(input); err != nil {
		t.Fatal(err)
	}
	input.Title += " updated"
	if _, _, err := server.UpdateInput(input.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	act, _, err := server.GetInput(input.ID)
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
	if _, _, err := server.UpdateInput(input.NewUpdateParams()); err == nil {
		t.Fatal("input id is required")
	}

	input.ID = "h"
	if _, _, err := server.UpdateInput(input.NewUpdateParams()); err == nil {
		t.Fatal(`no input whose id is "h"`)
	}

	input.Type = ""
	if _, _, err := server.UpdateInput(input.NewUpdateParams()); err == nil {
		t.Fatal("input type is required")
	}
	input.Type = act.Type
	input.Attributes = nil
	if _, _, err := server.UpdateInput(input.NewUpdateParams()); err == nil {
		t.Fatal("input attributes is required")
	}
	input.Attributes = act.Attributes
	input.Title = ""
	if _, _, err := server.UpdateInput(input.NewUpdateParams()); err == nil {
		t.Fatal("input title is required")
	}

	input.Title = act.Title
	switch input.Type {
	case graylog.InputTypeBeats:
		attrs, ok := input.Attributes.(*graylog.InputBeatsAttrs)
		if !ok {
			t.Fatal("input.Attributes's type assertion is failure")
		}
		attrs.BindAddress = ""
		input.Attributes = attrs
		if _, _, err := server.UpdateInput(input.NewUpdateParams()); err == nil {
			t.Fatal("input bind_address is required")
		}
	}
	if _, _, err := server.UpdateInput(nil); err == nil {
		t.Fatal("input is required")
	}
}

func TestDeleteInput(t *testing.T) {
	server, err := logic.NewLogic(nil)
	if err != nil {
		t.Fatal(err)
	}
	input := testutil.Input()
	if _, err = server.AddInput(input); err != nil {
		t.Fatal(err)
	}
	if _, err = server.DeleteInput(input.ID); err != nil {
		t.Fatal("Failed to DeleteInput", err)
	}

	if _, err := server.DeleteInput(""); err == nil {
		t.Fatal("input id is required")
	}

	if _, err := server.DeleteInput("h"); err == nil {
		t.Fatal(`no input whose id is "h"`)
	}
}
