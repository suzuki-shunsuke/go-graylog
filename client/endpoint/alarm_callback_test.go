package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog/v8/client/endpoint"
)

func TestEndpoints_AlarmCallbacks(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/alerts/callbacks", apiURL), ep.AlarmCallbacks())
}
