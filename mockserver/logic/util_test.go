package logic

import (
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
