package graylog

const (
	INPUT_TYPE_SYSLOG_TCP            string = "org.graylog2.inputs.syslog.tcp.SyslogTCPInput"
	INPUT_TYPE_AWS_CLOUD_TRAIL_INPUT string = "org.graylog.aws.inputs.cloudtrail.CloudTrailInput"
	INPUT_TYPE_AWS_FLOW_LOGS         string = "org.graylog.aws.inputs.flowlogs.FlowLogsInput"
	INPUT_TYPE_AWS_LOGS              string = "org.graylog.aws.inputs.cloudwatch.CloudWatchLogsInput"
	INPUT_TYPE_BEATS                 string = "org.graylog.plugins.beats.BeatsInput"
)

// InputConfiguration represents Input's configuration.
type InputConfiguration struct {
	// ex. 0.0.0.0
	BindAddress string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	Port        int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	// ex. 262144
	RecvBufferSize int `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`

	// The AWS region the Kinesis stream is running in.
	// aws flow logs, aws logs
	// ex. us-east-1
	AwsRegion string `json:"aws_region,omitempty"`
	// Role ARN with required permissions (cross account access)
	AwsAssumeRoleArn string `json:"aws_assume_role_arn"`
	// If enabled, no new messages will be read from this input until Graylog catches up with its message load.
	// This is typically useful for inputs reading from files or message queue systems like AMQP or Kafka.
	// If you regularly poll an external system, e.g. via HTTP, you normally want to leave this disabled.
	ThrottlingAllowed bool `json:"throttling_allowed"`
	// Access key of an AWS user with sufficient permissions. (See documentation)
	AwsAccessKey string `json:"aws_access_key,omitempty"`
	// The name of the Kinesis stream that receives your messages.
	// See README for instructions on how to connect messages to a Kinesis Stream.
	KinesisStreamName string `json:"kinesis_stream_name,omitempty"`
	// Secret key of an AWS user with sufficient permissions. (See documentation)
	AwsSecretKey string `json:"aws_secret_key,omitempty"`

	// The AWS region the SQS queue is in.
	// ex. us-east-1
	AwsSqsRegion string `json:"aws_sqs_region,omitempty"`
	// The AWS region the S3 bucket containing CloudTrail logs is in.
	// ex. us-east-1
	AwsS3Region string `json:"aws_s3_region,omitempty"`
	// The SQS queue that SNS is writing CloudTrail notifications to.
	// ex. "cloudtrail-notifications",
	AwsSqsQueueName string `json:"aws_sqs_queue_name"`
	// The source is a hostname derived from the received packet by default.
	// Set this if you want to override it with a custom string.
	OverrideSource string `json:"override_source"`

	// beats
	// Path to the TLS private key file
	TLSKeyFile string `json:"tls_key_file"`
	// Accept TLS connections
	TLSEnable bool `json:"tls_enable"`
	// The password for the encrypted key file.
	TLSKeyPassword string `json:"tls_key_password"`
	// Enable TCP keepalive packets
	TCPKeepAlive bool `json:"tcp_keepalive"`
	// TLS Client Auth Trusted Certs (File or Directory)
	TLSClientAuthCertFile string `json:"tls_client_auth_cert_file"`
	// Whether clients need to authenticate themselves in a TLS connection
	// disabled, optional, required
	TLSClientAuth string `json:"tls_client_auth"`
	// Path to the TLS certificate file
	TLSCertFile string `json:"tls_cert_file"`

	// CEF Amqp Input
	// Heartbeat interval in seconds (use 0 to disable heartbeat)
	HeartBeat int `json:"heartbeat"`
	// Timezone of the timestamps in CEF messages.
	// Set this to the local timezone if in doubt.
	// Format example: "+01:00" or "America/Chicago"
	Timezone string `json:"timezone"`
	// For advanced usage: AMQP prefetch count. Default is 100.
	Prefetch int `json:"prefetch"`
	// Binds the queue to the configured exchange. The exchange must already exist.
	ExchangeBind bool `json:"exchange_bind"`
	// Virtual host of the AMQP broker to use
	BrokerVHost string `json:"broker_vhost"`
	// Username to connect to AMQP broker
	BrokerUsername string `json:"broker_username"`
	// Locale to use for parsing the timestamps of CEF messages.
	// Set this to english if in doubt. Format example: "en" or "en_US"<Paste>
	Locale string `json:"locale"`
	// Port of the AMQP broker to use
	BrokerPort int `json:"broker_port"`
	// Number of parallel Queues
	ParallelQueues int `json:"parallel_queues"`
	// Password to connect to AMQP broker
	BrokerPassword string `json:"broker_password"`
	// Name of exchange to bind to.
	Exchange string `json:"exchange"`
	// Enable transport encryption via TLS. (requires valid TLS port setting)
	TLS bool `json:"tls"`
	// Routing key to listen for.
	RoutingKey string `json:"routing_key"`
	// Invalid messages will be discarded if disabled.
	RequeueInvalidMessages bool `json:"requeue_invalid_messages"`
	// Hostname of the AMQP broker to use
	BrokerHostname string `json:"broker_hostname"`
	// Name of queue that is created.
	Queue string `json:"queue"`
	// Use full field names in CEF messages (as defined in the CEF specification)
	UseFullNames bool `json:"use_full_names"`

	// CEF Kafka
	// Every topic that matches this regular expression will be consumed.
	TopicFilter string `json:"topic_filter"`
	// Wait for this time or the configured minimum size of a message batch before fetching.
	FetchWaitMax int `json:"fetch_wait_max"`
	// What to do when there is no initial offset in ZooKeeper or if an offset is out of range
	OffsetReset string `json:"offset_reset"`
	// Host and port of the ZooKeeper that is managing your Kafka cluster.
	Zookeeper string `json:"zookeeper"`
	// Wait for a message batch to reach at least this size or the configured maximum wait time before fetching.
	FetchMinBytes int `json:"fetch_min_bytes"`
	// Number of processor threads to spawn. Use one thread per Kafka topic partition.
	Threads int `json:"threads"`

	// CEF Tcp Input
	UseNullDelimiter bool `json:"use_null_delimiter"`
	MaxMessageSize   int  `json:"max_message_size"`

	// GELF AMQP
	// The maximum number of bytes after decompression.
	DecompressSizeLimit int `json:"decompress_size_limit"`

	// GELF HTTP
	// The server closes the connection after the given time in seconds after the last client write request. (use 0 to disable)
	IdleWriterTimeOut int `json:"idle_writer_timeout"`
	// The maximum HTTP chunk size in bytes (e. g. length of HTTP request body)
	MaxChunkSize int `json:"max_chunk_size"`
	// Input sends CORS headers to satisfy browser security policies
	EnableCORS bool `json:"enable_cors"`

	// JSON path from HTTP API
	// Add a comma separated list of additional HTTP headers.
	// For example: Accept: application/json, X-Requester: Graylog
	Headers string `json:"headers"`
	// Path to the value you want to extract from the JSON response.
	// Take a look at the documentation for a more detailed explanation.
	Path string `json:"path"`
	// HTTP resource returning JSON on GET
	TargetURL string `json:"target_url"`
	// Time between every collector run.
	// Select a time unit in the corresponding dropdown.
	// Example: Run every 5 minutes.
	Interval int `json:"interval"`
	// What to use as source field of the resulting message.
	Source   string `json:"source"`
	Timeunit string `json:"timeunit"`

	// net flow udp
	// Path to the YAML file containing Netflow 9 field definitions
	NetFlow9DefinitionsPath string `json:"netflow9_definitions_Path"`

	// syslog AMQP
	ForceRDNS            bool `json:"force_rdns"`
	StoreFullMessage     bool `json:"store_full_message"`
	ExpandStructuredData bool `json:"expand_structured_data"`

	// syslog kafka
	AllowOverrideDate bool `json:"allow_override_date"`
}

