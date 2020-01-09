package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog/v9/client/endpoint"
)

func TestEndpoints_Streams(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/streams", apiURL), ep.Streams())
}

func TestEndpoints_Stream(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/streams/%s", apiURL, ID), ep.Stream(ID))
}

func TestEndpoints_PauseStream(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/streams/%s/pause", apiURL, ID), ep.PauseStream(ID))
}

func TestEndpoints_ResumeStream(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/streams/%s/resume", apiURL, ID), ep.ResumeStream(ID))
}

func TestEndpoints_EnabledStreams(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/streams/enabled", apiURL), ep.EnabledStreams())
}
