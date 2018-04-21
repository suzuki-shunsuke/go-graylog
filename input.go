package graylog

const (
	INPUT_TYPE_SYSLOG_TCP            string = "org.graylog2.inputs.syslog.tcp.SyslogTCPInput"
	INPUT_TYPE_AWS_CLOUD_TRAIL_INPUT string = "org.graylog.aws.inputs.cloudtrail.CloudTrailInput"
	INPUT_TYPE_AWS_FLOW_LOGS         string = "org.graylog.aws.inputs.flowlogs.FlowLogsInput"
	INPUT_TYPE_AWS_LOGS              string = "org.graylog.aws.inputs.cloudwatch.CloudWatchLogsInput"
	INPUT_TYPE_BEATS                 string = "org.graylog.plugins.beats.BeatsInput"
)

var (
	// when update these fields variables, update also terraform graylog_input resource's document.
	InputAttributesIntFields []string = []string{
		"port", "recv_buffer_size", "heartbeat", "prefetch", "broker_port",
		"parallel_queues", "fetch_wait_max", "fetch_min_bytes", "threads",
		"max_message_size", "decompress_size_limit", "idle_writer_timeout",
		"max_chunk_size", "interval"}
	// when update these fields variables, update also terraform graylog_input resource's document.
	InputAttributesBoolFields []string = []string{
		"throttling_allowed", "tls_enable", "tcp_keepalive", "exchange_bind", "tls", "requeue_invalid_messages", "use_full_names", "use_null_delimiter", "enable_cors", "force_rdns", "store_full_message", "expand_structured_data", "allow_override_date"}
	// when update these fields variables, update also terraform graylog_input resource's document.
	InputAttributesStrFields []string = []string{
		"bind_address", "aws_region", "aws_assume_role_arn", "aws_access_key", "kinesis_stream_name", "aws_secret_key", "aws_sqs_region", "aws_s3_region", "aws_sqs_queue_name", "override_source", "tls_key_file", "tls_key_password", "tls_client_auth_cert_file", "tls_client_auth", "tls_cert_file", "timezone", "broker_vhost", "broker_username", "locale", "broker_password", "exchange", "routing_key", "broker_hostname", "queue", "topic_filter", "offset_reset", "zookeeper", "headers", "path", "target_url", "source", "timeunit", "netflow9_definitions_path"}
)

