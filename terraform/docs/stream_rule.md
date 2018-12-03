# graylog_stream_rule

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_stream_rule.go

```
resource "graylog_stream_rule" "test-terraform" {
  field = "tag"
  value = "${graylog_index_set.test-terraform.id}"
  stream_id = "${graylog_stream.test-terraform.id}"
  description = "test stream rule"
  type = 0
  inverted = false
}
```

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
field | string |
value | string |
description | string |
type | int |
stream_id | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
inverted | | bool |
