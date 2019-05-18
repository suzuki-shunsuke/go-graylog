# graylog_stream

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_stream.go

```
resource "graylog_stream" "test-terraform" {
  title = "test-terraform"
  index_set_id = "${graylog_index_set.test-terraform.id}"
  disabled = true
  matching_type = "AND"
  sync_pipelines = true
  pipelines = ["000000000000000000000000"]
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
pipelines | `[]` | []string | pipeline ids which connect to the stream

If `pipelines` is set pipeline connection is synchronized, otherwise pipeline connctions isn't synchronized.
There is no API to delete pipeline connections so we treat pipeline connections as not resource but attribute of streams.

## Attrs Reference

name | type | etc
--- | --- | ---
creator_user_id | string | computed
created_at | string | computed
