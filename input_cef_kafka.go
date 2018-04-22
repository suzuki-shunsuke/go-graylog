package graylog

func (attrs *InputCEFKafkaAttrs) InputType() string {
	return INPUT_TYPE_CEF_KAFKA
}

type InputCEFKafkaAttrs struct {
	ThrottlingAllowed bool   `json:"throttling_allowed,omitempty"`
	Locale            string `json:"locale,omitempty"`
	Threads           int    `json:"threads,omitempty"`
	Zookeeper         string `json:"zookeeper,omitempty"`
	Timezone          string `json:"timezone,omitempty"`
	UseFullNames      bool   `json:"use_full_names,omitempty"`
	TopicFilter       string `json:"topic_filter,omitempty"`
	FetchWaitMax      int    `json:"fetch_wait_max,omitempty"`
	FetchMinBytes     int    `json:"fetch_min_bytes,omitempty"`
	OffsetReset       string `json:"offset_reset,omitempty"`
}
