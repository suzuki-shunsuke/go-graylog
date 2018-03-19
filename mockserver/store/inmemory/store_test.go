package inmemory_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/inmemory"
)

func TestNewStore(t *testing.T) {
	store := inmemory.NewStore("")
	if store == nil {
		t.Fatal("store is nil")
	}
}

func TestSave(t *testing.T) {
	store := inmemory.NewStore("")
	if err := store.Save(); err != nil {
		t.Fatal(err)
	}
	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	store = inmemory.NewStore(tmpfile.Name())
	if err := store.Save(); err != nil {
		t.Fatal(err)
	}
}

func TestLoad(t *testing.T) {
	store := inmemory.NewStore("")
	if err := store.Load(); err != nil {
		t.Fatal(err)
	}
	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	store = inmemory.NewStore(tmpfile.Name())
	if err := store.Save(); err != nil {
		t.Fatal(err)
	}
	if err := store.Load(); err != nil {
		t.Fatal(err)
	}
}
