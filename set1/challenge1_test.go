package set1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
func TestHex2B64(t *testing.T) {
	b64, err := Hex2B64("f75")
	assert.Nil(t, err)
	assert.Equal(t, b64, "91")
}
*/

func TestHex2B64(t *testing.T) {
	cases := []struct {
		In  string
		Out string
	}{
		{"f75", "91=="},
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
		{"", ""},
		{"34", "NA=="},
	}

	for _, c := range cases {
		b64, err := Hex2B64(c.In)
		assert.Nil(t, err)
		assert.Equal(t, b64, c.Out)
	}
}

func TestParsePoint(t *testing.T) {
	cases := []struct {
		In  string
		Out int64
	}{
		{"42", 66},
		{"0", 0},
		{"c0fee", 790510},
	}

	for _, c := range cases {
		val, err := parsePoint([]byte(c.In))
		assert.Nil(t, err)
		assert.Equal(t, val, c.Out)
	}
}
