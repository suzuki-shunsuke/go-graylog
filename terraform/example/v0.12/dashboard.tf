resource "graylog_dashboard" "test" {
  title       = "test"
  description = "test"
}

resource "graylog_dashboard_widget" "test" {
  description  = "Stream search result count change"
  dashboard_id = graylog_dashboard.test.id
  type         = "STREAM_SEARCH_RESULT_COUNT"
  stream_search_result_count_configuration {
    timerange {
      type  = "relative"
      range = 400
    }
    lower_is_better = true
    trend           = true
    stream_id       = "000000000000000000000001"
    query           = ""
  }
  cache_time = 10
}

resource "graylog_dashboard_widget" "test2" {
  description  = "Quick values"
  dashboard_id = graylog_dashboard.test.id
  type         = "QUICKVALUES"
  quick_values_configuration {
    timerange {
      type  = "relative"
      range = 300
    }
    stream_id        = "000000000000000000000001"
    query            = ""
    field            = "status"
    show_data_table  = true
    show_pie_chart   = true
    limit            = 5
    sort_order       = "desc"
    stacked_fields   = ""
    data_table_limit = 60
  }
  cache_time = 10
}

resource "graylog_dashboard_widget_positions" "test" {
  dashboard_id = graylog_dashboard_widget.test.dashboard_id
  positions {
    widget_id = graylog_dashboard_widget.test.id
    row       = 0
    col       = 0
    height    = 1
    width     = 1
  }
  positions {
    widget_id = graylog_dashboard_widget.test2.id
    row       = 0
    col       = 1
    height    = 2
    width     = 2
  }
}

# resource "graylog_dashboard_widget" "test4" {
#   description  = "Quick values"
#   dashboard_id = "5d1aa378d497d2000e79e62f"
#   type         = "QUICKVALUES_HISTOGRAM"
#   quick_values_histogram_configuration {
#     timerange {
#       type  = "relative"
#       range = 28800
#     }
#     stream_id      = "000000000000000000000001"
#     query          = "status:200"
#     field          = "status"
#     limit          = 5
#     sort_order     = "desc"
#     stacked_fields = ""
#   }
#   cache_time = 10
# }
# 
# resource "graylog_dashboard_widget" "test5" {
#   description  = "Search result graph"
#   dashboard_id = "5d1aa378d497d2000e79e62f"
#   type         = "SEARCH_RESULT_CHART"
#   search_result_chart_configuration {
#     timerange {
#       type  = "relative"
#       range = 28800
#     }
#     query    = "status:200"
#     interval = "minute"
#   }
#   cache_time = 10
# }
# 
# resource "graylog_dashboard_widget" "test6" {
#   description  = "GET image success rate per day"
#   dashboard_id = "5d1aa378d497d2000e79e62f"
#   type         = "FIELD_CHART"
#   field_chart_configuration {
#     timerange {
#       type  = "relative"
#       range = 3600
#     }
#     query         = "status:200"
#     interval      = "day"
#     range_type    = "relative"
#     field         = "status"
#     relative      = 3600
#     valuetype     = "mean"
#     renderer      = "line"
#     interpolation = "linear"
#   }
#   cache_time = 10
# }
# 
# resource "graylog_dashboard_widget" "test7" {
#   description  = "Statistical value"
#   dashboard_id = "5d1aa378d497d2000e79e62f"
#   type         = "STATS_COUNT"
#   stats_count_configuration {
#     timerange {
#       type  = "relative"
#       range = 300
#     }
#     query           = ""
#     field           = "status"
#     lower_is_better = false
#     trend           = false
#     stats_function  = "cardinality"
#     stream_id       = graylog_stream.test-terraform.id
#   }
#   cache_time = 10
# }
