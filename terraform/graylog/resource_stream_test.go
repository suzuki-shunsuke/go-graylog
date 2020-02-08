package graylog

import (
	"net/http"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"

	"github.com/suzuki-shunsuke/go-graylog/v10/testdata"
)

func TestAccStream(t *testing.T) {
	setEnv()

	stream := testdata.Stream()

	tc := &testCase{
		t:          t,
		Name:       "stream",
		CreatePath: "/api/streams",
		GetPath:    "/api/streams/" + stream.ID,

		CreateReqBodyMap:   testdata.CreateStreamReqBodyMap(),
		UpdateReqBodyMap:   testdata.UpdateStreamReqBodyMap(),
		CreatedDataPath:    "stream/stream.json",
		UpdatedDataPath:    "stream/updated_stream.json",
		CreateRespBodyPath: "stream/create_response.json",
		UpdateRespBodyPath: "stream/stream.json",
		CreateTFPath:       "stream/create.tf",
		UpdateTFPath:       "stream/update.tf",
	}

	defaultTransport := http.DefaultClient.Transport
	defer func() {
		http.DefaultClient.Transport = defaultTransport
	}()
	transport, testCase, err := tc.Get()
	require.Nil(tc.t, err)
	tp := transport.(*flute.Transport)
	route := tp.Services[0].Routes[1]
	route.Matcher.Path = route.Tester.Path
	tp.Services[0].Routes[1] = route
	tp.Services[0].Routes = append(tp.Services[0].Routes, flute.Route{
		Name: "Resume stream",
		Matcher: &flute.Matcher{
			Method: "POST",
			Path:   tc.GetPath + "/resume",
		},
		Tester: &flute.Tester{
			PartOfHeader: getTestHeader(),
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 204,
			},
		},
	}, flute.Route{
		Name: "Pause stream",
		Matcher: &flute.Matcher{
			Method: "POST",
			Path:   tc.GetPath + "/pause",
		},
		Tester: &flute.Tester{
			PartOfHeader: getTestHeader(),
		},
		Response: &flute.Response{
			Base: http.Response{
				StatusCode: 204,
			},
		},
	})
	http.DefaultClient.Transport = transport
	resource.Test(tc.t, testCase)
}
