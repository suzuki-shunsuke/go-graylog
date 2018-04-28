package graylog

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/suzuki-shunsuke/go-set"
)

var (
	InputAttributesIntFieldSet  = set.NewStrSet()
	InputAttributesBoolFieldSet = set.NewStrSet()
	InputAttributesStrFieldSet  = set.NewStrSet()
	inputAttrsList              = []NewInputAttributes{
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
	inputAttrsMap = map[string]NewInputAttributes{}
)

// NewInputAttributes is the constructor of InputAttributes.
type NewInputAttributes func() InputAttributes

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
				InputAttributesStrFieldSet.Add(tag)
			case reflect.Int:
				InputAttributesIntFieldSet.Add(tag)
			case reflect.Bool:
				InputAttributesBoolFieldSet.Add(tag)
			default:
				panic(fmt.Sprintf("invalid type: %v", f.Type.Kind()))
			}
		}
	}
}

// NewInputAttrs returns a new InputAttributes.
func NewInputAttrs(t string) InputAttributes {
	a, ok := inputAttrsMap[t]
	if !ok {
		return &InputUnknownAttrs{inputType: t}
	}
	return a()
}

// InputAttributes represents Input Attributes.
// A receiver must be a pointer.
type InputAttributes interface {
	InputType() string
}
