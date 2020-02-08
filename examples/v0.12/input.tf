resource "graylog_input" "gelf_udp" {
  title  = "gelf udp 2"
  type   = "org.graylog2.inputs.gelf.udp.GELFUDPInput"
  global = "true"

  attributes {
    bind_address          = "0.0.0.0"
    port                  = 12201
    recv_buffer_size      = 262144
    decompress_size_limit = 8388608
  }
}

resource "graylog_input_static_fields" "gelf_udp" {
  input_id = graylog_input.gelf_udp.id
  fields = {
    foo = "bar"
  }
}

resource "graylog_input" "json_path" {
  title  = "json path"
  type   = "org.graylog2.inputs.misc.jsonpath.JsonPathInput"
  global = "true"

  attributes {
    interval           = 30
    path               = "$.userId"
    throttling_allowed = true
    target_url         = "http://jsonplaceholder.typicode.com/posts/1"
    source             = "id"
    timeunit           = "SECONDS"
  }
}

resource "graylog_input" "kube_events" {
  title = "Kube Events Input"
  type  = "org.graylog2.inputs.raw.tcp.RawTCPInput"

  global = true

  attributes {
    bind_address = "0.0.0.0"
    port         = 5555
  }
}
