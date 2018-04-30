package graylog

const (
	// InputTypeCEFUDP is one of input types.
	InputTypeCEFUDP string = "org.graylog.plugins.cef.input.CEFUDPInput"
)

// NewInputCEFUDPAttrs is the constructor of InputCEFUDPAttrs.
func NewInputCEFUDPAttrs() InputAttrs {
	return &InputCEFUDPAttrs{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputCEFUDPAttrs) InputType() string {
	return InputTypeCEFUDP
}

// InputCEFUDPAttrs represents CEF UDP Input's attributes.
type InputCEFUDPAttrs struct {
	Locale         string `json:"locale,omitempty"`
	UseFullNames   bool   `json:"use_full_names,omitempty"`
	Timezone       string `json:"timezone,omitempty"`
	BindAddress    string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	Port           int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	RecvBufferSize int    `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
}
