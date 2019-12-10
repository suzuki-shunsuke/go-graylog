resource "graylog_event_definition" "test" {
  title = "new-event-definition"
  description = ""
  priority = 1
  alert = true
  config = <<EOF
{
  "type": "aggregation-v1",
  "query": "test",
  "streams": [
    "${graylog_stream.test.id}"
  ],
  "search_within_ms": 60000,
  "execute_every_ms": 60000,
  "group_by": [],
  "series": [],
  "conditions": null
}
EOF
  field_spec = <<EOF
{
  "test": {
    "data_type": "string",
    "providers": [
      {
        "type": "template-v1",
        "template": "test",
        "require_values": false
      }
    ]
  }
}
EOF
  
  notification_settings {
    grace_period_ms = 0
    backlog_size = 0
  }

  notifications {
    notification_id = graylog_event_notification.http.id
  }
}
