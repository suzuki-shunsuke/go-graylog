# graylog_alert_condition

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_alert_condition.go

```
resource "graylog_alert_condition" "test-terraform" {
  type = "field_content_value"
  stream_id = "${graylog_stream.test-terraform.id}"
  in_grace = false
  title = "test"
  parameters = {
    backlog = 1
    repeat_notifications = false
    field = "message"
    query = "*"
    grace = 0
    value = "hoge hoge"
  }
}
```

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
type | string |
title | string |
parameters | |
parameters.backlog | int |
parameters.grace | int |
parameters.value | int |
parameters.field | string |
stream_id | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
in_grace | bool |
parameters.repeat_notifications | bool |
parameters.query | string |
