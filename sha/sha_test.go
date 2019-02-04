package sha

import (
	"github.com/dbalan/cryptopals/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPreprocessing(t *testing.T) {
	input := []byte("The quick brown fox jumps over the lazy dog")

	out := preprocess(input)
	assert.Equal(t, 0, (len(out)*8)%512)
}

func TestBEEncode(t *testing.T) {
	outp := BEEncodeUint64(0xdeadbeef)
	expected := []byte{byte(0), byte(0), byte(0), byte(0), byte(222), byte(173), byte(190), byte(239)}

	assert.Equal(t, expected, outp)
}

func TestSHA(t *testing.T) {
	input := "The quick brown fox jumps over the lazy dog"
	expected := common.DecodeB64([]byte("L9ThxnotKPzthJ7hu3bnORuT6xI="))

	assert.Equal(t, expected, SHA([]byte(input)))
}

func TestPackUint32(t *testing.T) {
	resp := packUint32([]byte{byte(0x54), byte(0x68), byte(0x65), byte(0x20)}...)

	assert.Equal(t, uint32(1416127776), resp)
}
