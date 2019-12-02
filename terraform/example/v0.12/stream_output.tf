resource "graylog_stream_output" "stdout" {
  stream_id = graylog_stream.test.id
  output_ids = [
    graylog_output.stdout.id,
    graylog_output.gelf.id,
  ]
}
