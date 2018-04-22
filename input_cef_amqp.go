package graylog

func (attrs *InputCEFAMQPAttrs) InputType() string {
	return INPUT_TYPE_CEF_AMQP
}

type InputCEFAMQPAttrs struct {
	Heartbeat              int    `json:"heartbeat,omitempty"`
	Exchange               string `json:"exchange,omitempty"`
	Timezone               string `json:"timezone,omitempty"`
	BrokerPassword         string `json:"broker_password,omitempty"`
	ParallelQueues         int    `json:"parallel_queues,omitempty"`
	Prefetch               int    `json:"prefetch,omitempty"`
	BrokerPort             int    `json:"broker_port,omitempty"`
	Locale                 string `json:"locale,omitempty"`
	ExchangeBind           bool   `json:"exchange_bind,omitempty"`
	RequeueInvalidMessages bool   `json:"requeue_invalid_messages,omitempty"`
	UseFullNames           bool   `json:"use_full_names,omitempty"`
	TLS                    bool   `json:"tls,omitempty"`
	BrokerHostname         string `json:"broker_hostname,omitempty"`
	ThrottlingAllowed      bool   `json:"throttling_allowed,omitempty"`
	Queue                  string `json:"queue,omitempty"`
	BrokerVHost            string `json:"broker_vhost,omitempty"`
	BrokerUsername         string `json:"broker_username,omitempty"`
	RoutingKey             string `json:"routing_key,omitempty"`
}
