package endpoint_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/suzuki-shunsuke/go-graylog/client/endpoint/v8"
)

func TestEndpoints_Inputs(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/system/inputs", apiURL), ep.Inputs())
}

func TestEndpoints_Input(t *testing.T) {
	ep, err := endpoint.NewEndpoints(apiURL)
	require.Nil(t, err)
	require.Equal(t, fmt.Sprintf("%s/system/inputs/%s", apiURL, ID), ep.Input(ID))
}
