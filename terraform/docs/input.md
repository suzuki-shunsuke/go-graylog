# graylog_input

https://github.com/suzuki-shunsuke/terraform-provider-graylog/blob/master/resource_input.go

```
resource "graylog_input" "test" {
  title = "terraform test"
  type = "org.graylog2.inputs.syslog.udp.SyslogUDPInput"
  configuration = {
    bind_address = "0.0.0.0"
    port = 514
    recv_buffer_size = 262144
  }
}
```

## Argument Reference

### Required Argument

name | type | description
--- | --- | ---
title | string |
type | string |
configuration | |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
global | "" | string |
node | "" | string |
configuration.bind_address | string |
configuration.port | int |
configuration.recv_buffer_size | int |
configuration.heartbeat | int |
configuration.prefetch | int |
configuration.broker_port | int |
configuration.parallel_queues | int |
configuration.fetch_wait_max | int |
configuration.fetch_min_bytes | int |
configuration.threads | int |
configuration.max_message_size | int |
configuration.decompress_size_limit | int |
configuration.idle_writer_timeout | int |
configuration.max_chunk_size | int |
configuration.interval | int |
configuration.throttling_allowed | bool |
configuration.tls_enable | bool |
configuration.tcp_keepalive | bool |
configuration.exchange_bind | bool |
configuration.tls | bool |
configuration.requeue_invalid_messages | bool |
configuration.use_full_names | bool |
configuration.use_null_delimiter | bool |
configuration.enable_cors | bool |
configuration.force_rdns | bool |
configuration.store_full_message | bool |
configuration.expand_structured_data | bool |
configuration.allow_override_date | bool |
configuration.aws_region | string |
configuration.aws_assume_role_arn | string |
configuration.aws_access_key | string |
configuration.kinesis_stream_name | string |
configuration.aws_secret_key | string |
configuration.aws_sqs_region | string |
configuration.aws_s3_region | string |
configuration.aws_sqs_queue_name | string |
configuration.override_source | string |
configuration.tls_key_file | string |
configuration.tls_key_password | string |
configuration.tls_client_auth | string |
configuration.tls_client_auth_cert_file | string |
configuration.tls_cert_file | string |
configuration.timezone | string |
configuration.broker_vhost | string |
configuration.broker_username | string |
configuration.locale | string |
configuration.broker_password | string |
configuration.exchange | string |
configuration.routing_key | string |
configuration.broker_hostname | string |
configuration.queue | string |
configuration.topic_filter | string |
configuration.offset_reset | string |
configuration.zookeeper | string |
configuration.headers | string |
configuration.path | string |
configuration.target_url | string |
configuration.source | string |
configuration.timeunit | string |
configuration.netflow9_definitions_path | string |

## Attributes Reference

name | type | etc
--- | --- | ---
input_id | string | computed
created_at | string | computed
creator_user_id | string | computed
