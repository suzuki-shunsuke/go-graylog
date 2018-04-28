package graylog

const (
	InputTypeSyslogKafka string = "org.graylog2.inputs.syslog.kafka.SyslogKafkaInput"
)

// NewInputSyslogKafkaAttrs is the constructor of InputSyslogKafkaAttrs.
func NewInputSyslogKafkaAttrs() InputAttributes {
	return &InputSyslogKafkaAttrs{}
}

// InputType is the implementation of the InputAttributes interface.
func (attrs InputSyslogKafkaAttrs) InputType() string {
	return InputTypeSyslogKafka
}

// InputSyslogKafkaAttrs represents SyslogKafka Input's attributes.
type InputSyslogKafkaAttrs struct {
	ForceRDNS            bool   `json:"force_rdns,omitempty"`
	StoreFullMessage     bool   `json:"store_full_message,omitempty"`
	ExpandStructuredData bool   `json:"expand_structured_data,omitempty"`
	AllowOverrideDate    bool   `json:"allow_override_date,omitempty"`
	ThrottlingAllowed    bool   `json:"throttling_allowed,omitempty"`
	OverrideSource       string `json:"override_source,omitempty"`
	TopicFilter          string `json:"topic_filter,omitempty"`
	FetchWaitMax         int    `json:"fetch_wait_max,omitempty"`
	OffsetReset          string `json:"offset_reset,omitempty"`
	Zookeeper            string `json:"zookeeper,omitempty"`
	FetchMinBytes        int    `json:"fetch_min_bytes,omitempty"`
	Threads              int    `json:"threads,omitempty"`
}