// InputAttributes represents Input's configuration.
type InputAttributes struct {
	// ex. 0.0.0.0
	BindAddress *string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	Port        *int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	// ex. 262144
	RecvBufferSize *int `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`

	// The AWS region the Kinesis stream is running in.
	// aws flow logs, aws logs
	// ex. us-east-1
	AwsRegion *string `json:"aws_region,omitempty"`
	// Role ARN with required permissions (cross account access)
	AwsAssumeRoleArn *string `json:"aws_assume_role_arn,omitempty"`
	// If enabled, no new messages will be read from this input until Graylog catches up with its message load.
	// This is typically useful for inputs reading from files or message queue systems like AMQP or Kafka.
	// If you regularly poll an external system, e.g. via HTTP, you normally want to leave this disabled.
	ThrottlingAllowed *bool `json:"throttling_allowed,omitempty"`
	// Access key of an AWS user with sufficient permissions. (See documentation)
	AwsAccessKey *string `json:"aws_access_key,omitempty"`
	// The name of the Kinesis stream that receives your messages.
	// See README for instructions on how to connect messages to a Kinesis Stream.
	KinesisStreamName *string `json:"kinesis_stream_name,omitempty"`
	// Secret key of an AWS user with sufficient permissions. (See documentation)
	AwsSecretKey *string `json:"aws_secret_key,omitempty"`

	// The AWS region the SQS queue is in.
	// ex. us-east-1
	AwsSQSRegion *string `json:"aws_sqs_region,omitempty"`
	// The AWS region the S3 bucket containing CloudTrail logs is in.
	// ex. us-east-1
	AwsS3Region *string `json:"aws_s3_region,omitempty"`
	// The SQS queue that SNS is writing CloudTrail notifications to.
	// ex. "cloudtrail-notifications",
	AwsSQSQueueName *string `json:"aws_sqs_queue_name,omitempty"`
	// The source is a hostname derived from the received packet by default.
	// Set this if you want to override it with a custom string.
	OverrideSource *string `json:"override_source,omitempty"`

	// beats
	// Path to the TLS private key file
	TLSKeyFile *string `json:"tls_key_file,omitempty"`
	// Accept TLS connections
	TLSEnable *bool `json:"tls_enable,omitempty"`
	// The password for the encrypted key file.
	TLSKeyPassword *string `json:"tls_key_password,omitempty"`
	// Enable TCP keepalive packets
	TCPKeepAlive *bool `json:"tcp_keepalive,omitempty"`
	// TLS Client Auth Trusted Certs (File or Directory)
	TLSClientAuthCertFile *string `json:"tls_client_auth_cert_file,omitempty"`
	// Whether clients need to authenticate themselves in a TLS connection
	// disabled, optional, required
	TLSClientAuth *string `json:"tls_client_auth,omitempty"`
	// Path to the TLS certificate file
	TLSCertFile *string `json:"tls_cert_file,omitempty"`

	// CEF Amqp Input
	// Heartbeat interval in seconds (use 0 to disable heartbeat)
	HeartBeat *int `json:"heartbeat,omitempty"`
	// Timezone of the timestamps in CEF messages.
	// Set this to the local timezone if in doubt.
	// Format example: "+01:00" or "America/Chicago"
	Timezone *string `json:"timezone,omitempty"`
	// For advanced usage: AMQP prefetch count. Default is 100.
	Prefetch *int `json:"prefetch,omitempty"`
	// Binds the queue to the configured exchange. The exchange must already exist.
	ExchangeBind *bool `json:"exchange_bind,omitempty"`
	// Virtual host of the AMQP broker to use
	BrokerVHost *string `json:"broker_vhost,omitempty"`
	// Username to connect to AMQP broker
	BrokerUsername *string `json:"broker_username,omitempty"`
	// Locale to use for parsing the timestamps of CEF messages.
	// Set this to english if in doubt. Format example: "en" or "en_US"<Paste>
	Locale *string `json:"locale,omitempty"`
	// Port of the AMQP broker to use
	BrokerPort *int `json:"broker_port,omitempty"`
	// Number of parallel Queues
	ParallelQueues *int `json:"parallel_queues,omitempty"`
	// Password to connect to AMQP broker
	BrokerPassword *string `json:"broker_password,omitempty"`
	// Name of exchange to bind to.
	Exchange *string `json:"exchange,omitempty"`
	// Enable transport encryption via TLS. (requires valid TLS port setting)
	TLS *bool `json:"tls,omitempty"`
	// Routing key to listen for.
	RoutingKey *string `json:"routing_key,omitempty"`
	// Invalid messages will be discarded if disabled.
	RequeueInvalidMessages *bool `json:"requeue_invalid_messages,omitempty"`
	// Hostname of the AMQP broker to use
	BrokerHostname *string `json:"broker_hostname,omitempty"`
	// Name of queue that is created.
	Queue *string `json:"queue,omitempty"`
	// Use full field names in CEF messages (as defined in the CEF specification)
	UseFullNames *bool `json:"use_full_names,omitempty"`

	// CEF Kafka
	// Every topic that matches this regular expression will be consumed.
	TopicFilter *string `json:"topic_filter,omitempty"`
	// Wait for this time or the configured minimum size of a message batch before fetching.
	FetchWaitMax *int `json:"fetch_wait_max,omitempty"`
	// What to do when there is no initial offset in ZooKeeper or if an offset is out of range
	OffsetReset *string `json:"offset_reset,omitempty"`
	// Host and port of the ZooKeeper that is managing your Kafka cluster.
	Zookeeper *string `json:"zookeeper,omitempty"`
	// Wait for a message batch to reach at least this size or the configured maximum wait time before fetching.
	FetchMinBytes *int `json:"fetch_min_bytes,omitempty"`
	// Number of processor threads to spawn. Use one thread per Kafka topic partition.
	Threads *int `json:"threads,omitempty"`

	// CEF Tcp Input
	UseNullDelimiter *bool `json:"use_null_delimiter,omitempty"`
	MaxMessageSize   *int  `json:"max_message_size,omitempty"`

	// GELF AMQP
	// The maximum number of bytes after decompression.
	DecompressSizeLimit *int `json:"decompress_size_limit,omitempty"`

	// GELF HTTP
	// The server closes the connection after the given time in seconds after the last client write request. (use 0 to disable)
	IdleWriterTimeOut *int `json:"idle_writer_timeout,omitempty"`
	// The maximum HTTP chunk size in bytes (e. g. length of HTTP request body)
	MaxChunkSize *int `json:"max_chunk_size,omitempty"`
	// Input sends CORS headers to satisfy browser security policies
	EnableCORS *bool `json:"enable_cors,omitempty"`

	// JSON path from HTTP API
	// Add a comma separated list of additional HTTP headers.
	// For example: Accept: application/json, X-Requester: Graylog
	Headers *string `json:"headers,omitempty"`
	// Path to the value you want to extract from the JSON response.
	// Take a look at the documentation for a more detailed explanation.
	Path *string `json:"path,omitempty"`
	// HTTP resource returning JSON on GET
	TargetURL *string `json:"target_url,omitempty"`
	// Time between every collector run.
	// Select a time unit in the corresponding dropdown.
	// Example: Run every 5 minutes.
	Interval *int `json:"interval,omitempty"`
	// What to use as source field of the resulting message.
	Source   *string `json:"source,omitempty"`
	Timeunit *string `json:"timeunit,omitempty"`

	// net flow udp
	// Path to the YAML file containing Netflow 9 field definitions
	NetFlow9DefinitionsPath *string `json:"netflow9_definitions_path,omitempty"`

	// syslog AMQP
	ForceRDNS            *bool `json:"force_rdns,omitempty"`
	StoreFullMessage     *bool `json:"store_full_message,omitempty"`
	ExpandStructuredData *bool `json:"expand_structured_data,omitempty"`

	// syslog kafka
	AllowOverrideDate *bool `json:"allow_override_date,omitempty"`
}

