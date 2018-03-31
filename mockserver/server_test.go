package mockserver_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver/store/inmemory"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestServerLoad(t *testing.T) {
	server, _, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	if err := server.Load(); err != nil {
		t.Fatal(err)
	}
	server.SetStore(inmemory.NewStore("hoge"))
	if err := server.Load(); err != nil {
		t.Fatal(err)
	}
}

func TestServerSave(t *testing.T) {
	server, _, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	server.SetStore(inmemory.NewStore(tmpfile.Name()))
	if err := server.Save(); err != nil {
		t.Fatal(err)
	}
	if err := server.Load(); err != nil {
		t.Fatal(err)
	}
}
