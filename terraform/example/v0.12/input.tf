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
