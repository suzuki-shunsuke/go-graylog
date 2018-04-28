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
	id := input.ID
	t.Run("normal", func(t *testing.T) {
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
	})
	t.Run("id is required", func(t *testing.T) {
		input := testutil.Input()
		if _, _, err := server.UpdateInput(input.NewUpdateParams()); err == nil {
			t.Fatal("input id is required")
		}
	})
	t.Run("invalid id", func(t *testing.T) {
		input := testutil.Input()
		input.ID = "h"
		if _, _, err := server.UpdateInput(input.NewUpdateParams()); err == nil {
			t.Fatal(`no input whose id is "h"`)
		}
	})
	t.Run("type is required", func(t *testing.T) {
		input := testutil.Input()
		input.ID = id
		input.Type = ""
		if _, _, err := server.UpdateInput(input.NewUpdateParams()); err == nil {
			t.Fatal("input type is required")
		}
	})
	t.Run("attributes is required", func(t *testing.T) {
		input := testutil.Input()
		input.ID = id
		input.Attributes = nil
		if _, _, err := server.UpdateInput(input.NewUpdateParams()); err == nil {
			t.Fatal("input attributes is required")
		}
	})
	t.Run("type is required", func(t *testing.T) {
		input := testutil.Input()
		input.ID = id
		input.Title = ""
		if _, _, err := server.UpdateInput(input.NewUpdateParams()); err == nil {
			t.Fatal("input title is required")
		}
	})
	t.Run("beats's bind_address is required", func(t *testing.T) {
		input := testutil.Input()
		input.ID = id
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
	})
	t.Run("input is required", func(t *testing.T) {
		if _, _, err := server.UpdateInput(nil); err == nil {
			t.Fatal("input is required")
		}
	})
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
