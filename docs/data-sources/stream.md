# Data source graylog_stream

* [Source code](../../graylog/terraform/data_source_stream.go)

```hcl
data "graylog_stream" "test" {
  title = "test"
}
```

## Required Argument

One of `stream_id` or `title` must be set.
If `title` is specified, the title must be unique in all streams.

## Attributes

name | type | description
--- | --- | ---
title | string |
stream_id | string |
index_set_id | string |
disabled | bool |
matching_type | string |
remove_matches_from_default_stream | bool |
is_default | bool |
creator_user_id | string |
created_at | string |
