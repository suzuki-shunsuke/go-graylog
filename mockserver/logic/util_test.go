package logic

import (
	"reflect"
	"testing"
)

func TestRandStringBytesMaskImprSrc(t *testing.T) {
	a1 := randStringBytesMaskImprSrc(24)
	a2 := randStringBytesMaskImprSrc(24)
	if len(a1) != 24 {
		t.Fatalf("len(a1) == %d, wanted 24", len(a1))
	}
	if a1 == a2 {
		t.Fatalf("a1 == a2 == %s", a1)
	}
}

func TestAddToStringArray(t *testing.T) {
	act := addToStringArray([]string{}, "foo")
	exp := []string{"foo"}
	if !reflect.DeepEqual(act, exp) {
		t.Fatalf(`addToStringArray([]string{}, "foo") == %v, wanted %v`, act, exp)
	}

	act = addToStringArray([]string{"foo", "bar"}, "foo")
	exp = []string{"foo", "bar"}
	if !reflect.DeepEqual(act, exp) {
		t.Fatalf(
			`addToStringArray([]string{"foo", "bar"}, "foo") == %v, wanted %v`,
			act, exp)
	}
}

func TestRemoveFromStringArray(t *testing.T) {
	act := removeFromStringArray([]string{}, "foo")
	exp := []string{}
	if !reflect.DeepEqual(act, exp) {
		t.Fatalf(
			`removeFromStringArray([]string{}, "foo") == %v, wanted %v`, act, exp)
	}

	act = removeFromStringArray([]string{"foo", "bar"}, "foo")
	exp = []string{"bar"}
	if !reflect.DeepEqual(act, exp) {
		t.Fatalf(
			`removeFromStringArray([]string{"foo", "bar"}, "foo") == %v, wanted %v`,
			act, exp)
	}
}
