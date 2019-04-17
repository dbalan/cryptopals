package set6

import (
	"fmt"
	"math/big"

	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/rsa"
)

const secret = "VGhhdCdzIHdoeSBJIGZvdW5kIHlvdSBkb24ndCBwbGF5" +
	"IGFyb3VuZCB3aXRoIHRoZSBGdW5reSBDb2xkIE1lZGluYQ=="

// takes a ciphertext, returns true of PT is an even number.
func ParityOracle() (*big.Int, func(*big.Int) bool) {
	pub, priv, n, err := rsa.GenKeyPair(1024)
	if err != nil {
		panic(err)
	}

	msg := common.DecodeB64([]byte(secret))

	actualCT := rsa.Encrypt(msg, pub, n)

	oracle := func(ct *big.Int) bool {
		msg, err := rsa.Decrypt(ct, priv, n)
		if err != nil {
			panic(err)
		}

		return isEven(msg)
	}

	return actualCT, oracle
}

func isEven(msg []byte) bool {
	hx := fmt.Sprintf("%x", msg)
	m := new(big.Int)
	m.SetString(hx, 16)
	m.Mod(m, big.NewInt(2))
	return m.Cmp(big.NewInt(0)) == 0
}
