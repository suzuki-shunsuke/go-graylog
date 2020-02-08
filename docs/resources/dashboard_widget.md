# graylog_dashboard_widget

* [Example](../../examples/v0.12/dashboard.tf)
* [Source code](../../graylog/terraform/resource_dashboard_widget.go)

## Supported types

* STREAM_SEARCH_RESULT_COUNT
* QUICKVALUES
* QUICKVALUES_HISTOGRAM
* SEARCH_RESULT_CHART
* FIELD_CHART
* STATS_COUNT

## `json_configuration`

From v10.0.0, the attribute `json_configuration` is added to support any type of dashboard widget.
`json_configuration` should be JSON string.

Please see the [Example](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/example/v0.12/dashboard.tf).

## Common required arguments

name | type | description
--- | --- | ---
type | string |
description | string |
dashboard_id | string |

## Common optional arguments

name | type | default | description
--- | --- | --- | ---
cache_time | int | |

## STREAM_SEARCH_RESULT_COUNT

```hcl
resource "graylog_dashboard_widget" "test" {
  description = "Stream search result count"
  dashboard_id = "5b6586000000000000000000"
  type = "STREAM_SEARCH_RESULT_COUNT"
  stream_search_result_count_configuration {
    timerange {
      type = "relative"
      range = 300
    }
    stream_id = "5b3983000000000000000000"
    query = ""
  }
  cache_time = 10
}
```

### Required arguments

name | type | description
--- | --- | ---
stream_search_result_count_configuration | object |
stream_search_result_count_configuration.timerange | object |

### Optional arguments

name | type | default | description
--- | --- | --- | ---
stream_search_result_count_configuration.stream_id | string | |
stream_search_result_count_configuration.query | string | |
stream_search_result_count_configuration.lower_is_better | bool | |
stream_search_result_count_configuration.trend | bool | |

## QUICKVALUES

```hcl
resource "graylog_dashboard_widget" "test" {
  description = "Quick values"
  dashboard_id = "5b6586000000000000000000"
  type = "QUICKVALUES"
  quick_values_configuration {
    timerange {
      type = "relative"
      range = 300
    }
    stream_id = "5b3983000000000000000000"
    query = ""
    field = "status"
    limit = 5
    sort_order = "desc"
    stacked_fields = ""
    data_table_limit = 50
  }
  cache_time = 10
}
```

### Required arguments

name | type | description
--- | --- | ---
quick_values_configuration | object |
quick_values_configuration.timerange | object |

### Optional arguments

name | type | default | description
--- | --- | --- | ---
quick_values_configuration.stream_id | string | |
quick_values_configuration.query | string | |
quick_values_configuration.interval | string | |
quick_values_configuration.field | string | |
quick_values_configuration.sort_order | string | |
quick_values_configuration.stacked_fields | string | |
quick_values_configuration.show_data_table | bool | |
quick_values_configuration.show_pie_chart | bool | |
quick_values_configuration.limit | int | |
quick_values_configuration.data_table_limit | int | |

## QUICKVALUES_HISTOGRAM

```hcl
resource "graylog_dashboard_widget" "test" {
  description = "Quick values"
  dashboard_id = "5b6586000000000000000000"
  type = "QUICKVALUES_HISTOGRAM"
  quick_values_histogram_configuration {
    timerange {
      type = "relative"
      range = 28800
    }
    stream_id = "5b3983000000000000000000"
    query = "status:200"
    field = "status"
    limit = 5
    sort_order = "desc"
    stacked_fields = ""
  }
  cache_time = 10
}
```

### Required arguments

name | type | description
--- | --- | ---
quick_values_histogram_configuration | object |
quick_values_histogram_configuration.timerange | object |

### Optional arguments

name | type | default | description
--- | --- | --- | ---
quick_values_histogram_configuration.stream_id | string | |
quick_values_histogram_configuration.query | string | |
quick_values_histogram_configuration.field | string | |
quick_values_histogram_configuration.sort_order | string | |
quick_values_histogram_configuration.stacked_fields | string | |
quick_values_histogram_configuration.limit | int | |

## SEARCH_RESULT_CHART

```hcl
resource "graylog_dashboard_widget" "test" {
  description = "Search result graph"
  dashboard_id = "5b6586000000000000000000"
  type = "SEARCH_RESULT_CHART"
  search_result_chart_configuration {
    timerange {
      type = "relative"
      range = 28800
    }
    query = "status:200"
    interval = "minute"
  }
  cache_time = 10
}
```

### Required arguments

name | type | description
--- | --- | ---
search_result_chart_configuration | object |
search_result_chart_configuration.timerange | object |

### Optional arguments

name | type | default | description
--- | --- | --- | ---
search_result_chart_configuration.query | string | |
search_result_chart_configuration.interval | string | |

## FIELD_CHART

```hcl
resource "graylog_dashboard_widget" "test" {
  description = "GET image success rate per day"
  dashboard_id = "5b6586000000000000000000"
  type = "FIELD_CHART"
  field_chart_configuration {
    timerange {
      type = "relative"
      range = 3600
    }
    query = "status:200"
    interval = "day"
    range_type = "relative"
    field = "status"
    relative = 3600
    valuetype = "mean"
    renderer = "line"
    interpolation = "linear"
  }
  cache_time = 10
}
```

### Required arguments

name | type | description
--- | --- | ---
field_chart_configuration | object |
field_chart_configuration.timerange | object |

### Optional arguments

name | type | default | description
--- | --- | --- | ---
field_chart_configuration.stream_id | string | |
field_chart_configuration.query | string | |
field_chart_configuration.interval | string | |
field_chart_configuration.field | string | |
field_chart_configuration.valuetype | string | |
field_chart_configuration.renderer | string | |
field_chart_configuration.interpolation | string | |
field_chart_configuration.range_type | string | |
field_chart_configuration.relative | int | |

## STATS_COUNT

```hcl
resource "graylog_dashboard_widget" "test" {
  description = "Statistical value"
  dashboard_id = "5b6586000000000000000000"
  type = "STATS_COUNT"
  stats_count_configuration {
    timerange {
      type = "relative"
      range = 300
    }
    query = ""
    field = "status"
    stats_function = "cardinality"
    stream_id = "5b3983000000000000000000"
  }
  cache_time = 10
}
```

### Required arguments

name | type | description
--- | --- | ---
stats_count_configuration | object |
stats_count_configuration.timerange | object |

### Optional arguments

name | type | default | description
--- | --- | --- | ---
stats_count_configuration.stream_id | string | |
stats_count_configuration.query | string | |
stats_count_configuration.field | string | |
stats_count_configuration.lower_is_better | bool | |
stats_count_configuration.trend | bool | |

## Attributes

name | type | description | etc
--- | --- | --- | ---
creator_user_id | string | | computed
