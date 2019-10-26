# graylog_extractor

* https://docs.graylog.org/en/latest/pages/extractors.html
* [Example](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/example/v0.12/extractor.tf)
* [Source code](https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_extractor.go)

## How to import

Specify `<input id>/<extractor id>` as ID.

```console
$ terraform import graylog_extractor.test 5bb1b4b5c9e77bbbbbbbbbbb/5c4acaefc9e77bbbbbbbbbbb
```

## Argument Reference

### Common Required Argument

name | type | description
--- | --- | ---
input_id | string |
type | string |
title | string |
cursor_strategy | string |
source_field | string |
condition_type | string |
extractor_config | object{} |
converters[].type | string |
converters[].config | object{} |

### Common Optional Argument

name | type | default | description
--- | --- | --- | ---
converters | list | [] |
target_field | string | "" |
condition_value | string | "" |
order | int | 0 |
converters[].config.date_format | string | "" |
converters[].config.time_zone | string | "" |
converters[].config.locale | string | "" |

## type: grok 

### Required Argument

name | type | description
--- | --- | ---
grok_type_extractor_config | object |
grok_type_extractor_config.grok_pattern | string |

### Optional Argument

None.

## type: json

## Required Argument

name | type | description
--- | --- | ---
json_type_extractor_config | object |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
json_type_extractor_config.list_separator | | string |
json_type_extractor_config.kv_separator | | string |
json_type_extractor_config.key_prefix | | string |
json_type_extractor_config.key_separator | | string |
json_type_extractor_config.replace_key_whitespace | | bool |
json_type_extractor_config.key_whitespace_replacement | | string |

## type: regex

## Required Argument

name | type | description
--- | --- | ---
regex_type_extractor_config | object |
regex_type_extractor_config.regex_value | string |

### Optional Argument

None.

## other types

We provide some additional attributes.

* `general_int_extractor_config`
* `general_bool_extractor_config`
* `general_float_extractor_config`
* `general_string_extractor_config`

### Required Argument

None.

### Optional Argument

name | default | type | description
--- | --- | --- | ---
general_int_extractor_config | {} | map[string]int |
general_bool_extractor_config | {} | map[string]bool |
general_float_extractor_config | {} | map[string]float64 |
general_string_extractor_config | {} | map[string]string |
