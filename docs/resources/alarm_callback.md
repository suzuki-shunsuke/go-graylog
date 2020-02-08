# graylog_alarm_callback

* [Example](../../examples/v0.12/alarm_callback.tf)
* [Source code](../../graylog/terraform/resource_alarm_callback.go)

## How to import

Specify `<stream id>/<alarm callback id>` as ID.

```console
$ terraform import graylog_alarm_callback.test 5bb1b4b5c9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```

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

### Required Argument

name | type | description
--- | --- | ---
http_configuration | |
http_configuration.url | string |

### Optional Argument

None.

## type: EmailAlarmCallback 

`org.graylog2.alarmcallbacks.EmailAlarmCallback`

### Required Argument

name | type | description
--- | --- | ---
email_configuration | |
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
slack_configuration.link_names | true | bool |
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

```hcl
resource "graylog_alarm_callback" "hipchat" {
  type = "org.graylog2.alarmcallbacks.hipchat.HipChatAlarmCallback"
  stream_id = "000000000000000000000001"
  title = "test"
  general_string_configuration {
    color = "yellow"
    api_url = "https://api.hipchat.com"
    message_template = "test template"
    api_token = "test"
    graylog_base_url = "http://localhost:9000"
    room = "test"
  }
  general_bool_configuration {
    notify = "true"
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
