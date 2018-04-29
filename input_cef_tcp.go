package graylog

const (
	InputTypeCEFTCP string = "org.graylog.plugins.cef.input.CEFTCPInput"
)

// NewInputCEFTCPAttrs is the constructor of InputCEFTCPAttrs.
func NewInputCEFTCPAttrs() InputAttrs {
	return &InputCEFTCPAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputCEFTCPAttrs) InputType() string {
	return InputTypeCEFTCP
}

// InputCEFTCPAttrs represents CEF TCP Input's attributes.
type InputCEFTCPAttrs struct {
	UseNullDelimiter      bool   `json:"use_null_delimiter,omitempty"`
	UseFullNames          bool   `json:"use_full_names,omitempty"`
	TLSEnable             bool   `json:"tls_enable,omitempty"`
	TCPKeepAlive          bool   `json:"tcp_keepalive,omitempty"`
	MaxMessageSize        int    `json:"max_message_size,omitempty"`
	Port                  int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	RecvBufferSize        int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
	Timezone              string `json:"timezone,omitempty"`
	Locale                string `json:"locale,omitempty"`
	BindAddress           string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	TLSKeyFile            string `json:"tls_key_file,omitempty"`
	TLSClientAuth         string `json:"tls_client_auth,omitempty"`
	TLSKeyPassword        string `json:"tls_key_password,omitempty"`
	TLSClientAuthCertFile string `json:"tls_client_auth_cert_file,omitempty"`
	TLSCertFile           string `json:"tls_cert_file,omitempty"`
}
