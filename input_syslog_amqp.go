package graylog

const (
	InputTypeSyslogAMQP string = "org.graylog2.inputs.syslog.amqp.SyslogAMQPInput"
)

// NewInputSyslogAMQPAttrs is the constructor of InputSyslogAMQPAttrs.
func NewInputSyslogAMQPAttrs() InputAttrs {
	return &InputSyslogAMQPAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputSyslogAMQPAttrs) InputType() string {
	return InputTypeSyslogAMQP
}

// InputSyslogAMQPAttrs represents SyslogAMQP Input's attributes.
type InputSyslogAMQPAttrs struct {
	Heartbeat              int    `json:"heartbeat,omitempty"`
	Prefetch               int    `json:"prefetch,omitempty"`
	BrokerPort             int    `json:"broker_port,omitempty"`
	ParallelQueues         int    `json:"parallel_queues,omitempty"`
	BrokerVHost            string `json:"broker_vhost,omitempty"`
	BrokerUsername         string `json:"broker_username,omitempty"`
	BrokerPassword         string `json:"broker_password,omitempty"`
	Exchange               string `json:"exchange,omitempty"`
	OverrideSource         string `json:"override_source,omitempty"`
	RoutingKey             string `json:"routing_key,omitempty"`
	BrokerHostname         string `json:"broker_hostname,omitempty"`
	Queue                  string `json:"queue,omitempty"`
	ExchangeBind           bool   `json:"exchange_bind,omitempty"`
	ForceRDNS              bool   `json:"force_rdns,omitempty"`
	StoreFullMessage       bool   `json:"store_full_message,omitempty"`
	ExpandStructuredData   bool   `json:"expand_structured_data,omitempty"`
	ThrottlingAllowed      bool   `json:"throttling_allowed,omitempty"`
	TLS                    bool   `json:"tls,omitempty"`
	AllowOverrideDate      bool   `json:"allow_override_date,omitempty"`
	RequeueInvalidMessages bool   `json:"requeue_invalid_messages,omitempty"`
}
