# data "graylog_dashboard" "test" {
#   title = "test"
# }

resource "graylog_dashboard" "test" {
  title       = "test"
  description = "test"
}

resource "graylog_dashboard_widget" "test" {
  description  = "Stream search result count change"
  # dashboard_id = data.graylog_dashboard.test.id
  dashboard_id = graylog_dashboard.test.id
  type         = "STREAM_SEARCH_RESULT_COUNT"
  stream_search_result_count_configuration {
    timerange {
      type  = "relative"
      range = 400
    }
    lower_is_better = true
    trend           = true
    stream_id       = graylog_stream.test.id
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
    stream_id        = graylog_stream.test.id
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
  dashboard_id = graylog_dashboard_widget.test2.dashboard_id
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
