package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog/v8/client/endpoint"
)

func TestEndpoints_StreamRules(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/streams/%s/rules", apiURL, ID), ep.StreamRules(ID))
}

func TestEndpoints_StreamRuleTypes(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/streams/%s/rules/types", apiURL, ID), ep.StreamRuleTypes(ID))
}

func TestEndpoints_StreamRule(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/streams/%s/rules/%s", apiURL, ID, ID), ep.StreamRule(ID, ID))
}
