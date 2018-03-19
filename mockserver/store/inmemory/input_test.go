package inmemory_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/inmemory"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHasInput(t *testing.T) {
	store := inmemory.NewStore("")
	ok, err := store.HasInput("foo")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("input foo should not exist")
	}
}

func TestGetInput(t *testing.T) {
	store := inmemory.NewStore("")
	input, err := store.GetInput("foo")
	if err != nil {
		t.Fatal(err)
	}
	if input != nil {
		t.Fatal("input foo should not exist")
	}
}

func TestGetInputs(t *testing.T) {
	store := inmemory.NewStore("")
	inputs, err := store.GetInputs()
	if err != nil {
		t.Fatal(err)
	}
	if inputs != nil && len(inputs) != 0 {
		t.Fatal("inputs should be nil or empty array")
	}
}

func TestAddInput(t *testing.T) {
	store := inmemory.NewStore("")
	input := testutil.Input()
	if err := store.AddInput(input); err != nil {
		t.Fatal(err)
	}
	r, err := store.GetInput(input.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("input is nil")
	}
}

func TestUpdateInput(t *testing.T) {
	store := inmemory.NewStore("")
	input := testutil.Input()
	if err := store.AddInput(input); err != nil {
		t.Fatal(err)
	}
	r, err := store.GetInput(input.ID)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("input is nil")
	}
	input.Title += " changed"
	if err := store.UpdateInput(input); err != nil {
		t.Fatal(err)
	}
	r, err = store.GetInput(input.ID)
	if err != nil {
		t.Fatal(err)
	}
	if input.Title != r.Title {
		t.Fatalf(`input.Title = "%s", wanted "%s"`, r.Title, input.Title)
	}
}

func TestDeleteInput(t *testing.T) {
	store := inmemory.NewStore("")
	if err := store.DeleteInput("foo"); err != nil {
		t.Fatal(err)
	}
	input := testutil.Input()
	if err := store.AddInput(input); err != nil {
		t.Fatal(err)
	}
	if err := store.DeleteInput(input.ID); err != nil {
		t.Fatal(err)
	}
	ok, err := store.HasInput(input.ID)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal("input should be deleted")
	}
}
