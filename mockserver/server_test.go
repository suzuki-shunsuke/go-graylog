package mockserver_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/mockserver"
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
	server.SetStore(mockserver.NewInMemoryStore("hoge"))
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
	server.SetStore(mockserver.NewInMemoryStore(tmpfile.Name()))
	if err := server.Save(); err != nil {
		t.Fatal(err)
	}
	if err := server.Load(); err != nil {
		t.Fatal(err)
	}
}

func TestServerHandleNotFound(t *testing.T) {
	server, _, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	endpoint := fmt.Sprintf("%s/dummy", server.GetEndpoint())
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}
	hc := &http.Client{}
	if _, err := hc.Do(req); err != nil {
		t.Fatal(err)
	}
}
