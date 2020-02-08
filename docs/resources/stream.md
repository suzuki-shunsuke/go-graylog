# graylog_stream

* [Example](../../examples/v0.12/stream.tf)
* [Source code](../../graylog/terraform/resource_stream.go)

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
description | | string |
remove_matches_from_default_stream | | bool |
is_default | | bool |

## Attrs Reference

name | type | etc
--- | --- | ---
creator_user_id | string | computed
created_at | string | computed
