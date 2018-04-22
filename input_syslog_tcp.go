package graylog

// InputType is the implementation of the InputAttributes interface.
func (attrs *InputSyslogTCPAttrs) InputType() string {
	return InputTypeSyslogTCP
}

// InputSyslogTCPAttrs
type InputSyslogTCPAttrs struct {
	Port           int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	BindAddress    string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	RecvBufferSize int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
}
