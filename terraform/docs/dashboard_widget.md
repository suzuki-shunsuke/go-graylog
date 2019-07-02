# graylog_dashboard_widget

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_dashboard_widget.go

```hcl
resource "graylog_dashboard_widget" "test" {
  description = "Stream search result count"
  dashboard_id = "5b6586000000000000000000"
  type = "STREAM_SEARCH_RESULT_COUNT"
  config {
    timerange {
      type = "relative"
      range = 300
    }
    lower_is_better = true
    trend = true
    stream_id = "5b3983000000000000000000"
    query = ""
  }
  cache_time = 10
}
```

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
type | string |
description | string |
dashboard_id | string |
config | object |
config.timerange | object |
config.timerange.type | string |
config.timerange.range | int |

### Optional Argument

name | type | default | description
--- | --- | --- | ---
config.stream_id | string | | 
config.lower_is_better | bool | false | 
config.trend | bool | false | 
config.query | string | "" | 
cache_time | int | 0 | 

## Attributes Reference

name | type | description | etc
--- | --- | --- | ---
creator_user_id | string | | computed
