package graylog

const (
	// InputTypeGELFTCP is one of input types.
	InputTypeGELFTCP string = "org.graylog2.inputs.gelf.tcp.GELFTCPInput"
)

// NewInputGELFTCPAttrs is the constructor of InputGELFTCPAttrs.
func NewInputGELFTCPAttrs() InputAttrs {
	return &InputGELFTCPAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputGELFTCPAttrs) InputType() string {
	return InputTypeGELFTCP
}

// InputGELFTCPAttrs represents GELF TCP Input's attributes.
type InputGELFTCPAttrs struct {
	MaxMessageSize        int    `json:"max_message_size,omitempty"`
	DecompressSizeLimit   int    `json:"decompress_size_limit,omitempty"`
	Port                  int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	RecvBufferSize        int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
	BindAddress           string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	OverrideSource        string `json:"override_source,omitempty"`
	TLSKeyFile            string `json:"tls_key_file,omitempty"`
	TLSKeyPassword        string `json:"tls_key_password,omitempty"`
	TLSClientAuthCertFile string `json:"tls_client_auth_cert_file,omitempty"`
	TLSClientAuth         string `json:"tls_client_auth,omitempty"`
	TLSCertFile           string `json:"tls_cert_file,omitempty"`
	UseNullDelimiter      bool   `json:"use_null_delimiter"`
	TLSEnable             bool   `json:"tls_enable"`
	TCPKeepAlive          bool   `json:"tcp_keepalive"`
}
