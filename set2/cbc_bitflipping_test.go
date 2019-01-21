package set2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCBCOracle(t *testing.T) {
	key, err := randBytes(16)

	assert.Nil(t, err)

	ct := cbcOracle("hello", key)
	resp := decryptOracle(ct, key)
	assert.Equal(t, false, resp)
}

func TestIsAdmin(t *testing.T) {
	cases := []struct {
		In  string
		Out bool
	}{
		{"something=what;admin=false;ehat=wer", false},
		{"admin=true", true},
		{"", false},
		{"admin=true;what=when", true},
	}

	for _, cs := range cases {
		assert.Equal(t, cs.Out, isAdmin(cs.In))

	}
}
