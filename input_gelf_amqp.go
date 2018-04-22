package graylog

// InputType is the implementation of the InputAttributes interface.
func (attrs *InputGELFAMQPAttrs) InputType() string {
	return InputTypeGELFAMQP
}

// InputGELFAMQPAttrs
type InputGELFAMQPAttrs struct {
	ExchangeBind           bool   `json:"exchange_bind,omitempty"`
	Prefetch               int    `json:"prefetch,omitempty"`
	Heartbeat              int    `json:"heartbeat,omitempty"`
	BrokerVHost            string `json:"broker_vhost,omitempty"`
	BrokerUsername         string `json:"broker_username,omitempty"`
	DecompressSizeLimit    int    `json:"decompress_size_limit,omitempty"`
	Queue                  string `json:"queue,omitempty"`
	RoutingKey             string `json:"routing_key,omitempty"`
	ThrottlingAllowed      bool   `json:"throttling_allowed,omitempty"`
	TLS                    bool   `json:"tls,omitempty"`
	OverrideSource         string `json:"override_source,omitempty"`
	BrokerHostname         string `json:"broker_hostname,omitempty"`
	RequeueInvalidMessages bool   `json:"requeue_invalid_messages,omitempty"`
	Exchange               string `json:"exchange,omitempty"`
	BrokerPort             int    `json:"broker_port,omitempty"`
	ParallelQueues         int    `json:"parallel_queues,omitempty"`
	BrokerPassword         string `json:"broker_password,omitempty"`
}
