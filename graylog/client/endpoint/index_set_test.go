package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/client/endpoint"
)

func TestEndpoints_IndexSets(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/system/indices/index_sets", apiURL), ep.IndexSets())
}

func TestEndpoints_IndexSet(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/system/indices/index_sets/%s", apiURL, ID), ep.IndexSet(ID))
}

func TestEndpoints_SetDefaultIndexSet(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/system/indices/index_sets/%s/default", apiURL, ID), ep.SetDefaultIndexSet(ID))
}

func TestEndpoints_IndexSetsStats(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/system/indices/index_sets/stats", apiURL), ep.IndexSetsStats())
}

func TestEndpoints_IndexSetStats(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/system/indices/index_sets/%s/stats", apiURL, ID), ep.IndexSetStats(ID))
}
