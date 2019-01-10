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
