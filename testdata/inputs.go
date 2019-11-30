package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v8"
)

var (
	Inputs = &graylog.InputsBody{
		Inputs: []graylog.Input{
			{
				Title: "gelf udp 2",
				Attrs: &graylog.InputGELFUDPAttrs{
					DecompressSizeLimit: 8388608,
					OverrideSource:      "",
					BindAddress:         "0.0.0.0",
					Port:                12201,
					RecvBufferSize:      262144,
				},
				ID:            "5d84c1aa2ab79c000d35d6d9",
				Global:        true,
				Node:          "",
				CreatedAt:     "2019-09-20T12:10:18.010Z",
				CreatorUserID: "admin",
				StaticFields: map[string]string{
					"foo": "bar",
				},
			},
		},
		Total: 1,
	}
)
