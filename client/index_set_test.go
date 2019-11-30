package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v8"
	"github.com/suzuki-shunsuke/go-graylog/v8/client"
	"github.com/suzuki-shunsuke/go-graylog/v8/testdata"
	"github.com/suzuki-shunsuke/go-graylog/v8/testutil"
	"github.com/suzuki-shunsuke/go-ptr"
)

func TestClient_GetIndexSets(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/index_sets.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "GET",
								Path:   "/api/system/indices/index_sets",
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
								Query: url.Values{
									"skip":  []string{"0"},
									"limit": []string{"0"},
									"stats": []string{"false"},
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: bodyStr,
							},
						},
					},
				},
			},
		},
	})

	is, _, total, _, err := cl.GetIndexSets(ctx, 0, 0, false)
	require.Nil(t, err)
	require.Equal(t, testdata.IndexSets.IndexSets, is)
	require.Equal(t, testdata.IndexSets.Total, total)
}

func TestClient_GetIndexSet(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/index_set.json")
	require.Nil(t, err)
	bodyStr := string(buf)

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method: "GET",
								Path:   "/api/system/indices/index_sets/" + testdata.IndexSet.ID,
								PartOfHeader: http.Header{
									"Content-Type":   []string{"application/json"},
									"X-Requested-By": []string{"go-graylog"},
									"Authorization":  nil,
								},
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: bodyStr,
							},
						},
					},
				},
			},
		},
	})

	_, _, err = cl.GetIndexSet(ctx, "")
	require.NotNil(t, err)

	is, _, err := cl.GetIndexSet(ctx, testdata.IndexSet.ID)
	require.Nil(t, err)
	require.Equal(t, testdata.IndexSet, is)
}

func TestClient_CreateIndexSet(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	// nil check
	if _, err := client.CreateIndexSet(ctx, nil); err == nil {
		t.Fatal("index set is nil")
	}
	u, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	is := testutil.IndexSet(u.String())
	// success
	if _, err := client.CreateIndexSet(ctx, is); err != nil {
		t.Fatal(err)
	}
	testutil.WaitAfterCreateIndexSet(server)
	// clean
	defer func() {
		if _, err := client.DeleteIndexSet(ctx, is.ID); err != nil {
			t.Fatal(err)
		}
		testutil.WaitAfterDeleteIndexSet(server)
	}()
}

func TestClient_UpdateIndexSet(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	u, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	is, f, err := testutil.GetIndexSet(ctx, client, server, u.String())
	if err != nil {
		t.Fatal(err)
	}
	if f != nil {
		defer f(is.ID)
	}
	// success
	if _, _, err := client.UpdateIndexSet(ctx, is.NewUpdateParams()); err != nil {
		t.Fatal(err)
	}
	// id required
	prms := is.NewUpdateParams()
	prms.ID = ""
	if _, _, err := client.UpdateIndexSet(ctx, prms); err == nil {
		t.Fatal("index set id is required")
	}
	// nil check
	if _, _, err := client.UpdateIndexSet(ctx, nil); err == nil {
		t.Fatal("index set is required")
	}
	// invalid id
	prms.ID = "h"
	if _, _, err := client.UpdateIndexSet(ctx, prms); err == nil {
		t.Fatal("index set should no be found")
	}
}

func TestClient_DeleteIndexSet(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}

	// id required
	if _, err := client.DeleteIndexSet(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	// invalid id
	if _, err := client.DeleteIndexSet(ctx, "h"); err == nil {
		t.Fatal(`no index set with id "h" is found`)
	}
}

func TestClient_SetDefaultIndexSet(t *testing.T) {
	ctx := context.Background()
	server, client, err := testutil.GetServerAndClient()
	if err != nil {
		t.Fatal(err)
	}
	if server != nil {
		defer server.Close()
	}
	iss, _, _, _, err := client.GetIndexSets(ctx, 0, 0, false)
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
		u, err := uuid.NewV4()
		if err != nil {
			t.Fatal(err)
		}
		is = testutil.IndexSet(u.String())
		is.Default = false
		is.Writable = true
		if _, err := client.CreateIndexSet(ctx, is); err != nil {
			t.Fatal(err)
		}
		testutil.WaitAfterCreateIndexSet(server)
		defer func(id string) {
			if _, err := client.DeleteIndexSet(ctx, id); err != nil {
				t.Fatal(err)
			}
			testutil.WaitAfterDeleteIndexSet(server)
		}(is.ID)
	}
	is, _, err = client.SetDefaultIndexSet(ctx, is.ID)
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	defer func(id string) {
		if _, _, err = client.SetDefaultIndexSet(ctx, id); err != nil {
			t.Fatal(err)
		}
	}(defIs.ID)
	if !is.Default {
		t.Fatal("updatedIndexSet.Default == false")
	}
	if _, _, err := client.SetDefaultIndexSet(ctx, ""); err == nil {
		t.Fatal("index set id is required")
	}
	if _, _, err := client.SetDefaultIndexSet(ctx, "h"); err == nil {
		t.Fatal(`no index set whose id is "h"`)
	}

	prms := is.NewUpdateParams()
	prms.Writable = ptr.PBool(false)

	if _, _, err := client.UpdateIndexSet(ctx, prms); err == nil {
		t.Fatal("Default index set must be writable.")
	}
}
