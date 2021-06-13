package stringutils

import (
	"testing"
)

func TestFromASCIIZ(t *testing.T) {
	b := []byte{0x54, 0x65, 0x73, 0x74, 0x0, 0x0, 0x0, 0x0}
	res := FromASCIIZ(b)
	if string(res) != "Test" {
		t.Fail()
	}
}

func TestFromUTF16LE(t *testing.T) {
	b := []byte{
		0xff, 0xfe, 0x14, 0x4, 0x32, 0x4, 0x30, 0x4, 0x20, 0x0, 0x3a, 0x4, 0x3e, 0x4, 0x3b, 0x4,
		0x4c, 0x4, 0x46, 0x4, 0x30, 0x4}
	if res, err := FromUTF16LE(b); err != nil || string(res) != "Два кольца" {
		t.Fail()
	}
}

func TestFromUTF16BE(t *testing.T) {

}
