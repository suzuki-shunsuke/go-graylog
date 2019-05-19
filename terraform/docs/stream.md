# graylog_stream

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_stream.go

```hcl
resource "graylog_stream" "test-terraform" {
  title = "test-terraform"
  index_set_id = "${graylog_index_set.test-terraform.id}"
  disabled = true
  matching_type = "AND"
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

## Attrs Reference

name | type | etc
--- | --- | ---
creator_user_id | string | computed
created_at | string | computed
