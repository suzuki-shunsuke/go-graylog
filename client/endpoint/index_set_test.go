package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

func TestIndexSets(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/system/indices/index_sets", apiURL), ep.IndexSets())
}

func TestIndexSet(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/system/indices/index_sets/%s", apiURL, ID), ep.IndexSet(ID))
}

func TestSetDefaultIndexSet(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/system/indices/index_sets/%s/default", apiURL, ID), ep.SetDefaultIndexSet(ID))
}

func TestIndexSetsStats(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/system/indices/index_sets/stats", apiURL), ep.IndexSetsStats())
}

func TestIndexSetStats(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/system/indices/index_sets/%s/stats", apiURL, ID), ep.IndexSetStats(ID))
}
