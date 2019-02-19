package common

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestEGCD(t *testing.T) {
	cases := []struct {
		A, B, GCD *big.Int
	}{
		{big.NewInt(123), big.NewInt(0), big.NewInt(123)},
		{big.NewInt(0), big.NewInt(123), big.NewInt(123)},
		{big.NewInt(270), big.NewInt(192), big.NewInt(6)},
	}

	for _, cs := range cases {
		assert.Equal(t, cs.GCD, EGCD(cs.A, cs.B))
	}
}

func TestInvMod(t *testing.T) {
	p, err := InvMod(big.NewInt(17), big.NewInt(3120))
	assert.Nil(t, err)
	assert.Equal(t, big.NewInt(2753), p)
}
