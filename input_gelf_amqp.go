package graylog

const (
	// InputTypeGELFAMQP is one of input types.
	InputTypeGELFAMQP string = "org.graylog2.inputs.gelf.amqp.GELFAMQPInput"
)

// NewInputGELFAMQPAttrs is the constructor of InputGELFAMQPAttrs.
func NewInputGELFAMQPAttrs() InputAttrs {
	return &InputGELFAMQPAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputGELFAMQPAttrs) InputType() string {
	return InputTypeGELFAMQP
}

// InputGELFAMQPAttrs represents GELF AMQP Input's attributes.
type InputGELFAMQPAttrs struct {
	ExchangeBind           bool   `json:"exchange_bind"`
	ThrottlingAllowed      bool   `json:"throttling_allowed"`
	TLS                    bool   `json:"tls"`
	RequeueInvalidMessages bool   `json:"requeue_invalid_messages"`
	BrokerVHost            string `json:"broker_vhost,omitempty"`
	BrokerUsername         string `json:"broker_username,omitempty"`
	Queue                  string `json:"queue,omitempty"`
	RoutingKey             string `json:"routing_key,omitempty"`
	OverrideSource         string `json:"override_source,omitempty"`
	BrokerHostname         string `json:"broker_hostname,omitempty"`
	Exchange               string `json:"exchange,omitempty"`
	BrokerPassword         string `json:"broker_password,omitempty"`
	Prefetch               int    `json:"prefetch,omitempty"`
	Heartbeat              int    `json:"heartbeat,omitempty"`
	DecompressSizeLimit    int    `json:"decompress_size_limit,omitempty"`
	BrokerPort             int    `json:"broker_port,omitempty"`
	ParallelQueues         int    `json:"parallel_queues,omitempty"`
}
