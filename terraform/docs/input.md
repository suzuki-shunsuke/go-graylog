# graylog_input

https://github.com/suzuki-shunsuke/go-graylog/blob/master/terraform/graylog/resource_input.go

```hcl
resource "graylog_input" "test" {
  title = "terraform test"
  type = "org.graylog2.inputs.syslog.udp.SyslogUDPInput"
  attributes = {
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
attributes | |

### Optional Argument

name | default | type | description
--- | --- | --- | ---
global | false | bool |
node | "" | string |
attributes.bind_address | string |
attributes.port | int |
attributes.recv_buffer_size | int |
attributes.heartbeat | int |
attributes.prefetch | int |
attributes.broker_port | int |
attributes.parallel_queues | int |
attributes.fetch_wait_max | int |
attributes.fetch_min_bytes | int |
attributes.threads | int |
attributes.max_message_size | int |
attributes.decompress_size_limit | int |
attributes.idle_writer_timeout | int |
attributes.max_chunk_size | int |
attributes.interval | int |
attributes.throttling_allowed | bool |
attributes.tls_enable | bool |
attributes.tcp_keepalive | bool |
attributes.exchange_bind | bool |
attributes.tls | bool |
attributes.requeue_invalid_messages | bool |
attributes.use_full_names | bool |
attributes.use_null_delimiter | bool |
attributes.enable_cors | bool |
attributes.force_rdns | bool |
attributes.store_full_message | bool |
attributes.expand_structured_data | bool |
attributes.allow_override_date | bool |
attributes.aws_region | string |
attributes.aws_assume_role_arn | string |
attributes.aws_access_key | string |
attributes.kinesis_stream_name | string |
attributes.aws_secret_key | string |
attributes.aws_sqs_region | string |
attributes.aws_s3_region | string |
attributes.aws_sqs_queue_name | string |
attributes.override_source | string |
attributes.tls_key_file | string |
attributes.tls_key_password | string |
attributes.tls_client_auth | string |
attributes.tls_client_auth_cert_file | string |
attributes.tls_cert_file | string |
attributes.timezone | string |
attributes.broker_vhost | string |
attributes.broker_username | string |
attributes.locale | string |
attributes.broker_password | string |
attributes.exchange | string |
attributes.routing_key | string |
attributes.broker_hostname | string |
attributes.queue | string |
attributes.topic_filter | string |
attributes.offset_reset | string |
attributes.zookeeper | string |
attributes.headers | string |
attributes.path | string |
attributes.target_url | string |
attributes.source | string |
attributes.timeunit | string |
attributes.netflow9_definitions_path | string |

## Attrs Reference

name | type | etc
--- | --- | ---
input_id | string | computed
created_at | string | computed
creator_user_id | string | computed
