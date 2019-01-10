package set1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixedXOR(t *testing.T) {
	cases := []struct {
		In1 string
		In2 string
		Out string
	}{
		{"c0fee", "c0fee", "00000"},
		{"1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965", "746865206b696420646f6e277420706c6179"},
	}

	for _, c := range cases {
		out, err := FixedXOR(c.In1, c.In2)
		assert.Nil(t, err)
		assert.Equal(t, c.Out, out)
	}
}
