package graylog

func (attrs *InputGELFKafkaAttrs) InputType() string {
	return INPUT_TYPE_GELF_KAFKA
}

type InputGELFKafkaAttrs struct {
	OverrideSource      string `json:"override_source,omitempty"`
	DecompressSizeLimit int    `json:"decompress_size_limit,omitempty"`
	TopicFilter         string `json:"topic_filter,omitempty"`
	ThrottlingAllowed   bool   `json:"throttling_allowed,omitempty"`
	FetchWaitMax        int    `json:"fetch_wait_max,omitempty"`
	FetchMinBytes       int    `json:"fetch_min_bytes,omitempty"`
	OffsetReset         string `json:"offset_reset,omitempty"`
	Threads             int    `json:"threads,omitempty"`
	Zookeeper           string `json:"zookeeper,omitempty"`
}