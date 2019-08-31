package graylog

const (
	// InputTypeSyslogUDP is one of input types.
	InputTypeSyslogUDP string = "org.graylog2.inputs.syslog.udp.SyslogUDPInput"
)

// NewInputSyslogUDPAttrs is the constructor of InputSyslogUDPAttrs.
func NewInputSyslogUDPAttrs() InputAttrs {
	return &InputSyslogUDPAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputSyslogUDPAttrs) InputType() string {
	return InputTypeSyslogUDP
}

// InputSyslogUDPAttrs represents SyslogUDP Input's attributes.
type InputSyslogUDPAttrs struct {
	BindAddress            string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	Port                   int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	RecvBufferSize         int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
	TCPKeepAlive           bool   `json:"tcp_keepalive"`
	TLSEnable              bool   `json:"tls_enable"`
	ThrottlingAllowed      bool   `json:"throttling_allowed"`
	EnableCORS             bool   `json:"enable_cors"`
	UseNullDelimiter       bool   `json:"use_null_delimiter"`
	ExchangeBind           bool   `json:"exchange_bind"`
	ForceRDNS              bool   `json:"force_rdns"`
	StoreFullMessage       bool   `json:"store_full_message"`
	ExpandStructuredData   bool   `json:"expand_structured_data"`
	AllowOverrideDate      bool   `json:"allow_override_date"`
	RequeueInvalidMessages bool   `json:"requeue_invalid_messages"`
	UseFullNames           bool   `json:"use_full_names"`
	TLS                    bool   `json:"tls"`
}