// Input represents Graylog Input.
type Input struct {
	// required
	// Select a name of your new input that describes it.
	Title string `json:"title,omitempty" v-create:"required" v-update:"required"`
	Type  string `json:"type,omitempty" v-create:"required" v-update:"required"`
	// https://github.com/Graylog2/graylog2-server/issues/3480
	Attributes *InputAttributes `json:"attributes,omitempty" v-create:"required" v-update:"required"`

	// ex. "5a90d5c2c006c60001efc368"
	ID string `json:"id,omitempty" v-create:"isdefault" v-update:"required,objectid"`

	// Should this input start on all nodes
	Global *bool `json:"global,omitempty"`
	// On which node should this input start
	// ex. "2ad6b340-3e5f-4a96-ae81-040cfb8b6024"
	Node string `json:"node,omitempty"`
	// ex. 2018-02-24T03:02:26.001Z
	CreatedAt string `json:"created_at,omitempty" v-create:"isdefault"`
	// ex. "admin"
	CreatorUserID string `json:"creator_user_id,omitempty" v-create:"isdefault"`
	// ContextPack `json:"context_pack,omitempty"`
	// StaticFields `json:"static_fields,omitempty"`
}

type InputsBody struct {
	Inputs []Input `json:"inputs"`
	Total  int     `json:"total"`
}
