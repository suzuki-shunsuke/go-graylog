package graylog

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"
)

type testCase struct {
	t          *testing.T
	Name       string
	CreatePath string
	GetPath    string

	ConvertReqBody func(body io.Reader) (map[string]interface{}, error)

	CreateReqBodyPath string
	CreateReqBodyMap  map[string]interface{}

	UpdateReqBodyPath string
	UpdateReqBodyMap  map[string]interface{}

	CreatedDataPath string
	UpdatedDataPath string

	CreateRespBodyPath string
	UpdateRespBodyPath string

	CreateTFPath string
	UpdateTFPath string

	Store *bodyStore
}

func getReqBody(t *testing.T, path string, m map[string]interface{}) (map[string]interface{}, error) {
	if path != "" {
		buf, err := ioutil.ReadFile("../../testdata/" + path)
		require.Nil(t, err)
		var body map[string]interface{}
		require.Nil(t, json.Unmarshal(buf, &body))
		return body, nil
	}
	return m, nil
}

func (tc *testCase) Test() {
	defaultTransport := http.DefaultClient.Transport
	defer func() {
		http.DefaultClient.Transport = defaultTransport
	}()
	transport, testCase, err := tc.Get()
	require.Nil(tc.t, err)
	http.DefaultClient.Transport = transport
	resource.Test(tc.t, testCase)
}

func (tc *testCase) genTestBody(exp map[string]interface{}, bodyString string, store *bodyStore) func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
	return func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
		if tc.ConvertReqBody != nil {
			data, err := tc.ConvertReqBody(req.Body)
			require.Nil(t, err, "failed to convert a request body; route: "+route.Name)
			assert.Equal(t, exp, data, "request body should match; route: "+route.Name)
		} else {
			body := map[string]interface{}{}
			require.Nil(t, json.NewDecoder(req.Body).Decode(&body), "failed to unmarshal request body as JSON; route: "+route.Name)
			assert.Equal(t, exp, body, "request body should match; route: "+route.Name)
		}
		store.Set(bodyString)
	}
}

func (tc *testCase) Get() (http.RoundTripper, resource.TestCase, error) {
	createReqBody, err := getReqBody(tc.t, tc.CreateReqBodyPath, tc.CreateReqBodyMap)
	require.Nil(tc.t, err)
	updateReqBody, err := getReqBody(tc.t, tc.UpdateReqBodyPath, tc.UpdateReqBodyMap)
	require.Nil(tc.t, err)
	if tc.Store == nil {
		tc.Store = newBodyStore("")
	}

	createTF, err := ioutil.ReadFile("../../testdata/" + tc.CreateTFPath)
	require.Nil(tc.t, err)

	updateTF, err := ioutil.ReadFile("../../testdata/" + tc.UpdateTFPath)
	require.Nil(tc.t, err)

	createdData, err := ioutil.ReadFile("../../testdata/" + tc.CreatedDataPath)
	require.Nil(tc.t, err)

	updatedData, err := ioutil.ReadFile("../../testdata/" + tc.UpdatedDataPath)
	require.Nil(tc.t, err)

	createRespBody := ""
	if tc.CreateRespBodyPath != "" {
		x, err := ioutil.ReadFile("../../testdata/" + tc.CreateRespBodyPath)
		require.Nil(tc.t, err)
		createRespBody = string(x)
	}

	updateRespBody := ""
	if tc.UpdateRespBodyPath != "" {
		x, err := ioutil.ReadFile("../../testdata/" + tc.UpdateRespBodyPath)
		require.Nil(tc.t, err)
		updateRespBody = string(x)
	}

	return &flute.Transport{
			T: tc.t,
			Services: []flute.Service{
				{
					Endpoint: "http://example.com",
					Routes: []flute.Route{
						{
							Name: "Get " + tc.Name,
							Matcher: &flute.Matcher{
								Method: "GET",
							},
							Tester: &flute.Tester{
								Path:         tc.GetPath,
								PartOfHeader: getTestHeader(),
							},
							Response: &flute.Response{
								Response: func(req *http.Request) (*http.Response, error) {
									return &http.Response{
										StatusCode: 200,
										Body:       ioutil.NopCloser(strings.NewReader(tc.Store.Get())),
									}, nil
								},
							},
						},
						{
							Name: "Create " + tc.Name,
							Matcher: &flute.Matcher{
								Method: "POST",
							},
							Tester: &flute.Tester{
								Path:         tc.CreatePath,
								PartOfHeader: getTestHeader(),
								Test:         tc.genTestBody(createReqBody, string(createdData), tc.Store),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 201,
								},
								BodyString: createRespBody,
							},
						},
						{
							Name: "Update " + tc.Name,
							Matcher: &flute.Matcher{
								Method: "PUT",
							},
							Tester: &flute.Tester{
								Path:         tc.GetPath,
								PartOfHeader: getTestHeader(),
								Test:         tc.genTestBody(updateReqBody, string(updatedData), tc.Store),
							},
							Response: &flute.Response{
								Base: http.Response{
									StatusCode: 201,
								},
								BodyString: updateRespBody,
							},
						},
						{
							Name: "Delete " + tc.Name,
							Matcher: &flute.Matcher{
								Method: "DELETE",
							},
							Tester: &flute.Tester{
								Path:         tc.GetPath,
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
		}, resource.TestCase{
			Providers: getTestProviders(),
			Steps: []resource.TestStep{
				{
					Config: string(createTF),
				},
				{
					Config: string(updateTF),
				},
			},
		}, nil
}

func getTestProviders() map[string]terraform.ResourceProvider {
	return map[string]terraform.ResourceProvider{
		"graylog": Provider(),
	}
}

func getTestHeader() http.Header {
	return http.Header{
		"Content-Type":   []string{"application/json"},
		"X-Requested-By": []string{"terraform-provider-graylog"},
		"Authorization":  nil,
	}
}

type bodyStore struct {
	body  string
	mutex *sync.RWMutex
}

func newBodyStore(body string) *bodyStore {
	return &bodyStore{
		body:  body,
		mutex: &sync.RWMutex{},
	}
}

func (store *bodyStore) Get() string {
	store.mutex.RLock()
	a := store.body
	store.mutex.RUnlock()
	return a
}

func (store *bodyStore) Set(body string) {
	store.mutex.Lock()
	store.body = body
	store.mutex.Unlock()
}
