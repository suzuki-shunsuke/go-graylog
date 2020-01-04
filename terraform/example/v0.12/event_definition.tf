resource "graylog_event_definition" "test" {
  title       = "new-event-definition"
  description = ""
  priority    = 1
  alert       = true
  config      = <<EOF
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
  field_spec  = <<EOF
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
    backlog_size    = 0
  }

  notifications {
    notification_id = graylog_event_notification.http.id
  }
}

# https://github.com/suzuki-shunsuke/go-graylog/issues/170#issuecomment-564817360
# https://www.terraform.io/docs/providers/random/r/uuid.html
resource "random_uuid" "event_definition_test2_series0" {}

resource "graylog_event_definition" "test2" {
  title    = "new-event-definition 2"
  priority = 2
  config   = <<EOF
{
  "type": "aggregation-v1",
  "query": "test",
  "streams": [
    "${graylog_stream.test.id}"
  ],
  "search_within_ms": 60000,
  "execute_every_ms": 60000,
  "group_by": [
    "alert"
  ],
  "series": [
    {
      "id": "${random_uuid.event_definition_test2_series0.result}",
      "function": "avg",
      "field": "alert"
    }
  ],
  "conditions": {
    "expression": {
      "expr": ">",
      "left": {
        "expr": "number-ref",
        "ref": "${random_uuid.event_definition_test2_series0.result}"
      },
      "right": {
        "expr": "number",
        "value": 0
      }
    }
  }
}
EOF

  notification_settings {
    grace_period_ms = 0
    backlog_size    = 0
  }
}
