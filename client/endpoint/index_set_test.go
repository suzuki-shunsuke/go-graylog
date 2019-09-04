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
	require.Equal(t, fmt.Sprintf("%s/%s", apiURL, "system/indices/index_sets"), ep.IndexSets())
}

func TestIndexSet(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/%s/%s", apiURL, "system/indices/index_sets", ID), ep.IndexSet(ID))
}

func TestSetDefaultIndexSet(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/%s/%s/default", apiURL, "system/indices/index_sets", ID), ep.SetDefaultIndexSet(ID))
}

func TestIndexSetsStats(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/%s", apiURL, "system/indices/index_sets/stats"), ep.IndexSetsStats())
}

func TestIndexSetStats(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/%s/%s/stats", apiURL, "system/indices/index_sets", ID), ep.IndexSetStats(ID))
}
