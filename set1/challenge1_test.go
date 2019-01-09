package set1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	bytes := convert(1901)
	fmt.Println(string(bytes))
}

func TestHex2Base64(t *testing.T) {
	b64 := hex2base64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	assert.Equal(t, b64, "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")
}
