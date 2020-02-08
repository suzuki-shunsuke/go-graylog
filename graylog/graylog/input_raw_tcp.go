package graylog

const (
	// InputTypeRawTCP is one of input types.
	InputTypeRawTCP string = "org.graylog2.inputs.raw.tcp.RawTCPInput"
)

// NewInputRawTCPAttrs is the constructor of InputRawTCPAttrs.
func NewInputRawTCPAttrs() InputAttrs {
	return &InputRawTCPAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputRawTCPAttrs) InputType() string {
	return InputTypeRawTCP
}

// InputRawTCPAttrs represents raw TCP Input's attributes.
type InputRawTCPAttrs struct {
	TLSEnable             bool   `json:"tls_enable"`
	TCPKeepAlive          bool   `json:"tcp_keepalive"`
	UseNullDelimiter      bool   `json:"use_null_delimiter"`
	Port                  int    `json:"port,omitempty"`
	RecvBufferSize        int    `json:"recv_buffer_size,omitempty"`
	NumberWorkerThreds    int    `json:"number_worker_threads,omitempty"`
	MaxMessageSize        int    `json:"max_message_size,omitempty"`
	BindAddress           string `json:"bind_address,omitempty"`
	TLSCertFile           string `json:"tls_cert_file,omitempty"`
	TLSKeyFile            string `json:"tls_key_file,omitempty"`
	TLSKeyPassword        string `json:"tls_key_password,omitempty"`
	TLSClientAuth         string `json:"tls_client_auth,omitempty"`
	TLSClientAuthCertFile string `json:"tls_client_auth_cert_file,omitempty"`
	OverrideSource        string `json:"override_source,omitempty"`
}
