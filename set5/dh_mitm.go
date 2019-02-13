package set5

import (
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/set2"
	"github.com/dbalan/cryptopals/sha"
	"math/big"
)

type Recv interface {
	PrimeEx(p, g *big.Int)
	Kex(pub *big.Int) *big.Int
	Exchange(ct, iv []byte) (msg, niv []byte)
}

type PersonB struct {
	p, g, pub, priv, theirPub, skey *big.Int
}

func newB() *PersonB {
	return &PersonB{}
}

func (b *PersonB) PrimeEx(p, g *big.Int) {
	b.p = p
	b.g = g
	pub, priv := keypair(p, g)
	b.pub = pub
	b.priv = priv
}

func (b *PersonB) Kex(pub *big.Int) *big.Int {
	b.theirPub = pub
	skey := sessionKey(pub, b.priv, b.p)
	b.skey = skey
	return b.pub
}

func (b *PersonB) Exchange(ct, iv []byte) ([]byte, []byte) {
	// decrypt!
	msg := decryptWithSK(b.skey, ct, iv)
	nct, niv := encryptWithSK(b.skey, msg)
	return nct, niv
}

func (b *PersonB) Skey() *big.Int {
	return b.skey
}

func getKeySK(skey *big.Int) []byte {
	keybytes := []byte(skey.Text(16))
	key := []byte(common.EncodeHexString(sha.SHA(keybytes))[0:16])
	return key
}

func encryptWithSK(skey *big.Int, msg []byte) ([]byte, []byte) {
	// get an AES key
	key := getKeySK(skey)

	iv, err := common.RandBytes(16)
	if err != nil {
		panic(err)
	}

	ct, err := set2.EncAES128CBC(msg, iv, key)
	if err != nil {
		panic(err)
	}

	return ct, iv
}

func decryptWithSK(skey *big.Int, ct, iv []byte) []byte {
	key := getKeySK(skey)
	msg, err := set2.DecAES128CBC(ct, iv, key)
	if err != nil {
		panic("error: decrypt" + err.Error())
	}
	return msg
}

func communicate(recv Recv) {
	// personA
	msg := []byte("hello world")
	p, g := primes()
	recv.PrimeEx(p, g)

	pub, priv := keypair(p, g)
	pubtheir := recv.Kex(pub)
	skey := sessionKey(pubtheir, priv, p)

	ct, iv := encryptWithSK(skey, msg)
	nct, niv := recv.Exchange(ct, iv)
	rmsg := decryptWithSK(skey, nct, niv)

	if !common.EqualBlocks(rmsg, msg) {
		panic("msgs don't match")
	}
}
