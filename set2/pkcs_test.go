package set2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPKCS7Padding(t *testing.T) {
	data := "YELLOW SUBMARINE"
	l := 20

	padded, err := PKCS7Padding([]byte(data), l)
	assert.Nil(t, err)

	assert.Equal(t, 20, len(padded))
	assert.Equal(t, []byte("YELLOW SUBMARINE\x04\x04\x04\x04"), padded)
}

func TestPKCS7StripPadding(t *testing.T) {
	padded := []byte("YELLOW SUBMARINE\x04\x04\x04\x04")
	data := []byte("YELLOW SUBMARINE")

	unpadded := PKCS7StripPadding(padded)
	assert.Equal(t, data, unpadded)

}

func TestValidateStripPKCS7Padding(t *testing.T) {
	inp := []byte("ICE ICE BABY")
	errCases := [][]byte{
		append(inp, byte(5), byte(5), byte(5)),
		append(inp, byte(1), byte(2), byte(3), byte(4)),
	}

	for _, cs := range errCases {
		_, err := validateStripPKCS7Padding(cs)
		assert.NotNil(t, err)
	}

	padded, err := PKCS7Padding(inp, 20)
	assert.Nil(t, err)
	out, err := validateStripPKCS7Padding(padded)
	assert.Nil(t, err)
	assert.Equal(t, inp, out)

}
