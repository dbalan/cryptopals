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
