package graylog

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestMockServerLoad(t *testing.T) {
	server, _, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	if err := server.Load(); err != nil {
		t.Fatal(err)
	}
	server.dataPath = "hoge"
	if err := server.Load(); err != nil {
		t.Fatal(err)
	}
}

func TestMockServerSave(t *testing.T) {
	server, _, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	server.dataPath = tmpfile.Name()
	if err := server.Save(); err != nil {
		t.Fatal(err)
	}
	if err := server.Load(); err != nil {
		t.Fatal(err)
	}
}

func TestMockServerHandleNotFound(t *testing.T) {
	server, _, err := getServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	endpoint := fmt.Sprintf("%s/dummy", server.endpoint)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}
	hc := &http.Client{}
	if _, err := hc.Do(req); err != nil {
		t.Fatal(err)
	}
}
