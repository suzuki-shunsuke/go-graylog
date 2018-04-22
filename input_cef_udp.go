package graylog

func (attrs *InputCEFUDPAttrs) InputType() string {
	return InputTypeCEFUDP
}

type InputCEFUDPAttrs struct {
	Locale         string `json:"locale,omitempty"`
	UseFullNames   bool   `json:"use_full_names,omitempty"`
	Timezone       string `json:"timezone,omitempty"`
	BindAddress    string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	Port           int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	RecvBufferSize int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
}
