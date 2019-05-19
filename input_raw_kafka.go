package graylog

const (
	// InputTypeRawKafka is one of input types.
	InputTypeRawKafka string = "org.graylog2.inputs.raw.kafka.RawKafkaInput"
)

// NewInputRawKafkaAttrs is the constructor of InputRawKafkaAttrs.
func NewInputRawKafkaAttrs() InputAttrs {
	return &InputRawKafkaAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputRawKafkaAttrs) InputType() string {
	return InputTypeRawKafka
}

// InputRawKafkaAttrs represents RawKafka Input's attributes.
type InputRawKafkaAttrs struct {
	TopicFilter       string `json:"topic_filter,omitempty"`
	FetchWaitMax      int    `json:"fetch_wait_max,omitempty"`
	OffsetReset       string `json:"offset_reset,omitempty"`
	Zookeeper         string `json:"zookeeper,omitempty"`
	ThrottlingAllowed bool   `json:"throttling_allowed,omitempty"`
	FetchMinBytes     int    `json:"fetch_min_bytes,omitempty"`
	Threads           int    `json:"threads,omitempty"`
	OverrideSource    string `json:"override_source,omitempty"`
}
