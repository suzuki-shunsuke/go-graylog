# graylog_pipeline_connection

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_pipeline_connection.go

```
resource "graylog_pipeline_connection" "test" {
  stream_id      = "000000000000000000000000"
  pipeline_ids = ["000000000000000000000000"]
}
```

## Import

Specify the stream id as ID.

```console
$ terraform import graylog_pipeline_connection.test <stream id>
```

## Argument Reference

### Required Argument

name | type | etc
--- | --- | ---
stream_id | string |
pipeline_ids | []string |

### Optional Argument

None.

## Note

This resource treats the stream id as the resource id,
because there is no Graylog API to operate resource by connection pipeline id.
So please make the stream id unique in all `graylog_pipeline_connection` resources.
