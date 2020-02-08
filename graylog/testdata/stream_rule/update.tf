resource "graylog_stream_rule" "test" {
  field       = "tag"
  value       = "4"
  stream_id   = "5d84c1a92ab79c000d35d6ca"
  description = "updated description"
  type        = 1
  inverted    = false
}
