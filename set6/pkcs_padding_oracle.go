package set6

import (
	"bytes"
	"fmt"
	"github.com/dbalan/cryptopals/common"
	"github.com/dbalan/cryptopals/rsa"
	"math/big"
	"os"
)

const size = 256

/* impliment ceil division using floor div
 */
func CeilDiv(x, y *big.Int) *big.Int {
	res := new(big.Int).Add(x, y)
	res.Sub(res, big.NewInt(1))
	res.Div(res, y)
	return res
}

/* padds the messege with pkcs
 */
func PKCSPad(msg []byte, n int) []byte {
	pad := []byte{0, 2}
	ps := n - 3 - len(msg)
	padding, err := common.RandBytes(ps)
	if err != nil {
		fmt.Fprint(os.Stderr, "genrandom failed\n")
		panic(err)
	}
	pad = append(pad, padding...)
	pad = append(pad, 0)
	pad = append(pad, msg...)
	return pad
}

/* pkcsunpad
 * (does not check for the conformity of msg)
 */
func PKCSUnpad(eb []byte) ([]byte, error) {
	padout := bytes.Split(eb, []byte{byte(0)})
	if len(padout) != 2 {
		return nil, fmt.Errorf("got more than two 00 bytes")
	}
	return padout[1], nil
}

func PKCSOracle(m []byte) (*big.Int, *big.Int, *big.Int, func(*big.Int) bool) {
	pub, priv, n, err := rsa.GenKeyPair(size)

	if err != nil {
		fmt.Fprintf(os.Stderr, "generating key failed\n")
		panic(err)
	}

	bytesize := (2 * size) / 8
	paddedMsg := PKCSPad(m, bytesize)

	nlim := new(big.Int).Exp(big.NewInt(2), big.NewInt(504), nil)
	blim := new(big.Int).Exp(big.NewInt(2), big.NewInt(512), nil)

	//
	if n.Cmp(nlim) < 0 || n.Cmp(blim) >= 0 {
		panic("n out of bounds")
	}
	actualCT := rsa.Encrypt(paddedMsg, pub, n)

	oracle := func(ct *big.Int) bool {
		c := new(big.Int).Set(ct)
		eb := rsa.DecryptRaw(c, priv, n)
		if len(eb) == bytesize-1 {
			// check if eb is PKCS conforming
			return (eb[0] == 2)
		} else {
			if eb[0] == 0 && eb[1] == 2 {
				panic("never hold")
			}

		}
		return false
	}
	return actualCT, pub, n, oracle
}

type Ivl struct {
	A *big.Int
	B *big.Int
}

func testConforming(s, ct, n, pub *big.Int, oracle func(*big.Int) bool) bool {
	c := new(big.Int).Exp(s, pub, n)
	c.Mul(ct, c)
	c.Mod(c, n)
	return oracle(c)
}

func PKCSOracleAttack(ct, pub, n *big.Int, oracle func(*big.Int) bool) ([]byte, error) {
	one := big.NewInt(1)
	two := big.NewInt(2)

	// specific constants
	k := big.NewInt(512 / 8)
	k28 := new(big.Int).Sub(k, two)
	k28.Mul(k28, big.NewInt(8))

	B := new(big.Int).Exp(two, k28, nil)
	B2 := new(big.Int).Mul(B, two)
	B3 := new(big.Int).Mul(B, big.NewInt(3))
	B3_1 := new(big.Int).Sub(B3, one)

	// Decrypt steps

	// 1. Blinding
	// C already is PKCS Conforming, set S0 = 1
	if !oracle(ct) {
		// assertion check
		return nil, fmt.Errorf("c is not pkcs conforming")
	}

	// S0 = 1, and Init M
	M := []Ivl{
		Ivl{B2, B3_1},
	}
	// this could be renamed as first, but paper calls it i
	i := 1

	// S1
	s := new(big.Int).Div(n, B3)
	for {
		if len(M) == 0 {
			// this should never arise, we should always
			// have atleast one interval
			return nil, fmt.Errorf("Zero intervals found")
		}
		if !((i > 1) && (len(M) == 1)) {
			if (i > 1) && (len(M) > 1) {
				// step 2b
				s.Add(s, one)
			}

			// step 1 and 2b (same steps, but s is different)
			for !testConforming(s, ct, n, pub, oracle) {
				s.Add(s, one)
			}
			i += 1

		} else {
			inter := M[0]

			if inter.A.Cmp(inter.B) == 0 {
				// step 4
				fmt.Printf("Found M: %x\n", inter.A)
				// unpad msg
				return PKCSUnpad(inter.A.Bytes())
			}

			r := new(big.Int).Mul(inter.B, s)
			r.Sub(r, B2)
			r.Mul(r, two)
			r = CeilDiv(r, n)

			// incrememnt r until we find an s
		step2c:
			for {
				// find s
				s = new(big.Int).Mul(r, n)
				s.Add(s, B2)
				s = CeilDiv(s, inter.B)

				slim := new(big.Int).Mul(r, n)
				slim.Add(slim, B3)
				slim = CeilDiv(slim, inter.A)

				for s.Cmp(slim) < 0 {
					if testConforming(s, ct, n, pub, oracle) {
						break step2c
					}
					s.Add(s, one)
				}

				r.Add(r, one)
			}
		}

		// Step 3
		MNew := []Ivl{}
		for _, inter := range M {
			// ceil division from floor
			// x + n-1/n <= x/n
			r := new(big.Int).Mul(inter.A, s)
			r.Sub(r, B3)
			r.Add(r, one)
			r = CeilDiv(r, n)

			rlim := new(big.Int).Mul(inter.B, s)
			rlim.Sub(rlim, B2)
			rlim.Div(rlim, n)

			for r.Cmp(rlim) <= 0 {
				// new a
				anew := new(big.Int).Mul(r, n)
				anew.Add(anew, B2)
				anew = CeilDiv(anew, s)

				if inter.A.Cmp(anew) > 0 {
					anew = inter.A
				}

				// new b
				bnew := new(big.Int).Mul(r, n)
				bnew.Add(bnew, B3_1)
				bnew.Div(bnew, s)

				if inter.B.Cmp(bnew) < 0 {
					bnew = inter.B
				}

				if anew.Cmp(bnew) <= 0 {
					MNew = append(MNew, Ivl{anew, bnew})
				}
				r.Add(r, one)
			}
		}
		M = MNew
	}
}
