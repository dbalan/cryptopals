package dsa

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

// Checking if inverses work the way I think they would.
func TestInvers(t *testing.T) {
	p, g, _ := getDSAParams()
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
	r, s, err := Sign(msg, x)
	assert.Nil(t, err)
	sig := Verify(msg, r, s, y)
	assert.True(t, sig)
}
