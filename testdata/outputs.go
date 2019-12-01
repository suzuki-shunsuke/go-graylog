package testdata

import (
	"github.com/suzuki-shunsuke/go-graylog/v8"
)

var (
	Outputs = &graylog.OutputsBody{
		Outputs: []graylog.Output{
			{
				ID:            "5de3772faf66c6001353cccb",
				Title:         "stdout",
				Type:          "org.graylog2.outputs.LoggingOutput",
				CreatorUserID: "admin",
				CreatedAt:     "2019-12-01T08:17:51.102Z",
				Configuration: map[string]interface{}{
					"prefix": "Writing message (updated): ",
				},
			},
			{
				ID:            "5de37740af66c6001353cce0",
				Title:         "gelf",
				Type:          "org.graylog2.outputs.GelfOutput",
				CreatorUserID: "admin",
				CreatedAt:     "2019-12-01T08:18:08.696Z",
				Configuration: map[string]interface{}{
					"connect_timeout":          1000,
					"hostname":                 "localhost",
					"max_inflight_sends":       512,
					"port":                     12201,
					"protocol":                 "TCP",
					"queue_size":               512,
					"reconnect_delay":          500,
					"tcp_keep_alive":           false,
					"tcp_no_delay":             false,
					"tls_trust_cert_chain":     "",
					"tls_verification_enabled": false,
				},
			},
			{
				ID:            "5de37825af66c6001353cdec",
				Title:         "slack",
				Type:          "org.graylog2.plugins.slack.output.SlackMessageOutput",
				CreatorUserID: "admin",
				CreatedAt:     "2019-12-01T08:21:57.480Z",
				Configuration: map[string]interface{}{
					"add_details":    true,
					"channel":        "#channel",
					"color":          "#FF0000",
					"custom_message": "##########\nDate: ${check_result.triggeredAt}\nStream ID: ${stream.id}\nStream title: ${stream.title}\nStream description: ${stream.description}\n${if stream_url}Stream URL: ${stream_url}${end}\n##########\n",
					"graylog2_url":   "",
					"icon_emoji":     "",
					"icon_url":       "",
					"link_names":     true,
					"notify_channel": false,
					"proxy_address":  "",
					"short_mode":     false,
					"user_name":      "Graylog",
					"webhook_url":    "http://example.com",
				},
			},
		},
		Total: 3,
	}
)
