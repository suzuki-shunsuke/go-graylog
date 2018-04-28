package graylog

const (
	InputTypeGELFHTTP string = "org.graylog2.inputs.gelf.http.GELFHttpInput"
)

// NewInputGELFHTTPAttrs is the constructor of InputGELFHTTPAttrs.
func NewInputGELFHTTPAttrs() InputAttributes {
	return &InputGELFHTTPAttrs{}
}

// InputType is the implementation of the InputAttributes interface.
func (attrs InputGELFHTTPAttrs) InputType() string {
	return InputTypeGELFHTTP
}

// InputGELFHTTPAttrs represents GELF HTTP Input's attributes.
type InputGELFHTTPAttrs struct {
	IdleWriterTimeOut     int    `json:"idle_writer_timeout,omitempty"`
	RecvBufferSize        int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
	MaxChunkSize          int    `json:"max_chunk_size,omitempty"`
	Port                  int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	DecompressSizeLimit   int    `json:"decompress_size_limit,omitempty"`
	TLSClientAuthCertFile string `json:"tls_client_auth_cert_file,omitempty"`
	BindAddress           string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	TLSCertFile           string `json:"tls_cert_file,omitempty"`
	TLSKeyFile            string `json:"tls_key_file,omitempty"`
	TLSKeyPassword        string `json:"tls_key_password,omitempty"`
	TLSClientAuth         string `json:"tls_client_auth,omitempty"`
	OverrideSource        string `json:"override_source,omitempty"`
	TCPKeepAlive          bool   `json:"tcp_keepalive,omitempty"`
	EnableCORS            bool   `json:"enable_cors,omitempty"`
	TLSEnable             bool   `json:"tls_enable,omitempty"`
}
