resource "graylog_stream" "test" {
  title       = "test"
  description = "test"

  index_set_id                       = "5d84bfbe2ab79c000d35d4a9"
  disabled                           = false
  matching_type                      = "AND"
  remove_matches_from_default_stream = true
}
