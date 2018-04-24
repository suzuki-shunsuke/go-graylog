package graylog

// InputType is the implementation of the InputAttributes interface.
func (attrs *InputAWSLogsAttrs) InputType() string {
	return InputTypeAWSLogs
}

// InputAWSLogsAttrs represents AWS logs Input's attributes.
type InputAWSLogsAttrs struct {
	AWSRegion         string `json:"aws_region,omitempty"`
	AWSAssumeRoleArn  string `json:"aws_assume_role_arn,omitempty"`
	AWSAccessKey      string `json:"aws_access_key,omitempty"`
	AWSSecretKey      string `json:"aws_secret_key,omitempty"`
	KinesisStreamName string `json:"kinesis_stream_name,omitempty"`
	ThrottlingAllowed bool   `json:"throttling_allowed,omitempty"`
	OverrideSource    string `json:"override_source,omitempty"`
}
