package graylog

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/suzuki-shunsuke/go-set"
)

var (
	InputAttrsIntFieldSet  = set.NewStrSet()
	InputAttrsBoolFieldSet = set.NewStrSet()
	InputAttrsStrFieldSet  = set.NewStrSet()
	inputAttrsList         = []NewInputAttrs{
		NewInputAWSFlowLogsAttrs,
		NewInputBeatsAttrs,
		NewInputAWSLogsAttrs,
		NewInputCEFAMQPAttrs,
		NewInputCEFKafkaAttrs,
		NewInputCEFTCPAttrs,
		NewInputCloudTrailAttrs,
		NewInputFakeHTTPMessageAttrs,
		NewInputGELFAMQPAttrs,
		NewInputGELFHTTPAttrs,
		NewInputGELFKafkaAttrs,
		NewInputGELFUDPAttrs,
		NewInputNetFlowUDPAttrs,
		NewInputRawAMQPAttrs,
		NewInputSyslogAMQPAttrs,
		NewInputSyslogKafkaAttrs,
		NewInputSyslogTCPAttrs,
		NewInputSyslogUDPAttrs,
	}
	inputAttrsMap = map[string]NewInputAttrs{}
)

// NewInputAttrs is the constructor of InputAttrs.
type NewInputAttrs func() InputAttrs

func init() {
	for _, attrs := range inputAttrsList {
		a := attrs()
		inputAttrsMap[a.InputType()] = attrs
		ts := reflect.ValueOf(a).Elem().Type()
		n := ts.NumField()
		for i := 0; i < n; i++ {
			f := ts.Field(i)
			tag := strings.Split(f.Tag.Get("json"), ",")[0]
			switch f.Type.Kind() {
			case reflect.String:
				InputAttrsStrFieldSet.Add(tag)
			case reflect.Int:
				InputAttrsIntFieldSet.Add(tag)
			case reflect.Bool:
				InputAttrsBoolFieldSet.Add(tag)
			default:
				panic(fmt.Sprintf("invalid type: %v", f.Type.Kind()))
			}
		}
	}
}

// NewInputAttrsByType returns a new InputAttrs.
func NewInputAttrsByType(t string) InputAttrs {
	a, ok := inputAttrsMap[t]
	if !ok {
		return &InputUnknownAttrs{inputType: t}
	}
	return a()
}

// InputAttrs represents Input Attributes.
// A receiver must be a pointer.
type InputAttrs interface {
	InputType() string
}
