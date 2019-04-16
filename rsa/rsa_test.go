package rsa

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

const primeSize = 128

func TestPrimes(t *testing.T) {
	p, q, err := genprimes(primeSize)
	assert.Nil(t, err)
	assert.NotEqual(t, p, q)
}

func TestSimpleFlow(t *testing.T) {
	pub, priv, n, err := GenKeyPair(primeSize)
	assert.Nil(t, err)
	pt := big.NewInt(2323)

	ct := op(pt, pub, n)
	ptBk := op(ct, priv, n)
	assert.Equal(t, pt.Text(16), ptBk.Text(16))

}

func TestLargerKeySize(t *testing.T) {
	_, _, _, err := GenKeyPair(1024)
	assert.Nil(t, err)
}

func TestRSA(t *testing.T) {
	msg := []byte("hello world")
	pub, priv, n, err := GenKeyPair(primeSize)
	assert.Nil(t, err)

	ct := Encrypt(msg, pub, n)
	gmsg, err := Decrypt(ct, priv, n)
	assert.Nil(t, err)
	assert.Equal(t, msg, gmsg)

}

func TestRSASigning(t *testing.T) {
	msg := []byte("hello word")
	pub, priv, n, err := GenKeyPair(primeSize)
	assert.Nil(t, err)

	sig := Sign(msg, priv, n)
	verify := VerifySign(msg, sig, pub, n)
	assert.Equal(t, true, verify)
	verify = VerifySign(msg, "512ba0e938d862261a9914c7f5370dab3d7c1695", pub, n)
	assert.NotEqual(t, true, verify)
}
