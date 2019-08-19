resource "graylog_alarm_callback" "test" {
  type      = "org.graylog2.alarmcallbacks.HTTPAlarmCallback"
  stream_id = graylog_stream.test.id
  title     = "test"

  http_configuration {
    url = "https://example.com"
  }
}

