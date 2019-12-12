package graylog

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/suzuki-shunsuke/go-set"
)

var (
	// InputAttrsIntFieldSet is the set of int fields of all type of input attributes.
	InputAttrsIntFieldSet = set.NewStrSet()
	// InputAttrsBoolFieldSet is the set of bool fields of all type of input attributes.
	InputAttrsBoolFieldSet = set.NewStrSet()
	// InputAttrsStrFieldSet is the set of string fields of all type of input attributes.
	InputAttrsStrFieldSet = set.NewStrSet()
	inputAttrsList        = []NewInputAttrs{
		NewInputAWSFlowLogsAttrs,
		NewInputAWSCloudWatchLogsAttrs,
		NewInputAWSCloudTrailAttrs,
		NewInputBeatsAttrs,
		NewInputCEFAMQPAttrs,
		NewInputCEFKafkaAttrs,
		NewInputCEFTCPAttrs,
		NewInputCEFUDPAttrs,
		NewInputFakeHTTPMessageAttrs,
		NewInputGELFAMQPAttrs,
		NewInputGELFHTTPAttrs,
		NewInputGELFKafkaAttrs,
		NewInputGELFTCPAttrs,
		NewInputGELFUDPAttrs,
		NewInputJSONPathAttrs,
		NewInputNetFlowUDPAttrs,
		NewInputRawAMQPAttrs,
		NewInputRawKafkaAttrs,
		NewInputSyslogAMQPAttrs,
		NewInputSyslogKafkaAttrs,
		NewInputSyslogTCPAttrs,
		NewInputSyslogUDPAttrs,
	}
	// initialize at init function for preventing initialization loop
	attrsSet = &inputAttrsSet{}
)

type (
	// GetUnknownTypeInputAttrsIntf returns an unknown type InputAttrs.
	GetUnknownTypeInputAttrsIntf func(map[string]NewInputAttrs, string) InputAttrs

	// GetInputAttrsByTypeIntf returns a given type InputAttrs.
	GetInputAttrsByTypeIntf func(map[string]NewInputAttrs, string) InputAttrs

	// NewInputAttrs is the constructor of InputAttrs.
	NewInputAttrs func() InputAttrs

	// InputAttrs represents Input Attributes.
	// A receiver must be a pointer.
	InputAttrs interface {
		InputType() string
	}

	inputAttrsSet struct {
		data           map[string]NewInputAttrs
		GetUnknownType GetUnknownTypeInputAttrsIntf
		GetByType      GetInputAttrsByTypeIntf
	}
)

func init() {
	attrsSet = &inputAttrsSet{
		data: map[string]NewInputAttrs{},
		GetUnknownType: func(data map[string]NewInputAttrs, t string) InputAttrs {
			return &InputUnknownAttrs{inputType: t}
		},
		GetByType: func(data map[string]NewInputAttrs, t string) InputAttrs {
			if data == nil {
				return attrsSet.GetUnknownType(data, t)
			}
			a, ok := data[t]
			if !ok {
				return attrsSet.GetUnknownType(data, t)
			}
			return a()
		},
	}
	if err := SetInputAttrs(inputAttrsList...); err != nil {
		panic(err)
	}
	for _, attrs := range inputAttrsList {
		a := attrs()
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

// SetFuncGetUnknownTypeInputAttrs customizes NewInputAttrsByType's behavior.
func SetFuncGetUnknownTypeInputAttrs(f GetUnknownTypeInputAttrsIntf) {
	attrsSet.GetUnknownType = f
}

// GetFuncGetUnknownTypeInputAttrs returns the global GetUnknownTypeInputAttrs function.
// Mainly this is used to prevent global pollution at test.
//
//   f := graylog.GetFuncGetUnknownTypeInputAttrs()
//   // change the global function temporary
//   defer graylog.SetFuncGetUnknownTypeInputAttrs(f)
//   graylog.SetFuncGetUnknownTypeInputAttrs(customFunc)
func GetFuncGetUnknownTypeInputAttrs() GetUnknownTypeInputAttrsIntf {
	if attrsSet == nil {
		return nil
	}
	return attrsSet.GetUnknownType
}

// SetFuncGetInputAttrsByType customizes NewInputAttrsByType's behavior.
func SetFuncGetInputAttrsByType(f GetInputAttrsByTypeIntf) {
	attrsSet.GetByType = f
}

// GetFuncGetInputAttrsByType returns the global GetInputAttrsByType function.
// Mainly this is used to prevent global pollution at test.
//
//   f := graylog.GetFuncGetInputAttrsByType()
//   // change the global function temporary
//   defer graylog.SetFuncGetInputAttrsByType(f)
//   graylog.SetFuncGetInputAttrsByType(customFunc)
func GetFuncGetInputAttrsByType() GetInputAttrsByTypeIntf {
	if attrsSet == nil {
		return nil
	}
	return attrsSet.GetByType
}

// NewInputAttrsByType returns a new InputAttrs.
func NewInputAttrsByType(t string) InputAttrs {
	return attrsSet.GetByType(attrsSet.data, t)
}

// SetInputAttrs sets InputAttrs.
// You can add the custom InputAttrs and override existing InputAttrs.
func SetInputAttrs(args ...NewInputAttrs) error {
	if attrsSet == nil {
		attrsSet = &inputAttrsSet{data: map[string]NewInputAttrs{}}
	}
	if attrsSet.data == nil {
		attrsSet.data = map[string]NewInputAttrs{}
	}
	for _, f := range args {
		a := f()
		if reflect.TypeOf(a).Kind() != reflect.Ptr {
			return errors.New("NewInputAttrs must return pointer")
		}
		attrsSet.data[a.InputType()] = f
	}
	return nil
}
