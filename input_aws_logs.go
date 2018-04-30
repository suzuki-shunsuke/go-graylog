package graylog

const (
	// InputTypeAWSCloudWatchLogs is one of input types.
	InputTypeAWSCloudWatchLogs string = "org.graylog.aws.inputs.cloudwatch.CloudWatchLogsInput"
)

// NewInputAWSCloudWatchLogsAttrs is the constructor of InputAWSCloudWatchLogsAttrs.
func NewInputAWSCloudWatchLogsAttrs() InputAttrs {
	return &InputAWSCloudWatchLogsAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputAWSCloudWatchLogsAttrs) InputType() string {
	return InputTypeAWSCloudWatchLogs
}

// InputAWSCloudWatchLogsAttrs represents AWS logs Input's attributes.
type InputAWSCloudWatchLogsAttrs struct {
	AWSRegion         string `json:"aws_region,omitempty"`
	AWSAssumeRoleArn  string `json:"aws_assume_role_arn,omitempty"`
	AWSAccessKey      string `json:"aws_access_key,omitempty"`
	AWSSecretKey      string `json:"aws_secret_key,omitempty"`
	KinesisStreamName string `json:"kinesis_stream_name,omitempty"`
	ThrottlingAllowed bool   `json:"throttling_allowed,omitempty"`
	OverrideSource    string `json:"override_source,omitempty"`
}
