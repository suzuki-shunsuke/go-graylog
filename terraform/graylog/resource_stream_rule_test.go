package graylog

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog/v10/testdata"
	"github.com/suzuki-shunsuke/go-graylog/v10/testutil"
)

func TestAccStreamRule(t *testing.T) {
	setEnv()

	rule := testdata.StreamRule()

	tc := &testCase{
		t:          t,
		Name:       "stream rule",
		CreatePath: "/api/streams/" + rule.StreamID + "/rules",
		GetPath:    "/api/streams/" + rule.StreamID + "/rules/" + rule.ID,

		CreateReqBodyMap:   testutil.ConvertIntToFloat64OfMap(testdata.CreateStreamRuleReqBodyMap()),
		UpdateReqBodyMap:   testutil.ConvertIntToFloat64OfMap(testdata.UpdateStreamRuleReqBodyMap()),
		CreatedDataPath:    "stream_rule/stream_rule.json",
		UpdatedDataPath:    "stream_rule/updated_stream_rule.json",
		CreateRespBodyPath: "stream_rule/create_response.json",
		UpdateRespBodyPath: "stream_rule/create_response.json",
		CreateTFPath:       "stream_rule/create.tf",
		UpdateTFPath:       "stream_rule/update.tf",
	}

	tc.Test()
}
