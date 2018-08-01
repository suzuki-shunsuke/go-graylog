# graylog_stream

https://github.com/suzuki-shunsuke/terraform-provider-graylog/blob/master/resource_stream.go

```
resource "graylog_stream" "test-terraform" {
  title = "test-terraform"
  index_set_id = "${graylog_index_set.test-terraform.id}"
  disabled = true
  matching_type = "AND"
  rule {
    type = 1
    field = "foo"
    value = "bar"
    description = "foo bar"
	}
  rule {
    type = 1
    field = "bar"
    value = "foo"
    description = "bar foo"
    inverted = true
  }
}
```

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
title | string |
index_set_id | string |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
disabled | | bool |
matching_type | | string |
remove_matches_from_default_stream | | bool |
is_default | | bool |
rule | | |
rule.type | | int | matching type (check graylog API URL `/api/streams/<stream_id>/rules/types` for the list of possible types.)
rule.field | | string | field to check
rule.value | | string | value to match
rule.description | | string | rule description
rule.inverted | | bool | inverts the rule

## Attrs Reference

name | type | etc
--- | --- | ---
creator_user_id | string | computed
created_at | string | computed
