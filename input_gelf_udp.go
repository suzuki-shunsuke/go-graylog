package graylog

func (attrs *InputGELFUDPAttrs) InputType() string {
	return InputTypeGELFUDP
}

type InputGELFUDPAttrs struct {
	DecompressSizeLimit int    `json:"decompress_size_limit,omitempty"`
	OverrideSource      string `json:"override_source,omitempty"`
	BindAddress         string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	Port                int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	RecvBufferSize      int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
}
