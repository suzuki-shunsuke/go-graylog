package graylog

const (
	// InputTypeAWSCloudTrail is one of input types.
	InputTypeAWSCloudTrail string = "org.graylog.aws.inputs.cloudtrail.CloudTrailInput"
)

// NewInputAWSCloudTrailAttrs is the constructor of InputAWSCloudTrailAttrs.
func NewInputAWSCloudTrailAttrs() InputAttrs {
	return &InputAWSCloudTrailAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputAWSCloudTrailAttrs) InputType() string {
	return InputTypeAWSCloudTrail
}

// InputAWSCloudTrailAttrs represents aws cloud trail Input's attributes.
type InputAWSCloudTrailAttrs struct {
	CreatorUserID     string `json:"creator_user_id,omitempty" v-create:"isdefault"`
	AWSAssumeRoleArn  string `json:"aws_assume_role_arn,omitempty"`
	AWSAccessKey      string `json:"aws_access_key,omitempty"`
	AWSSecretKey      string `json:"aws_secret_key,omitempty"`
	AWSSQSRegion      string `json:"aws_sqs_region,omitempty"`
	AWSSQSQueueName   string `json:"aws_sqs_queue_name,omitempty"`
	AWSS3Region       string `json:"aws_s3_region,omitempty"`
	ThrottlingAllowed bool   `json:"throttling_allowed,omitempty"`
	OverrideSource    string `json:"override_source,omitempty"`
}
