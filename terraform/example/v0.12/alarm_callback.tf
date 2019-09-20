resource "graylog_alarm_callback" "http" {
  type      = "org.graylog2.alarmcallbacks.HTTPAlarmCallback"
  stream_id = graylog_stream.test.id
  title     = "test"

  http_configuration {
    url = "https://example.com"
  }
}

resource "graylog_alarm_callback" "slack" {
  type = "org.graylog2.plugins.slack.callback.SlackAlarmCallback"
  stream_id = graylog_stream.test.id
  title = "test"
  slack_configuration {
    graylog2_url = "https://graylog.example.com"
    color = "#FF0000"
    webhook_url = "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
    user_name = "Graylog"
    backlog_items = 5
    channel = "#general"
    custom_message = "$${alert_condition.title}\\n\\n$${foreach backlog message}\\n<https://graylog.example.com/streams/$${stream.id}/search?rangetype=absolute&from=$${message.timestamp}&to=$${message.timestamp} | link> $${message.message}\\n$${end}"
  }
}
