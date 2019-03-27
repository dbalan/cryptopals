package dsa

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

// Checking if inverses work the way I think they would.
func TestInvers(t *testing.T) {
	p, g, _ := GetDefaultParams()
	gi := new(big.Int).Set(g)
	gi.Exp(g, big.NewInt(-1), p)
	gi.Mul(gi, g)
	gi.Mod(gi, p)
	assert.Equal(t, 0, gi.Cmp(big.NewInt(1)))
}

func TestDSASigning(t *testing.T) {
	x, y, err := KeyPair()
	assert.Nil(t, err)

	msg := []byte("jebus")
	r, s, err := Sign(msg, x, GetDefaultParams)
	assert.Nil(t, err)
	sig := Verify(msg, r, s, y, GetDefaultParams)
	assert.True(t, sig)
}

func TestDSAComputeKey(t *testing.T) {
	msg := []byte("foo")
	x, _, err := KeyPair()
	assert.Nil(t, err)

	r, s, k, err := signInternal(msg, x, GetDefaultParams)
	assert.Nil(t, err)

	hs := hsmsg(msg)
	newx := ComputeKey(k, r, s, hs)
	assert.Equal(t, 0, newx.Cmp(x), "mismatch x")
}
