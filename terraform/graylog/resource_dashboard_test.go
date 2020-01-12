package graylog

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v9/testdata"
)

func TestAccDashboard(t *testing.T) {
	setEnv()

	createRespBody, err := ioutil.ReadFile("../../testdata/dashboard/create_dashboard_response.json")
	require.Nil(t, err)

	getBody, err := ioutil.ReadFile("../../testdata/dashboard/dashboard.json")
	require.Nil(t, err)

	updatedGetBody, err := ioutil.ReadFile("../../testdata/dashboard/updated_dashboard.json")
	require.Nil(t, err)

	createTF, err := ioutil.ReadFile("../../testdata/dashboard/dashboard.tf")
	require.Nil(t, err)

	updateTF, err := ioutil.ReadFile("../../testdata/dashboard/update_dashboard.tf")
	require.Nil(t, err)

	store := newBodyStore("")

	ds := testdata.Dashboard()

	dsPath := "/api/dashboards/" + ds.ID

	defaultTransport := http.DefaultClient.Transport
	defer func() {
		http.DefaultClient.Transport = defaultTransport
	}()
	http.DefaultClient.Transport = &flute.Transport{
		T: t,
		Services: []flute.Service{
			{
				Endpoint: "http://example.com",
				Routes: []flute.Route{
					{
						Name: "get a dashboard",
						Matcher: &flute.Matcher{
							Method: "GET",
						},
						Tester: &flute.Tester{
							Path:         dsPath,
							PartOfHeader: getTestHeader(),
						},
						Response: &flute.Response{
							Response: func(req *http.Request) (*http.Response, error) {
								return &http.Response{
									StatusCode: 200,
									Body:       ioutil.NopCloser(strings.NewReader(store.Get())),
								}, nil
							},
						},
					},
					{
						Name: "create a dashboard",
						Matcher: &flute.Matcher{
							Method: "POST",
						},
						Tester: &flute.Tester{
							Path:         "/api/dashboards",
							PartOfHeader: getTestHeader(),
							Test: genTestBody(map[string]interface{}{
								"title":       "test",
								"description": "test",
							}, string(getBody), store),
						},
						Response: &flute.Response{
							Base: http.Response{
								StatusCode: 201,
							},
							BodyString: string(createRespBody),
						},
					},
					{
						Name: "update a dashboard",
						Matcher: &flute.Matcher{
							Method: "PUT",
						},
						Tester: &flute.Tester{
							Path:         dsPath,
							PartOfHeader: getTestHeader(),
							Test: genTestBody(map[string]interface{}{
								"title":       "updated title",
								"description": "updated description",
							}, string(updatedGetBody), store),
						},
						Response: &flute.Response{
							Base: http.Response{
								StatusCode: 204,
							},
						},
					},
					{
						Name: "delete a dashboard",
						Matcher: &flute.Matcher{
							Method: "DELETE",
						},
						Tester: &flute.Tester{
							Path:         dsPath,
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
	}

	resource.Test(t, resource.TestCase{
		Providers: getTestProviders(),
		Steps: []resource.TestStep{
			{
				Config: string(createTF),
			},
			{
				Config: string(updateTF),
			},
		},
	})
}
