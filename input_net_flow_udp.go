package graylog

func (attrs *InputNetFlowUDPAttrs) InputType() string {
	return INPUT_TYPE_NET_FLOW_UDP
}

type InputNetFlowUDPAttrs struct {
	NetFlow9DefinitionsPath string `json:"netflow9_definitions_path,omitempty"`
	OverrideSource          string `json:"override_source,omitempty"`
	BindAddress             string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	Port                    int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	RecvBufferSize          int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
}
