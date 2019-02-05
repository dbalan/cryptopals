package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackUint32(t *testing.T) {
	resp := PackUint32([]byte{byte(0x54), byte(0x68), byte(0x65), byte(0x20)}...)

	assert.Equal(t, uint32(1416127776), resp)
}
