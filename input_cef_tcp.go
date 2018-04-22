package graylog

func (attrs *InputCEFTCPAttrs) InputType() string {
	return INPUT_TYPE_CEF_TCP
}

type InputCEFTCPAttrs struct {
	MaxMessageSize        int    `json:"max_message_size,omitempty"`
	Timezone              string `json:"timezone,omitempty"`
	Locale                string `json:"locale,omitempty"`
	UseNullDelimiter      bool   `json:"use_null_delimiter,omitempty"`
	UseFullNames          bool   `json:"use_full_names,omitempty"`
	BindAddress           string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	Port                  int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	RecvBufferSize        int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
	TLSKeyFile            string `json:"tls_key_file,omitempty"`
	TLSEnable             bool   `json:"tls_enable,omitempty"`
	TLSClientAuth         string `json:"tls_client_auth,omitempty"`
	TLSKeyPassword        string `json:"tls_key_password,omitempty"`
	TCPKeepAlive          bool   `json:"tcp_keepalive,omitempty"`
	TLSClientAuthCertFile string `json:"tls_client_auth_cert_file,omitempty"`
	TLSCertFile           string `json:"tls_cert_file,omitempty"`
}
