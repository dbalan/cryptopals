package rsa

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestPrimes(t *testing.T) {
	p, q, err := genprimes()
	assert.Nil(t, err)
	assert.NotEqual(t, p, q)
}

func TestSimpleFlow(t *testing.T) {
	pub, priv, n, err := GenKeyPair()
	assert.Nil(t, err)
	pt := big.NewInt(2323)

	ct := op(pt, pub, n)
	ptBk := op(ct, priv, n)
	assert.Equal(t, pt.Text(16), ptBk.Text(16))

}

func TestRSA(t *testing.T) {
	msg := []byte("hello world")
	pub, priv, n, err := GenKeyPair()
	assert.Nil(t, err)

	ct := Encrypt(msg, pub, n)
	gmsg, err := Decrypt(ct, priv, n)
	assert.Nil(t, err)
	assert.Equal(t, msg, gmsg)

}
