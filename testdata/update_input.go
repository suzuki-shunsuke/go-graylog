package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v10"
	"github.com/suzuki-shunsuke/go-ptr"
)

func UpdateInput() graylog.InputUpdateParams {
	return graylog.InputUpdateParams{
		Title: "gelf udp 2",
		Attrs: &graylog.InputGELFUDPAttrs{
			DecompressSizeLimit: 8388608,
			BindAddress:         "0.0.0.0",
			Port:                12201,
			RecvBufferSize:      262144,
		},
		Global: ptr.PBool(true),
	}
}
