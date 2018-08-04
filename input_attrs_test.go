package graylog_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-graylog"
)

func TestNewInputAttrsByType(t *testing.T) {
	attrs := graylog.NewInputAttrsByType("hoge")
	if attrs.InputType() != "hoge" {
		t.Fatalf(`attrs.InputType() = "%s", wanted "hoge"`, attrs.InputType())
	}
}

type CustomInputAttrs struct {
	Type string
	Foo  string
}

func (attrs CustomInputAttrs) InputType() string {
	return attrs.Type
}

func TestSetFuncGetUnknownType(t *testing.T) {
	f := graylog.GetFuncGetUnknownTypeInputAttrs()
	defer graylog.SetFuncGetUnknownTypeInputAttrs(f)
	graylog.SetFuncGetUnknownTypeInputAttrs(func(data map[string]graylog.NewInputAttrs, t string) graylog.InputAttrs {
		return &CustomInputAttrs{Type: t, Foo: "foo"}
	})
	attrs := graylog.NewInputAttrsByType("hoge")
	a, ok := attrs.(*CustomInputAttrs)
	if !ok {
		t.Fatalf("attrs is not CustomInputAttrs")
	}
	if a.Foo != "foo" {
		t.Fatalf(`a.Foo = "%s", wanted "foo"`, a.Foo)
	}
}

func TestSetFuncGetInputAttrsByType(t *testing.T) {
	f := graylog.GetFuncGetInputAttrsByType()
	defer graylog.SetFuncGetInputAttrsByType(f)
	graylog.SetFuncGetInputAttrsByType(func(data map[string]graylog.NewInputAttrs, t string) graylog.InputAttrs {
		return &CustomInputAttrs{Type: t, Foo: "foo"}
	})
	attrs := graylog.NewInputAttrsByType(graylog.InputTypeBeats)
	a, ok := attrs.(*CustomInputAttrs)
	if !ok {
		t.Fatalf("attrs is not CustomInputAttrs")
	}
	if a.Foo != "foo" {
		t.Fatalf(`a.Foo = "%s", wanted "foo"`, a.Foo)
	}
}
