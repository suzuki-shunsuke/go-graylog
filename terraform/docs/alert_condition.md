# graylog_alert_condition

* http://docs.graylog.org/en/2.5/pages/streams/alerts.html#conditions
* https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_alert_condition.go

```hcl
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

## Breaking Changes

* v1 -> v2: https://github.com/suzuki-shunsuke/go-graylog/issues/76

## Argument Reference

### Common Required Argument

name | type | description
--- | --- | ---
type | string |
title | string |

### Common Optional Argument

name | default | type | description
--- | --- | --- | ---
in_grace | bool |

## type: field_content_value 

```hcl
resource "graylog_alert_condition" "test-terraform" {
  type = "field_content_value"
  stream_id = "${graylog_stream.test-terraform.id}"
  in_grace = false
  title = "test"
  field_content_value_parameters = {
    field = "message"
    value = "hoge hoge"
    backlog = 1
    repeat_notifications = false
    query = "*"
    grace = 0
  }
}
```

### Required Argument

name | type | description
--- | --- | ---
field_content_value_parameters.field | string |
field_content_value_parameters.value | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
field_content_value_parameters.grace | 0 | int |
field_content_value_parameters.backlog | 0 | int |
field_content_value_parameters.query | "" | string |
field_content_value_parameters.repeat_notifications | false | bool |

## type: field_value 

### Required Argument

name | type | description
--- | --- | ---
field_value_parameters | object |
field_value_parameters.field | string |
field_value_parameters.type | string |
field_value_parameters.threshold_type | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
field_value_parameters.grace | 0 | int |
field_value_parameters.backlog | 0 | int |
field_value_parameters.query | "" | string |
field_value_parameters.threshold | 0 | int |
field_value_parameters.time | 0 | int |
field_value_parameters.repeat_notifications | false | bool |

## type: message_count 

### Required Argument

name | type | description
--- | --- | ---
message_count_parameters.threshold_type | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
message_count_parameters.grace | 0 | int |
message_count_parameters.backlog | 0 | int |
message_count_parameters.query | "" | string |
message_count_parameters.threshold | 0 | int |
message_count_parameters.time | 0 | int |
message_count_parameters.repeat_notifications | false | bool |

## type: other third party's Alert Condition

We support only the above alert condition types officially,
but in order to support other alert condition types as much as possible,
we provide some additional attributes.

* `general_int_parameters`
* `general_bool_parameters`
* `general_float_parameters`
* `general_string_parameters`

### Required Argument

None.

### Optional Argument

name | default | type | description
--- | --- | --- | ---
general_int_parameters | {} | map[string]int |
general_bool_parameters | {} | map[string]bool |
general_float_parameters | {} | map[string]float64 |
general_string_parameters | {} | map[string]string |
