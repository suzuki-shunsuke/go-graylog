package graylog

func (attrs *InputAWSFlowLogsAttrs) InputType() string {
	return InputTypeAWSFlowLogs
}

type InputAWSFlowLogsAttrs struct {
	AWSRegion         string `json:"aws_region,omitempty"`
	AWSAssumeRoleArn  string `json:"aws_assume_role_arn,omitempty"`
	AWSAccessKey      string `json:"aws_access_key,omitempty"`
	AWSSecretKey      string `json:"aws_secret_key,omitempty"`
	KinesisStreamName string `json:"kinesis_stream_name,omitempty"`
	ThrottlingAllowed bool   `json:"throttling_allowed,omitempty"`
}
