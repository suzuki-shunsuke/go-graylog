package graylog

func (attrs *InputCloudTrailAttrs) InputType() string {
	return INPUT_TYPE_AWS_CLOUD_TRAIL_INPUT
}

type InputCloudTrailAttrs struct {
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
