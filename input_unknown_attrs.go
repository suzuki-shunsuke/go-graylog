package graylog

// InputUnknownAttrs represents unknown type's Input Attrs.
type InputUnknownAttrs struct {
	inputType string
	Data      map[string]interface{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputUnknownAttrs) InputType() string {
	return attrs.inputType
}
