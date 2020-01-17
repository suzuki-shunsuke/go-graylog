resource "graylog_stream" "test" {
  title       = "updated title"
  description = "updated description"

  index_set_id                       = "5d84bfbe2ab79c000d35d4a9"
  disabled                           = true
  matching_type                      = "AND"
  remove_matches_from_default_stream = true
}
