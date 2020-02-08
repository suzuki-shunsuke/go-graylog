package graylog

const (
	// InputTypeCEFAMQP is one of input types.
	InputTypeCEFAMQP string = "org.graylog.plugins.cef.input.CEFAmqpInput"
)

// NewInputCEFAMQPAttrs is the constructor of InputCEFAMQPAttrs.
func NewInputCEFAMQPAttrs() InputAttrs {
	return &InputCEFAMQPAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputCEFAMQPAttrs) InputType() string {
	return InputTypeCEFAMQP
}

// InputCEFAMQPAttrs represents CEF AMQP Input's attributes.
type InputCEFAMQPAttrs struct {
	Exchange               string `json:"exchange,omitempty"`
	Timezone               string `json:"timezone,omitempty"`
	BrokerPassword         string `json:"broker_password,omitempty"`
	Locale                 string `json:"locale,omitempty"`
	BrokerHostname         string `json:"broker_hostname,omitempty"`
	Queue                  string `json:"queue,omitempty"`
	BrokerVHost            string `json:"broker_vhost,omitempty"`
	BrokerUsername         string `json:"broker_username,omitempty"`
	RoutingKey             string `json:"routing_key,omitempty"`
	Heartbeat              int    `json:"heartbeat,omitempty"`
	ParallelQueues         int    `json:"parallel_queues,omitempty"`
	Prefetch               int    `json:"prefetch,omitempty"`
	BrokerPort             int    `json:"broker_port,omitempty"`
	ExchangeBind           bool   `json:"exchange_bind"`
	RequeueInvalidMessages bool   `json:"requeue_invalid_messages"`
	UseFullNames           bool   `json:"use_full_names"`
	TLS                    bool   `json:"tls"`
	ThrottlingAllowed      bool   `json:"throttling_allowed"`
}
