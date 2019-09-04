package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint"
)

func TestStreamRules(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/streams/%s/rules", apiURL, ID), ep.StreamRules(ID))
}

func TestStreamRuleTypes(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/streams/%s/rules/types", apiURL, ID), ep.StreamRuleTypes(ID))
}

func TestStreamRule(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/streams/%s/rules/%s", apiURL, ID, ID), ep.StreamRule(ID, ID))
}
