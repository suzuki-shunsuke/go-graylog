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

resource "graylog_extractor" "http_response_code" {
  input_id        = graylog_input.gelf_udp.id
  title           = "Apache http_response_code"
  type            = "regex"
  cursor_strategy = "copy"
  source_field    = "message"
  target_field    = "http_response_code"
  condition_type  = "regex"
  condition_value = "[1-5]\\d{2}"
  order           = 0
  
  converters {
    type = "numeric"
    config {}
  }

  regex_type_extractor_config {
    regex_value = "HTTP/1.[0-1]\" (\\d{3}) "
  }
}
