package graylog

const (
	// InputTypeAWSFlowLogs is one of input types.
	InputTypeAWSFlowLogs string = "org.graylog.aws.inputs.flowlogs.FlowLogsInput"
)

// NewInputAWSFlowLogsAttrs is the constructor of InputAWSFlowLogsAttrs.
func NewInputAWSFlowLogsAttrs() InputAttrs {
	return &InputAWSFlowLogsAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputAWSFlowLogsAttrs) InputType() string {
	return InputTypeAWSFlowLogs
}

// InputAWSFlowLogsAttrs represents AWS flow logs Input's attributes.
type InputAWSFlowLogsAttrs struct {
	AWSRegion         string `json:"aws_region,omitempty"`
	AWSAssumeRoleArn  string `json:"aws_assume_role_arn,omitempty"`
	AWSAccessKey      string `json:"aws_access_key,omitempty"`
	AWSSecretKey      string `json:"aws_secret_key,omitempty"`
	KinesisStreamName string `json:"kinesis_stream_name,omitempty"`
	ThrottlingAllowed bool   `json:"throttling_allowed,omitempty"`
}
