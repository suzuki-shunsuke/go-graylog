package graylog

func (attrs *InputRawAMQPAttrs) InputType() string {
	return INPUT_TYPE_RAW_AMQP
}

type InputRawAMQPAttrs struct {
	ParallelQueues         int    `json:"parallel_queues,omitempty"`
	Exchange               string `json:"exchange,omitempty"`
	RequeueInvalidMessages bool   `json:"requeue_invalid_messages,omitempty"`
	TLS                    bool   `json:"tls,omitempty"`
	RoutingKey             string `json:"routing_key,omitempty"`
	ExchangeBind           bool   `json:"exchange_bind,omitempty"`
	HeartBeat              int    `json:"heartbeat,omitempty"`
	BrokerPort             int    `json:"broker_port,omitempty"`
	BrokerHostname         string `json:"broker_hostname,omitempty"`
	Queue                  string `json:"queue,omitempty"`
	BrokerPassword         string `json:"broker_password,omitempty"`
	BrokerVHost            string `json:"broker_vhost,omitempty"`
	BrokerUsername         string `json:"broker_username,omitempty"`
	Prefetch               int    `json:"prefetch,omitempty"`
	OverrideSource         string `json:"override_source,omitempty"`
	ThrottlingAllowed      bool   `json:"throttling_allowed,omitempty"`
}
