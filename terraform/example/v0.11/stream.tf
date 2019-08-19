resource "graylog_stream" "test" {
  title         = "test"
  index_set_id  = "${graylog_index_set.test.id}"
  disabled      = true
  matching_type = "AND"
  description   = "test"
}
