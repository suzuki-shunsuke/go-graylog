package graylog

const (
	InputTypeCEFKafka string = "org.graylog.plugins.cef.input.CEFKafkaInput"
)

// NewInputCEFKafkaAttrs is the constructor of InputCEFKafkaAttrs.
func NewInputCEFKafkaAttrs() InputAttrs {
	return &InputCEFKafkaAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputCEFKafkaAttrs) InputType() string {
	return InputTypeCEFKafka
}

// InputCEFKafkaAttrs represents CEF Kafka Input's attributes.
type InputCEFKafkaAttrs struct {
	ThrottlingAllowed bool   `json:"throttling_allowed,omitempty"`
	UseFullNames      bool   `json:"use_full_names,omitempty"`
	Locale            string `json:"locale,omitempty"`
	Zookeeper         string `json:"zookeeper,omitempty"`
	Timezone          string `json:"timezone,omitempty"`
	TopicFilter       string `json:"topic_filter,omitempty"`
	OffsetReset       string `json:"offset_reset,omitempty"`
	Threads           int    `json:"threads,omitempty"`
	FetchWaitMax      int    `json:"fetch_wait_max,omitempty"`
	FetchMinBytes     int    `json:"fetch_min_bytes,omitempty"`
}
