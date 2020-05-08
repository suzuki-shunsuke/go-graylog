resource "graylog_output" "stdout" {
  title = "stdout"
  type  = "org.graylog2.outputs.LoggingOutput"

  configuration = jsonencode({
    prefix = "Writing message: "
  })
}

resource "graylog_output" "gelf" {
  title = "gelf"
  type  = "org.graylog2.outputs.GelfOutput"

  configuration = jsonencode({
    "hostname" : "localhost",
    "protocol" : "TCP",
    "connect_timeout" : 1000,
    "reconnect_delay" : 500,
    "queue_size" : 512,
    "port" : 12201,
    "max_inflight_sends" : 512,
    "tcp_no_delay" : false,
    "tcp_keep_alive" : false,
    "tls_trust_cert_chain" : "",
    "tls_verification_enabled" : false
  })
}

resource "graylog_output" "slack" {
  title = "slack"
  type  = "org.graylog2.plugins.slack.output.SlackMessageOutput"

  configuration = jsonencode({
    "icon_url" : "",
    "graylog2_url" : "",
    "link_names" : true,
    "color" : "#FF0000",
    "webhook_url" : "http://example.com",
    "icon_emoji" : "",
    "user_name" : "Graylog",
    "proxy_address" : "",
    "channel" : "#channel",
    "custom_message" : "message",
    "notify_channel" : false,
    "short_mode" : false,
    "add_details" : true
  })
}
