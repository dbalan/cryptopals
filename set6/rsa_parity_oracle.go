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
func ParityOracle() (*big.Int, *big.Int, *big.Int, func(*big.Int) bool) {
	pub, priv, n, err := rsa.GenKeyPair(1024)
	if err != nil {
		panic(err)
	}

	msg := common.DecodeB64([]byte(secret))

	actualCT := rsa.Encrypt(msg, pub, n)

	oracle := func(ct *big.Int) bool {
		pt := new(big.Int).Exp(ct, priv, n)
		pt.Mod(pt, big.NewInt(2))
		return pt.Cmp(big.NewInt(0)) == 0
	}

	return actualCT, pub, n, oracle
}

func times(n *big.Int) int {
	tmp := new(big.Int).Set(n)
	zero := big.NewInt(0)
	res := 0
	for {
		tmp.Rsh(tmp, 1)
		if tmp.Cmp(zero) <= 0 {
			break
		}
		res++
	}
	return res
}

func parityOracleAttack(e, n, ct *big.Int,
	oracle func(*big.Int) bool) {

	// set bounds
	lower := big.NewInt(0)
	upper := new(big.Int).Set(n)

	// generate a multiplier
	multi := big.NewInt(2)
	multi.Exp(multi, e, n)

	for i := 0; i < times(n); i++ {
		// multiply ct
		ct.Mul(ct, multi)
		ct.Mod(ct, n)

		if oracle(ct) {
			// less than half of upper
			upper.Add(upper, lower)
			upper.Div(upper, big.NewInt(2))
		} else {
			// more than half of lower
			lower.Add(upper, lower)
			lower.Div(lower, big.NewInt(2))
		}

		// we probably found it
		if upper.Cmp(lower) == 0 {
			break
		}

		printMsg(upper)
	}
	printMsg(upper)
}

func printMsg(pt *big.Int) {
	var msg []byte
	_, err := fmt.Sscanf(pt.Text(16), "%x", &msg)
	if err != nil {
		// We're gonna have EOF Errors
		fmt.Println("decoding errors")
		return
	}
	fmt.Println("Decrypted:", string(msg))
}
