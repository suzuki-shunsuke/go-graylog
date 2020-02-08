package graylog

const (
	// InputTypeGELFKafka is one of input types.
	InputTypeGELFKafka string = "org.graylog2.inputs.gelf.kafka.GELFKafkaInput"
)

// NewInputGELFKafkaAttrs is the constructor of InputGELFKafkaAttrs.
func NewInputGELFKafkaAttrs() InputAttrs {
	return &InputGELFKafkaAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputGELFKafkaAttrs) InputType() string {
	return InputTypeGELFKafka
}

// InputGELFKafkaAttrs represents GELF Kafka Input's attributes.
type InputGELFKafkaAttrs struct {
	OverrideSource      string `json:"override_source,omitempty"`
	DecompressSizeLimit int    `json:"decompress_size_limit,omitempty"`
	TopicFilter         string `json:"topic_filter,omitempty"`
	ThrottlingAllowed   bool   `json:"throttling_allowed"`
	FetchWaitMax        int    `json:"fetch_wait_max,omitempty"`
	FetchMinBytes       int    `json:"fetch_min_bytes,omitempty"`
	OffsetReset         string `json:"offset_reset,omitempty"`
	Threads             int    `json:"threads,omitempty"`
	Zookeeper           string `json:"zookeeper,omitempty"`
}
