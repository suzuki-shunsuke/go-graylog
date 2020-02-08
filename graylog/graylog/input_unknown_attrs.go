package graylog

import (
	"encoding/json"
)

// InputUnknownAttrs represents unknown type's Input Attrs.
type InputUnknownAttrs struct {
	inputType string
	Data      map[string]interface{}
}

// InputType is the implementation of the InputAttrs interface.
func (attrs InputUnknownAttrs) InputType() string {
	return attrs.inputType
}

func (attrs *InputUnknownAttrs) MarshalJSON() ([]byte, error) {
	return json.Marshal(attrs.Data)
}
