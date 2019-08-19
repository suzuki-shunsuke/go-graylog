resource "graylog_stream" "test" {
  title         = "test"
  index_set_id  = data.graylog_index_set.default.id
  disabled      = true
  matching_type = "AND"
  description   = "test"
}

# data "graylog_stream" "test" {
#   title = "test"
# }
