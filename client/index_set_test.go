package client_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v10/client"
	"github.com/suzuki-shunsuke/go-graylog/v10/testdata"
)

func TestClient_GetIndexSets(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/index_set/index_sets.json")
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
								Method:       "GET",
								Path:         "/api/system/indices/index_sets",
								PartOfHeader: getTestHeader(),
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
	require.Equal(t, testdata.IndexSets().IndexSets, is)
	require.Equal(t, testdata.IndexSets().Total, total)
}

func TestClient_GetIndexSet(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	buf, err := ioutil.ReadFile("../testdata/index_set/index_set.json")
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
								Method:       "GET",
								Path:         "/api/system/indices/index_sets/" + testdata.IndexSet().ID,
								PartOfHeader: getTestHeader(),
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

	is, _, err := cl.GetIndexSet(ctx, testdata.IndexSet().ID)
	require.Nil(t, err)
	require.Equal(t, testdata.IndexSet(), is)
}

func TestClient_CreateIndexSet(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	reqBuf, err := ioutil.ReadFile("../testdata/index_set/create_index_set.json")
	require.Nil(t, err)

	respBuf, err := ioutil.ReadFile("../testdata/index_set/index_set.json")
	require.Nil(t, err)

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method:         "POST",
								Path:           "/api/system/indices/index_sets",
								PartOfHeader:   getTestHeader(),
								BodyJSONString: string(reqBuf),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: string(respBuf),
							},
						},
					},
				},
			},
		},
	})

	// nil check
	if _, err := cl.CreateIndexSet(ctx, nil); err == nil {
		t.Fatal("index set is nil")
	}
	is := testdata.CreateIndexSet()
	if _, err := cl.CreateIndexSet(ctx, &is); err != nil {
		t.Fatal(err)
	}
}

func TestClient_UpdateIndexSet(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	reqBuf, err := ioutil.ReadFile("../testdata/index_set/update_index_set.json")
	require.Nil(t, err)

	respBuf, err := ioutil.ReadFile("../testdata/index_set/index_set.json")
	require.Nil(t, err)

	is := testdata.CreateIndexSet()

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method:         "PUT",
								Path:           "/api/system/indices/index_sets/" + testdata.IndexSet().ID,
								PartOfHeader:   getTestHeader(),
								BodyJSONString: string(reqBuf),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: string(respBuf),
							},
						},
					},
				},
			},
		},
	})

	// success
	is.ID = testdata.IndexSet().ID
	if _, _, err := cl.UpdateIndexSet(ctx, testdata.IndexSetUpdateParams()); err != nil {
		t.Fatal(err)
	}
	// id required
	prms := is.NewUpdateParams()
	prms.ID = ""
	if _, _, err := cl.UpdateIndexSet(ctx, prms); err == nil {
		t.Fatal("index set id is required")
	}
	// nil check
	if _, _, err := cl.UpdateIndexSet(ctx, nil); err == nil {
		t.Fatal("index set is required")
	}
}

func TestClient_DeleteIndexSet(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	is := testdata.IndexSet()

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method:       "DELETE",
								Path:         "/api/system/indices/index_sets/" + is.ID,
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 204,
								},
							},
						},
					},
				},
			},
		},
	})

	// id required
	if _, err := cl.DeleteIndexSet(ctx, ""); err == nil {
		t.Fatal("id is required")
	}
	// invalid id
	if _, err := cl.DeleteIndexSet(ctx, is.ID); err != nil {
		t.Fatal(err)
	}
}

func TestClient_SetDefaultIndexSet(t *testing.T) {
	ctx := context.Background()

	cl, err := client.NewClient("http://example.com/api", "admin", "admin")
	require.Nil(t, err)

	respBuf, err := ioutil.ReadFile("../testdata/index_set/index_set.json")
	require.Nil(t, err)

	is := testdata.IndexSet()

	cl.SetHTTPClient(&http.Client{
		Transport: &flute.Transport{
			T: t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Tester: &flute.Tester{
								Method:       "PUT",
								Path:         "/api/system/indices/index_sets/" + is.ID + "/default",
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 200,
								},
								BodyString: string(respBuf),
							},
						},
					},
				},
			},
		},
	})

	_, _, err = cl.SetDefaultIndexSet(ctx, is.ID)
	if err != nil {
		t.Fatal("Failed to UpdateIndexSet", err)
	}
	if _, _, err := cl.SetDefaultIndexSet(ctx, ""); err == nil {
		t.Fatal("index set id is required")
	}
}
