resource "graylog_dashboard_widget" "http_response_codes" {
  dashboard_id       = "5de4fcf7a1de1800127e2fbc"
  description        = "updated description"
  type               = "STACKED_CHART"
  json_configuration = <<EOF
{
  "interval": "minute",
  "timerange": {
    "type": "relative",
    "range": 86400
  },
  "renderer": "area",
  "interpolation": "linear",
  "stream_id": "000000000000000000000003",
  "series": [
    {
      "query": "labels_app: nginx-ingress AND response:[200 TO 399]",
      "field": "response",
      "statistical_function": "count"
    },
    {
      "query": "labels_app: nginx-ingress AND response:[500 TO 599]",
      "field": "response",
      "statistical_function": "count"
    },
    {
      "query": "labels_app: nginx-ingress AND response:[400 TO 499]",
      "field": "response",
      "statistical_function": "count"
    }
  ]
}
EOF
}
