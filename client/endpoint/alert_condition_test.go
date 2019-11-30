package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint/v8"
)

func TestEndpoints_AlertConditions(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/alerts/conditions", apiURL), ep.AlertConditions())
}
