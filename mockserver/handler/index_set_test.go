package handler_test

import (
	"reflect"
	"testing"

	"github.com/suzuki-shunsuke/go-ptr"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/testutil"
)

func TestHandleGetIndexSets(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	indexSets, _, _, _, err := client.GetIndexSets(0, 0, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(indexSets) == 0 {
		t.Fatal("len(indexSets) == 0")
	}
	// TODO run by nobody
}

func TestHandleGetIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()

	is := testutil.IndexSet("hoge")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	act, _, err := client.GetIndexSet(is.ID)
	if err != nil {
		t.Fatal("Failed to GetIndexSet", err)
	}
	if !reflect.DeepEqual(*act, *is) {
		t.Fatalf("client.GetIndexSet() == %v, wanted %v", act, is)
	}
	if _, _, err := client.GetIndexSet(""); err == nil {
		t.Fatal("index set id is required")
	}
	if _, _, err := client.GetIndexSet("h"); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}
}

func TestHandleCreateIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := testutil.IndexSet("hoge")
	if _, err := client.CreateIndexSet(is); err != nil {
		t.Fatal(err)
	}
	if is.ID == "" {
		t.Fatal("IndexSet's id is empty")
	}
	is.ID = ""
	if _, err := client.CreateIndexSet(is); err == nil {
		t.Fatal("index prefix should conflict")
	}

	pc := &plainClient{Name: client.Name(), Password: client.Password()}
	resp, err := pc.Post(client.Endpoints().IndexSets(), "hoge")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode = %d, wanted 400", resp.StatusCode)
	}
}

func TestHandleUpdateIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	is := testutil.IndexSet("hoge")
	if _, err = server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	id := is.ID
	prms := is.NewUpdateParams()
	if _, _, err := client.UpdateIndexSet(prms); err != nil {
		t.Fatal(err)
	}
	prms.ID = ""
	if _, _, err := client.UpdateIndexSet(prms); err == nil {
		t.Fatal("index set id is required")
	}
	prms.ID = "h"
	if _, _, err := client.UpdateIndexSet(prms); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}
	prms.ID = id
	title := prms.Title
	prms.Title = ""
	if _, _, err := client.UpdateIndexSet(prms); err == nil {
		t.Fatal("title is required")
	}
	prms.Title = title
	prms.IndexPrefix = "graylog"
	if _, _, err := server.UpdateIndexSet(prms); err == nil {
		t.Fatal("index prefix should be conflict")
	}
	if _, _, err := client.UpdateIndexSet(nil); err == nil {
		t.Fatal("index set is nil")
	}

	pc := &plainClient{Name: client.Name(), Password: client.Password()}
	u, err := client.Endpoints().IndexSet(id)
	if err != nil {
		t.Fatal(err)
	}
	resp, _ := pc.Put(u.String(), "hoge")
	if resp.StatusCode != 400 {
		t.Fatalf("resp.StatusCode = %d, wanted 400", resp.StatusCode)
	}
}

func TestHandleDeleteIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	defer server.Close()
	indexSets, _, _, err := server.GetIndexSets(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	indexSet := indexSets[0]
	if _, err = client.DeleteIndexSet(indexSet.ID); err == nil {
		t.Fatal("default index set should not be deleted")
	}
	is := testutil.IndexSet("hoge")
	if _, err := server.AddIndexSet(is); err != nil {
		t.Fatal(err)
	}
	if _, err = client.DeleteIndexSet(is.ID); err != nil {
		t.Fatal("Failed to DeleteIndexSet", err)
	}
	if _, err = client.DeleteIndexSet(""); err == nil {
		t.Fatal("index set id is required")
	}
	if _, err = client.DeleteIndexSet("h"); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}
}

func TestHandleSetDefaultIndexSet(t *testing.T) {
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	iss, _, _, _, err := client.GetIndexSets(0, 0, false)
	if err != nil {
		t.Fatal(err)
	}
	var defIs, is *graylog.IndexSet
	for _, i := range iss {
		if i.Default {
			defIs = &i
		} else {
			is = &i
		}
	}
	if is == nil {
		is = testutil.IndexSet("hoge")
		is.Default = false
		is.Writable = true
		if _, err := client.CreateIndexSet(is); err != nil {
			t.Fatal(err)
		}
		testutil.WaitAfterCreateIndexSet(server)
		defer func(id string) {
			if _, err := client.DeleteIndexSet(id); err != nil {
				t.Fatal(err)
			}
			testutil.WaitAfterDeleteIndexSet(server)
		}(is.ID)
	}
	is, _, err = client.SetDefaultIndexSet(is.ID)
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	defer func(id string) {
		if _, _, err = client.SetDefaultIndexSet(id); err != nil {
			t.Fatal(err)
		}
	}(defIs.ID)
	if !is.Default {
		t.Fatal("updatedIndexSet.Default == false")
	}
	if _, _, err := client.SetDefaultIndexSet(""); err == nil {
		t.Fatal("index set id is required")
	}
	if _, _, err := client.SetDefaultIndexSet("h"); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}

	prms := is.NewUpdateParams()
	prms.Writable = ptr.PBool(false)
	if _, _, err := client.UpdateIndexSet(prms); err == nil {
		t.Fatal("Default index set must be writable.")
	}
}
