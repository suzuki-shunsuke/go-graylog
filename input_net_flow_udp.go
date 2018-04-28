package graylog

const (
	InputTypeNetFlowUDP string = "org.graylog.plugins.netflow.inputs.NetFlowUdpInput"
)

// NewInputNetFlowUDPAttrs is the constructor of InputNetFlowUDPAttrs.
func NewInputNetFlowUDPAttrs() InputAttrs {
	return &InputNetFlowUDPAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputNetFlowUDPAttrs) InputType() string {
	return InputTypeNetFlowUDP
}

// InputNetFlowUDPAttrs represents net flow UDP Input's attributes.
type InputNetFlowUDPAttrs struct {
	NetFlow9DefinitionsPath string `json:"netflow9_definitions_path,omitempty"`
	OverrideSource          string `json:"override_source,omitempty"`
	BindAddress             string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	Port                    int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	RecvBufferSize          int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
}
