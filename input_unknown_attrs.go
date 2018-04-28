package graylog

// InputUnknownAttrs represents unknown type's Input Attributes.
type InputUnknownAttrs struct {
	inputType string
	Data      map[string]interface{}
}

// InputType is the implementation of the InputAttributes interface.
func (attrs InputUnknownAttrs) InputType() string {
	return attrs.inputType
}
