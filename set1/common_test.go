package set1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeHexString(t *testing.T) {
	cases := []struct {
		In  string
		Out []byte
	}{
		{"ff", []byte{byte(0xff)}},
		{"f", []byte{byte(0xf)}},
		{"cbc", []byte{byte(0xc), byte(0xbc)}},
	}

	for _, c := range cases {
		val := decodeHexString(c.In)
		assert.Equal(t, val, c.Out)
	}
}

func TestHexPrettyPrint(t *testing.T) {
	cases := []struct {
		In  byte
		Out string
	}{
		{byte(0xb), "0b"},
		{byte(0x45), "45"},
		{byte(0x2a), "2a"},
		{byte(0xa2), "a2"},
	}

	for _, c := range cases {
		out := hexPrettyPrint(c.In)
		assert.Equal(t, c.Out, out)
	}
}
