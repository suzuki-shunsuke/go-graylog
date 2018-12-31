# graylog_alarm_callback

* http://docs.graylog.org/en/2.5/pages/streams/alerts.html#notifications
* https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_alarm_callback.go

## Argument Reference

### Common Required Argument

name | type | description
--- | --- | ---
type | string |
title | string |
stream_id | string |

### Common Optional Argument

None.

## type: HTTPAlarmCallback 

`org.graylog2.alarmcallbacks.HTTPAlarmCallback`

```
resource "graylog_alarm_callback" "test-terraform" {
  type = "org.graylog2.alarmcallbacks.HTTPAlarmCallback"
  stream_id = "${graylog_stream.test-terraform.id}"
  title = "test"
  http_configuration = {
    url = "https://example.com"
  }
}
```

### Required Argument

name | type | description
--- | --- | ---
http_configuration | |
http_configuration.url | string |

### Optional Argument

None.

## type: EmailAlarmCallback 

`org.graylog2.alarmcallbacks.EmailAlarmCallback`

```
resource "graylog_alarm_callback" "test-terraform" {
  type = "org.graylog2.alarmcallbacks.EmailAlarmCallback"
  stream_id = "${graylog_stream.test-terraform.id}"
  title = "test"
  email_configuration = {
    sender = "graylog@example.org"
    subject = "Graylog alert for stream: $${stream.title}: $${check_result.resultDescription}"
    user_receivers = [
      "username"
    ]
    email_receivers = [
      "graylog@example.com"
    ]
    body = "##########\\nAlert Description: $${check_result.resultDescription}\\nDate: $${check_result.triggeredAt}\\nStream ID: $${stream.id}\\nStream title: $${stream.title}\\nStream description: $${stream.description}\\nAlert Condition Title: $${alertCondition.title}\\n$${if stream_url}Stream URL: $${stream_url}$${end}\\n\\nTriggered condition: $${check_result.triggeredCondition}\\n##########\\n\\n$${if backlog}Last messages accounting for this alert:\\n$${foreach backlog message}$${message}\\n\\n$${end}$${else}<No backlog>\\n$${end}\\n"
  }
}
```

### Required Argument

name | type | description
--- | --- | ---
email_configuration | string |
email_configuration.sender | string |
email_configuration.subject | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
email_configuration.body | "" | string |
email_configuration.user_receivers | [] | []string |
email_configuration.email_receivers | [] | []string |

## type: SlackAlarmCallback 

`org.graylog2.plugins.slack.callback.SlackAlarmCallback`

```
resource "graylog_alarm_callback" "test-terraform" {
  type = "org.graylog2.plugins.slack.callback.SlackAlarmCallback"
  stream_id = "${graylog_stream.test-terraform.id}"
  title = "test"
  slack_configuration = {
    graylog2_url = "https://graylog.example.com"
    color = "#FF0000"
    webhook_url = "https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX"
    user_name = "Graylog"
    backlog_items = 5
    channel = "#general"
    custom_message = "$${alert_condition.title}\\n\\n$${foreach backlog message}\\n<https://graylog.example.com/streams/$${stream.id}/search?rangetype=absolute&from=$${message.timestamp}&to=$${message.timestamp} | link> $${message.message}\\n$${end}"
  }
}
```

### Required Argument

name | type | description
--- | --- | ---
slack_configuration | |
slack_configuration.color | string |
slack_configuration.webhook_url | string |
slack_configuration.channel | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
slack_configuration.icon_url | "" | string |
slack_configuration.graylog2_url | "" | string |
slack_configuration.icon_emoji | "" | string |
slack_configuration.user_name | "" | string |
slack_configuration.proxy_address | "" | string |
slack_configuration.custom_message | "" | string |
slack_configuration.backlog_items | 0 | int |
slack_configuration.link_names | false | bool |
slack_configuration.notify_channel | false | bool |

## type: other third party's Alarm Callback

We support only the above alarm callback types officially,
but in order to support other alarm callback types as much as possible,
we provide some additional attributes.

* `general_int_configuration`
* `general_bool_configuration`
* `general_float_configuration`
* `general_string_configuration`

For example, you can use the [HipChat Plugin](https://marketplace.graylog.org/addons/e316cbfc-663f-4718-aa54-8aff97749449) although we don't support it officially.

```
resource "graylog_alarm_callback" "hipchat" {
  type = "org.graylog2.alarmcallbacks.hipchat.HipChatAlarmCallback"
  stream_id = "000000000000000000000001"
  title = "test"
  general_string_configuration = {
    color = "yellow"
    api_url = "https://api.hipchat.com"
    message_template = "test template"
    api_token = "test"
    graylog_base_url = "http://localhost:9000"
    room = "test"
  }
  general_bool_configuration = {
    notify = true
  }
}
```

### Required Argument

None.

### Optional Argument

name | default | type | description
--- | --- | --- | ---
general_int_configuration | {} | map[string]int |
general_bool_configuration | {} | map[string]bool |
general_float_configuration | {} | map[string]float64 |
general_string_configuration | {} | map[string]string |
