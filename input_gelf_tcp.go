package graylog

func (attrs *InputGELFTCPAttrs) InputType() string {
	return InputTypeGELFTCP
}

type InputGELFTCPAttrs struct {
	MaxMessageSize        int    `json:"max_message_size,omitempty"`
	DecompressSizeLimit   int    `json:"decompress_size_limit,omitempty"`
	UseNullDelimiter      bool   `json:"use_null_delimiter,omitempty"`
	BindAddress           string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	Port                  int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	RecvBufferSize        int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
	OverrideSource        string `json:"override_source,omitempty"`
	TLSKeyFile            string `json:"tls_key_file,omitempty"`
	TLSEnable             bool   `json:"tls_enable,omitempty"`
	TLSKeyPassword        string `json:"tls_key_password,omitempty"`
	TCPKeepAlive          bool   `json:"tcp_keepalive,omitempty"`
	TLSClientAuthCertFile string `json:"tls_client_auth_cert_file,omitempty"`
	TLSClientAuth         string `json:"tls_client_auth,omitempty"`
	TLSCertFile           string `json:"tls_cert_file,omitempty"`
}