// Input represents Graylog Input.
type Input struct {
	// required
	// Select a name of your new input that describes it.
	Title         string              `json:"title,omitempty" v-create:"required" v-update:"required"`
	Type          string              `json:"type,omitempty" v-create:"required" v-update:"required"`
	Configuration *InputConfiguration `json:"configuration,omitempty" v-create:"required" v-update:"required"`

	// ex. "5a90d5c2c006c60001efc368"
	ID string `json:"id,omitempty" v-create:"isdefault" v-update:"required"`

	// Should this input start on all nodes
	Global bool `json:"global,omitempty"`
	// On which node should this input start
	// ex. "2ad6b340-3e5f-4a96-ae81-040cfb8b6024"
	Node string `json:"node,omitempty"`
	// ex. 2018-02-24T03:02:26.001Z
	CreatedAt string `json:"created_at,omitempty" v-create:"isdefault" v-update:"isdefault"`
	// ex. "admin"
	CreatorUserID string `json:"creator_user_id,omitempty" v-create:"isdefault" v-update:"isdefault"`
	// ContextPack `json:"context_pack,omitempty"`
	// StaticFields `json:"static_fields,omitempty"`
}

type InputsBody struct {
	Inputs []Input `json:"inputs"`
	Total  int     `json:"total"`
}
