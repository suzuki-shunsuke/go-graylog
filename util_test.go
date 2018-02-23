package graylog

import (
	"reflect"
	"testing"
)

func TestAddToStringArray(t *testing.T) {
	act := addToStringArray([]string{}, "foo")
	exp := []string{"foo"}
	if !reflect.DeepEqual(act, exp) {
		t.Errorf(`addToStringArray([]string{}, "foo") == %v, wanted %v`, act, exp)
	}

	act = addToStringArray([]string{"foo", "bar"}, "foo")
	exp = []string{"foo", "bar"}
	if !reflect.DeepEqual(act, exp) {
		t.Errorf(`addToStringArray([]string{"foo", "bar"}, "foo") == %v, wanted %v`, act, exp)
	}
}

func TestRemoveFromStringArray(t *testing.T) {
	act := removeFromStringArray([]string{}, "foo")
	exp := []string{}
	if !reflect.DeepEqual(act, exp) {
		t.Errorf(`removeFromStringArray([]string{}, "foo") == %v, wanted %v`, act, exp)
	}

	act = removeFromStringArray([]string{"foo", "bar"}, "foo")
	exp = []string{"bar"}
	if !reflect.DeepEqual(act, exp) {
		t.Errorf(`removeFromStringArray([]string{"foo", "bar"}, "foo") == %v, wanted %v`, act, exp)
	}
}
