resource "graylog_pipeline" "test" {
  source = <<EOF
pipeline "test"
  stage 0 match either
end
EOF

  description = "test"
}

resource "graylog_pipeline_rule" "test" {
  source = <<EOF
rule "test"
when
    to_long($message.status) < 500
then
    set_field("status_01", 1);
end
EOF

  description = "test"
}

resource "graylog_pipeline_connection" "test" {
  stream_id    = graylog_stream.test.id
  pipeline_ids = [graylog_pipeline.test.id]
}
