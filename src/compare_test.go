package stringutils

import (
	"testing"
)

func TestJaroWinklerDistance(t *testing.T) {
	if JaroWinklerDistance("", "") != 0. {
		t.Fail()
	}
	if JaroWinklerDistance("ABC", "ABC") != 1. {
		t.Fail()
	}
	if JaroWinklerDistance("ABCDEFGH", "ABCD") != .9 {
		t.FailNow()
	}
}
