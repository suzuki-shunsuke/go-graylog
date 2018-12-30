# graylog_alert_condition

* http://docs.graylog.org/en/2.5/pages/streams/alerts.html#conditions
* https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_alert_condition.go

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

`parameters`'s fields depend on alert condition's type.

### Common Required Argument

name | type | description
--- | --- | ---
type | string |
title | string |
parameters | |

### Common Optional Argument

name | default | type | description
--- | --- | --- | ---
in_grace | bool |

## type: field_content_value 

### Required Argument

name | type | description
--- | --- | ---
parameters.field | string |
parameters.value | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
parameters.grace | 0 | int |
parameters.backlog | 0 | int |
parameters.query | "" | string |
parameters.repeat_notifications | false | bool |

## type: field_value 

### Required Argument

name | type | description
--- | --- | ---
parameters.field | string |
parameters.type | string |
parameters.threshold_type | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
parameters.grace | 0 | int |
parameters.backlog | 0 | int |
parameters.query | "" | string |
parameters.threshold | 0 | int |
parameters.time | 0 | int |

## type: message_count 

### Required Argument

name | type | description
--- | --- | ---
parameters.threshold_type | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
parameters.grace | 0 | int |
parameters.backlog | 0 | int |
parameters.query | "" | string |
parameters.threshold | 0 | int |
parameters.time | 0 | int |
