resource "graylog_stream_rule" "test" {
  field       = "tag"
  value       = "4"
  stream_id   = graylog_stream.test.id
  description = "test"
  type        = 1
  inverted    = false
}

