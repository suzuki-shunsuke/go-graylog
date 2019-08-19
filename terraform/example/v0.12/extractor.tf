resource "graylog_extractor" "test_grok" {
  input_id        = graylog_input.gelf_udp.id
  title           = "test_grok"
  type            = "grok"
  cursor_strategy = "copy"
  source_field    = "message"
  target_field    = "none"
  condition_type  = "none"
  condition_value = ""
  order           = 0

  grok_type_extractor_config {
    grok_pattern = "%%%{DATA}"
  }
}

resource "graylog_extractor" "test_json" {
  input_id        = graylog_input.gelf_udp.id
  title           = "test"
  type            = "json"
  cursor_strategy = "copy"
  source_field    = "message"
  target_field    = "none"
  condition_type  = "none"
  condition_value = ""
  order           = 0

  json_type_extractor_config {
    list_separator             = ", "
    kv_separator               = "="
    key_prefix                 = "visit_"
    key_separator              = "_"
    replace_key_whitespace     = false
    key_whitespace_replacement = "_"
  }
}

resource "graylog_extractor" "test_regex" {
  input_id        = graylog_input.gelf_udp.id
  title           = "test_regex"
  type            = "regex"
  cursor_strategy = "copy"

  source_field    = "message"
  condition_type  = "none"
  condition_value = ""
  order           = 0

  regex_type_extractor_config {
    regex_value = ".*"
  }

  converters {
    type = "date"

    config {
      date_format = "yyyy/MM/ddTHH:mm:ss"
      time_zone   = "Japan"
      locale      = "en"
    }
  }
}
