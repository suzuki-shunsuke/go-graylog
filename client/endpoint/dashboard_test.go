package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog/v8/client/endpoint"
)

func TestEndpoints_Dashboards(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/dashboards", apiURL), ep.Dashboards())
}

func TestEndpoints_Dashboard(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/dashboards/%s", apiURL, ID), ep.Dashboard(ID))
}
