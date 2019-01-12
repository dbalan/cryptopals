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

func TestB64Frag(t *testing.T) {
	cs := []struct {
		In  string
		Out string
	}{
		{"TWFu", "Man"},
		{"MTI=", "12"},
		{"MQ==", "1"},
		{"S2V5Ym9hcmRJbnRlcnJ1cHQ=", "KeyboardInterrupt"},
	}
	for _, c := range cs {
		decoded := base64decode([]byte(c.In))
		assert.Equal(t, []byte(c.Out), decoded)
	}

}

func TestFullSerialization(t *testing.T) {
	input := "0e3647e8592d35514a081243582536ed3de6734059001e3f535ce6271032"

	b64, err := Hex2B64(input)
	assert.Nil(t, err)
	decoded := base64decode([]byte(b64))
	pretty := encodeHexString(decoded)
	assert.Equal(t, input, pretty)
}
