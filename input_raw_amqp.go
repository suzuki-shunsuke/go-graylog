package graylog

const (
	InputTypeRawAMQP string = "org.graylog2.inputs.raw.amqp.RawAMQPInput"
)

// NewInputRawAMQPAttrs is the constructor of InputRawAMQPAttrs.
func NewInputRawAMQPAttrs() InputAttrs {
	return &InputRawAMQPAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputRawAMQPAttrs) InputType() string {
	return InputTypeRawAMQP
}

// InputRawAMQPAttrs represents raw AMQP Input's attributes.
type InputRawAMQPAttrs struct {
	ParallelQueues         int    `json:"parallel_queues,omitempty"`
	HeartBeat              int    `json:"heartbeat,omitempty"`
	BrokerPort             int    `json:"broker_port,omitempty"`
	Prefetch               int    `json:"prefetch,omitempty"`
	RequeueInvalidMessages bool   `json:"requeue_invalid_messages,omitempty"`
	TLS                    bool   `json:"tls,omitempty"`
	ExchangeBind           bool   `json:"exchange_bind,omitempty"`
	ThrottlingAllowed      bool   `json:"throttling_allowed,omitempty"`
	Exchange               string `json:"exchange,omitempty"`
	RoutingKey             string `json:"routing_key,omitempty"`
	BrokerHostname         string `json:"broker_hostname,omitempty"`
	Queue                  string `json:"queue,omitempty"`
	BrokerPassword         string `json:"broker_password,omitempty"`
	BrokerVHost            string `json:"broker_vhost,omitempty"`
	BrokerUsername         string `json:"broker_username,omitempty"`
	OverrideSource         string `json:"override_source,omitempty"`
}
